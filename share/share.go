package share

import (
	"github.com/hezhis/easyrpc/codec"
	"github.com/hezhis/easyrpc/protocol"
)

// Trace is a flag to write a trace log or not.
// You should not enable this flag for product environment and enable it only for test.
// It writes trace log with logger Debug level.
var Trace bool

// Codecs are codecs supported by rpcx. You can add customized codecs in Codecs.
var Codecs = map[protocol.SerializeType]codec.Codec{
	protocol.SerializeNone: nil,
	protocol.ProtoBuffer:   &codec.PBCodec{},
	protocol.GobBuffer:     &codec.GobCodec{},
}
