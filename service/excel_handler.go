package service

import (
	"github.com/tealeg/xlsx"
)

type ExcelDataSheet struct {
	Rows  [][]string `json:"rows"`
	Sheet string     `json:"sheet"`
}

type ExcelDataArgs struct {
	OutFilePath string           `json:"out_file_path"`
	Sheets      []ExcelDataSheet `json:"sheets"`
}

func GenExcelFile(req *ExcelDataArgs) error {
	file := xlsx.NewFile()
	for i := 0; i < len(req.Sheets); i++ {
		sheet, err := file.AddSheet(req.Sheets[i].Sheet)
		if err != nil {
			return err
		}
		for j := 0; j < len(req.Sheets[i].Rows); j++ {
			row := sheet.AddRow()
			for k := 0; k < len(req.Sheets[i].Rows[j]); k++ {
				row.AddCell().Value = req.Sheets[i].Rows[j][k]
			}
		}
	}
	return file.Save(req.OutFilePath)
}
