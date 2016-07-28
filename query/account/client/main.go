package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	// "strconv"
	"strings"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"

	banerwaiglobal "github.com/banerwai/global"
	thriftclient "github.com/banerwai/micros/query/account/client/thrift"
	"github.com/banerwai/micros/query/account/service"
	thriftaccount "github.com/banerwai/micros/query/account/thrift/gen-go/account"

	"github.com/banerwai/global/bean"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:39100", "Address for Thrift server")
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

	_instances := strings.Split(*thriftAddr, ",")
	_instances_random_index := banerwaicrypto.GetRandomItNum(len(_instances))

	method, s1 := flag.Arg(0), flag.Arg(1)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.AccountService

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
	cli := thriftaccount.NewAccountServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {

	case "ping":
		v := svc.Ping()
		logger.Log("method", "Ping", s1, "v", v, "took", time.Since(begin))

	case "account":
		_user_id := s1
		v := svc.GetAccountByUserId(_user_id)
		var _account bean.Account
		json.Unmarshal([]byte(v), &_account)
		fmt.Println(_account)
		logger.Log("method", "GetAccountByUserId", "_user_id", _user_id, "took", time.Since(begin))

	case "billing":
		_id := s1
		v := svc.GetBillingById(_id)
		var _billing bean.Billing
		json.Unmarshal([]byte(v), &_billing)
		fmt.Println(_billing)
		logger.Log("method", "GetBillingById", "took", time.Since(begin))

	case "deal":
		_user_id := s1
		v := svc.GetDealBillingByUserId(_user_id, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)
		var _billings []bean.Billing
		json.Unmarshal([]byte(v), &_billings)
		for _, _billing := range _billings {
			fmt.Println(_billing.Operate, _billing.Amount)
		}
		logger.Log("method", "GetDealBillingByUserId", "took", time.Since(begin))

	case "all":
		_user_id := s1
		v := svc.GetBillingByUserId(_user_id, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)
		var _billings []bean.Billing
		json.Unmarshal([]byte(v), &_billings)
		for _, _billing := range _billings {
			fmt.Println(_billing.Operate, _billing.Amount)
		}
		logger.Log("method", "GetBillingByUserId", "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
