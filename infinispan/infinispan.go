package infinispan

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strings"
)

type Server struct {
	Host string
	Port uint16
}

//Client represents a Client connection to an Hot Rod server
type Client struct {
	Servers   []Server
	CacheName string

	connection net.Conn
	buf        [1024]byte
}

//NewClientJSON creates a new client from a JSON file
func NewClientJSON(conf string) (*Client, error) {
	var c *Client
	if err := json.NewDecoder(strings.NewReader(conf)).Decode(&c); err != nil {
		//return &Client{Servers: []&Server{Host: "127.0.0.1", Port: 11222}, CacheName: ""}, err
		return nil, err
	}
	log.Printf("Connecting to host %s:%d\n", c.Servers[0].Host, c.Servers[0].Port)
	return c.connect()
}

func (c *Client) connect() (*Client, error) {

	server := fmt.Sprintf("%s:%d", c.Servers[0].Host, c.Servers[0].Port)
	conn, err := net.Dial("tcp", server)
	c.connection = conn
	return c, err
}

//Close Hot Rod Connection
func (c *Client) Close() error {
	if c.connection != nil {
		return c.connection.Close()
	}
	return nil
}

//Get gets a key
func (c *Client) Get(key []byte) (*ResponseGet, error) {
	get := createGet(key, <-id, c.CacheName)
	c.connection.Write(get)
	status, err := bufio.NewReader(c.connection).Read(c.buf[:1024])
	if err != nil {
		return &ResponseGet{}, err
	}
	p := NewBuffer(c.buf[:status])
	return p.DecodeGetResponse()
}

//Put puts an object with a key
func (c *Client) Put(key []byte, object []byte) (*ResponsePut, error) {
	return c.PutWithLifespanAndMaxidle(key, object, "0", "0")
}

//PutWithLifespan puts an object with a key and a lifespan
func (c *Client) PutWithLifespan(key []byte, object []byte, lifespan string) (*ResponsePut, error) {
	return c.PutWithLifespanAndMaxidle(key, object, lifespan, "0")
}

//PutWithMaxidle puts an object with a key and a maxidle
func (c *Client) PutWithMaxidle(key []byte, object []byte, maxidle string) (*ResponsePut, error) {
	return c.PutWithLifespanAndMaxidle(key, object, "0", maxidle)
}

//PutWithLifespanAndMaxidle puts an object with a key and a lifespan/maxidle
func (c *Client) PutWithLifespanAndMaxidle(key []byte, object []byte, lifespan string, maxidle string) (*ResponsePut, error) {
	if put, createErr := createPut(key, object, <-id, c.CacheName, lifespan, maxidle); createErr == nil {
		c.connection.Write(put)
		if status, ioErr := bufio.NewReader(c.connection).Read(c.buf[:1024]); ioErr == nil {
			p := NewBuffer(c.buf[:status])
			return p.DecodePutResponse()
		} else {
			return &ResponsePut{}, ioErr
		}
	} else {
		return &ResponsePut{}, createErr
	}

}
