package csvstat

import "github.com/urfave/cli/v2"

var CSVStatCMD = &cli.Command{
	Name:    "csvstat",
	Aliases: []string{"stat"},
	Usage:   "打印CSV文件中每一列的描述性统计信息",
}
