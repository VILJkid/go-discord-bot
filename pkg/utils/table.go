package utils

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

func PrintInTable(header, data []string) {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(header)
	table.SetAutoWrapText(false)
	table.SetAlignment(tablewriter.ALIGN_CENTER)
	table.SetBorder(true)
	table.Append(data)
	table.Render()
}
