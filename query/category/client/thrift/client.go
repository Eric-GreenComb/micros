package thrift

import (
	"github.com/go-kit/kit/log"

	"github.com/banerwai/micros/query/category/service"
	thriftcategory "github.com/banerwai/micros/query/category/thrift/gen-go/category"
)

// New returns an AddService that's backed by the Thrift client.
func New(cli *thriftcategory.CategoryServiceClient, logger log.Logger) service.CategoryService {
	return &client{cli, logger}
}

type client struct {
	*thriftcategory.CategoryServiceClient
	log.Logger
}

func (c client) Ping() string {
	reply, err := c.CategoryServiceClient.Ping()
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) LoadCategory(path string) bool {
	reply, err := c.CategoryServiceClient.LoadCategory(path)
	if err != nil {
		c.Logger.Log("err", err)
		return false
	}
	return reply
}

func (c client) GetCategories() string {
	reply, _ := c.CategoryServiceClient.GetCategories()
	return reply
}

func (c client) GetSubCategories(serialnumber int32) string {
	_subs, _ := c.CategoryServiceClient.GetSubCategories(serialnumber)
	return _subs
}
