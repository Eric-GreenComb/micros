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

	thriftclient "github.com/banerwai/micros/command/account/client/thrift"
	"github.com/banerwai/micros/command/account/service"
	thriftaccount "github.com/banerwai/micros/command/account/thrift/gen-go/account"

	banerwaicrypto "github.com/banerwai/gommon/crypto"

	"github.com/banerwai/global/bean"
	"github.com/banerwai/global/constant"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:36100", "Address for Thrift server")
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

	method, a1 := flag.Arg(0), flag.Arg(1)

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
		logger.Log("method", "Ping", "v", v, "took", time.Since(begin))

	case "create":
		var _obj bean.Account
		_obj.ID = bson.ObjectIdHex(*_defaultObjectID)
		_obj.UserID = bson.ObjectIdHex(*_defaultObjectID)
		_obj.Email = a1

		b, _ := json.Marshal(_obj)
		v := svc.CreateAccount(string(b))
		logger.Log("method", "CreateAccount", "v", v, "took", time.Since(begin))

	case "bill":
		var _obj bean.Billing
		_obj.ID = bson.ObjectIdHex(*_defaultObjectID)
		_obj.UserID = bson.ObjectIdHex(*_defaultObjectID)
		_obj.PayUserID = bson.ObjectIdHex(*_defaultObjectID)
		_obj.ServiceID = bson.ObjectIdHex(*_defaultObjectID)
		_obj.LinkID = bson.ObjectIdHex(*_defaultObjectID)

		_obj.Operate = 1
		_obj.Currency = constant.CurrencyCNY
		_obj.Amount = 2000
		_obj.PayType = constant.PayTypeBankRemittance

		b, _ := json.Marshal(_obj)
		v := svc.CreateBilling(string(b))
		logger.Log("method", "CreateBilling", "v", v, "took", time.Since(begin))

	case "deal":

		_billingID := a1
		v := svc.DealBilling(_billingID)
		logger.Log("method", "DealBilling", "v", v, "took", time.Since(begin))

	case "gen":
		_userID := *_defaultObjectID
		v := svc.GenAccount(_userID)
		logger.Log("method", "GenAccount", "v", v, "took", time.Since(begin))

	// case "update":

	// 	var _obj bean.Profile
	// 	_obj.Id = bson.ObjectIdHex(_profile_id)
	// 	_obj.UserID = bson.ObjectIdHex(*_defaultObjectId)
	// 	_obj.Name = "Test1"
	// 	_obj.JobTitle = "this is a title1"
	// 	_obj.CategoryNumber = 10100
	// 	_obj.SerialNumber = 10101
	// 	_obj.HourRate = 21234
	// 	_obj.WorkHours = 40

	// 	var lsAgencyMembers []bean.AgencyMember
	// 	var AgencyMember01 bean.AgencyMember
	// 	AgencyMember01.Email = "ministor@126.com"
	// 	AgencyMember01.Manager = true
	// 	lsAgencyMembers = append(lsAgencyMembers, AgencyMember01)
	// 	_obj.AgencyMembers = lsAgencyMembers

	// 	b, _ := json.Marshal(_obj)
	// 	v := svc.UpdateProfile(_profile_id, string(b))
	// 	logger.Log("method", "UpdateProfile", "v", v, "took", time.Since(begin))

	// case "status":
	// 	v := svc.UpdateProfileStatus(_profile_id, false)
	// 	logger.Log("method", "UpdateProfileStatus", "v", v, "took", time.Since(begin))

	// case "base":
	// 	_map_update := make(map[string]string)
	// 	_map_update["freelancer_name"] = "freelancer_name"
	// 	_map_update["job_title"] = "job_title"
	// 	_map_update["hour_rate"] = "1601234"
	// 	_map_update["portfolio_nums"] = "4"

	// 	v := svc.UpdateProfileBase(_profile_id, _map_update)
	// 	logger.Log("method", "UpdateProfileBase", "v", v, "took", time.Since(begin))

	// case "member":

	// 	var lsAgencyMembers []bean.AgencyMember

	// 	var _obj1 bean.AgencyMember
	// 	_obj1.Email = "ministor@126.com"
	// 	_obj1.Manager = false
	// 	lsAgencyMembers = append(lsAgencyMembers, _obj1)

	// 	b, _ := json.Marshal(lsAgencyMembers)

	// 	v := svc.UpdateProfileAgencyMembers(_profile_id, string(b))
	// 	logger.Log("method", "UpdateProfileAgencyMembers", "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
