package logger

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"github.com/Fitray/auth_service/internal/config"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Logger struct {
	logger *zap.Logger
	file   *os.File
}

func NewLogger(
	loggerConfig config.LoggerConfig,
	rootPath string,
) (*Logger, error) {
	level := zap.NewAtomicLevel()
	if err := level.UnmarshalText([]byte(loggerConfig.Level)); err != nil {
		return nil, fmt.Errorf("invalid logger level: %w", err)
	}

	var file *os.File

	if loggerConfig.Folder != "" {
		path := filepath.Join(
			rootPath,
			loggerConfig.Folder,
		)

		if err := os.MkdirAll(path, 0755); err != nil {
			return nil, fmt.Errorf("create log directory: %w", err)
		}

		filename := time.Now().UTC().Format(loggerConfig.Format) + ".log"
		path = filepath.Join(path, filename)

		f, err := os.OpenFile(
			path,
			os.O_CREATE|os.O_WRONLY|os.O_APPEND,
			0644,
		)
		if err != nil {
			return nil, fmt.Errorf("open log file: %w", err)
		}

		file = f
	}

	if file == nil {
		return nil, fmt.Errorf("failed to open file")
	}

	encoderConfig := zap.NewProductionEncoderConfig()

	encoderConfig.TimeKey = "time"
	encoderConfig.LevelKey = "level"
	encoderConfig.MessageKey = "message"
	encoderConfig.CallerKey = "caller"

	encoderConfig.EncodeLevel = zapcore.LowercaseLevelEncoder
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeDuration = zapcore.StringDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder

	encoder := zapcore.NewConsoleEncoder(encoderConfig)

	cores := []zapcore.Core{
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(os.Stdout),
			level,
		),
		zapcore.NewCore(
			encoder,
			zapcore.AddSync(file),
			level,
		),
	}

	core := zapcore.NewTee(cores...)

	log := zap.New(
		core,
		zap.AddCaller(),
		zap.AddStacktrace(zap.PanicLevel),
		zap.AddStacktrace(zap.DPanicLevel),
		zap.AddStacktrace(zap.FatalLevel),
	)

	return &Logger{
		logger: log,
		file:   file,
	}, nil
}

func (l *Logger) Logger() *zap.Logger {
	return l.logger
}

func (l *Logger) With(fields ...zap.Field) *zap.Logger {
	return l.logger.With(fields...)
}

func (l *Logger) Close() {
	_ = l.logger.Sync()
	l.file.Close()
}
