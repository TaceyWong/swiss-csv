package main

import (
	"log"
	"os"
	"scsv/cli/csvclean"
	"scsv/cli/csvcut"
	"scsv/cli/csvformat"
	"scsv/cli/csvgrep"
	"scsv/cli/csvjoin"
	"scsv/cli/csvlook"
	"scsv/cli/csvsort"
	"scsv/cli/csvsql"
	"scsv/cli/csvstack"
	"scsv/cli/csvstat"
	"scsv/cli/in2csv"
	"scsv/cli/sql2csv"

	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "scsv",
		Usage: "Swiss CSV,一套用于转换和使用CSV的实用工具",
		Action: func(ctx *cli.Context) error {
			cli.ShowAppHelp(ctx)
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
