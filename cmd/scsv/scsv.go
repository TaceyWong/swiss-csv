package main

import (
	"log"
	"os"
	"scsv/cmd"
	"scsv/cmd/csvclean"
	"scsv/cmd/csvcut"
	"scsv/cmd/csvformat"
	"scsv/cmd/csvgrep"
	"scsv/cmd/csvjoin"
	"scsv/cmd/csvlook"
	"scsv/cmd/csvsort"
	"scsv/cmd/csvsql"
	"scsv/cmd/csvstack"
	"scsv/cmd/csvstat"
	"scsv/cmd/in2csv"
	"scsv/cmd/sql2csv"

	"github.com/urfave/cli/v2"
)

func init() {
	cli.AppHelpTemplate = cmd.AppHelpTemplate
	cli.CommandHelpTemplate = cmd.CommandHelpTemplate
	cli.SubcommandHelpTemplate = cmd.SubcommandHelpTemplate
	cli.MarkdownDocTemplate = cmd.MarkdownDocTemplate
	cli.VersionFlag = &cli.BoolFlag{
		Name:    "version",
		Aliases: []string{"v"},
		Usage:   "显示版本信息",
	}
	cli.HelpFlag = &cli.BoolFlag{
		Name:    "help",
		Aliases: []string{"h"},
		Usage:   "显示帮助信息",
	}

}

func main() {
	app := &cli.App{
		Name:            "scsv",
		Usage:           "Swiss CSV,一套用于转换和使用CSV的实用工具",
		Version:         "0.1.0",
		HideHelpCommand: true,
		Authors:         []*cli.Author{{Name: "Tacey Wong", Email: "xinyong.wang@qq.com"}},
		Copyright:       "© 2022 Tacey Wong",
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelpAndExit(ctx, 0)
			return nil
		},
	}
	app.Commands = []*cli.Command{
		csvclean.CSVCleanCMD,
		csvcut.CSVCutCMD,
		csvstack.CSVStackCMD,
		csvstat.CSVStatCMD,
		csvjoin.CSVJoinCMD,
		csvformat.CSVFormatCMD,
		csvgrep.CSVGrepCMD,
		csvjoin.CSVJoinCMD,
		csvlook.CSVLookCMD,
		csvsort.CSVSortCMD,
		csvsql.CSVSQLCMD,
		in2csv.In2CSVCMD,
		sql2csv.SQL2CSVCMD,
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
