package utils

import (
	"errors"
	"go.uber.org/zap"
	"os"
	"server/global"
)

func PathExists(path string) (bool, error) {
	fi, err := os.Stat(path)
	if err == nil {
		if fi.IsDir() {
			return true, nil
		}

		return false, errors.New("存在同名文件")
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return false, err
}

func CreateDir(dirs ...string) error {
	for _, v := range dirs {
		exist, err := PathExists(v)
		if nil != err {
			return err
		}

		if !exist {
			global.Logger.Debug("create directory" + v)
			if err := os.MkdirAll(v, os.ModePerm); nil != err {
				global.Logger.Error("create directory"+v, zap.Any("error : ", err))
				return err
			}
		}
	}

	return nil
}
