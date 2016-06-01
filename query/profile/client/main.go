package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"

	banerwaiglobal "github.com/banerwai/global"
	thriftclient "github.com/banerwai/micros/query/profile/client/thrift"
	"github.com/banerwai/micros/query/profile/service"
	thriftprofile "github.com/banerwai/micros/query/profile/thrift/gen-go/profile"

	"github.com/banerwai/global/bean"
	banerwaicrypto "github.com/banerwai/gommon/crypto"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:39050", "Address for Thrift server")
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
	cli := thriftprofile.NewProfileServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {

	case "ping":
		v := svc.Ping()
		logger.Log("method", "Ping", "v", v, "took", time.Since(begin))

	case "prof":
		profile_id := s1
		v := svc.GetProfile(profile_id)
		var _profile bean.Profile
		json.Unmarshal([]byte(v), &_profile)
		fmt.Println(_profile)
		logger.Log("method", "GetProfile", "profile_id", profile_id, "v", v, "took", time.Since(begin))

	case "profs":
		_userid := s1
		v := svc.GetProfilesByUserId(_userid)
		var _profiles []bean.Profile
		json.Unmarshal([]byte(v), &_profiles)
		fmt.Println(_profiles)
		logger.Log("method", "GetProfilesByUserId", "v", v, "took", time.Since(begin))

	case "cat":
		_cat_id, _ := strconv.ParseInt(s1, 10, 64)
		v := svc.GetProfilesByCategory(_cat_id, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)
		var _profiles []bean.Profile
		json.Unmarshal([]byte(v), &_profiles)
		for _, _prof := range _profiles {
			fmt.Println(_prof.JobTitle, _prof.SerialNumber)
		}
		logger.Log("method", "GetProfilesByUserId", "took", time.Since(begin))

	case "subcat":
		_subcat_id, _ := strconv.ParseInt(s1, 10, 64)
		v := svc.GetProfilesBySubCategory(_subcat_id, 1464785330, banerwaiglobal.Pagination_PAGESIZE_Web)
		var _profiles []bean.Profile
		json.Unmarshal([]byte(v), &_profiles)
		for _, _prof := range _profiles {
			fmt.Println(_prof.SerialNumber, _prof.LastActiveTime.Unix())
		}
		logger.Log("method", "GetProfilesByUserId", "took", time.Since(begin))

	case "search":
		_key := s1

		option_mmap := make(map[string]int64)

		option_mmap["freelancer_type"] = 0
		option_mmap["job_success"] = 0

		key_mmap := make(map[string]string)
		key_mmap["job_title"] = _key

		v := svc.SearchProfiles(option_mmap, key_mmap, time.Now().Unix(), banerwaiglobal.Pagination_PAGESIZE_Web)

		var _profiles []bean.Profile
		json.Unmarshal([]byte(v), &_profiles)
		fmt.Println(_profiles)

		logger.Log("method", "SearchProfiles", "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
