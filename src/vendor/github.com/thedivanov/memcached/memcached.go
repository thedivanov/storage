package memcached

import (
	"bufio"
	"bytes"
	"fmt"
	"net"
	"time"
)

var (
	crlf      = []byte("\r\n")
	resultEnd = []byte("END")
)

type Conn struct {
	conn net.Conn
}

func NewConn(MemcachedAddr string) (*Conn, error) {
	conn, err := net.Dial("tcp", MemcachedAddr)
	if err != nil {
		return nil, err
	}
	return &Conn{
		conn: conn,
	}, nil
}

type Item struct {
	Key        string
	Value      string
	Flags      uint32
	Expiration int32
}

func (c *Conn) Set(key string, val string, expire int64) error {
	var exp int
	if expire == 0 {
		exp = 0
	} else {
		exp = int(time.Unix(expire, 0).Sub(time.Now()).Seconds())
	}
	_, err := c.conn.Write([]byte(fmt.Sprintf("set %s 0 %d %d%s%s%s", key, exp, len(val), crlf, val, crlf)))
	return err
}

func (c *Conn) Get(key string) (Item, error) {
	_, err := c.conn.Write([]byte(fmt.Sprintf("get %s%s", key, crlf)))
	if err != nil {
		return Item{}, err
	}
	scanner := bufio.NewScanner(c.conn)
	var item Item

	for scanner.Scan() {
		if bytes.Equal(scanner.Bytes(), resultEnd) {
			break
		} else if bytes.Count(scanner.Bytes(), []byte(" ")) == 3 {
			fmt.Sscanf(scanner.Text(), "VALUE %s %d %d", &item.Key, &item.Expiration, &item.Flags)
			if scanner.Scan() {
				item.Value = scanner.Text()
			} else {
				return Item{}, fmt.Errorf("Memcached response corrupt")
			}
		}
	}
	if err := scanner.Err(); err != nil {
		return Item{}, err
	}
	if item.Value == "" {
		return Item{}, fmt.Errorf("Not found")
	}

	return item, nil
}

func (c *Conn) Delete(key string) {
	c.conn.Write([]byte(fmt.Sprintf("delete %s%s", key, crlf)))
}

func (c *Conn) Close() error {
	return c.conn.Close()
}
