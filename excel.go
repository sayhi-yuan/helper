package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/xuri/excelize/v2"
)

// Excel相关的操作

const defaultSheetName = "Sheet1"

// ExcelReadGin ctx gin的上下文，name 表单中的文件名，sn sheet名
func ExcelReadGin(ctx *gin.Context, name string, sn ...string) (dataList [][][]string, err error) {
	fh, err := ctx.FormFile(name)
	if err != nil {
		return
	}

	f, err := fh.Open()
	if err != nil {
		return
	}
	defer f.Close()

	ef, err := excelize.OpenReader(f)
	if err != nil {
		return
	}

	sheetList := []string{}
	if len(sn) == 0 {
		sheetList = append(sheetList, defaultSheetName)
	} else {
		sheetList = append(sheetList, sn...)
	}

	for _, sheetName := range sheetList {
		rows, rErr := ef.GetRows(sheetName)
		if rErr != nil {
			err = rErr
			return
		}

		max := 0
		for _, row := range rows {
			if len(row) > max {
				max = len(row)
			}
		}

		for i, row := range rows {
			if len(row) < max {
				rows[i] = append(row, make([]string, max-len(row))...)
			}
		}

		dataList = append(dataList, rows)
	}

	return
}
