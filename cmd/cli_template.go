package cmd

import (
	cs "github.com/mitchellh/colorstring"
)

var AppHelpTemplate = cs.Color(`
[light_red][bold]{{$v := offset .Name 6}}{{wrap .Name 3}}[reset]{{if .Usage}} - [light_blue]{{wrap .Usage $v}}[reset]{{end}}

[light_green]使用方法[reset]:
   [light_gray]{{if .UsageText}}{{wrap .UsageText 3}}{{else}}{{.HelpName}} {{if .VisibleFlags}}[全局选项]{{end}}{{if .Commands}} 命令 [命令选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}{{end}}[reset]{{if .Version}}{{if not .HideVersion}}
[light_green]版本[reset]:
   [red]{{.Version}}{{end}}{{end}}[reset]{{if .Description}}
[light_green]描述[reset]:
   {{wrap .Description 3}}{{end}}{{if .VisibleCommands}}
[light_green]命令[reset]:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{range .VisibleCommands}}
     [green]{{join .Names ", "}}[reset]{{"\t"}}[light_gray]{{.Usage}}[reset]{{end}}{{else}}{{range .VisibleCommands}}
   [green]{{join .Names ", "}}[reset]{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{end}}{{if .VisibleFlagCategories}}
[light_green]全局选项[reset]:{{range .VisibleFlagCategories}}
   {{if .Name}}{{.Name}}
   {{end}}{{range .Flags}}{{.}}
   {{end}}{{end}}{{else}}{{if .VisibleFlags}}
[light_green]全局选项[reset]:
   {{range $index, $option := .VisibleFlags}}{{if $index}}
   {{end}}{{wrap $option.String 6}}{{end}}{{end}}{{end}}{{if len .Authors}}
[light_green]作者[reset]{{with $length := len .Authors}}{{if ne 1 $length}}S{{end}}{{end}}:
   {{range $index, $author := .Authors}}{{if $index}}
   {{end}}{{$author}}{{end}}{{end}}
[light_green]版权信息[reset]:{{if .Copyright}}
   {{wrap .Copyright 3}}{{end}}
`)

var CommandHelpTemplate = cs.Color(`
[light_red]{{$v := offset .HelpName 6}}{{wrap .HelpName 3}}[reset]{{if .Usage}} - [light_blue]{{wrap .Usage $v}}[reset]{{end}}

[light_green]使用方法[reset]:
   [light_gray]{{if .UsageText}}{{wrap .HelpName 3}} {{.UsageText}}{{else}}{{.HelpName}}{{if .VisibleFlags}} [命令选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}{{end}}[reset]{{if .Category}}
[light_green]分类[reset]:
   {{.Category}}{{end}}{{if .Description}}
[light_green]描述[reset]:
   {{wrap .Description 3}}{{end}}{{if .VisibleFlagCategories}}
[light_green]选项[reset]:{{range .VisibleFlagCategories}}
   {{if .Name}}{{.Name}}
   {{end}}{{range .Flags}}{{.}}
   {{end}}{{end}}{{else}}{{if .VisibleFlags}}
[light_green]选项[reset]:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}{{end}}
`)

var SubcommandHelpTemplate = cs.Color(`
[light_blue]{{.HelpName}}[reset] - [light_gray]{{.Usage}}[reset]

[light_green]使用方法[reset]:
   {{if .UsageText}}{{wrap .UsageText 3}}{{else}}{{.HelpName}} 命令{{if .VisibleFlags}} [命令选项]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[参数...]{{end}}{{end}}{{if .Description}}
[light_green]描述[reset]:
   {{wrap .Description 3}}{{end}}
[light_green]命令[reset]:{{range .VisibleCategories}}{{if .Name}}
   {{.Name}}:{{range .VisibleCommands}}
     {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{else}}{{range .VisibleCommands}}
   {{join .Names ", "}}{{"\t"}}{{.Usage}}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
[light_green]选项[reset]:
   {{range .VisibleFlags}}{{.}}
   {{end}}{{end}}
`)

var MarkdownDocTemplate = `{{if gt .SectionNum 0}}% {{ .App.Name }} {{ .SectionNum }}

{{end}}# 名称

{{ .App.Name }}{{ if .App.Usage }} - {{ .App.Usage }}{{ end }}

# 概要

{{ .App.Name }}
{{ if .SynopsisArgs }}
` + "```" + `
{{ range $v := .SynopsisArgs }}{{ $v }}{{ end }}` + "```" + `
{{ end }}{{ if .App.Description }}
# 描述

{{ .App.Description }}
{{ end }}
**使用方法**:

` + "```" + `{{ if .App.UsageText }}
{{ .App.UsageText }}
{{ else }}
{{ .App.Name }} [GLOBAL OPTIONS] command [COMMAND OPTIONS] [ARGUMENTS...]
{{ end }}` + "```" + `
{{ if .GlobalArgs }}
# 全局选项
{{ range $v := .GlobalArgs }}
{{ $v }}{{ end }}
{{ end }}{{ if .Commands }}
# 命令
{{ range $v := .Commands }}
{{ $v }}{{ end }}{{ end }}`

var FishCompletionTemplate = `# {{ .App.Name }} fish shell completion

function __fish_{{ .App.Name }}_no_subcommand --description '测试是否有子命令'
    for i in (commandline -opc)
        if contains -- $i{{ range $v := .AllCommands }} {{ $v }}{{ end }}
            return 1
        end
    end
    return 0
end

{{ range $v := .Completions }}{{ $v }}
{{ end }}`
