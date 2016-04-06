package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/apache/thrift/lib/go/thrift"

	"github.com/go-kit/kit/log"

	thriftclient "github.com/banerwai/micros/user/client/thrift"
	"github.com/banerwai/micros/user/service"
	thriftuser "github.com/banerwai/micros/user/thrift/gen-go/user"
)

func main() {
	var (
		thriftAddr       = flag.String("thrift.addr", "localhost:6006", "Address for Thrift server")
		thriftProtocol   = flag.String("thrift.protocol", "binary", "binary, compact, json, simplejson")
		thriftBufferSize = flag.Int("thrift.buffer.size", 0, "0 for unbuffered")
		thriftFramed     = flag.Bool("thrift.framed", false, "true to enable framing")
	)
	flag.Parse()
	if len(os.Args) < 4 {
		fmt.Fprintf(os.Stderr, "\n%s [flags] method arg1 arg2\n\n", filepath.Base(os.Args[0]))
		flag.Usage()
		os.Exit(1)
	}

	method, s1, s2, s3 := flag.Arg(0), flag.Arg(1), flag.Arg(2), flag.Arg(3)

	var logger log.Logger
	logger = log.NewLogfmtLogger(os.Stdout)
	logger = log.NewContext(logger).With("caller", log.DefaultCaller)

	var svc service.UserService

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
	cli := thriftuser.NewUserServiceClientFactory(trans, protocolFactory)
	svc = thriftclient.New(cli, logger)

	begin := time.Now()
	switch method {

	case "create":
		email := s1
		usernameraw := s2
		pwd := s3
		v := svc.CreateUser(email, usernameraw, pwd)
		logger.Log("method", "CreateUser", "email", email, "usernameraw", usernameraw, "pwd", pwd, "v", v, "took", time.Since(begin))

	case "chpwd":
		email := s1
		oldpwd := s2
		newpwd := s3
		v := svc.UpdatePwd(email, oldpwd, newpwd)
		logger.Log("method", "UpdatePwd", "email", email, "oldpwd", oldpwd, "newpwd", newpwd, "v", v, "took", time.Since(begin))

	case "active":
		token := s1
		v := svc.ActiveUser(token)
		logger.Log("method", "ActiveUser", "token", token, "v", v, "took", time.Since(begin))

	case "count":
		v := svc.CountUser()
		logger.Log("method", "CountUser", "v", v, "took", time.Since(begin))

	default:
		logger.Log("err", "invalid method "+method)
		os.Exit(1)
	}
}
