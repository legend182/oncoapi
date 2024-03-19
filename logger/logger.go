package logger

import (
	"oncoapi/setting"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func Init(cfg *setting.LogConf,mode string)(err error){
	var lg *zap.Logger
	encode:=getEncode()
	writesync := getWriteSync(cfg.Filename,cfg.Maxsize,cfg.MaxBackups,cfg.MaxBackups)
	var l zapcore.Level
	err = l.UnmarshalText([]byte(cfg.Level))
	if err!=nil {
		return
	}
	var core zapcore.Core
	if mode == "dev"{
		// 开发模式需要输出到控制台方便查看
		consoleEncoder := zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig())
		core = zapcore.NewTee(
			zapcore.NewCore(encode,writesync,l),
			zapcore.NewCore(consoleEncoder,zapcore.Lock(os.Stdout),zapcore.DebugLevel),
		)
	}else{
		core = zapcore.NewCore(encode,writesync,l)
	}
	lg = zap.New(core,zap.AddCaller(),zap.AddStacktrace(l))
	zap.ReplaceGlobals(lg)
	zap.L().Info("init logger success")
	return
}

// 设置编码器
func getEncode() zapcore.Encoder{
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.TimeKey = "time"
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoderConfig.EncodeDuration = zapcore.SecondsDurationEncoder
	encoderConfig.EncodeCaller = zapcore.ShortCallerEncoder
	return zapcore.NewJSONEncoder(encoderConfig)
}
// 设置写入器
func getWriteSync(filename string, maxSize, maxBackup, maxAge int) zapcore.WriteSyncer{
	lumberjcak := &lumberjack.Logger{
		Filename:   filename,
		MaxSize:    maxSize,
		MaxBackups: maxBackup,
		MaxAge:     maxAge,
	}
	return zapcore.AddSync(lumberjcak)
}