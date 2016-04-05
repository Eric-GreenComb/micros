package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"

	banerwaiglobal "github.com/banerwai/gather/global"
	thriftclient "github.com/banerwai/micros/profile/client/thrift"
	"github.com/banerwai/micros/profile/service"
	thriftprofile "github.com/banerwai/micros/profile/thrift/gen-go/profile"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:6005", "Address for Thrift server")
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")
	)
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "\n%s [flags] method arg1 arg2\n\n", filepath.Base(os.Args[0]))
		flag.Usage()
		os.Exit(1)
	}

	method, s1 := flag.Arg(0), flag.Arg(1)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.ProfileService

	var protocolFactory thrift.TProtocolFactory
	switch *thriftProtocol {
	case "compact":
		protocolFactory = thrift.NewTCompactProtocolFactory()
	case "simplejson":
		protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
	case "json":
		protocolFactory = thrift.NewTJSONProtocolFactory()
	case "binary", "":
		protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
	default:
		logger.Log("protocol", *thriftProtocol, "err", "invalid protocol")
		os.Exit(1)
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
	transportSocket, err := thrift.NewTSocket(*thriftAddr)
	if err != nil {
		logger.Log("during", "thrift.NewTSocket", "err", err)
		os.Exit(1)
	}
	trans := transportFactory.GetTransport(transportSocket)
	defer trans.Close()
	if err := trans.Open(); err != nil {
		logger.Log("during", "thrift transport.Open", "err", err)
		os.Exit(1)
	}
	cli := thriftprofile.NewProfileServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {

	case "get":
		profile_id := s1
		v := svc.GetProfile(profile_id)
		logger.Log("method", "GetProfile", "profile_id", profile_id, "v", v, "took", time.Since(begin))

	case "search":
		var profile_search thriftprofile.ProfileSearchCondition
		_serial_number, _ := strconv.Atoi(s1)
		profile_search.SerialNumber = int32(_serial_number)
		profile_search.HoursBilled = -1
		profile_search.AvailableHours = -1
		v := svc.SearchProfiles(&profile_search, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)
		logger.Log("method", "SearchProfiles", "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
