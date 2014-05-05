package main

import (
	"time"

	"github.com/garyburd/redigo/redis"
	"github.com/youtube/vitess/go/pools"
)

type resourceConn struct {
	redis.Conn
}

func (r *resourceConn) Close() {
	r.Conn.Close()
}

func newPool(server string) *pools.ResourcePool {
	f := func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", server)
		return &resourceConn{c}, err
	}
	return pools.NewResourcePool(f, 3, 20, time.Minute)
}
