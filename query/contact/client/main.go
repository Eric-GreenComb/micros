package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"gopkg.in/mgo.v2/bson"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"

	thriftclient "github.com/banerwai/micros/query/contact/client/thrift"
	"github.com/banerwai/micros/query/contact/service"
	thriftcontact "github.com/banerwai/micros/query/contact/thrift/gen-go/contact"

	"github.com/banerwai/global/bean"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:39090", "Address for Thrift server")
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

	case "tpl":
		tplname := s1
		v := svc.GetContactTpl(tplname)
		logger.Log("method", "GetContactTpl", "tplname", tplname, "v", v, "took", time.Since(begin))

	case "get":
		contactid := s1
		v := svc.GetContact(contactid)

		if len(v) == 0 {
			fmt.Println("error : return is null")
		}

		var _contact bean.Contact
		err := json.Unmarshal([]byte(v), &_contact)
		if err != nil {
			fmt.Println("error : ", err)
		}

		fmt.Println(_contact)

		logger.Log("method", "GetContact", "took", time.Since(begin))

	case "client":
		clientemail := s1
		v := svc.GetClientContact(clientemail)

		var _contacts []bean.Contact
		err := json.Unmarshal([]byte(v), &_contacts)
		if err != nil {
			fmt.Println("error:", err)
		}

		for _, _contact := range _contacts {
			fmt.Println(_contact.ContactContent)
		}

		logger.Log("method", "GetClientContact", "_contacts", len(_contacts), "took", time.Since(begin))

	case "freelancer":
		freelanceremail := s1
		v := svc.GetFreelancerContact(freelanceremail)

		var _contacts []bean.Contact
		err := json.Unmarshal([]byte(v), &_contacts)
		if err != nil {
			fmt.Println("error:", err)
		}

		for _, _contact := range _contacts {
			fmt.Println(_contact.ContactContent)
		}

		logger.Log("method", "GetFreelancerContact", "_contacts", len(_contacts), "took", time.Since(begin))

	case "check":
		_contactid := s1
		v := svc.GetContactSignStatus(_contactid)
		mmap := bson.M{}
		bson.Unmarshal([]byte(v), mmap)
		if mmap["client_signup"] == true && mmap["freelancer_signup"] == true {
			fmt.Println(true)
		}
		logger.Log("method", "GetContactSignStatus", "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
