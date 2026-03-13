package logger

import (
	"context"
	"os"
	"path/filepath"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
)

var Log *zap.Logger

// Define a private type for context keys to avoid collisions
type ctxKey struct{}

func Init() {
	// 1. Define log file path
	logPath := "logs/app.log"
	logDir := filepath.Dir(logPath)
	if err := os.MkdirAll(logDir, 0755); err != nil {
		panic(err)
	}

	// 2. Log Rotation Configuration (The Optimizer)
	// This prevents the log file from growing too large
	fileWriter := zapcore.AddSync(&lumberjack.Logger{
		Filename:   logPath,
		MaxSize:    10,   // Max megabytes before rotation
		MaxBackups: 5,    // Max number of old log files to keep
		MaxAge:     30,   // Max days to retain old log files
		Compress:   true, // Compress old log files (gzip)
	})

	// 3. Encoder Configuration for Console (Human Readable)
	consoleConfig := zap.NewProductionEncoderConfig()
	consoleConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder // Adds colors to levels
	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)

	// 4. Encoder Configuration for File (Machine Readable JSON)
	fileConfig := zap.NewProductionEncoderConfig()
	fileConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	fileEncoder := zapcore.NewJSONEncoder(fileConfig)

	// 5. Create Core with Tee (Splitting output to multiple destinations)
	core := zapcore.NewTee(
		// High performance console output
		zapcore.NewCore(consoleEncoder, zapcore.AddSync(os.Stdout), zap.InfoLevel),
		// Optimized file output with rotation
		zapcore.NewCore(fileEncoder, fileWriter, zap.InfoLevel),
	)

	// 6. Initialize the logger with Caller information (shows file name and line number)
	Log = zap.New(core,
		zap.AddCaller(),
	)
}

// FromCtx retrieves the logger from context or returns the global logger if not found
func FromCtx(ctx context.Context) *zap.Logger {
	if l, ok := ctx.Value(ctxKey{}).(*zap.Logger); ok {
		return l
	}
	return Log
}

// WithCtx injects a specific logger into the context
func WithCtx(ctx context.Context, l *zap.Logger) context.Context {
	return context.WithValue(ctx, ctxKey{}, l)
}
