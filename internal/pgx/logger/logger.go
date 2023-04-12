package logger

import (
	"capstone/util"
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/jackc/pgx/v4"
)

// Colors
const (
	Reset       = "\033[0m"
	Red         = "\033[31m"
	Green       = "\033[32m"
	Yellow      = "\033[33m"
	Blue        = "\033[34m"
	Magenta     = "\033[35m"
	Cyan        = "\033[36m"
	White       = "\033[37m"
	BlueBold    = "\033[34;1m"
	MagentaBold = "\033[35;1m"
	RedBold     = "\033[31;1m"
	YellowBold  = "\033[33;1m"
)

// Writer log writer interface
type Writer interface {
	Printf(string, ...interface{})
}

// Config logger config
type Config struct {
	SlowThreshold time.Duration
	Colorful      bool
	LogLevel      pgx.LogLevel
}

// New initialize logger
func New(writer Writer, config Config) pgx.Logger {
	var (
		traceStr     = "%s\n[%.3fms] [rows:%v] %s"
		traceWarnStr = "%s %s\n[%.3fms] [rows:%v] %s"
		traceErrStr  = "%s %s\n[%.3fms] [rows:%v] %s"
	)

	if config.Colorful {
		traceStr = Green + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
		traceWarnStr = Green + "%s " + Yellow + "%s\n" + Reset + RedBold + "[%.3fms] " + Yellow + "[rows:%v]" + Magenta + " %s" + Reset
		traceErrStr = RedBold + "%s " + MagentaBold + "%s\n" + Reset + Yellow + "[%.3fms] " + BlueBold + "[rows:%v]" + Reset + " %s"
	}

	return &logger{
		Writer:       writer,
		Config:       config,
		traceStr:     traceStr,
		traceWarnStr: traceWarnStr,
		traceErrStr:  traceErrStr,
	}
}

type logger struct {
	Writer
	Config
	traceStr, traceErrStr, traceWarnStr string
}

func (l logger) Log(ctx context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	switch msg {
	case "Exec", "Query":
		var (
			sql     string        = ""
			rows    string        = "-"
			err     error         = nil
			elapsed time.Duration = 0
		)

		if data["sql"] != nil {
			sql = data["sql"].(string)
		}

		if sql != "" {
			sql = strings.ReplaceAll(sql, "\n", " ")
			sql = strings.ReplaceAll(sql, "\t", " ")
			sql = strings.ReplaceAll(sql, "    ", "")

			args := []interface{}{}
			if data["args"] != nil {
				args = data["args"].([]interface{})
			}

			sql = ExplainSQL(sql, `'`, args...)
		}

		if data["rowCount"] != nil {
			rows = util.ToString(data["rowCount"])
		}

		if data["err"] != nil {
			err = data["err"].(error)
		}

		if data["time"] != nil {
			elapsed = data["time"].(time.Duration)
		}

		switch {
		case err != nil && l.LogLevel >= pgx.LogLevelError:
			l.Printf(l.traceErrStr, util.FileWithLineNum(), err, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case elapsed > l.SlowThreshold && l.SlowThreshold != 0 && l.LogLevel >= pgx.LogLevelWarn:
			slowLog := fmt.Sprintf("SLOW SQL >= %v", l.SlowThreshold)
			l.Printf(l.traceWarnStr, util.FileWithLineNum(), slowLog, float64(elapsed.Nanoseconds())/1e6, rows, sql)
		case l.LogLevel == pgx.LogLevelInfo:
			l.Printf(l.traceStr, util.FileWithLineNum(), float64(elapsed.Nanoseconds())/1e6, rows, sql)
		}
	}
}
