package mongo

import (
	"fmt"

	"github.com/go-martini/martini"
	"labix.org/v2/mgo"
)

var dbInfo MongoInfo
var mgoSession *mgo.Session

func init() {
	dbInfo.Init()
	fmt.Println(dbInfo)

	var err error

	mgoSession, err = mgo.Dial(dbInfo.ConnString)
	if err != nil {
		fmt.Println(err.Error())
	}
	mgoSession.SetMode(mgo.Monotonic, true)
}

func GetMongo() (*mgo.Session, *mgo.Database) {
	return mgoSession, mgoSession.DB(dbInfo.Database)
}

// Middleware handler for mongodb
func Mongo() martini.Handler {

	return func(c martini.Context) {
		reqSession := mgoSession.Clone()
		c.Map(reqSession.DB(dbInfo.Database))
		defer reqSession.Close()

		c.Next()
	}
}
