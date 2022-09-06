package csvformat

import "github.com/urfave/cli/v2"

var CSVFormatCMD = &cli.Command{
	Name:    "csvformat",
	Aliases: []string{"format"},
	Usage:   "将CSV文件转换为自定义输出格式",
}
