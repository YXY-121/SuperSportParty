package pkg

import (
	"os"

	"github.com/sirupsen/logrus"
)

func PathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}
func CreateFile(file string) {
	_, err := os.Create(file)
	if err != nil {
		logrus.Errorf("mkdir fail [%v] \n", err)
		return
	}
	logrus.Debugln("mkdir successfully ")

}
