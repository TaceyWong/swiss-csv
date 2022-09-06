package csvlook

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVLookCMD = &cli.Command{
	Name:    "csvlook",
	Aliases: []string{"look"},
	Usage:   "在控制台中将CSV文件呈现为与markdown兼容、固定宽度的表",
}

func init() {
	CSVLookCMD.Flags = append(CSVLookCMD.Flags, cmd.PublicFlags...)
	CSVLookCMD.Flags = append(CSVLookCMD.Flags,
		&cli.Int64Flag{Name: "max-rows", Usage: "在截断数据之前显示的`最大行数`"},
		&cli.Int64Flag{Name: "max-columns", Usage: "在截断数据之前显示的`最大列数`"},
		&cli.Int64Flag{Name: "max-column-width", Usage: "将所有列截断为最多的`宽度`,其余的将被省略号代替"},
	)
}
