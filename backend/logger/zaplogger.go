package logger

import (
	"fmt"
	"io"
	"os"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"github.com/Walker088/rigel_ledger_server/backend/config"
)

const (
	fileOutputPath = "./logs"
	logFileName    = "rigelledger.log"
	errFileName    = "rigelledger_err.log"
)

type WriteSyncer struct {
	io.Writer
}

func (ws WriteSyncer) Sync() error {
	return nil
}

func getWriteSyncer(filename string) zapcore.WriteSyncer {
	var ioWriter = &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    10, // MB
		MaxBackups: 3,  // number of backups
		MaxAge:     14, //days
		LocalTime:  true,
		Compress:   false, // disabled by default
	}
	var sw = WriteSyncer{
		ioWriter,
	}
	return sw
}

func New() *zap.SugaredLogger {
	if _, err := os.Stat(fileOutputPath); os.IsNotExist(err) {
		os.MkdirAll(fileOutputPath, 0700)
	}

	cfg := config.GetLoggerConfig()
	var logger *zap.Logger

	fn := fmt.Sprintf("%s/%s", fileOutputPath, logFileName)
	consoleEnc := zapcore.NewConsoleEncoder(cfg.EncoderConfig)
	fileEnc := zapcore.NewJSONEncoder(cfg.EncoderConfig)
	core := zapcore.NewTee(
		zapcore.NewCore(consoleEnc, zapcore.AddSync(os.Stdout), zap.NewAtomicLevelAt(zap.InfoLevel)),
		zapcore.NewCore(fileEnc, zapcore.AddSync(getWriteSyncer(fn)), zap.NewAtomicLevelAt(zap.InfoLevel)),
	)

	logger = zap.New(core)
	return logger.Sugar()
}
