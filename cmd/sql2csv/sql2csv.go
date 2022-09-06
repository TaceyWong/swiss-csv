package sql2csv

import "github.com/urfave/cli/v2"

var SQL2CSVCMD = &cli.Command{
	Name:  "sql2csv",
	Usage: "在数据库上执行SQL查询,并将结果输出到CSV文件",
}
