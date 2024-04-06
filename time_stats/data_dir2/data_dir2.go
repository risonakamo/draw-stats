// V2 of functions for dealing with the data dir.
// provides functions for retrieving list of data, and auto-retrieving data

package datadir_v2

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path/filepath"

	"github.com/sirupsen/logrus"
	"gopkg.in/yaml.v3"
)

// read target metadata file v2 into obj
func ReadMetadataFileV2(filepath string) MetadataYamlV2 {
    var data []byte
    var e error
    data,e=os.ReadFile(filepath)

    if e!=nil {
        logrus.Warn("failed to read metadata file")
        logrus.Warn("the error:",e)
        return MetadataYamlV2{}
    }

    var parsedData MetadataYamlV2
    yaml.Unmarshal(data,&parsedData)

    return parsedData
}

// given a target data file info and an output location, attempt to auto fetch
// the data file and place it into the output dir.
func FetchDataFile(datainfo DataFileInfo2,outputDir string) error {
    if len(datainfo.SheetsUrl)==0 {
        fmt.Println("did not fetch datafile from url, no url")
        return nil
    }

    var extractedSheetInfo SheetsUrlInfo=extractSheetsInfo(datainfo.SheetsUrl)

    var sheetUrl string=fmt.Sprintf(
        "https://docs.google.com/spreadsheets/d/%s/export?format=tsv&gid=%s",
        extractedSheetInfo.MainSheetid,
        extractedSheetInfo.SubSheetId,
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

// try to update the target filename, using the metadatafile.
// set skipIfExists to skip if it exists, useful for trying to initialise the
// file instead of updating it
func TryUpdateDatafile(
    metadataFile string,
    filename string,
    dataDir string,
    skipIfExists bool,
) error {
    var e error

    if skipIfExists {
        _,e=os.Stat(filepath.Join(dataDir,filename))

        if !os.IsNotExist(e) {
            return nil
        }
    }

    // read metadata file for urls
    var datafiles MetadataYamlV2=ReadMetadataFileV2(metadataFile)

    // given the filename, find the info for the target file
    var datafile DataFileInfo2
    datafile,e=FindDataFile(filename,datafiles)

    if e!=nil {
        panic(e)
    }

    // try to fetch the file
    e=FetchDataFile(datafile,dataDir)

    if e!=nil {
        panic(e)
    }

    return nil
}