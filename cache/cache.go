package cache

import (
	"bytes"
	"encoding/gob"
	"fmt"

	"github.com/gomodule/redigo/redis"
)

type Cache interface {
	Has(string) (bool, error)
	Get(string) (interface{}, error)
	Set(string, interface{}, ...int) error
	Forget(string) error
	EmptyByMatch(string) error
	Empty() error
}

type RedisCache struct {
	Conn   *redis.Pool
	Prefix string
}

type Entry map[string]interface{}

func encode(item Entry) ([]byte, error) {
	b := bytes.Buffer{}
	e := gob.NewEncoder(&b)
	err := e.Encode(item)
	if err != nil {
		return nil, err
	}

	return b.Bytes(), nil
}

func decode(str string) (Entry, error) {
	item := Entry{}
	b := bytes.Buffer{}
	b.Write([]byte(str))
	d := gob.NewDecoder(&b)
	err := d.Decode(&item)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (c *RedisCache) Has(str string) (bool, error) {
	key := fmt.Sprintf("%s:%s", c.Prefix, str)
	conn := c.Conn.Get()
	defer conn.Close()

	ok, err := redis.Bool(conn.Do("EXISTS", key))
	if err != nil {
		return false, err
	}

	return ok, nil
}

func (c *RedisCache) Get(str string) (interface{}, error) {
	return "", nil
}

func (c *RedisCache) Set(str string, data interface{}, ttl ...int) error {
	return nil
}

func (c *RedisCache) Forget(str string) error {
	return nil
}

func (c *RedisCache) EmptyByMatch(str string) error {
	return nil
}

func (c *RedisCache) Empty() error {
	return nil
}
