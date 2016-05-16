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

	thriftclient "github.com/banerwai/micros/command/resume/client/thrift"
	"github.com/banerwai/micros/command/resume/service"
	thriftresume "github.com/banerwai/micros/command/resume/thrift/gen-go/resume"

	banerwaicrypto "github.com/banerwai/gommon/crypto"

	"github.com/banerwai/global/bean"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:36070", "Address for Thrift server")
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")

		_defaultObjectId = flag.String("default.user.ojbectid", "5707cb10ae6faa1d1071a189", "default user ojbectid")
	)
	flag.Parse()
	if len(os.Args) < 1 {
		fmt.Fprintf(os.Stderr, "\n%s [flags] method arg1 arg2\n\n", filepath.Base(os.Args[0]))
		flag.Usage()
		os.Exit(1)
	}

	_instances := strings.Split(*thriftAddr, ",")
	_instances_random_index := banerwaicrypto.GetRandomItNum(len(_instances))

	method := flag.Arg(0)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.ResumeService

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
	cli := thriftresume.NewResumeServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {
	case "ping":
		v := svc.Ping()
		logger.Log("method", "Ping", "v", v, "took", time.Since(begin))

	case "create":
		var _obj bean.Resume
		_obj.Id = bson.ObjectIdHex(*_defaultObjectId)
		_obj.AuthEmail = "ministor@126.com"
		_obj.UserID = bson.ObjectIdHex(*_defaultObjectId)

		_obj.Phone = "12345678901"

		var lsToolandArchs []bean.ToolandArch

		var _tool1 bean.ToolandArch
		_tool1.ToolLevel = 5
		_tool1.ToolTitle = "Java"
		lsToolandArchs = append(lsToolandArchs, _tool1)

		var _tool2 bean.ToolandArch
		_tool2.ToolLevel = 2
		_tool2.ToolTitle = "Go"
		lsToolandArchs = append(lsToolandArchs, _tool2)

		_obj.ToolandArchs = lsToolandArchs

		b, _ := json.Marshal(_obj)
		v := svc.AddResume(string(b))
		logger.Log("method", "AddResume", "v", v, "took", time.Since(begin))

	case "update":
		var _obj bean.Resume
		_obj.Id = bson.ObjectIdHex(*_defaultObjectId)
		_obj.AuthEmail = "ministor@126.com"
		_obj.UserID = bson.ObjectIdHex(*_defaultObjectId)

		_obj.Phone = "12345678901"

		var lsToolandArchs []bean.ToolandArch

		var _tool1 bean.ToolandArch
		_tool1.ToolLevel = 5
		_tool1.ToolTitle = "Java+"
		lsToolandArchs = append(lsToolandArchs, _tool1)

		var _tool2 bean.ToolandArch
		_tool2.ToolLevel = 2
		_tool2.ToolTitle = "Go+"
		lsToolandArchs = append(lsToolandArchs, _tool2)

		_obj.ToolandArchs = lsToolandArchs

		b, _ := json.Marshal(_obj)
		v := svc.UpdateResume(*_defaultObjectId, string(b))
		logger.Log("method", "UpdateResumeBase", "v", v, "took", time.Since(begin))

	case "updatebase":
		_map_update := make(map[string]string)
		_map_update["auth_email"] = "ministor@126.com"
		_map_update["phone"] = "13811111111"
		v := svc.UpdateResumeBase(*_defaultObjectId, _map_update)
		logger.Log("method", "UpdateResumeBase", "v", v, "took", time.Since(begin))

	case "updatetools":
		var lsToolandArchs []bean.ToolandArch

		var _tool1 bean.ToolandArch
		_tool1.ToolLevel = 5
		_tool1.ToolTitle = "Java++"
		lsToolandArchs = append(lsToolandArchs, _tool1)

		var _tool2 bean.ToolandArch
		_tool2.ToolLevel = 2
		_tool2.ToolTitle = "Go++"
		lsToolandArchs = append(lsToolandArchs, _tool2)

		b, _ := json.Marshal(lsToolandArchs)

		v := svc.UpdateResumeToolandArchs(*_defaultObjectId, string(b))
		logger.Log("method", "UpdateResumeToolandArchs", "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
