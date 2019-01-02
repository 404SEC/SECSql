package logic

import (
	"database/sql"
	"fmt"

	"log"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	_ "github.com/lib/pq"
	_ "github.com/mattn/go-oci8"
	_ "github.com/mattn/go-sqlite3"
)

func DatabaseExec(IDb *sql.DB, ExecStr string) string {

	//	IDb, err := sql.Open(typ, Conn)
	reCl := ""

	reStr := strings.Split(ExecStr, "\r\n")
	for _, v := range reStr {
		mexStr := strings.TrimSpace(v)
		d := strings.Split(mexStr, " ")
		if strings.ToLower(d[0]) == "select" || strings.ToLower(d[0]) == "show" || strings.ToLower(d[0]) == "desc" {
			if reCl == "" {
				reCl = reCl + ISelect(IDb, mexStr)
			} else {
				reCl = reCl + "," + ISelect(IDb, mexStr)
			}

		} else {
			if IExec(IDb, mexStr) {
				reCl = "ture"
			} else {
				reCl = "false"
			}

		}
	}
	return reCl
}

func IExec(IDB *sql.DB, str string) bool {

	ret, err := IDB.Exec(str)
	if err != nil {
		log.Println(err.Error())
		return false
	}
	_, err = ret.RowsAffected()
	if err != nil {
		return false
	}
	return true
}
func ISelect(IDB *sql.DB, str string) string {
	ret := "["
	rows, err := IDB.Query(str)
	if err != nil {
		log.Println(err.Error())
		return ret
	}
	return ITOjson(rows)
}
func ITOjson(rows *sql.Rows) string {
	columns, err1 := rows.Columns()
	if err1 != nil {
		log.Println(err1.Error())
		return ""
	}
	values := make([]string, len(columns))

	scanArgs := make([]interface{}, len(values))
	for i := range values {
		scanArgs[i] = &values[i]
	}

	list := "["

	for rows.Next() {
		rows.Scan(scanArgs...)
		row := "{"
		var value string
		for i, col := range values {
			if col == "" {
				value = "NULL"
			} else {
				value = string(col)
			}

			columName := strings.ToLower(columns[i])

			cell := fmt.Sprintf(`"%v":"%v"`, columName, value)
			row = row + cell + ","
		}
		row = row[0 : len(row)-1]
		row += "}"
		list = list + row + ","

	}
	list = list[0 : len(list)-1]
	list += "]"
	return list
}
