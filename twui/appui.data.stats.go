package twui

import "html/template"

type AppUIDataStats struct {
	Wrap  bool
	Items []AppUIDataStatsItem
}

type AppUIDataStatsItem struct {
	Title        string
	Image        template.HTML
	ValueSymbol  string
	Value        float64
	ChangeSymbol string
	Change       float64
}
