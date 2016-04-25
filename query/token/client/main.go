package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"

	thriftclient "github.com/banerwai/micros/query/token/client/thrift"
	"github.com/banerwai/micros/query/token/service"
	thrifttoken "github.com/banerwai/micros/query/token/thrift/gen-go/token"

	"github.com/banerwai/gather/command/bean"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:9040", "Address for Thrift server")
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")
	)
	flag.Parse()
	if len(os.Args) < 3 {
		fmt.Fprintf(os.Stderr, "\n%s [flags] method arg1 arg2\n\n", filepath.Base(os.Args[0]))
		flag.Usage()
		os.Exit(1)
	}

	_instances := strings.Split(*thriftAddr, ",")
	_instances_random_index := banerwaicrypto.GetRandomItNum(len(_instances))

	method, s1, s2 := flag.Arg(0), flag.Arg(1), flag.Arg(2)
	_i64, _ := strconv.ParseInt(s2, 10, 64)
	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.TokenService

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
	transportSocket, err := thrift.NewTSocket(_instances[_instances_random_index])
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
	cli := thrifttoken.NewTokenServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {

	case "ver":

		_overHours := getOverHours(_i64)

		v := svc.VerifyToken(s1, _i64, _overHours)
		logger.Log("method", "VerifyToken", "s1", s1, "s2", _i64, "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}

func getOverHours(ttype int64) float64 {
	switch ttype {

	// 0 - 2.0
	case bean.TokenPwd:
		return bean.PwdOverHours

	// 1 - 48.0
	case bean.TokenActiveEmail:
		return bean.ActiveEmailOverHours

	// 2 - 2.0
	case bean.TokenUpdateEmail:
		return bean.UpdateEmailOverHours
	}
	return 0
}
