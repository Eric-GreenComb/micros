package main

import (
	"github.com/banerwai/micros/category/bean"
	"github.com/go-kit/kit/endpoint"
	"golang.org/x/net/context"
)

type endpoints struct {
	getSubCategoriesEndpoint   endpoint.Endpoint
	loadCategoriesFileEndpoint endpoint.Endpoint
	getCategoriesEndpoint      endpoint.Endpoint
}

func makeEndpoints(s CategoryService) endpoints {
	return endpoints{
		loadCategoriesFileEndpoint: makeLoadCategoriesFileEndpoint(s),
		getSubCategoriesEndpoint:   makeGetSubCategoriesEndpoint(s),
		getCategoriesEndpoint:      makeGetCategoriesEndpoint(s),
	}
}

type getSubCategoriesRequest struct {
	CategoryID string
}

type getSubCategoriesResponse struct {
	SubCategories []bean.SubCategory `json:"subcategories,omitempty"`
	Err           error              `json:"err,omitempty"`
}

func (r getSubCategoriesResponse) error() error { return r.Err }

func makeGetSubCategoriesEndpoint(s CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		req := request.(getSubCategoriesRequest)
		a, e := s.GetSubCategories(ctx, req.CategoryID)
		return getSubCategoriesResponse{SubCategories: a, Err: e}, nil
	}
}

type loadCategoriesFileRequest struct {
}

type loadCategoriesFileResponse struct {
	Success string `json:"success,omitempty"`
	Err     error  `json:"err,omitempty"`
}

func (r loadCategoriesFileResponse) error() error { return r.Err }

func makeLoadCategoriesFileEndpoint(s CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(loadCategoriesFileRequest)
		e := s.LoadCategoriesFile(ctx, "categories.json")
		return loadCategoriesFileResponse{Success: "load json file", Err: e}, nil
	}
}

type getCategoriesRequest struct {
}

type getCategoriesResponse struct {
	Categories []bean.Category `json:"categories,omitempty"`
	Err        error           `json:"err,omitempty"`
}

func (r getCategoriesResponse) error() error { return r.Err }

func makeGetCategoriesEndpoint(s CategoryService) endpoint.Endpoint {
	return func(ctx context.Context, request interface{}) (response interface{}, err error) {
		// req := request.(getCategoriesRequest)
		a, e := s.GetCategories(ctx)
		return getCategoriesResponse{Categories: a, Err: e}, nil
	}
}
