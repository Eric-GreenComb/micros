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

	banerwaiglobal "github.com/banerwai/global/constant"
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
	_instancesRandomIndex := banerwaicrypto.GetRandomItNum(len(_instances))

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
	transportSocket, err := thrift.NewTSocket(_instances[_instancesRandomIndex])
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
		_userID := s1
		v := svc.GetAccountByUserID(_userID)
		var _account bean.Account
		json.Unmarshal([]byte(v), &_account)
		fmt.Println(_account)
		logger.Log("method", "GetAccountByUserID", "_userID", _userID, "took", time.Since(begin))

	case "billing":
		_id := s1
		v := svc.GetBillingByID(_id)
		var _billing bean.Billing
		json.Unmarshal([]byte(v), &_billing)
		fmt.Println(_billing)
		logger.Log("method", "GetBillingByID", "took", time.Since(begin))

	case "deal":
		_userID := s1
		_timestamp := time.Now().Unix()
		fmt.Println(_timestamp)
		v := svc.GetDealBillingByUserID(_userID, _timestamp, banerwaiglobal.DefaultPageSize)
		var _billings []bean.Billing
		json.Unmarshal([]byte(v), &_billings)
		for _, _billing := range _billings {
			fmt.Println(_billing.Operate, _billing.Amount)
		}
		logger.Log("method", "GetDealBillingByUserID", "took", time.Since(begin))

	case "all":
		_userID := s1
		v := svc.GetBillingByUserID(_userID, time.Now().Unix(), banerwaiglobal.DefaultPageSize)
		var _billings []bean.Billing
		json.Unmarshal([]byte(v), &_billings)
		for _, _billing := range _billings {
			fmt.Println(_billing.Operate, _billing.Amount)
		}
		logger.Log("method", "GetBillingByUserID", "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
