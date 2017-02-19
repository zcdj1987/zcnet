package config

import (
	"github.com/Sirupsen/logrus"
	"os"
	"reflect"
)

const ()

var log = logrus.New()

func Read_config(typ string, path string, parameter interface{}) bool {
	var s bool
	if !check_typ(typ) || !check_path(path) || !check_parameter(parameter) {
		return false
	}
	switch typ {
	case "json":
		s = read_json(path, parameter)
	case "xlsx":
		s = read_xlsx(path, parameter)
	}
	return s
}

// ********************************************************************************
// ********************************************************************************
// ********************************************************************************

func check_typ(typ string) bool {
	if typ == "json" || typ == "xlsx" {
		return true
	}
	log.Errorln("文件类型存在错误")
	return false
}

func check_path(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	log.Errorln("文件路径存在错误")
	return false
}

func check_parameter(parameter interface{}) bool {
	if parameter != nil {
		sn := reflect.TypeOf(parameter).String()[0]
		if sn == 42 {
			return true
		}
	}
	log.Errorln("传入参数并非指针")
	return false
}
