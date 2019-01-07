package models

import (
	l "github.com/404SEC/SECSql/logic"
)

func GetAllConn(n string) (string, error) {
	DB, err := Open()
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	var Msql string
	if n == "0" {
		Msql = `SELECT * FROM Connect ORDER BY Types limit 20`
	} else if n == "-1" {
		Msql = `SELECT * FROM Connect ORDER BY Types `
	} else {

		Msql = `SELECT * FROM Connect ORDER BY Types limit ` + n + ` ,20`

	}
	rows, err := DB.Query(Msql)
	ret := " "
	if err != nil {
		return ret, err
	}
	ret, err = l.ITOjson(rows)
	if err != nil {
		return ret, err
	}
	rows.Close()
	defer DB.Close()

	return ret, nil
}
func GetConn(n string) (string, error) {
	DB, err := Open()
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	rows, err := DB.Query(`SELECT * FROM Connect where ID= ? `, n)

	ret := " "
	if err != nil {
		return ret, err
	}
	ret, err = l.ITOjson(rows)
	if err != nil {
		return ret, err
	}
	rows.Close()
	defer DB.Close()
	return ret, nil

}
func DelConnect(ID string) (string, error) {
	DB, err := Open()
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	stmt, _ := DB.Prepare(`DELETE FROM Connect WHERE ID = ?`)

	ret, err := stmt.Exec(ID)
	if err != nil {
		Logger.Error(err.Error())

		return `{"code":"false"}`, err
	}
	_, err = ret.RowsAffected()
	if err != nil {
		return `{"code":"false"}`, err
	}
	defer stmt.Close()
	defer DB.Close()

	return `{"code":"true"}`, err
}
func AddConnect(ta PostConnect) (string, error) {
	DB, err := Open()
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	stmt, err := DB.Prepare(`INSERT INTO Connect (Name,Types,ConnStr,Info) VALUES(?,?,?,?) `)
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	ret, err := stmt.Exec(ta.Name, ta.Types, ta.ConnStr, ta.Info)
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	_, err = ret.RowsAffected()
	if err != nil {
		return `{"code":"false"}`, err
	}
	defer stmt.Close()
	defer DB.Close()

	return `{"code":"true"}`, nil
}
func MotiConnect(ta PostConnect) (string, error) {
	DB, err := Open()
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	stmt, _ := DB.Prepare(`UPDATE Connect SET Name=? Types=? ConnStr=? Info=?) WHERE ID = ?`)
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	ret, err := stmt.Exec(ta.Name, ta.Types, ta.ConnStr, ta.Info, ta.ID)
	if err != nil {
		Logger.Error(err.Error())
		return `{"code":"false"}`, err
	}
	_, err = ret.RowsAffected()
	if err != nil {
		return `{"code":"false"}`, err
	}
	defer stmt.Close()
	defer DB.Close()

	return `{"code":"true"}`, err
}
