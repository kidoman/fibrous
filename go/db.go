package main

import "github.com/garyburd/redigo/redis"

type DB struct {
	p *redis.Pool
}

func newDB(p *redis.Pool) *DB {
	return &DB{p}
}

func (db *DB) conn() redis.Conn {
	return db.p.Get()
}

func (db *DB) LoadUser(id int) (*User, error) {
	c := db.conn()
	defer c.Close()

	name, err := redis.String(c.Do("GET", UserKey(id)))
	if err != nil {
		return nil, err
	}

	return &User{ID: id, Name: name}, nil
}

func (db *DB) SaveUser(u *User) error {
	c := db.conn()
	defer c.Close()

	_, err := c.Do("SET", u.Key(), u.Name)
	return err
}

func (db *DB) Close() error {
	return db.p.Close()
}
