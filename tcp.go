package tcp

import (
	"context"
	"fmt"
	"go.k6.io/k6/js/common"
	"go.k6.io/k6/js/modules"
	"net"
)

func init() {
	modules.Register("k6/x/tcp", new(Tcp))
}

type Tcp struct {}

type Option struct {
	Host	string	`json:"host"`
	Port	int		`json:"port"`
}

type Client struct {
	ctx 			*context.Context
	opts 			*Option
	conn 			net.Conn
	eventHandlers	map[string][]EventHandler
	msgChannel		chan []byte
}

type EventHandler func(...interface{})

func (t *Tcp) XClient(ctx *context.Context) interface{}  {
	rt := common.GetRuntime(*ctx)
	client := &Client{
		ctx: ctx,
		msgChannel: make(chan []byte),
		eventHandlers: make(map[string][]EventHandler),
	}
	return common.Bind(rt, client, ctx)
}

func (c *Client) Connect(opts *Option) error {
	c.opts = opts
	address := fmt.Sprintf("%s:%d", c.opts.Host, c.opts.Port)
	conn, err := net.Dial("tcp", address)
	if err != nil {
		return err
	}
	c.conn = conn
	go c.readConn()
	return nil
}

func (c *Client) On(event string, handler EventHandler) {
	c.eventHandlers[event] = append(c.eventHandlers[event], handler)
}

func (c *Client) handleEvent(event string, args ...interface{}) {
	if handlers, ok := c.eventHandlers[event]; ok {
		for _, handler := range handlers {
			handler(args...)
		}
	}
}

func (c *Client) readConn() {
	for {
		buf := make([]byte, 4096)
		n, err := c.conn.Read(buf[:])
		if err != nil {
			c.handleEvent("error", err)
			break
		}
		c.handleEvent("data", buf[:n], "data123")
	}
}

func (c *Client) Write(data []byte) {
	//var dstByteArr []byte
	//fmt.Println(reflect.TypeOf(data).String())
	//switch data.(type) {
	//case string:
	//	dstByteArr = []byte(data.(string))
	//	break
	//case []byte:
	//	dstByteArr = data.([]byte)
	//	break
	//default:
	//	return
	//}
	_, err := c.conn.Write(data)
	if err != nil {
		c.handleEvent("error", err)
	}
}
