package datadir_v2

import (
	"testing"

	"github.com/k0kubun/pp/v3"
)

func Test_extractSheetInfo(t *testing.T) {
    var res SheetsUrlInfo=extractSheetsInfo("https://docs.google.com/spreadsheets/d/1reD2OvNyl5Fkvs4LXNESuhRNTOQQtrtqU31njzGR-RY/edit#gid=1780809564")
    pp.Print(res)

    var res2 SheetsUrlInfo=extractSheetsInfo("https://www.google.com/search?q=sff+pc&source=lmns&bih=1034&biw=1798&prmd=sivnmbtz&hl=en&sa=X&ved=2ahUKEwiqi4O7-42FAxXaK2IAHdXWD14Q0pQJKAB6BAgBEAI")
    pp.Print(res2)
}