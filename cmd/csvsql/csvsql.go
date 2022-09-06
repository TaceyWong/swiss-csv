package csvsql

import "github.com/urfave/cli/v2"

var CSVSQLCMD = &cli.Command{
	Name:    "csvsql",
	Aliases: []string{"sql"},
	Usage: `为一个或多个CSV文件生成SQL语句
	或直接在数据库上执行这些语句,并执行一个或多个SQL查询。`,
}
