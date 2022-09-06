package csvstack

import "github.com/urfave/cli/v2"

var CSVStackCMD = &cli.Command{
	Name:    "csvstack",
	Aliases: []string{"stack"},
	Usage: `堆叠来自多个CSV文件的行,可以选择添加分组值。
	假定文件具有相同顺序的相同列。`,
}
