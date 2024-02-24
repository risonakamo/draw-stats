// types that have to do more with api instead of the time stats library

package time_stats

// typing for the metadata yml file. metadata file can be used to add additional information
// to the tsv data files in the data dir
// key: filename of a tsv data file, without the file extension
// val: display name for that particular data file
// todo: this might need to be upgraded to map[string]object if want more fields than just the
// display name
type MetadataFile map[string]string

// information about an available data file
// display name comes from metadata. if no metadata for the particular filename,
// then the filename is the displayname
type DataFileInfo struct {
    Filename string
    DisplayName string
}