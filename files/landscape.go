package saddle_points

const (
    tupleSize = 5
)

type landscape [tupleSize]tuple

func Landscape(data [25]int) (result landscape) {    
    rowLen := len(result[0])
    
    for i := 0; i < len(data); i++ {
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
    for i := 0; i < len(result); i++ {
        result[i] = l[i][num]
    }
    
    return
}
