# server goal
- server is to be stateless. there will be a data dir where all raw data will be sitting
- frontend should be able to select which data to display, and display multiple at a time
    - it won't need to get all data at the same time, but maybe can be optimised to get multiple at the same time?
        - maybe not, since this makes the api look weird
- while viewing a data, frontend can choose to "enter" a level of the data, usually by specifying a tag/tag value combo. thus, the frontend should be able to get data by providing a list of these filters

# apis
- get all data names
    - list of all data available so frontend can choose to view each one
- get data (single)
    - should accept filter list
    - gets the target data with filter applied
    - returns the stats for that data and all breakdowns for the data