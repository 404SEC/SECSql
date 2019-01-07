package controller

import (
	"encoding/json"
	"html/template"
	"io/ioutil"

	"net/http"

	m "github.com/404SEC/SECSql/test/WebAppTest/models"
)

func ConnHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		t, _ := template.ParseFiles(Conf.Template+"Conn.html", Conf.Template+"header.html", Conf.Template+"footer.html")

		var datas []m.Connect
		redata, err := m.GetAllConn("-1")
		if err == nil {
			er := json.Unmarshal([]byte(redata), &datas)
			if er != nil {
				Logger.Error(er.Error())
			}
		}
		data := struct {
			Ta []m.Connect
		}{
			Ta: datas,
		}

		t.Execute(w, data)
	} else {
		var mexec m.PostConnect
		result, err := ioutil.ReadAll(r.Body)
		var ret string
		if err != nil {
			Logger.Error(err.Error())
		}

		er := json.Unmarshal(result, &mexec)
		if er != nil {
			Logger.Error(err.Error())
		}
		if mexec.Action == "addconn" {
			ret, _ = m.AddConnect(mexec)

		} else if mexec.Action == "moticonn" {
			ret, _ = m.MotiConnect(mexec)
		} else if mexec.Action == "delconn" {
			ret, _ = m.DelConnect(mexec.ID)

		}
		w.Write([]byte(ret))
	}
}
