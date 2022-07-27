package easyrpc

import (
	"github.com/hezhis/easyrpc/protocol"
	"github.com/hezhis/easyrpc/share"
	logger "github.com/hezhis/go_log"
)

type (
	PostHandler func(msg *protocol.Message)

	Client struct {
		handler PostHandler
	}
)

func NewClient(handler PostHandler) *Client {
	client := &Client{
		handler: handler,
	}

	if nil == client.handler {
		return nil
	}

	return client
}

func (client *Client) call(serializeType protocol.SerializeType, servicePath, serviceMethod string, args interface{}) {
	codec := share.Codecs[serializeType]
	if codec == nil {
		logger.Error("easy tcp call error! undefined serialize type %d", serializeType)
		return
	}

	data, err := codec.Encode(args)
	if err != nil {
		logger.Error("easy tcp call encode args error! %v", err)
		return
	}

	req := protocol.NewMessage()
	req.SetSerializeType(serializeType)
	req.ServicePath = servicePath
	req.ServiceMethod = serviceMethod
	req.Payload = data

	if share.Trace {
		logger.Debug("client.send for %s.%s, args: %+v in case of client call", servicePath, serviceMethod, args)
	}
	client.handler(req)
}

func (client *Client) CallPb(servicePath, serviceMethod string, args interface{}) {
	client.call(protocol.ProtoBuffer, servicePath, serviceMethod, args)
}

func (client *Client) CallGob(servicePath, serviceMethod string, args interface{}) {
	client.call(protocol.GobBuffer, servicePath, serviceMethod, args)
}
