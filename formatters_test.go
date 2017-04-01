package epilog

import (
	"time"

	"github.com/go-st/logger"
	. "gopkg.in/check.v1"
)

type FormatterTestSuite struct{}

var (
	_ = Suite(&FormatterTestSuite{})
)

func (s *FormatterTestSuite) TestFormat(c *C) {
	formatter := NewTextFormatter("[:time:] (:foo:) :message: const")

	entryTime, _ := time.Parse("2006-01-02T15:04:05", "2015-09-17T16:00:00")

	entry := NewEntry(logger.LevelDebug, entryTime, "hello")
	entry.Fields["foo"] = "bar"
	result := formatter.Format(entry)

	c.Assert(string(result), Equals, "[2015-09-17T16:00:00.000000+00:00] (bar) hello const\n")
}

func (s *FormatterTestSuite) TestFormatWithTimeLayout(c *C) {
	formatter := NewTextFormatterWithTimeLayout("[:time:] :message:", time.RFC3339)

	entryTime, _ := time.Parse("2006-01-02T15:04:05", "2015-09-17T16:00:00")

	entry := NewEntry(logger.LevelDebug, entryTime, "hello world")
	entry.Fields["foo"] = "bar"
	result := formatter.Format(entry)

	c.Assert(string(result), Equals, "[2015-09-17T16:00:00Z] hello world\n")
}
