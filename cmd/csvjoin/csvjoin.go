package csvjoin

import "github.com/urfave/cli/v2"

var CSVJoinCMD = &cli.Command{
	Name:    "csvjoin",
	Aliases: []string{"join"},
	Usage:   "执行类似sql的Join来合并指定列上的CSV文件",
}
