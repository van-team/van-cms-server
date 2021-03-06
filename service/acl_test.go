package service

import (
	"context"
	. "gopkg.in/check.v1"
)

func (x *MySuite) TestAclFetch(c *C) {
	result, err := x.acl.Fetch(context.Background(), "resource", "0")
	if err != nil {
		c.Error(err)
	}
	c.Log(result)
	result, err = x.acl.Fetch(context.Background(), "resource", "1")
	if err != nil {
		c.Error(err)
	}
	c.Log(result)
}

func (x *MySuite) TestAclClear(c *C) {
	err := x.acl.Clear(context.Background())
	if err != nil {
		c.Error(err)
	}
}
