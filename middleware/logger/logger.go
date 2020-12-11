package logger

import (
	"os"
	"strings"
	"sync"
	"time"

	"github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	irislogger "github.com/kataras/iris/v12/middleware/logger"

)

const deleteFileOnExit = true

type Logger struct {
}

var (
	instance *Logger
	lock     *sync.Mutex = &sync.Mutex{}
)

func Instance() *Logger {
	if instance == nil {
		lock.Lock()
		defer lock.Unlock()
		if instance == nil {
			instance = &Logger{}
		}
	}
	return instance
}
func (a *Logger) New() context.Handler {

	return NewConsoleLogger()
}

func NewConsoleLogger(cfg ...irislogger.Config) context.Handler {

	config := irislogger.Config{
		//状态显示状态代码
		Status: true,
		// IP显示请求的远程地址
		IP: true,
		//方法显示http方法
		Method: true,
		// Path显示请求路径
		Path: true,
		// Query将url查询附加到Path。
		Query: true,
		//Columns：true，
		// 如果不为空然后它的内容来自`ctx.Values(),Get("logger_message")
		//将添加到日志中。
		MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		MessageHeaderKeys: []string{"User-Agent"},
	}
	if len(cfg) > 0 {
		config = cfg[0]
	}
	return irislogger.New(config)
}

func NewRequestLogger() context.Handler {
	r, close := newRequestLogger()
	defer close()
	return r
}

//根据日期获取文件名，文件日志以最常用的方式工作
//但这些只是好的命名方式。
func todayFilename() string {
	today := time.Now().Format("Jan 02 2006")
	return today + ".txt"
}

func newLogFile() *os.File {
	filename := todayFilename()
	//打开一个输出文件，如果重新启动服务器，它将追加到今天的文件中
	f, err := os.OpenFile(filename, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	return f
}

var excludeExtensions = [...]string{
	".js",
	".css",
	".jpg",
	".png",
	".ico",
	".svg",
}

func newRequestLogger() (h iris.Handler, close func() error) {
	close = func() error { return nil }
	c := irislogger.Config{
		Status:  true,
		IP:      true,
		Method:  true,
		Path:    true,
		Columns: true,
	}
	logFile := newLogFile()
	close = func() error {
		err := logFile.Close()
		if deleteFileOnExit {
			err = os.Remove(logFile.Name())
		}
		return err
	}
	c.LogFunc = func(now time.Time, latency time.Duration, status, ip, method, path string, message interface{}, headerMessage interface{}) {
		output := irislogger.Columnize(now.Format("2006/01/02 - 15:04:05"), latency, status, ip, method, path, message, headerMessage)
		logFile.Write([]byte(output))
	}
	
	//我们不想使用记录器，一些静态请求等
	c.AddSkipper(func(ctx iris.Context) bool {
		path := ctx.Path()
		for _, ext := range excludeExtensions {
			if strings.HasSuffix(path, ext) {
				return true
			}
		}
		return false
	})
	h = irislogger.New(c)

	golog.AddOutput(logFile)
	return
}
