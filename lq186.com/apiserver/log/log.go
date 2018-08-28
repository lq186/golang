package log

import (
	"fmt"
	"github.com/lq186/golang/lq186.com/apiserver/config"
	"github.com/phachon/go-logger"
	"os"
	"strings"
)

var (
	Log *go_logger.Logger
)

func init() {
	Log = go_logger.NewLogger()

	// 默认已经添加了 console 的输出，默认不显示颜色，如果需要修改，先删除掉 console
	Log.Detach("console")

	// 配置 console adapter
	console := &go_logger.ConsoleConfig{
		Color: true, // 文字是否显示颜色
		JsonFormat: false, // 是否格式化成 json 字符串
		Format: logFormat(),
	}

	// 添加输出到命令行
	// console: adapter name
	// level: go_logger.LOGGER_LEVEL_DEBUG
	// config: go_logger.NewConfigConsole(console)
	Log.Attach("console", logLevel(), console)

	// 配置 file adapter
	logPath, err := logPath()
	if err != nil {
		fmt.Println("Can not create file logger. More info: ", err)
	} else {
		fileConfig := &go_logger.FileConfig{
			Filename: logPath + "/all.log", // 日志输出的文件名, 不存在会自动创建
			// 如果想要将不同级别的日志单独输出到文件，配置 LevelFileName 参数
			LevelFileName: map[int]string{
				go_logger.LOGGER_LEVEL_ERROR: logPath + "/error.log", // 会将 error 级别的日志写入到 error.log 文件里
				go_logger.LOGGER_LEVEL_INFO:  logPath + "/info.log",  // 会将 info  级别的日志写入到 info.log  文件里
				go_logger.LOGGER_LEVEL_DEBUG: logPath + "/debug.log", // 会将 debug 级别的日志写入到 debug.log 文件里
			},
			MaxSize:    1024 * 1024, // 文件最大(kb) ，默认 0 不限制
			MaxLine:    100000,      // 文件最多多少行，默认 0 不限制
			DateSlice:  "d",         // 按日期切分文件，支持 "y"(年), "m"(月), "d"(日), "h"(小时), 默认 "" 不限制
			JsonFormat: true,        // 写入文件数据是否 json 格式化
			Format: logFormat(),
		}
		Log.Attach("file", logLevel(), fileConfig)
	}
}

func logLevel() int {
	level := strings.ToUpper(config.Config().Log.Level)
	if "DEBUG" == level {
		return go_logger.LOGGER_LEVEL_DEBUG
	} else if "INFO" == level {
		return go_logger.LOGGER_LEVEL_INFO
	} else if "WARNING" == level {
		return go_logger.LOGGER_LEVEL_WARNING
	} else if "ERROR" == level {
		return go_logger.LOGGER_LEVEL_ERROR
	} else {
		return go_logger.LOGGER_LEVEL_INFO
	}
}

func logFormat() string {
	logFormat := config.Config().Log.Format
	if "" == logFormat {
		logFormat = "%millisecond_format% (%function% %line%) [%level_string%] %body%"
	}
	return logFormat
}

func logPath() (string, error) {
	logPath := config.Config().Log.Path
	err := os.MkdirAll(logPath, 0644)
	return logPath, err
}
