package loggo

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

// IProcessor is an interface for processor
type IProcessor interface {
	Process(entry *Entry)
}

// CalleeProcessor adds packageName and filename to entry fields
type CalleeProcessor struct {
	shift int
}

// NewCalleeProcessor constructor for CalleeProcessor
func NewCalleeProcessor(shift int) *CalleeProcessor {
	return &CalleeProcessor{shift: shift}
}

// Process adds two fields to entry:
// _package - name of package where logger was called
// _file - file:line where logger was called
func (p *CalleeProcessor) Process(entry *Entry) {
	entry.Fields["_package"] = getPackageName(p.shift)
	entry.Fields["_file"] = getFileName(p.shift)
}

func getPackageName(shift int) string {
	v := "???"
	if pc, _, _, ok := runtime.Caller(shift + 5); ok {
		if f := runtime.FuncForPC(pc); f != nil {
			v = formatFuncName(f.Name())
		}
	}

	return v
}

func getFileName(shift int) string {
	_, file, line, ok := runtime.Caller(shift + 5)
	if !ok {
		file = "???"
		line = 0
	} else {
		file = filepath.Base(file)
	}

	return fmt.Sprintf("%s:%d", file, line)
}

func formatFuncName(f string) string {
	i := strings.LastIndex(f, "/")
	j := strings.Index(f[i+1:], ".")
	if j < 1 {
		return "???"
	}

	return f[:i+j+1]
}

// FieldsProcessor is a processor of fields list
type FieldsProcessor struct {
	fields map[string]interface{}
}

// NewFieldsProcessor creates new FieldsProcessor
func NewFieldsProcessor(fields map[string]interface{}) *FieldsProcessor {
	return &FieldsProcessor{
		fields: fields,
	}
}

// Process processes Entry
func (p *FieldsProcessor) Process(entry *Entry) {
	for key, value := range p.fields {
		entry.Fields[key] = value
	}
}
