package csvgrep

import "github.com/urfave/cli/v2"

var CSVGrepCMD = &cli.Command{
	Name:    "csvgrep",
	Aliases: []string{"grep"},
	Usage:   "搜索CSV文件(类似于Unix的grep命令,但用于表格数据)",
}
