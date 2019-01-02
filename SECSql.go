package SECSql

import (
	"database/sql"
	"log"

	l "github.com/404SEC/SECSql/logic"
)

var (
	SECSqlTyp     string //数据库连接ID
	SECSqlConn    string //数据库名
	SECSqlExecstr string //数据库执行语句
)

func ExecOnece(exeStr string) string {
	SECSqlExecstr = exeStr

	sqlRet := ""
	switch SECSqlTyp {
	case "hive":
		sqlRet = l.HiveExec(SECSqlConn, SECSqlExecstr)
	case "mysql", "postgres", "mssql", "sqlite3", "oracle":
		if SECSqlTyp == "oracle" {
			SECSqlTyp = "oci8"
		}

		IDb, err := sql.Open(SECSqlTyp, SECSqlConn)
		if err != nil {
			log.Println(err.Error())
		}
		err = IDb.Ping()
		if err != nil {
			log.Println(err.Error())
		}

		sqlRet = l.DatabaseExec(IDb, SECSqlExecstr)
		IDb.Close()
	}
	return sqlRet
}

func Open(typ string, conn string) {

	SECSqlConn = conn
	SECSqlTyp = typ

}
