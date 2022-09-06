package csvstat

import (
	"scsv/cmd"

	"github.com/urfave/cli/v2"
)

var CSVStatCMD = &cli.Command{}

func init() {
	CSVStatCMD = &cli.Command{
		Name:    "csvstat",
		Aliases: []string{"stat"},
		Usage:   "打印CSV文件中每一列的描述性统计信息",
		Action:  CSVStatAction,
	}
	CSVStatCMD.Flags = append(CSVStatCMD.Flags, cmd.PublicFlags...)
	CSVStatCMD.Flags = append(CSVStatCMD.Flags,
		&cli.BoolFlag{Name: "names", Aliases: []string{"n"},
			Usage: "输出输入CSV文件的列名和行索引并退出"},
		&cli.BoolFlag{Name: "csv", Usage: "结果以CSV输出,而不是文本"},
		&cli.StringFlag{Name: "columns", Aliases: []string{"c"},
			Usage: "要检查的列索引、名称或范围的逗号分隔列表(1,id,3-5)"},
		&cli.BoolFlag{Name: "type", Usage: "只输出数据类型"},
		&cli.BoolFlag{Name: "nulls", Usage: "只输出列是否包含空值"},
		&cli.BoolFlag{Name: "unique", Usage: "只输出唯一值的数量"},
		&cli.BoolFlag{Name: "min", Usage: "只输出最小值"},
		&cli.BoolFlag{Name: "max", Usage: "只输出最大值"},
		&cli.BoolFlag{Name: "sum", Usage: "只输出加和"},
		&cli.BoolFlag{Name: "mean", Usage: "只输出平均值"},
		&cli.BoolFlag{Name: "median", Usage: "只输出中间值"},
		&cli.BoolFlag{Name: "stdev", Usage: "只输出标准差"},
		&cli.BoolFlag{Name: "len", Usage: "只输出最长的值"},
		&cli.BoolFlag{Name: "freq", Usage: "只输出值频率列表"},
		&cli.Int64Flag{Name: "freq-count", Usage: "要显示的频率值的`最大数目`"},
		&cli.BoolFlag{Name: "count", Usage: "只输出总行数"},
	)

}

func CSVStatAction(ctx *cli.Context) error {
	return nil
}
