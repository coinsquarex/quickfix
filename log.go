package quickfix

import (
	"fmt"
	"path/filepath"
	"runtime"
	"strings"
)

//Log is a generic interface for logging FIX messages and events.
type Log interface {
	//log incoming fix message
	OnIncoming([]byte)

	//log outgoing fix message
	OnOutgoing([]byte)

	//log fix event
	OnEvent(string)

	//log fix event according to format specifier
	OnEventf(string, ...interface{})
}

//The LogFactory interface creates global and session specific Log instances
type LogFactory interface {
	//global log
	Create() (Log, error)

	//session specific log
	CreateSessionLog(sessionID SessionID) (Log, error)
}

func logWithTracef(format string, a ...interface{}) string {
	msg := fmt.Sprintf(format, a...)
	return logWithTracePCSkip(msg, 3)
}

func logWithTrace(msg string) string {
	return logWithTracePCSkip(msg, 3)
}

func logWithTracePCSkip(msg string, pcSkip int) string {
	pc := make([]uintptr, 10) // at least 1 entry needed
	runtime.Callers(pcSkip, pc)
	f := runtime.FuncForPC(pc[0])
	file, line := f.FileLine(pc[0])
	file = filepath.Base(file)
	funcName := strings.ReplaceAll(f.Name(), getModuleName(pc[0]), "")
	return fmt.Sprintf("%s (%s:%d%s)", msg, file, line, funcName)
}

func getModuleName(pc uintptr) string {
	parts := strings.Split(runtime.FuncForPC(pc).Name(), ".")
	pl := len(parts)
	if parts[pl-2][0] == '(' {
		return strings.Join(parts[0:pl-2], ".")
	}
	return strings.Join(parts[0:pl-1], ".")
}
