// functions dealing with the data dir

package time_stats

import (
	"os"
	"path/filepath"
	"strings"

	"gopkg.in/yaml.v3"
)

// scan the data dir, which should have a metadata yml, and try to find all
// datas.
func GetDataList(datadir string) []DataFileInfo {
    var files []string
    var e error
    files,e=filepath.Glob(filepath.Join(datadir,"*.tsv"))

    if e!=nil {
        panic(e)
    }

    var metadata MetadataFile=readMetadataFile(filepath.Join(datadir,"metadata.yml"))

    var fileInfos []DataFileInfo

    for i := range files {
        var baseName string=filepath.Base(files[i])
        var noExtensionName string=strings.TrimSuffix(baseName,filepath.Ext(baseName))
        var displayName string=noExtensionName

        var foundDisplayName string
        var found bool
        foundDisplayName,found=metadata[noExtensionName]

        if found {
            displayName=foundDisplayName
        }

        fileInfos=append(fileInfos,DataFileInfo {
            Filename: filepath.Base(baseName),
            DisplayName: displayName,
        })
    }

    return fileInfos
}

// try to parse metdata file
func readMetadataFile(metadataFile string) MetadataFile {
    var data []byte
    var e error
    data,e=os.ReadFile(metadataFile)

    if e!=nil {
        panic(e)
    }

    var parsedData MetadataFile=make(MetadataFile)
    yaml.Unmarshal(data,&parsedData)

    return parsedData
}