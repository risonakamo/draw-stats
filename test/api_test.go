// tests of api related functions

package test

import (
	"testing"
	"time-stats/time_stats"

	"github.com/k0kubun/pp/v3"
)

// test getting data list from data dir
func Test_datalist(t *testing.T) {
    var datalist []time_stats.DataFileInfo=time_stats.GetDataList("../data")

    pp.Print(datalist)
}