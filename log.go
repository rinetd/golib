package golib

import (
	"io"
	"os"
	"sync"
	"time"

	"github.com/rs/zerolog"
)

var log zerolog.Logger

func New() *zerolog.Logger {
	var once = &sync.Once{}
	once.Do(func() {
		var file, err = os.OpenFile("log_"+time.Now().Format("2006-01-02")+".log",
			os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
		if err != nil {
			return
		}

		log = zerolog.New(zerolog.ConsoleWriter{
			// Out: os.Stdout,
			// Out:     file,
			Out:     io.MultiWriter(os.Stdout, file),
			NoColor: true,
		}).With().Timestamp().Str("", "").Logger()
	})
	return &log
}
func Log() *zerolog.Logger {
	var once = &sync.Once{}
	once.Do(func() {
		//log = New()
	})
	return nil
}

func SetLevel(level zerolog.Level) {
	log.Level(level)
	zerolog.SetGlobalLevel(zerolog.DebugLevel)
}

// 同时输出文件和控制台
// io.MultiWriter(f, os.Stdout)
// log.New(io.MultiWriter(os.Stdout),
// func New(){
// var logWriter io.Writer
// 	// logWriter = fileLogger

// 	// if interactive write formated log to stderr, disable color on windows
// 	if service.Interactive() {
// 		logWriter = io.MultiWriter(zerolog.SyncWriter(zerolog.ConsoleWriter{
// 			Out:     os.Stderr,
// 			NoColor: runtime.GOOS == "windows",
// 		}), logWriter)
// 	}

// 	// minimize json footprint
// 	zerolog.TimestampFieldName = "T"
// 	zerolog.LevelFieldName = "L"
// 	zerolog.MessageFieldName = "M"

// 	// set loglevel
// 	switch config.Level {
// 	case "debug":
// 		zerolog.SetGlobalLevel(zerolog.DebugLevel)
// 	case "info":
// 		zerolog.SetGlobalLevel(zerolog.InfoLevel)
// 	case "warn":
// 		zerolog.SetGlobalLevel(zerolog.WarnLevel)
// 	case "error":
// 		zerolog.SetGlobalLevel(zerolog.ErrorLevel)
// 	case "disable":
// 		zerolog.SetGlobalLevel(zerolog.Disabled)
// 	default:
// 		return errors.New("unrecognized logging level '" + config.Level + "'")
// 	}

// 	// initialize logger
// 	logger = zerolog.New(logWriter).With().Timestamp().Logger()
// return nil
// }
