package cmd

import (
	"github.com/urfave/cli/v2"
)

var PublicFlags = []cli.Flag{
	&cli.StringFlag{Name: "delimiter", Aliases: []string{"d"}, Usage: "输入CSV文件的`分隔符`"},
	&cli.BoolFlag{Name: "tabs", Aliases: []string{"t"}, Usage: "输入CSV文件是否用制表符分隔。用[-d]覆盖"},
	&cli.StringFlag{Name: "quotechar", Aliases: []string{"q"}, Usage: "输入CSV文件中引用字符串的`引用字符`"},
	&cli.GenericFlag{
		Name:    "quoting",
		Aliases: []string{"u"},
		Value: &EnumValue{
			Enum:    []string{"0", "1", "2", "3"},
			Default: "3",
		},
		Usage: "输入CSV文件的引用风格`{0,1,2,3}`" + `
	(0最小引用1引用所有2引用非数字3无引用).`,
	},
	&cli.BoolFlag{Name: "no-doublequote", Aliases: []string{"b"}, Usage: "输入CSV文件中的双引用符号是否被双引"},
	&cli.StringFlag{Name: "escapechar", Aliases: []string{"p"}, Usage: "转义分隔符的`字符`"},
	&cli.Int64Flag{Name: "maxfieldsize", Aliases: []string{"z"},
		Usage: "输入CSV文件单一字段`最大长度`"},
	&cli.StringFlag{Name: "encoding", Aliases: []string{"e"},
		Usage: "指定输入CSV文件的`编码`"},
	&cli.StringFlag{Name: "local", Aliases: []string{"L"},
		Usage: "指定任意格式化数字的本地`代码`(比如中国大陆是zh_CN)"},
	&cli.BoolFlag{Name: "skipinitialspace", Aliases: []string{"S"},
		Usage: "忽略分隔符后的空白"},
	&cli.BoolFlag{Name: "blanks", Usage: "不将na、n/a、none、.、null转换成NULL"},
	&cli.StringFlag{Name: "date-format", Usage: "指定`日期格式`"},
	&cli.StringFlag{Name: "datetime-format", Usage: "指定`日期时间格式`"},
	&cli.BoolFlag{Name: "no-header-row", Aliases: []string{"H"}, Usage: "输入CSV文件无Header行,将会创建1.2.3...Header行"},
	&cli.Int64Flag{Name: "skip-lines", Aliases: []string{"K"},
		Usage: "跳过Header行之前的初始行数,比如注释、版权信息、空白行"},
	&cli.BoolFlag{Name: "verbose", Aliases: []string{"V"},
		Usage: "当错误发生打印详细溯回追踪信息"},
	&cli.BoolFlag{Name: "linenumbers", Aliases: []string{"l"},
		Usage: "输出加上一列行号"},
	&cli.BoolFlag{Name: "zero", Usage: "行号以0为基数开始(缺省为1)"},

	&cli.StringFlag{Name: "schema", Aliases: []string{"s"},
		Usage: "指定CSV格式schema"},
	&cli.StringFlag{Name: "key", Aliases: []string{"k"},
		Usage: "处理JSON时指定对象列表的顶部键"},

	&cli.StringFlag{Name: "write-sheets",
		Usage: "写入文件的Sheet名,-标识写入所有Sheet",
		Value: "-"},
	&cli.StringFlag{Name: "encoding-xls", Usage: "指定XLS文件的`编码`"},
	&cli.Int64Flag{Name: "snifflimit", Aliases: []string{"y"},
		Usage: "将CSVdialect `sniffing`限制为指定字节数(0关闭-1整个文件)"},
	&cli.BoolFlag{Name: "no-inference", Aliases: []string{"I"},
		Usage: "解析CSV输入时禁用类型推断(和--locale, --date-format, --datetime-format)"},
}
