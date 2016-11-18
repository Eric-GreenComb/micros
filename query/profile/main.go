package main

import (
	"flag"
	"fmt"
	stdlog "log"
	"math/rand"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/apache/thrift/lib/go/thrift"
	stdprometheus "github.com/prometheus/client_golang/prometheus"

	"github.com/go-kit/kit/log"
	kitprometheus "github.com/go-kit/kit/metrics/prometheus"

	"github.com/banerwai/micros/query/profile/service"
	thriftprofile "github.com/banerwai/micros/query/profile/thrift/gen-go/profile"

	banerwaiglobal "github.com/banerwai/global/constant"
	"github.com/banerwai/gommon/etcd"

	"gopkg.in/mgo.v2"
)

// Session 数据连接
var Session *mgo.Session

// ProfileCollection Profile表的Collection对象
var ProfileCollection *mgo.Collection

func main() {
	// Flag domain. Note that gRPC transitively registers flags via its import
	// of glog. So, we define a new flag set, to keep those domains distinct.
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		thriftAddr       = fs.String("thrift.addr", ":39050", "Address for Thrift server")
		thriftProtocol   = fs.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = fs.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = fs.Bool("thrift.framed", false, "true to enable framing")

		mongodbURL    = fs.String("mongodb.url", "127.0.0.1:27017", "mongodb url")
		mongodbDbname = fs.String("mongodb.dbname", "banerwai", "mongodb dbname")
	)
	flag.Usage = fs.Usage // only show our flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	var err error
	Session, err = mgo.Dial(*mongodbURL) //连接数据库
	if err != nil {
		panic(err)
	}
	defer Session.Close()
	Session.SetMode(mgo.Monotonic, true)

	ProfileCollection = Session.DB(*mongodbDbname).C("profiles") //数据库名称

	// package log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC).With("caller", log.DefaultCaller)
		stdlog.SetFlags(0)                             // flags are handled by Go kit's logger
		stdlog.SetOutput(log.NewStdlibAdapter(logger)) // redirect anything using stdlib log to us
	}

	// package metrics
	fieldKeys := []string{"method", "error"}
	requestCount := kitprometheus.NewCounterFrom(stdprometheus.CounterOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_count",
		Help:      "Number of requests received.",
	}, fieldKeys)
	requestLatency := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "request_latency_microseconds",
		Help:      "Total duration of requests in microseconds.",
	}, fieldKeys)
	countResult := kitprometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
		Namespace: "my_group",
		Subsystem: "string_service",
		Name:      "count_result",
		Help:      "The result of each count method.",
	}, []string{}) // no fields here

	// Business domain
	var svc service.ProfileService
	{
		svc = newInmemService()
		svc = loggingMiddleware{svc, logger}
		svc = instrumentingMiddleware{requestCount, requestLatency, countResult, svc}
	}

	// Mechanical stuff
	rand.Seed(time.Now().UnixNano())
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	client := etcd.ReigistryClient{
		etcd.RegistryConfig{
			ServiceName:  banerwaiglobal.EtcdKeyMicrosQueryProfile,
			InstanceName: *thriftAddr,
			BaseURL:      *thriftAddr,
		},
		etcd.KeysAPI,
	}
	client.Register()
	defer client.Unregister()

	// Transport: Thrift
	go func() {
		var protocolFactory thrift.TProtocolFactory
		switch *thriftProtocol {
		case "binary":
			protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
		case "compact":
			protocolFactory = thrift.NewTCompactProtocolFactory()
		case "json":
			protocolFactory = thrift.NewTJSONProtocolFactory()
		case "simplejson":
			protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
		default:
			errc <- fmt.Errorf("invalid Thrift protocol %q", *thriftProtocol)
			return
		}
		var transportFactory thrift.TTransportFactory
		if *thriftBufferSize > 0 {
			transportFactory = thrift.NewTBufferedTransportFactory(*thriftBufferSize)
		} else {
			transportFactory = thrift.NewTTransportFactory()
		}
		if *thriftFramed {
			transportFactory = thrift.NewTFramedTransportFactory(transportFactory)
		}
		transport, err := thrift.NewTServerSocket(*thriftAddr)
		if err != nil {
			errc <- err
			return
		}
		transportLogger := log.NewContext(logger).With("transport", "thrift")
		transportLogger.Log("addr", *thriftAddr)
		errc <- thrift.NewTSimpleServer4(
			thriftprofile.NewProfileServiceProcessor(thriftBinding{svc}),
			transport,
			transportFactory,
			protocolFactory,
		).Serve()
	}()

	logger.Log("fatal", <-errc)
}

func interrupt() error {
	c := make(chan os.Signal)
	signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	return fmt.Errorf("%s", <-c)
}

type loggingCollector struct{ log.Logger }
