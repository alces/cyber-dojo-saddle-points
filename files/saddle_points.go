package saddle_points

func Find(data landscape) (result []point) {
    hightest := pointSet{}
    lowest := pointSet{}
    
    for i := 0; i < tupleSize; i++ {
        rowMax := data.Row(i).Max()
        
        for j := 0; j < tupleSize; j++ {
            if data[j][i] == rowMax {
                hightest.Add(Point(j, i))
            }
        }

        colMin := data.Column(i).Min()
        
        for j := 0; j < tupleSize; j++ {
            if data[i][j] == colMin {
                lowest.Add(Point(i, j))
            }
        }
    }
    
    return hightest.Intersect(lowest).ToSlice()
}
