package infinispan

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
)

//Server represents an host and a port
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
		log.Printf("Reverting to default conf because of error in JSON: %s", conf)
		c = &Client{Servers: []Server{Server{Host: "127.0.0.1", Port: 11222}}, CacheName: ""}
	}
	return c.connect()
}

func (c *Client) connect() (*Client, error) {

	server := fmt.Sprintf("%s:%d", c.Servers[0].Host, c.Servers[0].Port)
	log.Printf("Connecting to host: %s\n", server)

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
func (c *Client) Get(key []byte) ([]byte, error) {
	get := createGet(key, <-id, c.CacheName)
	c.connection.Write(get)
	status, err := bufio.NewReader(c.connection).Read(c.buf[:1024])
	if err != nil {
		return []byte{}, err
	}
	p := NewBuffer(c.buf[:status])
	return p.DecodeResponse(GetResponse)
}

//Put puts an object with a key
func (c *Client) Put(key []byte, object []byte) ([]byte, error) {
	return c.realPut(key, object, "0", "0", false)
}

//PutWithOptions puts an object with a key and optional parameters
func (c *Client) PutWithOptions(key []byte, object []byte, opts map[string]string) ([]byte, error) {
	lifespan := opts["lifespan"]
	if lifespan == "" {
		lifespan = "0"
	}

	maxidle := opts["maxidle"]
	if maxidle == "" {
		maxidle = "0"
	}

	returnValues, err := strconv.ParseBool(opts["previous"])

	if err != nil {
		return c.realPut(key, object, lifespan, maxidle, false)
	}
	return c.realPut(key, object, lifespan, maxidle, returnValues)

}

func (c *Client) realPut(key []byte, object []byte, lifespan string, maxidle string, previous bool) ([]byte, error) {
	if put, createErr := createPut(key, object, <-id, c.CacheName, lifespan, maxidle, previous); createErr == nil {
		c.connection.Write(put)
		if status, ioErr := bufio.NewReader(c.connection).Read(c.buf[:1024]); ioErr == nil {
			p := NewBuffer(c.buf[:status])
			return p.DecodeResponse(PutResponse)
		} else {
			return []byte{}, ioErr
		}
	} else {
		return []byte{}, createErr
	}

}
