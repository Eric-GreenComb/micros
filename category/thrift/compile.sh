#!/usr/bin/env sh

# See also https://thrift.apache.org/tutorial/go

thrift -r --gen "go:package_prefix=github.com/banerwai/micros/category/thrift/gen-go/,thrift_import=github.com/apache/thrift/lib/go/thrift" category.thrift
