package log

import (
	"apiproject/config"
	"apiproject/pkg"
	"io"
	"os"
	"path"

	"github.com/sirupsen/logrus"
)

func SetupLog() {
	lever, err := logrus.ParseLevel(config.App.Log.Level)
	if err != nil {
		logrus.Errorf("logrus parse level error:[%v]\n", err)
		return
	}
	logrus.SetLevel(lever)
	//先判断该文件是否存在，如果不存在的话就创建
	logfile := path.Join(config.App.Log.LogPath, config.App.Log.LogFile)
	if ok, _ := pkg.PathExist(logfile); !ok {
		pkg.CreateFile(logfile)
	}
	logFile, err := os.OpenFile(logfile, os.O_APPEND|os.O_WRONLY|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}

	//设置双写入，文件以及控制台
	mw := io.MultiWriter(os.Stdout, logFile)
	logrus.SetOutput(mw)
	logrus.SetReportCaller(true)

}
