package log

import (
	"apiproject/config"

	"github.com/sirupsen/logrus"
)

//var Log = logrus.New()

func SetupLog() {
	lever, err := logrus.ParseLevel(config.App.Log.Level)
	if err != nil {
		logrus.Errorf("logrus parse level error:[%v]\n", err)
		return
	}
	logrus.SetLevel(lever)
	//设置双写入，文件以及控制台

}
