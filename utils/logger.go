package utils

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

/**
 * 获取日志
 * filePath 日志文件路径
 * level 日志级别
 * maxSize 每个日志文件保存的最大尺寸 单位：M
 * maxBackups 日志文件最多保存多少个备份
 * maxAge 文件最多保存多少天
 * compress 是否压缩
 * serviceName 服务名
 */
func NewLogger(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool, serviceName string) *zap.Logger {
	core := newCore(filePath, level, maxSize, maxBackups, maxAge, compress)
	return zap.New(core, zap.AddCaller(), zap.Development(), zap.Fields(zap.String("serviceName", serviceName)))
}

/**
 *  使用方法
	var MainLogger *zap.Logger
	var GatewayLogger *zap.Logger

	func init() {

		MainLogger = NewLogger("./logs/main.log", zapcore.InfoLevel, 128, 30, 7, true, "Main")
		GatewayLogger = NewLogger("./logs/gateway.log", zapcore.DebugLevel, 128, 30, 7, true, "Gateway")
	}

*/

/**
 * zapcore构造
 */
func newCore(filePath string, level zapcore.Level, maxSize int, maxBackups int, maxAge int, compress bool) zapcore.Core {
	//日志文件路径配置2
	hook := lumberjack.Logger{
		Filename:   filePath,   // 日志文件路径
		MaxSize:    maxSize,    // 每个日志文件保存的最大尺寸 单位：M
		MaxBackups: maxBackups, // 日志文件最多保存多少个备份
		MaxAge:     maxAge,     // 文件最多保存多少天
		Compress:   compress,   // 是否压缩
	}
	// 设置日志级别
	atomicLevel := zap.NewAtomicLevel()
	atomicLevel.SetLevel(level)
	//公用编码器
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "linenum",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeLevel:    zapcore.LowercaseLevelEncoder,  // 小写编码器
		EncodeTime:     customTimeEncoder,              // 用户自定义格式化时间
		EncodeDuration: zapcore.SecondsDurationEncoder, //
		EncodeCaller:   zapcore.FullCallerEncoder,      // 全路径编码器
		EncodeName:     zapcore.FullNameEncoder,
	}
	return zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),                                           // 编码器配置
		zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook)), // 打印到控制台和文件
		atomicLevel, // 日志级别
	)
}

func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
}

