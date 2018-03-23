package controllers

import (
	"gopkg.in/mgo.v2"

	"github.com/salar-hafezi/go-taskmanager/common"
)

// Context used for maintaining HTTP request context
type Context struct {
	MongoSession *mgo.Session
	User         string
}

// close session
func (c *Context) Close() {
	c.MongoSession.Close()
}

// return mgo.Collection for a given name
func (c *Context) DbCollection(name string) *mgo.Collection {
	return c.MongoSession.DB(common.AppConfig.Database).C(name)
}

// create a new Context for each HTTP request
func NewContext() *Context {
	session := common.GetSession().Copy()
	context := &Context{
		MongoSession: session,
	}
	return context
}
