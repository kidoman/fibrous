package main

import (
	"os"
	"strconv"
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

func redisConnParams() int {
	capStr := os.Getenv("REDIS_CAP")
	if capStr == "" {
		capStr = "20"
	}
	cap, err := strconv.Atoi(capStr)
	if err != nil {
		panic(err)
	}

	return cap
}

const redisMaxCap = 200

func newPool(server string) *pools.ResourcePool {
	f := func() (pools.Resource, error) {
		c, err := redis.Dial("tcp", server)
		return &resourceConn{c}, err
	}

	cap := redisConnParams()

	return pools.NewResourcePool(f, cap, redisMaxCap, time.Minute)
}
