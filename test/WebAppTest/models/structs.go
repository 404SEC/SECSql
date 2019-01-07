package models

type Config struct {
	IP         string
	Port       string
	DB         string
	DBStr      string
	Dev        string
	Cooexpires string
	Secure     string
	Ssl        string
	Timeouts   string
	Template   string
}
type Connect struct {
	ID      string
	Name    string
	Types   string
	ConnStr string
	Info    string
}
type PostConnect struct {
	Action  string
	ID      string
	Name    string
	Types   string
	ConnStr string
	Info    string
}

/*
CREATE TABLE Connect(
	ID serial not null,
	Name varchar(255),
	Types varchar(255),
	ConnStr varchar(255),
	Info   varchar(255),
	 PRIMARY KEY( ID , Name)
  );
*/
type Exec struct {
	ConnID  string //数据库连接ID
	Databse string //数据库名
	Tabl    string //表名
	Execstr string //数据库执行语句
	RunType string //执行命令控制，如点击数据库，点击表
	Limit   string //执行的limit
}
