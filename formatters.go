package loggo

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strings"
)

// IFormatter formatters convert entries to []byte and used by handlers
type IFormatter interface {
	Format(entry *Entry) []byte
}

// TextFormatter simple formatter
type TextFormatter struct {
	format string
}

// NewTextFormatter creates new TextFormatter
func NewTextFormatter(format string) *TextFormatter {
	return &TextFormatter{format: format}
}

// Format formats Entry
func (f *TextFormatter) Format(entry *Entry) []byte {
	result := f.format

	additionalBuf := &bytes.Buffer{}
	data := filterEntryFields(entry)
	if marshaledData, err := json.Marshal(data); err == nil {
		additionalBuf.Write(marshaledData)
	}

	replaces := make([]string, 0, 2+len(entry.Fields))
	replaces = append(
		replaces,
		":level:", entry.Level.String(),
		":time:", entry.Time.UTC().Format("2006-01-02T15:04:05.000000-07:00"),
		":message:", entry.Message,
		":additional:", additionalBuf.String(),
	)

	for key, value := range entry.Fields {
		replaces = append(replaces, fmt.Sprintf(":%s:", key), fmt.Sprintf("%s", value))
	}

	replacer := strings.NewReplacer(replaces...)

	buf := &bytes.Buffer{}

	replacer.WriteString(buf, result)

	buf.WriteByte('\n')

	return buf.Bytes()
}

func filterEntryFields(entry *Entry) map[string]interface{} {
	result := make(map[string]interface{}, len(entry.Fields))

	for key, value := range entry.Fields {
		if key[0] != '_' {
			result[key] = value
		}
	}

	return result
}
