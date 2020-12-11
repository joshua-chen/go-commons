package logger

import (
	"os"
	"reflect"
	"strings"
	"sync"
	"time"

	"github.com/joshua-chen/go-commons/config"
	utilspath	"github.com/joshua-chen/go-commons/utils/path"
	"github.com/kataras/golog"
	_ "github.com/kataras/golog"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/context"
	irislogger "github.com/kataras/iris/v12/middleware/logger"

)

const deleteFileOnExit = true
const DATE_FORMAT = "2006-01-02"

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
		//MessageContextKeys: []string{"logger_message"},
		//如果不为空然后它的内容来自`ctx.GetHeader（“User-Agent”）
		//MessageHeaderKeys: []string{"User-Agent"},
	}
	if cfg != nil && len(cfg) > 0 {
		config = cfg[0]
	}
	return irislogger.New(config)
}

func NewRequestLogger() (handler context.Handler, close func() error) {
	handler, close = newRequestLogger()
	//defer close()
	return
}
func NewRequestLoggerForGolog() (handler context.Handler, close func() error) {

	handler, close = newRequestLoggerForGolog()

	return
}

//根据日期获取文件名，文件日志以最常用的方式工作
//但这些只是好的命名方式。
func todayFilename(dir string) string {
	if !utilspath.PathExisted(dir) {
		err := utilspath.MakeDir(dir)
		if err != nil {
			panic(err)
		}
	}
	if !strings.HasSuffix(dir, "/") {
		dir = dir + "/"
	}
	today := dir + time.Now().Format(DATE_FORMAT)
	return today + ".txt"
}

func newLogFile(dir string) *os.File {
	filename := todayFilename(dir)
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
		Query:   true,
		Columns: true,
	}

	logFile := newLogFile("./logs")
	close = func() error {		 
		return closeFile(logFile)
	}
	//golog.AddOutput(logFile)
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
	return
}

func closeFile(logFile *os.File) error {
	err := logFile.Close()
	deleteFile := deleteFileOnExit
	if !reflect.DeepEqual(config.AppConfig.Log, &config.Log{}) {
		deleteFile = config.AppConfig.Log.DeleteFileOnExit
	}
	if deleteFile {
		err = os.Remove(logFile.Name())
	}
	return err
}
func newRequestLoggerForGolog() (h iris.Handler, close func() error) {
	//defer close()
	close = func() error { return nil }

	logFile := newLogFile("./logs/golog")
	close = func() error {		 
		return closeFile(logFile)
	}
	golog.AddOutput(logFile)

	handler := func(ctx context.Context) {

		ctx.Next()
	}
	return handler, close
}
