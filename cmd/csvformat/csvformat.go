package csvformat

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVFormatCMD = &cli.Command{
	Name:    "csvformat",
	Aliases: []string{"format"},
	Usage:   "将CSV文件转换为自定义输出格式",
}

func init() {
	CSVFormatCMD.Flags = append(CSVFormatCMD.Flags, cmd.PublicFlags...)
	CSVFormatCMD.Flags = append(CSVFormatCMD.Flags,
		&cli.StringFlag{Name: "out-delimiter", Aliases: []string{"D"},
			Usage: "输出CSV文件的`分隔字符`",
		},
		&cli.BoolFlag{Name: "out-tabs", Aliases: []string{"T"},
			Usage: "指定输出CSV文件以制表键分隔,覆盖-d"},
		&cli.StringFlag{Name: "out-quotechar", Aliases: []string{"Q"},
			Usage: "输出CSV文件中引用字符串的`字符`"},
		&cli.BoolFlag{Name: "out-no-doublequote", Aliases: []string{"B"},
			Usage: "输出CSV文件中双引号是否被双引"},
		&cli.StringFlag{Name: "out-escapechar", Aliases: []string{"P"},
			Usage: "用于转义输出CSV文件中的分隔符的`字符`"},
		&cli.StringFlag{Name: "out-lineterminator", Aliases: []string{"M"},
			Usage: "用于结束输出CSV文件中的行的`字符`"},
	)
}
