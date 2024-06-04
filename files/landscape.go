package saddle_points

const (
    tupleSize = 5
)

type landscape [tupleSize]tuple

func Landscape(data [25]int) (result landscape) {    
    for i := 0; i < tupleSize; i++ {
        c := i % rowLen
        r := i / rowLen
        
        result[c][r] = data[i]
    }
        
    return
}

func (l landscape) Column(num int) tuple {
    return l[num]
}

func (l landscape) Row(num int) (result tuple) {
    for i := 0; i < tupleSize; i++ {
        result[i] = l[i][num]
    }
    
    return
}
