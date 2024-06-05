package saddle_points

const (
    tupleSize = 5
)

type Landscape [tupleSize]tuple

func MakeLandscape(data [25]int) (result Landscape) {    
    for i := 0; i < len(data); i++ {
        c := i % tupleSize
        r := i / tupleSize
        
        result[c][r] = data[i]
    }
        
    return
}

func (l Landscape) column(num int) tuple {
    return l[num]
}

func (l Landscape) row(num int) (result tuple) {
    for i := 0; i < tupleSize; i++ {
        result[i] = l[i][num]
    }
    
    return
}


func (l Landscape) rowHighestPoints() pointSet {
    result := pointSet{}
    
    for i := 0; i < tupleSize; i++ {
        rowMax := l.row(i).max()
        
        for j := 0; j < tupleSize; j++ {
            if l[j][i] == rowMax {
                result.add(point(j, i))
            }
        }
    }
    
    return result
}