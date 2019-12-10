package stdout

import (
	"fmt"
	"os"
)

type Loggious struct {
	Desc    Descripter
	logfunc StdoutFunc
}

func (lg Loggious) Log(cmd string, stat int, desc string) {
	log := lg.logfunc(cmd, stat, desc)
	if lg.Desc == FILE {
		current, _ := os.Getwd()
		fp, err := os.OpenFile(current+"/loggious.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0755)
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(fp, log)
		defer fp.Close()
		return
	}
	if lg.Desc == CONSOLE {
		fmt.Println(log)
	}
	os.Exit(1)
}

func NewLoggious(d Descripter, lf StdoutFunc) *Loggious {
	return &Loggious{
		Desc:    d,
		logfunc: lf,
	}
}

type Descripter int

const (
	FILE Descripter = iota
	CONSOLE
)

type Logger interface {
	Log(cmd string, stat int, desc string) string
}

type StdoutFunc func(cmd string, stat int, desc string) string

func (l StdoutFunc) Log(cmd string, stat int, desc string) string {
	return l(cmd, stat, desc)
}
