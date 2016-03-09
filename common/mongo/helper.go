package mongo

import (
	"fmt"

	"github.com/banerwai/micros/common/etcd"
)

type MongoInfo struct {
	ConnString string `json:"connString"`
	Database   string `json:"database"`
}

func (self *MongoInfo) String() string {
	return fmt.Sprintf("[MongoInfo]\nConnString: %s\nDatabase: %s\n", self.ConnString, self.Database)
}

func (self *MongoInfo) Init() bool {
	self.ConnString, _ = etcd.GetValue("/banerwai/mongo/conn")
	self.Database, _ = etcd.GetValue("/banerwai/mongo/database")

	return true
}
