package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"labix.org/v2/mgo/bson"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"

	thriftclient "github.com/banerwai/micros/command/contact/client/thrift"
	"github.com/banerwai/micros/command/contact/service"
	thriftcontact "github.com/banerwai/micros/command/contact/thrift/gen-go/contact"

	"github.com/banerwai/global/bean"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
	bstrings "github.com/banerwai/gommon/strings"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:36090", "Address for Thrift server")
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

	method, _contactID := flag.Arg(0), flag.Arg(1)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.ContactService

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
	cli := thriftcontact.NewContactServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {
	case "ping":
		v := svc.Ping()
		logger.Log("method", "Ping", "v", v, "took", time.Since(begin))

	// case "prepare":
	// 	_tpl := prepareContactTpl()
	// 	_mmap := prepareContactParam()

	// 	b, _ := json.Marshal(_mmap)
	// 	fmt.Println(string(b))

	// 	_contact := bstrings.ParseTpl("default", _tpl, _mmap)
	// 	fmt.Println(_contact)
	case "create":
		var _obj bean.Contact
		_obj.ID = bson.NewObjectId()
		_obj.ClientEmail = "ministor@gmail.com"
		_obj.FreeLancerEmail = "ministor@126.com"

		_tpl := prepareContactTpl()
		_mmap := prepareContactParam()
		_bParam, _ := json.Marshal(_mmap)

		// _content := bstrings.ParseTpl("default", _tpl, _mmap)

		// _obj.ContactContent = _content

		_obj.ContactTpl = _tpl
		_obj.TplParam = string(_bParam)

		b, _ := json.Marshal(_obj)
		v := svc.CreateContact(string(b))
		logger.Log("method", "CreateContact", "v", v, "took", time.Since(begin))

	case "csign":
		v := svc.ClientSignContact(_contactID, true)
		logger.Log("method", "ClientSignContact", "v", v, "took", time.Since(begin))

	case "fsign":
		v := svc.FreelancerSignContact(_contactID, true)
		logger.Log("method", "FreelancerSignContact", "v", v, "took", time.Since(begin))

	case "deal":
		v := svc.DealContact(_contactID, true)
		logger.Log("method", "DealContact", "v", v, "took", time.Since(begin))

	case "update":
		_mapUpdate := make(map[string]string)

		_tpl := prepareContactTpl()
		_mmap := updateContactParam()
		_bParam, _ := json.Marshal(_mmap)

		_content := bstrings.ParseTpl("default", _tpl, _mmap)

		_mapUpdate["contact_content"] = _content
		_mapUpdate["contact_tpl"] = _tpl
		_mapUpdate["tpl_param"] = string(_bParam)

		v := svc.UpdateContact(_contactID, _mapUpdate)
		logger.Log("method", "UpdateContact", "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}

func prepareContactTpl() string {
	_tpl := `合同编号:{{.ContactNumber}}

# {{.ProjectName}}合同

**甲方 : {{.ClientName}}**

**乙方 : {{.FreeLancerName}}**

**预期工作时间：{{.ExpectedTime}}**

**费用标准 : {{.PayRate}}**

**支付方式 : {{.PaymentMethod}}**

引言： 依据《中华人民共和国合同法》的有关规定，（以下简称甲方）与（以下简称乙方）就甲方软件项目的服务外包合作事宜，经协商达成一致，签订本合同，以资共同遵守。
	`
	return _tpl
}

func prepareContactParam() map[string]string {
	_map := make(map[string]string)
	_map["ContactNumber"] = bson.NewObjectId().Hex()
	_map["ProjectName"] = "banerwai服务"
	_map["ClientName"] = "ministor@gmail.com"
	_map["FreeLancerName"] = "ministor@126.com"
	_map["ExpectedTime"] = "4weeks"
	_map["PayRate"] = "150Y / hour"
	_map["PaymentMethod"] = "week"
	return _map
}

func updateContactParam() map[string]string {
	_map := make(map[string]string)
	_map["ContactNumber"] = bson.NewObjectId().Hex()
	_map["ProjectName"] = "banerwai服务"
	_map["ClientName"] = "ministor@gmail.com"
	_map["FreeLancerName"] = "ministor@126.com"
	_map["ExpectedTime"] = "8weeks"
	_map["PayRate"] = "200Y / hour"
	_map["PaymentMethod"] = "day"
	return _map
}
