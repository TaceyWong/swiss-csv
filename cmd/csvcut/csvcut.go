package csvcut

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVCutCMD = &cli.Command{
	Name:    "csvcut",
	Aliases: []string{"cut"},
	Usage:   "过滤和截断CSV文件(类似于Unix的“cut”命令,但用于表格数据)",
}

func init() {
	CSVCutCMD.Flags = append(CSVCutCMD.Flags, cmd.PublicFlags...)
	CSVCutCMD.Flags = append(CSVCutCMD.Flags,
		&cli.BoolFlag{Name: "names", Aliases: []string{"n"},
			Usage: "显示输入CSV文件的列名和列索引名并退出"},
		&cli.StringFlag{Name: "columns", Aliases: []string{"c"},
			Usage: "要提取的列索引、名称或范围的逗号`分隔列表`,例如“1,id,3-5”。默认为所有列"},
		&cli.StringFlag{Name: "not-columns", Aliases: []string{"C"},
			Usage: "要排除的列索引、名称或范围的逗号`分隔列表`,例如“1,id,3-5”。默认为所有列"},
		&cli.BoolFlag{Name: "delete-empty-rows", Aliases: []string{"x"},
			Usage: "剪切后，删除完全为空的行。"},
	)
}
