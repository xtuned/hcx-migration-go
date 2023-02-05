package log

import "os"
import (
	"github.com/rs/zerolog"
)

func GetLogInstance() *zerolog.Logger {
	file, err := os.OpenFile("hcx.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0664)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	zerolog.TimeFieldFormat = zerolog.TimeFormatUnix
	consoleWriter := zerolog.ConsoleWriter{Out: os.Stdout}
	multi := zerolog.MultiLevelWriter(consoleWriter, file)
	logger := zerolog.New(multi).With().Timestamp().Logger()
	return &logger
}

//if err != nil {
//panic(err)
//}
//
//defer file.Close()
//
//multi := zerolog.MultiLevelWriter(os.Stdout, runLogFile)
//log.Logger = zerolog.New(multi).With().Timestamp().Logger()
