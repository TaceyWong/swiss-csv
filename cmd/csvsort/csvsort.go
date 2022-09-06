package csvsort

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVSortCMD = &cli.Command{
	Name:    "csvsort",
	Aliases: []string{"sort"},
	Usage:   "对CSV文件进行排序(像Unix的Sort命令,但用于表格数据)",
}

func init() {
	CSVSortCMD.Flags = append(CSVSortCMD.Flags, cmd.PublicFlags...)
	CSVSortCMD.Flags = append(CSVSortCMD.Flags,
		&cli.BoolFlag{Name: "names", Aliases: []string{"n"},
			Usage: "显示输入CSV文件的行名和列索引并退出"},
		&cli.StringFlag{Name: "columns", Aliases: []string{"c"},
			Usage: "排序依据的列索引、名称或范围的逗号分隔列表(1,id,3-5)"},
		&cli.BoolFlag{Name: "reverse", Aliases: []string{"r"},
			Usage: "递降序排序"},
	)
}
