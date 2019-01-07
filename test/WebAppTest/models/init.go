package models

import (
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	go_logger "github.com/phachon/go-logger"
)

var Conf Config
var Logger *go_logger.Logger

func init() {
	Logger = go_logger.NewLogger()
	Logger.SetAsync()
	Conf = GetConfig()

	if Conf.Dev == "1" {
		Logger.Detach("console")
		consoleConfig := &go_logger.ConsoleConfig{
			Color:      true,  // 命令行输出字符串是否显示颜色
			JsonFormat: false, // 命令行输出字符串是否格式化
			Format:     "",    // 如果输出的不是 json 字符串，JsonFormat: false, 自定义输出的格式
		}

		Logger.Attach("console", go_logger.LOGGER_LEVEL_DEBUG, consoleConfig)
	}
	fileConfig := &go_logger.FileConfig{
		//Filename: "./test.log", // The file name of the logger output, does not exist automatically
		// If you want to separate separate logs into files, configure LevelFileName parameters.
		LevelFileName: map[int]string{
			Logger.LoggerLevel("error"): "./log/error.log", // The error level log is written to the error.log file.
			Logger.LoggerLevel("info"):  "./log/info.log",  // The info level log is written to the info.log file.
			Logger.LoggerLevel("debug"): "./log/debug.log", // The debug level log is written to the debug.log file.
		},
		MaxSize:    1024 * 1024, // File maximum (KB), default 0 is not limited
		MaxLine:    100000,      // The maximum number of lines in the file, the default 0 is not limited
		DateSlice:  "d",         // Cut the document by date, support "Y" (year), "m" (month), "d" (day), "H" (hour), default "no".
		JsonFormat: true,        // Whether the file data is written to JSON formatting
		Format:     "",          // JsonFormat is false, logger message written to file format string
	}

	// add output to the file
	Logger.Attach("file", go_logger.LOGGER_LEVEL_DEBUG, fileConfig)

}
func Open() (*sql.DB, error) {
	DB, err := sql.Open(Conf.DB, Conf.DBStr)
	err = DB.Ping()
	if err != nil {
		Logger.Error(err.Error())
	}
	return DB, err
}
