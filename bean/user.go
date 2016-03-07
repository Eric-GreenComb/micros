package bean

import (
	"labix.org/v2/mgo/bson"
	"time"
)

type User struct {
	Id          bson.ObjectId `bson:"_id,omitempty"`           // 必须要设置bson:"_id" 不然mgo不会认为是主键
	Email       string        `Email json:"email"`             // 全是小写
	Verified    bool          `Verified json:"verified"`       // Email是否已验证过?
	Username    string        `Username json:"username"`       // 不区分大小写, 全是小写
	UsernameRaw string        `UsernameRaw json:"usernameraw"` // 可能有大小写
	Pwd         string        `bson:"Pwd" json:"-"`
	CreatedTime time.Time     `CreatedTime json:"createdtime"`
}
