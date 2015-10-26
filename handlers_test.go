package loggo

import (
	"bytes"
	"time"

	"bitbucket.org/lazadaweb/go-logger"
	. "gopkg.in/check.v1"
)

type HandlersTestSuite struct {
	rawFormatter IFormatter
}

var (
	_ = Suite(&HandlersTestSuite{})
)

func (s *HandlersTestSuite) SetUpSuite(c *C) {
	s.rawFormatter = NewTextFormatter(":message:")
}

func (s *HandlersTestSuite) TestStreamHandlerHandle(c *C) {
	buf := &bytes.Buffer{}
	handler := NewStreamHandler(logger.LevelDebug, s.rawFormatter, buf)
	handler.Handle(NewEntry(logger.LevelDebug, time.Now(), "hello"))
	handler.Handle(NewEntry(logger.LevelInfo, time.Now(), "man"))

	c.Assert(buf.String(), Equals, "hello\nman\n")
}

func (s *HandlersTestSuite) TestStreamHandlerHandleLowLevel(c *C) {
	buf := &bytes.Buffer{}
	handler := NewStreamHandler(logger.LevelInfo, s.rawFormatter, buf)
	handler.Handle(NewEntry(logger.LevelDebug, time.Now(), "hello"))
	handler.Handle(NewEntry(logger.LevelInfo, time.Now(), "man"))

	c.Assert(buf.String(), Equals, "man\n")
}

func (s *HandlersTestSuite) TestStreamHandlerCopy(c *C) {
	buf := &bytes.Buffer{}
	handler := NewStreamHandler(logger.LevelInfo, s.rawFormatter, buf)

	c.Assert(handler.Copy(), Equals, handler)
}

func (s *HandlersTestSuite) TestBufferHandlerHandle(c *C) {
	buf := &bytes.Buffer{}
	streamHandler := NewStreamHandler(logger.LevelDebug, s.rawFormatter, buf)

	handler := NewBufferHandler(streamHandler, logger.LevelWarning)
	handler.Handle(NewEntry(logger.LevelDebug, time.Now(), "debug"))
	handler.Handle(NewEntry(logger.LevelInfo, time.Now(), "info"))
	handler.Handle(NewEntry(logger.LevelWarning, time.Now(), "warning"))

	c.Assert(buf.String(), Equals, "debug\ninfo\nwarning\n")
}

func (s *HandlersTestSuite) TestBufferHandlerHandleLowLevel(c *C) {
	buf := &bytes.Buffer{}
	streamHandler := NewStreamHandler(logger.LevelDebug, s.rawFormatter, buf)

	handler := NewBufferHandler(streamHandler, logger.LevelWarning)
	handler.Handle(NewEntry(logger.LevelDebug, time.Now(), "debug"))
	handler.Handle(NewEntry(logger.LevelInfo, time.Now(), "info"))

	c.Assert(buf.String(), Equals, "")
}

func (s *HandlersTestSuite) TestBufferHandlerCopy(c *C) {
	handler := &handlerForCopy{}
	bufferHandler := NewBufferHandler(handler, logger.LevelWarning)
	copy := bufferHandler.Copy().(*BufferHandler)
	c.Assert(copy, Not(Equals), bufferHandler)
	c.Assert(copy.flushLevel, Equals, bufferHandler.flushLevel)
	c.Assert(copy.handler.(*handlerForCopy).original, Equals, handler)
}

func (s *HandlersTestSuite) TestMultiHandlerHandle(c *C) {
	buf := &bytes.Buffer{}
	streamHandler := NewStreamHandler(logger.LevelDebug, s.rawFormatter, buf)

	buf2 := &bytes.Buffer{}
	streamHandler2 := NewStreamHandler(logger.LevelInfo, s.rawFormatter, buf2)

	handler := NewMultiHandler(streamHandler, streamHandler2)
	handler.Handle(NewEntry(logger.LevelDebug, time.Now(), "debug"))
	handler.Handle(NewEntry(logger.LevelInfo, time.Now(), "info"))

	c.Assert(buf.String(), Equals, "debug\ninfo\n")
	c.Assert(buf2.String(), Equals, "info\n")
}

func (s *HandlersTestSuite) TestMultiHandlerCopy(c *C) {
	handler := &handlerForCopy{}
	handler2 := &handlerForCopy{}

	multiHandler := NewMultiHandler(handler, handler2)
	copy := multiHandler.Copy().(*MultiHandler)

	c.Assert(copy, Not(Equals), multiHandler)
	c.Assert(copy.handlers[0].(*handlerForCopy).original, Equals, handler)
	c.Assert(copy.handlers[1].(*handlerForCopy).original, Equals, handler2)
}
