package infinispan

import (
	"bufio"
	"fmt"
	"net"
)

type Client struct {
	server     string
	connection net.Conn
}

func NewClient(s string) *Client {
	return &Client{server: s}
}

func (c *Client) Connect() {
	conn, err := net.Dial("tcp", c.server)
	c.connection = conn
	if err != nil {

	}
}

func (c *Client) Close() {
	if c.connection != nil {
		c.connection.Close()
	}
}

func TestConnect() {

	c, err := net.Dial("tcp", "127.0.0.1:11222")

	if err != nil {
		fmt.Println(err)
		return
	}

	//put := createPut([]byte("2"), []byte("ugol"), <-id, DEFAULT_CACHE)
	get := createGet([]byte("2"), <-id, DefaultCache)
	fmt.Println(get)
	var buf [1024]byte

	w := bufio.NewWriter(c)
	w.Write(get)
	w.Flush()
	status, err := bufio.NewReader(c).Read(buf[0:256])
	fmt.Printf("Status: %d\n", status)
	fmt.Println(buf[:status])

	p := NewBuffer(buf[:status])
	res, err := p.DecodeGetResponse()
	fmt.Println(p.index)
	fmt.Println(res)
	if err != nil {
		fmt.Println(err)
	}

	c.Close()
}
