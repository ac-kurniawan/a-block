package log

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"
)

type ICustomLog interface {
	Error(ctx context.Context, message string)
	Info(ctx context.Context, message string)
	Warning(ctx context.Context, message string)
	Fatal(ctx context.Context, message string)
}

type LogModel struct {
	Level     string            `json:"level,omitempty"`
	Timestamp time.Time         `json:"timestamp"`
	Stack     string            `json:"stack,omitempty"`
	Context   map[string]string `json:"context,omitempty"`
	Message   string            `json:"message,omitempty"`
}

type CustomLog struct {
	IsJson bool
}

func (c CustomLog) print(ctx context.Context, level string, message string) {
	pc, _, _, ok := runtime.Caller(2)
	details := runtime.FuncForPC(pc)
	if ok && details != nil {
		if c.IsJson {
			contextMap := map[string]string{}
			if ctx.Value("requestId") != nil {
				contextMap["requestId"] = ctx.Value("requestId").(string)
			}
			if ctx.Value("customerId") != nil {
				contextMap["customerId"] = ctx.Value("customerId").(string)
			}
			logPrintable := LogModel{
				Level:     level,
				Timestamp: time.Now(),
				Stack:     details.Name(),
				Context:   contextMap,
				Message:   message,
			}
			json, _ := json.Marshal(logPrintable)
			fmt.Println(string(json))
		} else {
			fmt.Printf("[%s] - %s => %s\n", level, details.Name(), message)
		}
	}

}

func (c CustomLog) Error(ctx context.Context, message string) {
	c.print(ctx, "ERROR", message)
}

func (c CustomLog) Info(ctx context.Context, message string) {
	c.print(ctx, "INFO", message)
}

func (c CustomLog) Warning(ctx context.Context, message string) {
	c.print(ctx, "WARN", message)
}

func (c CustomLog) Fatal(ctx context.Context, message string) {
	panic(message)
}

func NewCustomLog(module CustomLog) ICustomLog {
	return &module
}
