package log

import (
	"bytes"
	context "context"
	"fmt"
	"io"
	"path/filepath"
	"runtime"

	stdlog "log"

	"google.golang.org/grpc"
)

var _client LogServiceClient

func InitClient(conn *grpc.ClientConn) error {
	_client = NewLogServiceClient(conn)
	Infow(context.Background(), "initialized logger")
	return nil
}

func Infow(ctx context.Context, message string, attributes ...interface{}) {
	logw(ctx, LogRequest_INFO, message, attributes...)
}

func Debugw(ctx context.Context, message string, attributes ...interface{}) {
	if IsDebug(ctx) {
		logw(ctx, LogRequest_DEBUG, message, attributes...)
	}
}

func Errorw(ctx context.Context, message string, attributes ...interface{}) {
	logw(ctx, LogRequest_ERROR, message, attributes...)
}

func logw(ctx context.Context, level LogRequest_LogLevel, message string, attributes ...interface{}) {
	if _client == nil {
		stdlog.Println("v1/log/_client not initialized, call InitClient, message:attrs:", message, attributes)
		return
	}
	attrs := []*LogRequest_KeyValuePair{}
	for i := 0; i < len(attributes)-1; i += 2 {
		attrs = append(attrs, &LogRequest_KeyValuePair{
			Key:   toString(attributes[i]),
			Value: toString(attributes[i+1]),
		})
	}
	in := &LogRequest{
		Level:      level,
		Message:    message,
		Attributes: attrs,
		Caller:     findCaller(),
	}
	_, err := _client.Log(ctx, in)
	if err != nil {
		// fallback
		fmt.Printf("[log.fallback] %s %s, %v", level.String(), message, attributes)
	}
}

func toString(i interface{}) string {
	if s, ok := i.(string); ok {
		return s
	}
	return fmt.Sprintf("%v", i)
}

func findCaller() string {
	pc, file, line, _ := runtime.Caller(3) // caller->Infow->logw->findCaller
	function := runtime.FuncForPC(pc).Name()
	_, function = filepath.Split(function)
	return fmt.Sprintf("%s:%d#%s", file, line, function)
}

type ErrorWithContext struct {
	Cause   error
	Message string
	Context []interface{}
}

func Wrap(err error, message string, kv ...interface{}) ErrorWithContext {
	return ErrorWithContext{Cause: err, Context: append(kv, "err", err)}
}

func (e ErrorWithContext) Error() string {
	b := new(bytes.Buffer)
	io.WriteString(b, e.Message)
	for i := 0; i < len(e.Context)-1; i += 2 {
		io.WriteString(b, fmt.Sprintf(", %v:%v", e.Context[i], e.Context[i+1]))
	}
	return b.String()
}

func LogWithError(ctx context.Context, err ErrorWithContext) {
	logw(ctx, LogRequest_ERROR, err.Message, append(err.Context, "err", err.Cause)...)
}

func ErrorWithLog(ctx context.Context, err error, message string, kv ...interface{}) error {
	logw(ctx, LogRequest_ERROR, message, append(kv, "err", err)...)
	return err
}

var GlobalDebug bool = false

func IsDebug(ctx context.Context) bool {
	if GlobalDebug {
		return true
	}
	// todo check x-cloud-debug
	return false
}
