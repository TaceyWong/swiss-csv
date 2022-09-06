package csvgrep

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVGrepCMD = &cli.Command{
	Name:    "csvgrep",
	Aliases: []string{"grep"},
	Usage:   "搜索CSV文件(类似于Unix的grep命令,但用于表格数据)",
}

func init() {
	CSVGrepCMD.Flags = append(CSVGrepCMD.Flags, cmd.PublicFlags...)
	CSVGrepCMD.Flags = append(CSVGrepCMD.Flags,
		&cli.BoolFlag{Name: "names", Aliases: []string{"n"},
			Usage: "显示输入CSV文件的行名和列索引并退出"},
		&cli.StringFlag{Name: "columns", Aliases: []string{"c"},
			Usage: "要检查的列索引、名称或范围的逗号分隔列表(1,id,3-5)"},
		&cli.StringFlag{Name: "match", Aliases: []string{"m"},
			Usage: "要搜索的`字符串`"},
		&cli.StringFlag{Name: "regex", Aliases: []string{"r"},
			Usage: "要匹配的`正则表达式`"},
		&cli.PathFlag{Name: "file", Aliases: []string{"f"},
			Usage: "一个文件路径，这个文件中任意一行(除去行分隔符)如果准确匹配cell值,该行就匹配"},
		&cli.BoolFlag{Name: "invert-match", Aliases: []string{"i"},
			Usage: "选择非匹模式"},
		&cli.BoolFlag{Name: "any-match", Aliases: []string{"a"},
			Usage: "选择列匹配的行，而非所有列"},
	)
}
