package resutil

import (
	"errors"
	"fmt"
	"github.com/tealeg/xlsx/v3"
)

func GetDetailsFromExcel(fp string, ressourceIdColName string, awsAccountColName string, awsRegionColName string, sheetName string, accountName string, accountRegion string, ress *Resources) {

	filename := fp
	sheet := sheetName
	wb, err := xlsx.OpenFile(filename)
	if err != nil {
		panic(err)
	}
	sh, ok := wb.Sheet[sheet]
	if !ok {
		panic(errors.New("Sheet not found"))
	}

	ridColCord := getColumnCordinate(sh, ressourceIdColName)
	awsAcctNameCord := getColumnCordinate(sh, awsAccountColName)
	awsRegion := getColumnCordinate(sh, awsRegionColName)
	resourceIds := getResourceIds(sh, ridColCord, awsAcctNameCord, awsRegion, accountName, accountRegion, sh.MaxRow)
	//ress := Resources{}
	for indx := range resourceIds {
		res := resource{}
		res.Resourcetype = sheetName
		res.ResourceID = resourceIds[indx]
		res.AwsRegion = accountRegion
		res.AwsAccountName = accountName

		ress.Resources = append(ress.Resources, res)

	}

}

func getColumnCordinate(s *xlsx.Sheet, colName string) int {

	resourceIdColCord := 0
	firstRow, err := s.Row(0)
	if err != nil {
		fmt.Println(err.Error())
	}

	for i := 0; i <= s.MaxCol; i++ {
		cellVal, err := firstRow.GetCell(i).FormattedValue()
		if err != nil {
			fmt.Println(err.Error())
		}
		if cellVal == colName {
			resourceIdColCord = i
			break
		}
	}

	return resourceIdColCord
}

func getResourceIds(s *xlsx.Sheet, acctIdColumnCord int, accountNameColumnCord int, awsRegionColumnCord int, accountName string, accountRegion string, totalRecords int) []string {
	var resourceIds []string
	for indx := 1; indx <= totalRecords; indx++ {
		rid := ""
		row, err := s.Row(indx)
		if err != nil {
			fmt.Println(err.Error())
		}

		cellAwsAccontName, _ := row.GetCell(accountNameColumnCord).FormattedValue()
		cellAwsRegion, _ := row.GetCell(awsRegionColumnCord).FormattedValue()

		if cellAwsAccontName == accountName && cellAwsRegion == accountRegion {
			rid, err = row.GetCell(acctIdColumnCord).FormattedValue()
			if err != nil {
				fmt.Println(err.Error())
			}
			resourceIds = append(resourceIds, rid)
		}

	}
	return resourceIds
}
