package config

import (
	"fmt"
	"strings"

	"github.com/jedib0t/go-pretty/v6/table"
	"github.com/jedib0t/go-pretty/v6/text"
)

type Renderer interface {
	RenderAsTable() string
}

func (cfg Config) Show() {
	fmt.Println(cfg.RenderAsTable())
}

func (cfg Config) RenderAsTable() string {
	return cfg.AppConfig.RenderAsTable() + "\n" +
		cfg.DBConfig.RenderAsTable()
}

func (cfg AppConfig) RenderAsTable() string {
	rows := []table.Row{
		{"Name", cfg.Name},
		{"Mode", cfg.Mode},
		{"Host", cfg.Host},
		{"Port", cfg.Port},
		{"Timezone", cfg.Timezone},
		{"Shutdown Timeout", cfg.ShutdownTimeout},
		{"Metrics Enabled", cfg.Metrics.Enabled},
	}
	if cfg.Metrics.Enabled {
		rows = append(rows,
			table.Row{"Metrics Port", cfg.Metrics.Port},
			table.Row{"Metrics Route", cfg.Metrics.Route})
	}
	rows = append(rows,
		table.Row{"Cache Life Window", cfg.Cache.LifeWindow},
		table.Row{"Cache Hard Max Size", cfg.Cache.HardMaxSize})
	title := AppCfgFile
	return renderTable(rows, title)
}

func (cfg DBConfig) RenderAsTable() string {
	rows := []table.Row{
		{"Host", cfg.Host},
		{"Port", cfg.Port},
		{"Database", cfg.Database},
		{"Pool Size", cfg.PoolSize},
	}
	title := DatabaseCfgFile
	return renderTable(rows, title)
}

func renderTable(rows []table.Row, title string) (tbl string) {
	tw := table.NewWriter()
	tw.AppendHeader(table.Row{"Key", "Value"})
	tw.AppendRows(rows)
	tw.SetIndexColumn(1)
	tw.SetTitle(title)
	stylePairs := [][]table.Style{
		{table.StyleRounded},
	}
	twOuter := table.NewWriter()
	for _, stylePair := range stylePairs {
		row := make(table.Row, 1)
		for idx, style := range stylePair {
			tw.SetStyle(style)
			tw.Style().Title.Align = text.AlignCenter
			row[idx] = tw.Render()
		}
		twOuter.AppendRow(row)
	}
	return tw.Render()
}

func renderTableMulti(header table.Row, rows []table.Row, title string, outerTitle string) (tbl string) {
	tw := table.NewWriter()
	tw.AppendHeader(header)
	tw.AppendRows(rows)
	tw.SetIndexColumn(1)
	tw.SetTitle(title)
	stylePairs := [][]table.Style{
		{table.StyleColoredYellowWhiteOnBlack},
	}
	twOuter := table.NewWriter()
	for _, stylePair := range stylePairs {
		row := make(table.Row, 1)
		for idx, style := range stylePair {
			tw.SetStyle(style)
			tw.Style().Title.Align = text.AlignCenter
			row[idx] = tw.Render()
		}
		twOuter.AppendRow(row)
	}
	twOuter.SetStyle(table.StyleLight)
	twOuter.Style().Title.Align = text.AlignCenter
	twOuter.SetTitle(outerTitle)
	twOuter.Style().Options.SeparateRows = true
	return twOuter.Render()
}

func outerTitle(outerTitle string) string {
	return strings.ToUpper(fileNameWithoutExtension(outerTitle) + " config")
}

func fileNameWithoutExtension(fileName string) string {
	if pos := strings.LastIndexByte(fileName, '.'); pos != -1 {
		return fileName[:pos]
	}
	return fileName
}
