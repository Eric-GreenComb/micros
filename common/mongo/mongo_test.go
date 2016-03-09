package mongo

import (
	"fmt"
	"labix.org/v2/mgo/bson"
	"testing"
)

type Person struct {
	NAME  string
	PHONE string
}
type Men struct {
	Persons []Person
}

func TestGetMongo(t *testing.T) {
	session, db := GetMongo()
	defer session.Close()

	collection := db.C("person") //如果该集合已经存在的话，则直接返回

	//*****集合中元素数目********
	countNum, err := collection.Count()
	if err != nil {
		panic(err)
	}
	fmt.Println("Things objects count: ", countNum)

	//*******插入元素*******
	temp := &Person{
		PHONE: "18811577546",
		NAME:  "zhangzheHero",
	}
	//一次可以插入多个对象 插入两个Person对象
	err = collection.Insert(&Person{"Ale", "+55 53 8116 9639"}, temp)
	if err != nil {
		panic(err)
	}

	//*****查询单条数据*******
	result := Person{}
	err = collection.Find(bson.M{"phone": "18811577546"}).One(&result)
	fmt.Println("Phone:", result.NAME, result.PHONE)

	//*****查询多条数据*******
	var personAll Men //存放结果
	iter := collection.Find(nil).Iter()
	for iter.Next(&result) {
		fmt.Printf("Result: %v\n", result)
		personAll.Persons = append(personAll.Persons, result)
	}

	//*******更新数据**********
	err = collection.Update(bson.M{"name": "ccc"}, bson.M{"$set": bson.M{"name": "ddd"}})
	err = collection.Update(bson.M{"name": "ddd"}, bson.M{"$set": bson.M{"phone": "12345678"}})
	err = collection.Update(bson.M{"name": "aaa"}, bson.M{"phone": "1245", "name": "bbb"})

	//******删除数据************
	//_, err = collection.RemoveAll(bson.M{"name": "Ale"})
	//_, err = collection.RemoveAll(bson.M{})
}
