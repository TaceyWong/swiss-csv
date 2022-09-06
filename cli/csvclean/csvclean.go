package csvclean

import "github.com/urfave/cli/v2"

var CSVCleanCMD = &cli.Command{
	Name:    "csvclean",
	Aliases: []string{"clean"},
	Usage:   "修复CSV文件中的常见错误",
}
