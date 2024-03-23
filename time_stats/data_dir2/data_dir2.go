// V2 of functions for dealing with the data dir.
// provides functions for retrieving list of data, and auto-retrieving data

package datadir_v2

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"gopkg.in/yaml.v3"
)

// read target metadata file v2 into obj
func ReadMetadataFileV2(filepath string) MetadataYamlV2 {
    var data []byte
    var e error
    data,e=os.ReadFile(filepath)

    if e!=nil {
        panic(e)
    }

    var parsedData MetadataYamlV2
    yaml.Unmarshal(data,&parsedData)

    return parsedData
}

// given a target data file info and an output location, attempt to auto fetch
// the data file and place it into the output dir.
func FetchDataFile(datainfo DataFileInfo2,outputDir string) error {
    if len(datainfo.MainSheetid)==0 || len(datainfo.SubSheetId)==0 {
        fmt.Println("did not fetch datafile from url, no url")
        return nil
    }

    // https://docs.google.com/spreadsheets/d/1reD2OvNyl5Fkvs4LXNESuhRNTOQQtrtqU31njzGR-RY/export?format=tsv&gid=1780809564
    var sheetUrl string=fmt.Sprintf(
        "https://docs.google.com/spreadsheets/d/%s/export?format=tsv&gid=%s",
        datainfo.MainSheetid,
        datainfo.SubSheetId,
    )

    var resp *http.Response
    var err error
    resp,err=http.Get(sheetUrl)

    if err!=nil {
        return err
    }

    defer resp.Body.Close()

    if resp.StatusCode!=200 {
        fmt.Println("response status code was not 200")
        fmt.Println("got this status instead:",resp.StatusCode)
        return fmt.Errorf("bad status code")
    }

    var outputFile string=filepath.Join(outputDir,datainfo.Filename)
    var wfile *os.File
    wfile,err=os.Create(outputFile)

    if err!=nil {
        return err
    }

    _,err=io.Copy(wfile,resp.Body)

    if err!=nil {
        return err
    }

    fmt.Println("saved file:",outputFile)
    return nil
}

// try to find a datafile by filename in a list of datafiles. return error if fails.
// if multiple datafiles with the same name exist, returns the first one
func FindDataFile(filename string,datafiles MetadataYamlV2) (DataFileInfo2,error) {
    for i := range datafiles {
        if datafiles[i].Filename==filename {
            return datafiles[i],nil
        }
    }

    return DataFileInfo2{},fmt.Errorf("failed to find datafile: %s",filename)
}