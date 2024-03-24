// types for implementing data dir v2

package datadir_v2

// interface for metadata v2. the file is a list of data file entrys.
// each data file entry represents one thing that should appear in the
// list for the user to select.
type MetadataYamlV2 []DataFileInfo2

// information about a single data file
type DataFileInfo2 struct {
    // should include the file extension
    Filename string `yaml:"filename" json:"filename"`
    DisplayName string `yaml:"displayName" json:"displayName"`

    // full url to google sheets page for this data file
    SheetsUrl string `yaml:"sheetUrl" json:"sheetUrl"`
}

// info extracted from a google sheets url
type SheetsUrlInfo struct {
    // the big main part of the sheet url. lets you access the sheet
    MainSheetid string

    // called the GID in the url. lets you choose one of the sub-sheets
    // that make up a main sheet
    SubSheetId string
}