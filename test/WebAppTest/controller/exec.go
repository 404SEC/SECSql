package controller

import (
	"encoding/json"
	"html/template"
	"io/ioutil"

	"net/http"

	SECSql "github.com/404SEC/SECSql"
	m "github.com/404SEC/SECSql/test/WebAppTest/models"
)

func ExecHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == "GET" {
		var datas []m.Connect
		redata, err := m.GetAllConn("-1")
		if err == nil && len(redata) > 0 {
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

		t, _ := template.ParseFiles(Conf.Template+"exec.html", Conf.Template+"header.html", Conf.Template+"footer.html")
		err = t.Execute(w, data)
		if err != nil {
			Logger.Error(err.Error())
		}

	} else {
		var mexec m.Exec
		result, err := ioutil.ReadAll(r.Body)
		if err != nil {
			Logger.Error(err.Error())
		}
		er := json.Unmarshal(result, &mexec)
		if er != nil {
			Logger.Error(er.Error())
		}

		var ret m.Connect
		Conn, err := m.GetConn(mexec.ConnID)
		if err != nil {
			return
		}

		err = json.Unmarshal([]byte(Conn[1:len(Conn)-1]), &ret)
		if err != nil {
			Logger.Error(err.Error())
		}
		mexec.ConnID = ret.ConnStr
		if mexec.RunType == "getdatabase" {
			mexec.Execstr = "show databases;"
		} else if mexec.RunType == "gettables" {
			mexec.Execstr = "use " + mexec.Databse + ";\r\n show tables;"
		} else if mexec.RunType == "getlimit" {
			mexec.Execstr = "use " + mexec.Databse + ";\r\n select * from " + mexec.Tabl + " limit 20"
		} else {
			mexec.Execstr = "use " + mexec.Databse + ";\r\n" + mexec.Execstr
		}

		DB := SECSql.Ini(ret.Types, ret.ConnStr)

		a, err := DB.ExecOnece(mexec.Execstr)
		if err != nil {
			Logger.Error(err.Error())
		}

		_, err = w.Write([]byte(a))
		if err != nil {
			Logger.Error(err.Error())
		}

	}

}
