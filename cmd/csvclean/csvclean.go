package csvclean

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVCleanCMD = &cli.Command{
	Name:    "csvclean",
	Aliases: []string{"clean"},
	Usage:   "修复CSV文件中的常见错误",
}

func init() {
	CSVCleanCMD.Flags = append(CSVCleanCMD.Flags, cmd.PublicFlags...)
	CSVCleanCMD.Flags = append(CSVCleanCMD.Flags,
		&cli.BoolFlag{Name: "dry-run", Aliases: []string{"n"},
			Usage: "不创建输出文件,结果信息将会打印到标准错误(STDERR)",
		},
	)
}
