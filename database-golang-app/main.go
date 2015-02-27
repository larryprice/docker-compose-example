package main

import (
	"fmt"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"os"
	"time"
)

type Ping struct {
	Id   bson.ObjectId `bson:"_id"`
	Time time.Time     `bson:"time"`
}

func main() {
	session, err := mgo.Dial(os.Getenv("DATABASE_PORT_27017_TCP_ADDR"))

	if err != nil {
		panic(err)
	}

	db := session.DB(os.Getenv("DB_NAME"))
	defer session.Close()

	ping := Ping{
		Id:   bson.NewObjectId(),
		Time: time.Now(),
	}
	db.C("pings").Insert(ping)

	pings := []Ping{}
	db.C("pings").Find(nil).All(&pings)

	fmt.Println(pings)
}
