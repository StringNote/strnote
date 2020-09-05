package main

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/labstack/echo/v4"
)

// Note はノートの記録形式です。
type Note struct {
	Name string `json:"name"`
	Mes  string `json:"mes"`
	UTC  string `json:"utc"`
}

// ノートを取得
func getNote(c echo.Context, uid string) (*Note, error) {
	data := Note{
		Name: "",
		Mes:  "",
		UTC:  "20200901000000",
	}
	if jsonstr, _ := getValue(c.Request(), uidPRE+uid); jsonstr != "" {
		err := json.Unmarshal([]byte(jsonstr), &data)
		if err != nil {
			logOutput(jsonstr)
			logOutput(err.Error())
			return &data, fmt.Errorf("invalid json %s : %s", uid, err.Error())
		}
		return &data, nil
	}
	return &data, fmt.Errorf("unknown uid %s", uid)
}

// ノートを保存
func setNote(c echo.Context, uid, name, mes string) error {
	if name == "" && mes == "" {
		// 変更なし
		return nil
	}
	if name == "" || mes == "" {
		if note, err := getNote(c, uid); err == nil {
			if name == "" {
				name = note.Name
			}
			if mes == "" {
				mes = note.Mes
			}
		}
	}
	data := Note{
		Name: name,
		Mes:  mes,
		UTC:  time2ymdhms(time.Now().UTC()),
	}
	mesjson, _ := json.Marshal(data)
	strjson := string(mesjson)
	return setValue(c.Request(), uidPRE+uid, strjson)
}
