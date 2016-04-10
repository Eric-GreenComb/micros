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

func (c client) SayHi(name string) string {
	reply, err := c.CategoryServiceClient.SayHi(name)
	if err != nil {
		c.Logger.Log("err", err)
		return ""
	}
	return reply
}

func (c client) GetDemoSubCategory(id string) thriftcategory.SubCategory {
	reply, err := c.CategoryServiceClient.GetDemoSubCategory(id)
	if err != nil {
		c.Logger.Log("err", err)
		return thriftcategory.SubCategory{}
	}
	return *reply
}

func (c client) GetDemoSubCategories(category_id string) []thriftcategory.SubCategory {
	var subs []thriftcategory.SubCategory
	_subs, err := c.CategoryServiceClient.GetDemoSubCategories(category_id)

	if err != nil {
		c.Logger.Log("err", err)
		return nil
	}

	for _, _sub := range _subs {
		subs = append(subs, *_sub)
	}

	return subs
}

func (c client) LoadCategory(path string) bool {
	reply, err := c.CategoryServiceClient.LoadCategory(path)
	if err != nil {
		c.Logger.Log("err", err)
		return false
	}
	return reply
}

func (c client) GetCategories() []*thriftcategory.Category {
	reply, _ := c.CategoryServiceClient.GetCategories()
	return reply
}

func (c client) GetSubCategories(serialnumber int32) []*thriftcategory.SubCategory {
	_subs, _ := c.CategoryServiceClient.GetSubCategories(serialnumber)
	return _subs
}
