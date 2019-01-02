# SECSql
database manager Suitable for multiple databases
```
go get github.com/404SEC/SECSql
```
test code
```
package main

import (
	"fmt"

	SECSql "github.com/404SEC/SECSql"
)

func main() {

	DB := SECSql.Ini("mysql", "root:123456@tcp(127.0.0.1:3306)/BDP?charset=utf8")
	a, err := DB.ExecOnece("use mysql;\r\nselect * from user")
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println(a)
}
```
Apply to : mysql,mssql, sqlite3, oracle, postgres, hive 

