package datadir_v2

import (
	"fmt"
	"regexp"
)

// https://docs.google.com/spreadsheets/d/1reD2OvNyl5Fkvs4LXNESuhRNTOQQtrtqU31njzGR-RY/edit#gid=1780809564

// extract from a google sheets url the info
func extractSheetsInfo(url string) SheetsUrlInfo {
    reg:=regexp.MustCompile(`spreadsheets/d\/(.*?)\/edit#gid=(\d*)`)

    match:=reg.FindStringSubmatch(url)

    if len(match)!=3 {
        fmt.Println("failed to extract sheet url, wrong number of matches")
        fmt.Println("url:",url)
        fmt.Println("matches:",match)
        panic("bad match")
    }

    return SheetsUrlInfo {
        MainSheetid: match[1],
        SubSheetId: match[2],
    }
}