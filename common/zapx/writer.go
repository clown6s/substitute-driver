package zapx

import (
	"fmt"

	"github.com/zeromicro/go-zero/core/logx"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

// ZapLogWriter 实现 logx.Writer 接口，将 go-zero 日志桥接到 zap
type ZapLogWriter struct{}

// NewZapLogWriter 创建实现 logx.Writer 的桥接 Writer
func NewZapLogWriter() *ZapLogWriter {
	return &ZapLogWriter{}
}

func (w *ZapLogWriter) Alert(v any) {
	w.log(zapcore.WarnLevel, v)
}

func (w *ZapLogWriter) Close() error {
	Sync()
	return nil
}

func (w *ZapLogWriter) Debug(v any, fields ...logx.LogField) {
	w.logFields(zapcore.DebugLevel, v, fields)
}

func (w *ZapLogWriter) Error(v any, fields ...logx.LogField) {
	w.logFields(zapcore.ErrorLevel, v, fields)
}

func (w *ZapLogWriter) Info(v any, fields ...logx.LogField) {
	w.logFields(zapcore.InfoLevel, v, fields)
}

func (w *ZapLogWriter) Severe(v any) {
	w.log(zapcore.FatalLevel, v)
}

func (w *ZapLogWriter) Slow(v any, fields ...logx.LogField) {
	w.logFields(zapcore.WarnLevel, v, fields)
}

func (w *ZapLogWriter) Stack(v any) {
	if Logger == nil {
		return
	}
	Logger.Desugar().Error(fmt.Sprintf("%v", v), zap.Stack("stack"))
}

func (w *ZapLogWriter) Stat(v any, fields ...logx.LogField) {
	w.logFields(zapcore.InfoLevel, v, fields)
}

func (w *ZapLogWriter) log(level zapcore.Level, v any) {
	if Logger == nil {
		return
	}
	msg := fmt.Sprintf("%v", v)
	switch level {
	case zapcore.DebugLevel:
		Logger.Debug(msg)
	case zapcore.InfoLevel:
		Logger.Info(msg)
	case zapcore.WarnLevel:
		Logger.Warn(msg)
	case zapcore.ErrorLevel:
		Logger.Error(msg)
	case zapcore.FatalLevel:
		Logger.Fatal(msg)
	}
}

func (w *ZapLogWriter) logFields(level zapcore.Level, v any, fields []logx.LogField) {
	if Logger == nil {
		return
	}

	// SugaredLogger.With / *w 方法接受 []any（即 []interface{}）
	kv := make([]any, 0, len(fields)*2)
	for _, f := range fields {
		kv = append(kv, f.Key, f.Value)
	}

	msg := fmt.Sprintf("%v", v)
	switch level {
	case zapcore.DebugLevel:
		Logger.Debugw(msg, kv...)
	case zapcore.InfoLevel:
		Logger.Infow(msg, kv...)
	case zapcore.WarnLevel:
		Logger.Warnw(msg, kv...)
	case zapcore.ErrorLevel:
		Logger.Errorw(msg, kv...)
	case zapcore.FatalLevel:
		Logger.Fatalw(msg, kv...)
	}
}

// WithOptions 创建带额外选项的 zap.Logger（非 SugaredLogger）
// 用于需要结构化日志字段的场景
func WithOptions(opts ...zap.Option) *zap.Logger {
	if Logger == nil {
		return zap.NewNop()
	}
	return Logger.Desugar().WithOptions(opts...)
}

// 确保 ZapLogWriter 实现了 logx.Writer 接口
var _ logx.Writer = (*ZapLogWriter)(nil)
