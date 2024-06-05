package saddle_points

func Find(data Landscape) (result []Point) {
    highest := data.rowHighestPoints()
    lowest := data.columnLowestPoints()
        
    return highest.intersect(lowest).toSlice()
}
