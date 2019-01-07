package logic

import (
	"fmt"
	"log"
	"strings"

	"github.com/dazheng/gohive"
)

func HiveExec(Conn string, ExecStr string) (string, error) {

	c1 := ""
	conn, err := gohive.Connect(Conn, gohive.DefaultOptions) // 无用户名、密码
	if err != nil {
		log.Println(err.Error())
		return `{"result":"false"}`, err
	}
	re := strings.Split(ExecStr, "\r\n")
	for _, v := range re {
		m := strings.TrimSpace(v)
		d := strings.Split(m, " ")
		if strings.ToLower(d[0]) == "select" || strings.ToLower(d[0]) == "show" || strings.ToLower(d[0]) == "desc" {
			rs, err := conn.Query(m)
			if err != nil {
				log.Println(err.Error())
				return `{"result":"false"}`, err
			}
			if c1 == "" {
				c1 = c1 + TOjson(rs)
			} else {
				c1 = c1 + "," + TOjson(rs)
			}

		} else {
			_, err = conn.Exec(m)
			if err != nil {

				c1 = `{"result":"false"}`
			}
		}
	}
	conn.Close()
	return c1, nil

}

func TOjson(rows gohive.RowSet) string {
	columns := rows.Columns()

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
