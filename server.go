package easyrpc

import (
	"github.com/hezhis/easyrpc/protocol"
	logger "github.com/hezhis/go_log"
)

type Handler func(ctx *Context)

type Server struct {
	router map[string]Handler
}

func NewServer(options ...OptionFn) *Server {
	s := &Server{router: make(map[string]Handler)}

	for _, opt := range options {
		opt(s)
	}

	return s
}

func (s *Server) AddHandler(servicePath, serviceMethod string, handler Handler) {
	s.router[servicePath+"."+serviceMethod] = handler
}

func (s *Server) DoCall(req *protocol.Message) {
	handler, ok := s.router[req.ServicePath+"."+req.ServiceMethod]
	if !ok {
		logger.Error("easy rpc handler not found! servicePath:%s, serviceMethod:%s", req.ServicePath, req.ServiceMethod)
		return
	}
	handler(NewContext(req))
}
