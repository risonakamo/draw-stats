// functions helping with api types

package time_stats

// convert list of tag filters to filter dict
func TagFiltersListToDict(filters []TagFilter) FilterDict {
    var dict FilterDict=make(FilterDict)

    for i := range filters {
        dict[TagType(filters[i].Tag)]=TagValue(filters[i].Value)
    }

    return dict
}