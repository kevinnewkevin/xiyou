package prpc

import "bytes"

const(
	NoneDispatchMatchError = "NoneDispatchMatchError"
	NoneMethodError = "NoneMethodError"
	NoneBufferError = "NoneBufferError"
	NoneProxyError = "NoneProxyError"
)

type StubSender interface {
	MethodBegin() *bytes.Buffer
	MethodEnd() error
}