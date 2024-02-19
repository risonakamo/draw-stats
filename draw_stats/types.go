package draw_stats

import "time"

// map of category stats. keyed by the category name
type CategoryStats map[string]CategoryStat

// single draw row event
type DrawRow struct {
	Item string
    Category string

	Start time.Time
    End time.Time

    Duration time.Duration
}

type DrawStats struct {
    Rows []DrawRow

    TotalTime time.Duration

    // stats per day. keyed by the date
    DayStats map[time.Time]DayStat
    // average time over all days
    AverageTimePerDay time.Duration

    // stats per single item. keyed by item name
    ItemStats map[string]ItemStat

    // stats per category. keyed by category name
    CategoryStats CategoryStats
}

type DayStat struct {
    Day time.Time
    Time time.Duration
}

type ItemStat struct {
    Item string
    TotalTime time.Duration

    // breakdown of categories for the item
    CategoryStats CategoryStats
}

type CategoryStat struct {
    Category string
    TotalTime string
}