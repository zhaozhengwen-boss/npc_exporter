// internal/logger/logger.go
package logger

import (
	"log/slog"
	"os"

	"gopkg.in/natefinch/lumberjack.v2"
)

var Logger *slog.Logger

func Init() {
	// 确保 logs 目录存在
	if err := os.MkdirAll("logs", 0755); err != nil {
		panic("failed to create logs directory: " + err.Error())
	}

	// 配置 lumberjack：按天轮转，保留 7 天，单文件最大 100MB
	logFile := &lumberjack.Logger{
		Filename:   "logs/npu-exporter.log", // 实际会按日期重命名
		MaxSize:    100,                     // MB
		MaxAge:     7,                       // 天
		MaxBackups: 7,                       // 最多保留 7 个旧文件
		LocalTime:  true,                    // 使用本地时间命名
		Compress:   false,                   // 不压缩
	}

	// 创建带时间、级别、文件名的日志处理器
	handler := slog.NewJSONHandler(logFile, &slog.HandlerOptions{
		Level:     slog.LevelInfo,
		AddSource: true,
	})

	Logger = slog.New(handler)
}
