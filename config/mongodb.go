package config

import (
	"cryptocurrencies-service/util"
	"github.com/joho/godotenv"
	"gopkg.in/mgo.v2"
	"log"
	"os"
)

type Connection interface {
	Close()
	DB() *mgo.Database
}

type conn struct {
	session *mgo.Session
}

func NewConnection() (Connection, error) {
	var c conn
	var err error

	envError := godotenv.Load(".env")
	if envError != nil {
		log.Fatal("Error loading .env file")
	}

	c.session, err = mgo.Dial(os.Getenv("MONGO_URI"))
	if err != nil {
		return nil, util.ErrNotConnecInDatabase
	}
	return &c, nil
}

func (c *conn) Close() {
	c.session.Close()
}

func (c *conn) DB() *mgo.Database {
	return c.session.DB("")
}
