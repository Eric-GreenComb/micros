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

	"github.com/banerwai/gather/command/bean"
	thriftclient "github.com/banerwai/micros/command/category/client/thrift"
	"github.com/banerwai/micros/command/category/service"
	thriftcategory "github.com/banerwai/micros/command/category/thrift/gen-go/category"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:6010", "Address for Thrift server")
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

	method, s1 := flag.Arg(0), flag.Arg(1)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.CategoryService

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
	cli := thriftcategory.NewCategoryServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {

	case "hi":
		name := s1
		v := svc.SayHi(name)
		logger.Log("method", "SayHi", "name", name, "v", v, "took", time.Since(begin))

	case "demosub":
		id := s1
		v := svc.GetDemoSubCategory(id)
		_sub := bean.SubCategory{v.ID, v.Serialnumber, v.Name, v.Desc}

		// _sub.ID = v.ID
		// _sub.Name = v.Name

		logger.Log("method", "GetSubCategory", "id", id, "v", _sub.ID, "took", time.Since(begin))

	case "demosubs":
		category_id := s1
		v := svc.GetDemoSubCategories(category_id)
		logger.Log("method", "GetSubCategory", "category_id", category_id, "v", fmt.Sprintf("%v", v), "took", time.Since(begin))

	case "load":
		path := s1
		v := svc.LoadCategory(path)
		logger.Log("method", "LoadCategory", "path", path, "v", v, "took", time.Since(begin))

	case "cats":
		v := svc.GetCategories()
		for _, _cat := range v {
			fmt.Println(strconv.Itoa(int(_cat.Serialnumber)) + " " + _cat.Name + "(" + _cat.Desc + ")")
			for _, _sub := range _cat.Subcategories {
				fmt.Println("    " + strconv.Itoa(int(_sub.Serialnumber)) + " " + _sub.Name + "(" + _sub.Desc + ")")
			}
		}

		logger.Log("method", "GetCategories", "v", len(v), "took", time.Since(begin))

	case "subs":
		_serialnumber, _ := strconv.Atoi(s1)
		v := svc.GetSubCategories(int32(_serialnumber))

		for _, _sub := range v {
			fmt.Println("    " + strconv.Itoa(int(_sub.Serialnumber)) + " " + _sub.Name + "(" + _sub.Desc + ")")
		}

		logger.Log("method", "GetSubCategories", "Serialnumber", _serialnumber, "v", len(v), "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}