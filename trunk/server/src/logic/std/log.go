package std

import (

	"os"
	"time"
	"path"
	"fmt"
	"sync"
	"runtime"
	"strings"
)

const (
	//levels
	LDebug       = 1 << iota
	LWarning
	LError
	LFatal
	LInformation

	LLongFile

	LShortFile

	allLevels = LDebug | LWarning | LError | LFatal  | LInformation

	LConsole

	allFlags = allLevels | LShortFile | LConsole

	LDefaultFlags = allFlags

)



type Logger struct {
	sync.Mutex
	flags int
	file  *os.File
	outpath string
}

func (this *Logger) Close() {
	if this.file != nil {
		this.file.Close()
	}
	this.file = nil
}

func (this* Logger) SetFlags(flags int){
	flags = flags & allFlags
	this.Lock()
	defer this.Unlock()
	this.flags = flags
}

func (this *Logger) GetFlags()int{
	this.Lock()
	defer this.Unlock()
	return this.flags
}

func (this* Logger) sprintf(l int, t time.Time , f string, a ...interface{}) string {

	s := []string{
		"debug",
		"warning",
		"error",
		"fatal",
		"info",
	}

	var(
		file string
		line int
		ok bool
	)

	if this.flags & (LLongFile|LShortFile) != 0 {

		_, file, line, ok = runtime.Caller(4)

		if !ok {
			file = "???"
			line = 0
		}else {
			if this.flags & LShortFile !=0 {
				i := strings.LastIndex(file, "/")
				if i == -1 {
					i = strings.LastIndex(file, "\\")
				}
				if i != -1{
					file = file[i+1:]
				}
			}
		}
	}

	return fmt.Sprintf("%s|%s|%s [%s:%d]\n",t.String(),s[l],fmt.Sprintf(f,a...),file,line)


}

func (this* Logger) Println(l int, f string, a ...interface{}){
	this.Lock()
	defer this.Unlock()

	if this.flags & l == 0 {
		return
	}

	l = l & allLevels
	if l != 0 {
		for i:=0; i <= LInformation; i++{
			if ((1<<uint(i)) & l) !=0 {
				l = i
				break
			}
		}
	}

	s := this.sprintf(l,time.Now(),f,a...)

	if this.flags & LConsole != 0{
		c := []int{32,34,33,31,36}
		fmt.Printf("\x1b[0;%dm%s\x1b[0m",c[l],s)
	}

	if this.file != nil{
		this.file.WriteString(s)
	}
}

func (this* Logger) Debug(format string, a ...interface{}){
	this.Println(LDebug,format, a...)
}

func (this* Logger) Warning(format string, a ...interface{}){
	this.Println(LWarning,format, a...)
}

func (this* Logger) Error(format string, a ...interface{}){
	this.Println(LError,format, a...)
}

func (this* Logger) Fatal(format string, a ...interface{}){
	this.Println(LFatal,format, a...)
	os.Exit(1)
}

func (this* Logger) Info(format string, a ...interface{}){
	this.Println(LInformation,format, a...)
}

func (this* Logger) Backup() error {
	var (
		file *os.File
		erro error
	)
	if this.outpath != ""{
		now := time.Now()
		filename := fmt.Sprintf("%s_%d.log", now.Format("2006_01_02_15_04_05"),now.Nanosecond())
		file, erro = os.Create(path.Join(this.outpath,filename))
		if erro != nil {
			return  erro
		}
	}

	if this.file != nil {
		this.file.Close()
	}

	this.file = file

	return  nil
}

func New(outpath string, flags int) (*Logger, error){
	l := &Logger{sync.Mutex{}, flags, nil , outpath}
	return l , l.Backup()
}

var defaultLogger ,_ = New("../",LDefaultFlags)

func LogPrintln(level int, format string, a ...interface{}){
	defaultLogger.Println(level,format,a...)
}

func SetLogFlags(flags int){
	defaultLogger.SetFlags(flags)
}

func GetLogFlags()int {
	return  defaultLogger.GetFlags()
}

func LogDebug(format string, a ...interface{}){
	defaultLogger.Debug(format, a...)
}

func LogWarning(format string, a ...interface{}){
	defaultLogger.Warning(format, a...)
}

func LogError(format string, a ...interface{}){
	defaultLogger.Error(format, a...)
}

func LogFatal(format string, a ...interface{}){
	defaultLogger.Fatal(format, a...)
}

func LogInfo(format string, a ...interface{}){
	defaultLogger.Info(format, a...)
}

func Log(a ...interface{}){
	defaultLogger.Debug("%s",fmt.Sprint(a...))
}

func LogBackup(){
	defaultLogger.Backup()
}
