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
	"github.com/go-kit/kit/metrics"
	"github.com/go-kit/kit/metrics/expvar"
	"github.com/go-kit/kit/metrics/prometheus"

	"github.com/banerwai/micros/query/workhistory/service"
	thriftworkhistory "github.com/banerwai/micros/query/workhistory/thrift/gen-go/workhistory"

	banerwaiglobal "github.com/banerwai/global"
	"github.com/banerwai/gommon/etcd"

	"labix.org/v2/mgo"
)

// 数据连接
var Session *mgo.Session

// WorkHistory表的Collection对象
var WorkHistoryCollection *mgo.Collection

func main() {
	// Flag domain. Note that gRPC transitively registers flags via its import
	// of glog. So, we define a new flag set, to keep those domains distinct.
	fs := flag.NewFlagSet("", flag.ExitOnError)
	var (
		thriftAddr       = fs.String("thrift.addr", ":39080", "Address for Thrift server")
		thriftProtocol   = fs.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = fs.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = fs.Bool("thrift.framed", false, "true to enable framing")

		mongodbUrl    = fs.String("mongodb.url", "127.0.0.1:27017", "mongodb url")
		mongodbDbname = fs.String("mongodb.dbname", "banerwai", "mongodb dbname")
	)
	flag.Usage = fs.Usage // only show our flags
	if err := fs.Parse(os.Args[1:]); err != nil {
		fmt.Fprintf(os.Stderr, "%v", err)
		os.Exit(1)
	}

	var err error
	Session, err = mgo.Dial(*mongodbUrl) //连接数据库
	if err != nil {
		panic(err)
	}
	defer Session.Close()
	Session.SetMode(mgo.Monotonic, true)

	WorkHistoryCollection = Session.DB(*mongodbDbname).C("workhistory") //数据库名称

	// package log
	var logger log.Logger
	{
		logger = log.NewLogfmtLogger(os.Stderr)
		logger = log.NewContext(logger).With("ts", log.DefaultTimestampUTC).With("caller", log.DefaultCaller)
		stdlog.SetFlags(0)                             // flags are handled by Go kit's logger
		stdlog.SetOutput(log.NewStdlibAdapter(logger)) // redirect anything using stdlib log to us
	}

	// package metrics
	var requestDuration metrics.TimeHistogram
	{
		requestDuration = metrics.NewTimeHistogram(time.Nanosecond, metrics.NewMultiHistogram(
			"request_duration_ns",
			expvar.NewHistogram("request_duration_ns", 0, 5e9, 1, 50, 95, 99),
			prometheus.NewSummary(stdprometheus.SummaryOpts{
				Namespace: "banerwai",
				Subsystem: "resume",
				Name:      "duration_ns",
				Help:      "Request duration in nanoseconds.",
			}, []string{"method"}),
		))
	}

	// Business domain
	var svc service.WorkHistoryService
	{
		svc = newInmemService()
		svc = loggingMiddleware{svc, logger}
		svc = instrumentingMiddleware{svc, requestDuration}
	}

	// Mechanical stuff
	rand.Seed(time.Now().UnixNano())
	errc := make(chan error)

	go func() {
		errc <- interrupt()
	}()

	client := etcd.EtcdReigistryClient{
		etcd.EtcdRegistryConfig{
			ServiceName:  banerwaiglobal.ETCD_KEY_MICROS_QUERY_WORKHISTORY,
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
			thriftworkhistory.NewWorkHistoryServiceProcessor(thriftBinding{svc}),
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
