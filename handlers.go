package loggo

import (
	"io"
	"os"
	"sync"

	"bitbucket.org/lazadaweb/go-logger"
)

// IHandler interface
type IHandler interface {
	Handle(entry *Entry)
	Copy() IHandler
	IsEnabledFor(level logger.Level) bool
}

// BufferHandler is a handler collecting entries in buffer
type BufferHandler struct {
	handler    IHandler
	flushLevel logger.Level
	buffer     []*Entry
	lock       sync.Mutex
}

// NewBufferHandler creates new BufferHandler
func NewBufferHandler(handler IHandler, flushLevel logger.Level) *BufferHandler {
	return &BufferHandler{
		handler:    handler,
		flushLevel: flushLevel,
		buffer:     make([]*Entry, 0, 128),
	}
}

// Handle processes Entry
func (h *BufferHandler) Handle(entry *Entry) {
	h.lock.Lock()
	defer h.lock.Unlock()

	h.buffer = append(h.buffer, entry)
	if entry.Level <= h.flushLevel {
		for _, e := range h.buffer {
			h.handler.Handle(e)
		}

		h.buffer = make([]*Entry, 0, 128)
	}
}

// Copy creates copy of current logger
func (h *BufferHandler) Copy() IHandler {
	return NewBufferHandler(h.handler.Copy(), h.flushLevel)
}

// IsEnabledFor checks whether level is enabled or not
func (h *BufferHandler) IsEnabledFor(level logger.Level) bool {
	return h.handler.IsEnabledFor(level)
}

// MultiHandler is a handler consisting of another handlers
type MultiHandler struct {
	handlers []IHandler
}

// NewMultiHandler creates new MultiHandler
func NewMultiHandler(handlers ...IHandler) *MultiHandler {
	return &MultiHandler{
		handlers: handlers,
	}
}

// Handle processes Entry
func (h *MultiHandler) Handle(entry *Entry) {
	for _, handler := range h.handlers {
		handler.Handle(entry)
	}
}

// Copy creates copy of current logger
func (h *MultiHandler) Copy() IHandler {
	handlers := make([]IHandler, len(h.handlers))
	for i, handler := range h.handlers {
		handlers[i] = handler.Copy()
	}

	return NewMultiHandler(handlers...)
}

// IsEnabledFor checks whether level is enabled or not
func (h *MultiHandler) IsEnabledFor(level logger.Level) bool {
	for _, handler := range h.handlers {
		if handler.IsEnabledFor(level) {
			return true
		}
	}

	return false
}

// StreamHandler is a handler using io.Writer as output source
type StreamHandler struct {
	level        logger.Level
	outputWriter io.Writer
	formatter    IFormatter
}

// NewStreamHandler returns new StreamHandler.
// If out is not passed - stdout will be used
func NewStreamHandler(level logger.Level, formatter IFormatter, outputWriter ...io.Writer) *StreamHandler {
	var outputW io.Writer

	if len(outputWriter) > 1 {
		panic("You can't pass more than one outputWriter")
	}

	if len(outputWriter) == 1 {
		outputW = outputWriter[0]
	} else {
		outputW = os.Stdout
	}

	return &StreamHandler{
		level:        level,
		outputWriter: outputW,
		formatter:    formatter,
	}
}

// Handle processes Entry
func (h *StreamHandler) Handle(entry *Entry) {
	if h.IsEnabledFor(entry.Level) {
		h.outputWriter.Write(h.formatter.Format(entry))
	}
}

// Copy creates copy of current logger
func (h *StreamHandler) Copy() IHandler {
	return h
}

// IsEnabledFor checks whether level is enabled or not
func (h *StreamHandler) IsEnabledFor(level logger.Level) bool {
	return level <= h.level
}
