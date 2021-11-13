package terminal

import (
	"github.com/pkg/errors"
	"github.com/rs/zerolog"
)

var (
	Logger    zerolog.Logger
	LogLevels = map[int]zerolog.Level{
		1: zerolog.ErrorLevel,
		2: zerolog.WarnLevel,
		3: zerolog.InfoLevel,
		4: zerolog.DebugLevel,
		5: zerolog.TraceLevel,
	}
)

func SetLogLevel(level int) error {
	if severity, ok := LogLevels[level]; ok {
		Logger = Logger.Level(severity)
		return nil
	}
	return errors.Errorf("The provided verbosity level '%d' is not in the range [1,4]", level)
}

func IsVerbose() bool {
	return Logger.GetLevel() > zerolog.ErrorLevel
}

func IsDebug() bool {
	return Logger.GetLevel() >= zerolog.TraceLevel
}

func GetLogLevel() int {
	severity := Logger.GetLevel()
	for level, sev := range LogLevels {
		if sev == severity {
			return level
		}
	}
	return 1
}

func init() {
	Logger = zerolog.New(Stderr).Level(zerolog.ErrorLevel)
}
