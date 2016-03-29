package infinispan

import (
	"bufio"
	"net"
)

//Connection represents a connection to an Hot Rod server
type Connection struct {
	server     string
	connection net.Conn
	buf        [1024]byte
}

//NewConnection creates a new client
func NewConnection(s string) (*Connection, error) {
	c := &Connection{server: s}
	return c.connect()
}

func (c *Connection) connect() (*Connection, error) {
	conn, err := net.Dial("tcp", c.server)
	c.connection = conn
	return c, err
}

//Close Hot Rod Connection
func (c *Connection) Close() error {
	if c.connection != nil {
		return c.connection.Close()
	}
	return nil
}

//Get gets a key
func (c *Connection) Get(key []byte) (*GetRes, error) {
	get := createGet(key, <-id, DefaultCache)
	c.connection.Write(get)
	status, err := bufio.NewReader(c.connection).Read(c.buf[:1024])
	if err != nil {
		return &GetRes{}, err
	}
	p := NewBuffer(c.buf[:status])
	return p.DecodeGetResponse()
}

//Put puts an object with a key
func (c *Connection) Put(key []byte, object []byte) (*PutRes, error) {
	put := createPut(key, object, <-id, DefaultCache)
	c.connection.Write(put)
	status, err := bufio.NewReader(c.connection).Read(c.buf[:1024])
	if err != nil {
		return &PutRes{}, err
	}
	p := NewBuffer(c.buf[:status])
	return p.DecodePutResponse()
}
