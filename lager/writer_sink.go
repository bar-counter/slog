package lager

import (
	"bytes"
	"github.com/bar-counter/slog/lager/color"
	"io"
	"sync"
)

//const logBufferSize = 1024

// A Sink represents a write destination for a Logger. It provides
// a thread-safe interface for writing logs
type Sink interface {
	//Log to the sink.  Best effort -- no need to worry about errors.
	Log(level LogLevel, payload []byte)
}

type writerSink struct {
	writer      io.Writer
	minLogLevel LogLevel
	name        string
	writeL      *sync.Mutex
}

// NewWriterSink is function which returns new struct object
func NewWriterSink(name string, writer io.Writer, minLogLevel LogLevel) Sink {
	return &writerSink{
		writer:      writer,
		minLogLevel: minLogLevel,
		writeL:      new(sync.Mutex),
		name:        name,
	}
}

func (sink *writerSink) Log(level LogLevel, log []byte) {
	if level < sink.minLogLevel {
		return
	}
	if sink.name == "stdout" {
		if bytes.Contains(log, []byte("WARN")) {
			log = bytes.Replace(log, []byte("WARN"), color.WarnByte, -1)
		} else if bytes.Contains(log, []byte("ERROR")) {
			log = bytes.Replace(log, []byte("ERROR"), color.ErrorByte, -1)
		} else if bytes.Contains(log, []byte("FATAL")) {
			log = bytes.Replace(log, []byte("FATAL"), color.FatalByte, -1)
		} else if bytes.Contains(log, []byte("DEBUG")) {
			log = bytes.Replace(log, []byte("DEBUG"), color.DebugByte, -1)
		}
	}
	log = append(log, '\n')
	sink.writeL.Lock()
	_, err := sink.writer.Write(log)
	if err != nil {
		println("writer_sink Log Failed to write log: " + err.Error())
	}
	sink.writeL.Unlock()
}
