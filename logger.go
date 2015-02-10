package logger

import (
	"os"
	"fmt"
	"path"
	"time"
	"runtime"
	"strings"
	
	"github.com/toolkits/color"
)

var (
	level int = 0
	gopath string
)

func init() {
	//输出系统所有环境变量的值
	for _, v := range os.Environ() {
		if ret := strings.Index(v, "GOPATH"); ret >= 0 {
			gopath = strings.Split(v, "=")[1]
			break
		}
	}
}
func SetLevelWithDefault(lv, defaultLv string) {
	err := SetLevel(lv)
	if err != nil {
		SetLevel(defaultLv)
	}
}

func SetLevel(lv string) error {
	if lv == "" {
		return fmt.Errorf("log level is blank")
	}

	l := strings.ToUpper(lv)

	switch l[0] {
	case 'T':
		level = 0
	case 'D':
		level = 1
	case 'I':
		level = 2
	case 'W':
		level = 3
	case 'E':
		level = 4
	case 'F':
		level = 5
	default:
		level = 6
	}

	if level == 6 {
		return fmt.Errorf("log level setting error")
	}

	return nil
}

// level: 0
func Trace(format string, v ...interface{}) {
	if level <= 0 {
		p(color.CyanBG("[T]")+" "+format, v...)
	}
}

// level: 1
func Debug(format string, v ...interface{}) {
	if level <= 1 {
		p(color.GreenBG("[D]")+" "+format, v...)
	}
}

// level: 2
func Info(format string, v ...interface{}) {
	if level <= 2 {
		p(color.BlueBG("[I]")+" "+format, v...)
	}
}

// level: 3
func Warn(format string, v ...interface{}) {
	if level <= 3 {
		p(color.YellowBG("[W]")+" "+format, v...)
	}
}

// level: 4
func Error(format string, v ...interface{}) {
	if level <= 4 {
		p(color.RedBG("[E]")+" "+format, v...)
	}
}

// level: 5
func Fatal(format string, v ...interface{}) {
	if level <= 5 {
		p(color.RedBG("[F]")+" "+format, v...)
	}
}

func p(format string, v ...interface{}) {
	_, filename, line, ok := runtime.Caller(2)
	if !ok {
		filename = "???"
		line = 0
	} else {
		filename = path.Base(filename)
	}
	v = append(v, filename)
	v = append(v, line)
	fmt.Printf(color.Cyan(time.Now().Format("2006/01/02 15:04:05"))+" "+format+" "+color.Blue("[%v:%v]")+"\n", v...)
}
