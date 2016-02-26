package main

import (
	"github.com/banerwai/micros/category/bean"
	"github.com/go-kit/kit/log"
	"golang.org/x/net/context"
	"time"
)

type loggingMiddleware struct {
	next   CategoryService
	logger log.Logger
}

func (mw loggingMiddleware) LoadCategoriesFile(ctx context.Context, file string) (err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "LoadCategoriesFile", "file", file, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.LoadCategoriesFile(ctx, file)
}

func (mw loggingMiddleware) GetSubCategories(ctx context.Context, categoryID string) (subcategories []bean.SubCategory, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetSubCategories", "categoryID", categoryID, "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetSubCategories(ctx, categoryID)
}

func (mw loggingMiddleware) GetCategories(ctx context.Context) (categories []bean.Category, err error) {
	defer func(begin time.Time) {
		mw.logger.Log("method", "GetCategories", "took", time.Since(begin), "err", err)
	}(time.Now())
	return mw.next.GetCategories(ctx)
}
