package db

import (
	"gopkg.in/mgo.v2"
	"log"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}

type conn struct {
	session *mgo.Session
}

func NewConnection() Connection {
	var c conn
	var err error

	url := "mongodb://localhost:27017"
	c.session, err = mgo.Dial(url)
	if err != nil {
		log.Panicln(err.Error())
	}
	return &c
}

func (c *conn) Close() {
	c.session.Close()
}

func (c *conn) DB() *mgo.Database {
	return c.session.DB("")
}
