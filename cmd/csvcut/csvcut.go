package csvcut

import "github.com/urfave/cli/v2"

var CSVCutCMD = &cli.Command{
	Name:    "csvcut",
	Aliases: []string{"cut"},
	Usage:   "过滤和截断CSV文件(类似于Unix的“cut”命令,但用于表格数据)",
}
