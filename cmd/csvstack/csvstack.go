package csvstack

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVStackCMD = &cli.Command{
	Name:    "csvstack",
	Aliases: []string{"stack"},
	Usage: `堆叠来自多个CSV文件的行,可以选择添加分组值。
	假定文件具有相同顺序的相同列。`,
}

func init() {
	CSVStackCMD.Flags = append(CSVStackCMD.Flags, cmd.PublicFlags...)
	CSVStackCMD.Flags = append(CSVStackCMD.Flags,
		&cli.StringFlag{Name: "groups", Aliases: []string{"g"},
			Usage: "每堆叠一个CSV要添加为“`分组`因子”的逗号分隔值列表,它们作为新列添加到输出中。可以使用-n标志为新列指定一个名称。",
		},
		&cli.StringFlag{Name: "group-name", Aliases: []string{"n"},
			Usage: "指定分组列名称，只有指定-g时才生效"},
		&cli.BoolFlag{Name: "filenames", Usage: "使用每个输入文件的文件名作为其分组值,此项将会导致-g选项失效"},
	)
}
