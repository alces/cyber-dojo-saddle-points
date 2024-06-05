package saddle_points

func Find(data Landscape) (result []Point) {
    hightest := pointSet{}
    lowest := pointSet{}
    
    for i := 0; i < tupleSize; i++ {
        rowMax := data.row(i).Max()
        
        for j := 0; j < tupleSize; j++ {
            if data[j][i] == rowMax {
                hightest.Add(point(j, i))
            }
        }

        colMin := data.column(i).Min()
        
        for j := 0; j < tupleSize; j++ {
            if data[i][j] == colMin {
                lowest.Add(point(i, j))
            }
        }
    }
    
    return hightest.Intersect(lowest).ToSlice()
}
