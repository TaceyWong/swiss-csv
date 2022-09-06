package csvsort

import "github.com/urfave/cli/v2"

var CSVSortCMD = &cli.Command{
	Name:    "csvsort",
	Aliases: []string{"sort"},
	Usage:   "对CSV文件进行排序(像Unix的Sort命令,但用于表格数据)",
}
