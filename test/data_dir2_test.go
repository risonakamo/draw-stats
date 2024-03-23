package test

import (
	"testing"
	datadir_v2 "time-stats/time_stats/data_dir2"

	"github.com/k0kubun/pp/v3"
)

const METADATA_FILEPATH string="../data/metadata_v2.yml"

// test reading the v2 metadata file
func Test_readMetadataV2(t *testing.T) {
    var res datadir_v2.MetadataYamlV2=datadir_v2.ReadMetadataFileV2(METADATA_FILEPATH)

    pp.Print(res)
}

// test reading the metadata file, then retrieving one of the items in the data file
func Test_fetchDataFile(t *testing.T) {
    var datafiles datadir_v2.MetadataYamlV2=datadir_v2.ReadMetadataFileV2(METADATA_FILEPATH)

    datadir_v2.FetchDataFile(
        datafiles[2],
        "../data",
    )
}