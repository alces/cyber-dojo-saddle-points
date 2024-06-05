package saddle_points

func Find(data Landscape) (result []Point) {
    hightest := data.rowHighestPoints()
    lowest := data.columnLowestPoints()
        
    return highest.intersect(lowest).toSlice()
}