//
//type ZapConfig struct {
//	Level         string `json:"Level"`         // 日志级别
//	Format        string `json:"Format"`        // 日志格式
//	Prefix        string `json:"Prefix"`        // 日志前缀
//	LogDir        string `json:"LogDir"`        // 日志目录
//	ShowLine      bool   `json:"ShowLine"`      // 是否显示行号
//	EncodeLevel   string `json:"EncodeLevel"`   // 日志级别编码
//	StacktraceKey string `json:"StacktraceKey"` // 堆栈名称
//	InConsole     bool   `json:"InConsole"`     // 是否在控制台打印
//	TimeFormat    string `json:"TimeFormat"`    // 时间格式
//}
//
//var (
//	zapConfig ZapConfig
//	logger    *zap.Logger
//)
//
//// initConfig 初始化配置
//// format:json
//func initConfig(format string, logDir string) {
//	zapConfig.Level = zap.InfoLevel.String()
//	zapConfig.Format = format
//	zapConfig.Prefix = ""
//	zapConfig.LogDir = logDir
//	zapConfig.ShowLine = true
//	zapConfig.EncodeLevel = "LowercaseLevelEncoder"
//	zapConfig.StacktraceKey = "Stacktrace"
//	zapConfig.InConsole = false
//	zapConfig.TimeFormat = "2006-01-02"
//}
//
//// CreateLogger 创建日志
//func CreateLogger(format string, logDir string) (logger *zap.Logger) {
//
//	initConfig(format, logDir)
//
//	//// 调试级别日志
//	//debugPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
//	//	return lvl == zapcore.DebugLevel
//	//})
//	//
//	//// 普通级别日志
//	//infoPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
//	//	return lvl == zapcore.InfoLevel
//	//})
//	//
//	//// 警告级别日志
//	//warnPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
//	//	return lvl == zapcore.WarnLevel
//	//})
//	//
//	//// 错误级别日志
//	//errorPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
//	//	return lvl >= zapcore.ErrorLevel
//	//})
//
//	// 全级别日志
//	allPriority := zap.LevelEnablerFunc(func(lvl zapcore.Level) bool {
//		return lvl >= zapcore.DebugLevel
//	})
//
//	today := time.Now().Format(zapConfig.TimeFormat)
//
//	cores := [...]zapcore.Core{
//		//getEncoderCore(path.Join(LogDir, "debug-"+ZapConfig.Prefix+"-"+today+".log"), debugPriority, ZapConfig),
//		//getEncoderCore(path.Join(LogDir, "info-"+ZapConfig.Prefix+"-"+today+".log"), infoPriority, ZapConfig),
//		//getEncoderCore(path.Join(LogDir, "warn-"+ZapConfig.Prefix+"-"+today+".log"), warnPriority, ZapConfig),
//		//getEncoderCore(path.Join(LogDir, "error-"+ZapConfig.Prefix+"-"+today+".log"), errorPriority, ZapConfig),
//		getEncoderCore(path.Join(zapConfig.LogDir, "all-"+zapConfig.Prefix+"-"+today+".log"), allPriority, zapConfig),
//	}
//
//	// 创建日志
//	logger = zap.New(zapcore.NewTee(cores[:]...))
//
//	// 输出文件名和行号
//	if zapConfig.ShowLine {
//		logger = logger.WithOptions(zap.AddCaller())
//	}
//
//	//// 警告级别输出堆栈
//	//logger = logger.WithOptions(zap.AddStacktrace(zapcore.WarnLevel))
//	zap.ReplaceGlobals(logger)
//	return logger
//}
//
//func getEncoderConfig(zapConfig ZapConfig) zapcore.EncoderConfig {
//	config := zapcore.EncoderConfig{
//		MessageKey:     "Message",
//		LevelKey:       "Level",
//		TimeKey:        "Time",
//		NameKey:        "Logger",
//		CallerKey:      "Caller",
//		StacktraceKey:  zapConfig.StacktraceKey,
//		LineEnding:     zapcore.DefaultLineEnding,
//		EncodeTime:     customTimeEncoder,
//		EncodeDuration: zapcore.SecondsDurationEncoder,
//		EncodeCaller:   zapcore.FullCallerEncoder,
//	}
//
//	switch zapConfig.EncodeLevel {
//	case "LowercaseLevelEncoder":
//		config.EncodeLevel = zapcore.LowercaseLevelEncoder
//	case "LowercaseColorLevelEncoder":
//		config.EncodeLevel = zapcore.LowercaseColorLevelEncoder
//	case "CapitalLevelEncoder":
//		config.EncodeLevel = zapcore.CapitalLevelEncoder
//	case "CapitalColorLevelEncoder":
//		config.EncodeLevel = zapcore.CapitalColorLevelEncoder
//	default:
//		config.EncodeLevel = zapcore.LowercaseLevelEncoder
//	}
//	return config
//}
//
//func getEncoder(zapConfig ZapConfig) zapcore.Encoder {
//	if "json" == zapConfig.Format {
//		return zapcore.NewJSONEncoder(getEncoderConfig(zapConfig))
//	}
//	return zapcore.NewConsoleEncoder(getEncoderConfig(zapConfig))
//}
//
//func getWriteSyncer(fileName string, zapConfig ZapConfig) zapcore.WriteSyncer {
//
//	logger := &lumberjack.Logger{
//		Filename:   fileName, // 日志文件路径
//		MaxSize:    10,       // 每个日志文件保存的最大尺寸 单位：M
//		MaxAge:     7,        // 文件最多保存多少天
//		MaxBackups: 200,      // 日志文件最多保存多少个备份
//		Compress:   true,     // 是否压缩
//	}
//
//	if zapConfig.InConsole {
//		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(logger))
//	}
//
//	return zapcore.AddSync(logger)
//}
//
//func getEncoderCore(fileName string, level zapcore.LevelEnabler, zapConfig ZapConfig) zapcore.Core {
//	writer := getWriteSyncer(fileName, zapConfig)
//	return zapcore.NewCore(getEncoder(zapConfig), writer, level)
//}
//
