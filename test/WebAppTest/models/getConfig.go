package models

import (
	"encoding/json"
	"io/ioutil"
)

func GetConfig() Config {
	JsonParse := NewJsonStruct()
	Confi := Config{}

	JsonParse.Load("conf/app.conf", &Confi)
	Logger.Info("Loading Config File...")

	return Confi
}

type JsonStruct struct {
}

func NewJsonStruct() *JsonStruct {
	return &JsonStruct{}
}

func (jst *JsonStruct) Load(filename string, v interface{}) {

	data, err := ioutil.ReadFile(filename)
	if err != nil {
		Logger.Error(err.Error())

		return
	}
	err = json.Unmarshal(data, v)
	if err != nil {
		Logger.Error(err.Error())

		return
	}
}
