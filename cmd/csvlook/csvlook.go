package csvlook

import "github.com/urfave/cli/v2"

var CSVLookCMD = &cli.Command{
	Name:    "csvlook",
	Aliases: []string{"look"},
	Usage:   "在控制台中将CSV文件呈现为与markdown兼容、固定宽度的表",
}
