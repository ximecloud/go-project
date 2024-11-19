package logger

import (
	"fmt"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"net"
	"os"
	"project/internal/adapter/config"
	"sync"
	"time"
)

type LogWriter struct {
	conn net.Conn
	addr string
	mu   sync.Mutex
}

func New(addr string) *LogWriter {
	writer := &LogWriter{
		addr: addr,
	}
	go writer.connect()
	return writer
}

func (r *LogWriter) connect() {
	for {
		r.mu.Lock()
		conn, err := net.Dial("tcp", r.addr)
		if err == nil {
			r.conn = conn
			fmt.Println("Connected to Logstash")
			r.mu.Unlock()
			return
		}
		r.mu.Unlock()
		time.Sleep(5 * time.Second)
	}
}

func (r *LogWriter) Write(p []byte) (n int, err error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	if r.conn == nil {
		return len(p), nil
	}

	n, err = r.conn.Write(p)
	if err != nil {
		fmt.Printf("Failed to write to Logstash: %v\n", err)
		if err := r.conn.Close(); err != nil {
			return 0, err
		}
		r.conn = nil
		go r.connect()
	}
	return n, err
}

func (r *LogWriter) Sync() error {
	return nil
}

var log *zap.Logger

func Initial() {
	writer := New(config.GetLogstash())

	encoderConfig := zap.NewProductionConfig().EncoderConfig
	encoderConfig.TimeKey = "timestamp"
	encoderConfig.MessageKey = "message"
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.StacktraceKey = ""

	consoleCore := zapcore.NewCore(
		zapcore.NewConsoleEncoder(encoderConfig),
		zapcore.AddSync(os.Stdout),
		zap.InfoLevel,
	)

	logstashCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		zapcore.AddSync(writer),
		zap.InfoLevel,
	)

	core := zapcore.NewTee(consoleCore, logstashCore)
	log = zap.New(core, zap.AddCaller(), zap.AddCallerSkip(1))
}

func Info(message string) {
	log.Info(message)
}
