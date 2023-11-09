package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func sameColumnColor(columnCount int, colorProperties []int) []tablewriter.Colors {
	columnColors := []tablewriter.Colors{}
	for i := 0; i < columnCount; i++ {
		columnColors = append(columnColors, colorProperties)
	}
	return columnColors
}

func printInTable(header []string) *tablewriter.Table {
	t := tablewriter.NewWriter(os.Stdout)
	t.SetHeader(header)
	colorProperties := []int{tablewriter.Bold, tablewriter.FgGreenColor}
	t.SetColumnColor(sameColumnColor(len(header), colorProperties)...)
	t.SetAutoWrapText(false)
	t.SetAlignment(tablewriter.ALIGN_CENTER)
	t.SetBorder(true)
	return t
}

func PrintInTable(header, data []string) {
	t := printInTable(header)
	t.Append(data)
	t.Render()
}

func PrintInTableBulk(header []string, data [][]string) {
	t := printInTable(header)
	t.AppendBulk(data)
	t.Render()
}
