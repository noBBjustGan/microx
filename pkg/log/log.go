package log

//
//import (
//	// "path"
//	"encoding/json"
//	"runtime"
//	"strconv"
//	"strings"
//	// "time"
//
//	"github.com/natefinch/lumberjack"
//	log "github.com/sirupsen/logrus"
//)
//
//func init() {
//	log.SetFormatter(&DFormatter{})
//}
//
////type PanicLevel = log.PanicLevel
//const (
//	PanicLevel log.Level = log.PanicLevel
//	FatalLevel log.Level = log.FatalLevel
//	ErrorLevel log.Level = log.ErrorLevel
//	WarnLevel  log.Level = log.WarnLevel
//	InfoLevel  log.Level = log.InfoLevel
//	DebugLevel log.Level = log.DebugLevel
//)
//
//// config need to be correct JSON as string: {"filename":"default.log","maxsize":100}
//func SetLogger(config string) {
//	type LogConfig struct {
//		Filename string `json:"filename"`
//		MaxSize  int    `json:"maxsize"`
//	}
//
//	logConfig := &LogConfig{}
//	if err := json.Unmarshal([]byte(config), logConfig); err != nil {
//		log.Fatal(err)
//	}
//
//	l := &lumberjack.Logger{
//		Filename:  logConfig.Filename,
//		MaxSize:   logConfig.MaxSize,
//		LocalTime: true,
//	}
//	log.SetOutput(l)
//}
//func SetLevel(logLevel log.Level) {
//	log.SetLevel(logLevel)
//}
//
//type DFormatter struct {
//	TimestampFormat string
//}
//
//func (f *DFormatter) Format(entry *log.Entry) ([]byte, error) {
//	timestampFormat := f.TimestampFormat
//	if timestampFormat == "" {
//		timestampFormat = "2006/01/02 15:04:05"
//	}
//
//	_, file, line, ok := runtime.Caller(9)
//	if !ok {
//		file = "???"
//		line = 0
//	}
//	// _, filename := path.Split(file)
//	msg := entry.Time.Format(timestampFormat) +
//		" " + strings.ToUpper(entry.Level.String()) +
//		" [" + file + ":" + strconv.Itoa(line) + "] " +
//		entry.Message + "\n"
//
//	return []byte(msg), nil
//}
//
//func Debugf(format string, args ...interface{}) {
//	log.Debugf(format, args...)
//}
//
//func Infof(format string, args ...interface{}) {
//	log.Infof(format, args...)
//}
//
//func Warnf(format string, args ...interface{}) {
//	log.Warnf(format, args...)
//}
//
//func Errorf(format string, args ...interface{}) {
//	log.Errorf(format, args...)
//}
//
//func Fatalf(format string, args ...interface{}) {
//	log.Fatalf(format, args...)
//}
//
//func Panicf(format string, args ...interface{}) {
//	log.Panicf(format, args...)
//}
//
//// 以下函数不可以写成这种形式
//// func Debug(args ...interface{}) {
//// 	log.Debug(args...)
//// }
//
//func Debug(args ...interface{}) {
//	debug(args...)
//}
//
//func debug(args ...interface{}) {
//	log.Debug(args...)
//}
//
//func Info(args ...interface{}) {
//	info(args...)
//}
//
//func info(args ...interface{}) {
//	log.Info(args...)
//}
//
//func Warn(args ...interface{}) {
//	warn(args...)
//}
//
//func warn(args ...interface{}) {
//	log.Warn(args...)
//}
//
//func Error(args ...interface{}) {
//	ferror(args...)
//}
//
//func ferror(args ...interface{}) {
//	log.Error(args...)
//}
//
//func Fatal(args ...interface{}) {
//	fatal(args...)
//}
//
//func fatal(args ...interface{}) {
//	log.Fatal(args...)
//}
//
//func Panic(args ...interface{}) {
//	panic(args...)
//}
//
//func panic(args ...interface{}) {
//	log.Panic(args...)
//}
