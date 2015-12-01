package epilog

import "github.com/go-st/logger"

func ExampleSimpleUsage() {
	logger := New("MyLogger", NewStreamHandler(logger.LevelDebug, NewTextFormatter("(:level:) :message:")))
	logger.Debug("hello debug")
	logger.Info("hello info")

	// Output:
	// (DEBUG) hello debug
	// (INFO) hello info
}

func ExampleBufferEmpty() {
	handler := NewStreamHandler(logger.LevelDebug, NewTextFormatter("(:level:) :message:"))
	logger := New("MyLogger", NewBufferHandler(handler, logger.LevelWarning))
	logger.Debug("hello debug")
	logger.Info("hello info")

	// Output:
}

func ExampleBuffer() {
	handler := NewStreamHandler(logger.LevelInfo, NewTextFormatter("(:level:) :message:"))
	logger := New("MyLogger", NewBufferHandler(handler, logger.LevelWarning))
	logger.Debug("hello debug")
	logger.Info("hello info")
	logger.Warning("hello warning")

	// Output:
	// (INFO) hello info
	// (WARNING) hello warning
}
