package infinispan

import (
	"bufio"
	"net"
)

//Client represents a Client connection to an Hot Rod server
type Client struct {
	server     string
	connection net.Conn
	buf        [1024]byte
}

//NewClient creates a new client
func NewClient(s string) (*Client, error) {
	c := &Client{server: s}
	return c.connect()
}

func (c *Client) connect() (*Client, error) {
	conn, err := net.Dial("tcp", c.server)
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
	get := createGet(key, <-id, DefaultCache)
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
	if put, createErr := createPut(key, object, <-id, DefaultCache, lifespan, maxidle); createErr == nil {
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
