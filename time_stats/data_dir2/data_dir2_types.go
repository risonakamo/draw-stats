// types for implementing data dir v2

package datadir_v2

// interface for metadata v2. the file is a list of data file entrys.
// each data file entry represents one thing that should appear in the
// list for the user to select.
type MetadataYamlV2 []DataFileInfo2

// information about a single data file
type DataFileInfo2 struct {
    // should include the file extension
    Filename string `yaml:"filename"`
    DisplayName string `yaml:"displayName"`

    // details for google sheets url. if either is empty, this feature
    // is disabled
    // the big main part of the sheet url. lets you access the sheet
    MainSheetid string `yaml:"mainSheetId"`
    // called the GID in the url. lets you choose one of the sub-sheets
    // that make up a main sheet
    SubSheetId string `yaml:"subSheetId"`
}