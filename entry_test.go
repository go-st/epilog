package loggo

import (
	"time"

	"bitbucket.org/lazadaweb/go-logger"
	. "gopkg.in/check.v1"
)

type EntryTestSuite struct{}

var (
	_ = Suite(&EntryTestSuite{})
)

func (s *EntryTestSuite) TestNewEntry(c *C) {
	now := time.Now()
	entry := NewEntry(logger.LevelWarning, now, "some message")

	c.Assert(entry.Level, Equals, logger.LevelWarning)
	c.Assert(entry.Message, Equals, "some message")
	c.Assert(entry.Time, Equals, now)
	c.Assert(entry.Fields, HasLen, 0)
}
