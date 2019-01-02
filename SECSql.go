package SECSql

import (
	"database/sql"
	"log"

	l "github.com/404SEC/SECSql/logic"
)

type SECSql struct {
	Typ  string //数据库连接ID
	Conn string //数据库名
}

func (SECSql *SECSql) ExecOnece(Execstr string) (string, error) {

	sqlRet := ""
	switch SECSql.Typ {
	case "hive":
		sqlRet, err := l.HiveExec(SECSql.Conn, Execstr)
		if err != nil {
			return sqlRet, err
		}
	case "mysql", "postgres", "mssql", "sqlite3", "oracle":
		if SECSql.Typ == "oracle" {
			SECSql.Typ = "oci8"
		}

		IDb, err := sql.Open(SECSql.Typ, SECSql.Conn)
		if err != nil {
			log.Println(err.Error())
		}
		err = IDb.Ping()
		if err != nil {
			log.Println(err.Error())
		}

		sqlRet, err := l.DatabaseExec(IDb, Execstr)
		if err != nil {
			return sqlRet, err
		}
		IDb.Close()
	}
	return sqlRet, nil
}

func Ini(typ string, conn string) SECSql {
	var DB SECSql
	DB.Conn = conn
	DB.Typ = typ
	return DB
}
