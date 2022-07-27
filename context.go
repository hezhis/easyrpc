package easyrpc

import (
	"fmt"
	"github.com/hezhis/easyrpc/protocol"
	"github.com/hezhis/easyrpc/share"
)

type Context struct {
	req *protocol.Message
}

// NewContext creates a server.Context for Handler.
func NewContext(req *protocol.Message) *Context {
	return &Context{req: req}
}

// Payload returns the  payload.
func (ctx *Context) Payload() []byte {
	return ctx.req.Payload
}

// ServicePath returns the ServicePath.
func (ctx *Context) ServicePath() string {
	return ctx.req.ServicePath
}

// ServiceMethod returns the ServiceMethod.
func (ctx *Context) ServiceMethod() string {
	return ctx.req.ServiceMethod
}

// Bind parses the body data and stores the result to v.
func (ctx *Context) Bind(v interface{}) error {
	req := ctx.req
	if v != nil {
		codec := share.Codecs[req.SerializeType()]
		if codec == nil {
			return fmt.Errorf("can not find codec for %d", req.SerializeType())
		}

		err := codec.Decode(req.Payload, v)
		if err != nil {
			return err
		}
	}
	return nil
}
