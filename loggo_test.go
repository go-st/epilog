package epilog

import (
	"testing"

	"github.com/go-st/logger"
	. "gopkg.in/check.v1"
)

func Test(t *testing.T) { TestingT(t) }

type testProcessor struct {
	entries []*Entry
}

func (p *testProcessor) Process(entry *Entry) {
	p.entries = append(p.entries, entry)
}

type testHandler struct {
	Level   logger.Level
	entries []*Entry
}

func (h *testHandler) Handle(entry *Entry) {
	h.entries = append(h.entries, entry)
}

func (h *testHandler) Copy() IHandler {
	return h
}

func (h *testHandler) IsEnabledFor(level logger.Level) bool {
	return level <= h.Level
}

type handlerForCopy struct {
	original IHandler
}

func (h *handlerForCopy) Handle(entry *Entry) {
}

func (h *handlerForCopy) Copy() IHandler {
	return &handlerForCopy{original: h}
}

func (h *handlerForCopy) IsEnabledFor(level logger.Level) bool {
	return h.original.IsEnabledFor(level)
}
