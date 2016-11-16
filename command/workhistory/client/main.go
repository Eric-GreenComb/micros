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

	thriftclient "github.com/banerwai/micros/command/workhistory/client/thrift"
	"github.com/banerwai/micros/command/workhistory/service"
	thriftworkhistory "github.com/banerwai/micros/command/workhistory/thrift/gen-go/workhistory"

	banerwaicrypto "github.com/banerwai/gommon/crypto"

	"github.com/banerwai/global/bean"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:36080", "Address for Thrift server")
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")

		_defaultObjectID = flag.String("default.user.ojbectid", "5707cb10ae6faa1d1071a189", "default user ojbectid")
	)
	flag.Parse()
	if len(os.Args) < 2 {
		fmt.Fprintf(os.Stderr, "\n%s [flags] method arg1 arg2\n\n", filepath.Base(os.Args[0]))
		flag.Usage()
		os.Exit(1)
	}

	_instances := strings.Split(*thriftAddr, ",")
	_instancesRandomIndex := banerwaicrypto.GetRandomItNum(len(_instances))

	method := flag.Arg(0)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.WorkHistoryService

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
	cli := thriftworkhistory.NewWorkHistoryServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {
	case "ping":
		v := svc.Ping()
		logger.Log("method", "Ping", "v", v, "took", time.Since(begin))

	case "upsert":

		var _obj bean.WorkHistory
		_obj.ID = bson.ObjectIdHex(*_defaultObjectID)
		_obj.ProfileID = bson.ObjectIdHex(*_defaultObjectID)

		var lsWorkHistoryAndFeedbacks []bean.WorkHistoryAndFeedback

		var _WorkHistoryAndFeedback1 bean.WorkHistoryAndFeedback
		_WorkHistoryAndFeedback1.Title = "ceshi"
		_WorkHistoryAndFeedback1.WorkPeriod = "2016.01-2016.04"
		_WorkHistoryAndFeedback1.WorkHours = 40

		var lsWorkFeedbacks []bean.WorkFeedback

		var _WorkFeedback1 bean.WorkFeedback
		_WorkFeedback1.WorkRate = 5
		_WorkFeedback1.Feedback = "perfect"
		lsWorkFeedbacks = append(lsWorkFeedbacks, _WorkFeedback1)

		var _WorkFeedback2 bean.WorkFeedback
		_WorkFeedback2.WorkRate = 5
		_WorkFeedback2.Feedback = "good job"

		lsWorkFeedbacks = append(lsWorkFeedbacks, _WorkFeedback2)

		_WorkHistoryAndFeedback1.WorkFeedbacks = lsWorkFeedbacks

		lsWorkHistoryAndFeedbacks = append(lsWorkHistoryAndFeedbacks, _WorkHistoryAndFeedback1)

		_obj.HistoryAndFeedbacks = lsWorkHistoryAndFeedbacks

		b, _ := json.Marshal(_obj)

		v := svc.UpdateWorkHistory(*_defaultObjectID, string(b))
		logger.Log("method", "UpdateWorkHistory", "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
