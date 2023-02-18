package csvjoin

import (
	cs "github.com/mitchellh/colorstring"
	"github.com/urfave/cli/v2"
)

var CSVJoinCMD = &cli.Command{
	Name:        "csvjoin",
	Aliases:     []string{"join"},
	Usage:       "执行类似sql的Join来合并指定列上的CSV文件",
	Description: cs.Color("[light_red]注意：Join操作需要将所有文件读到内存，！！不要在非常大的文件上执行此操作！！[reset]"),
}
