package saddle_points

type tuple [tupleSize]int

func (t tuple) max() (result int) {
    result = t[0]
    
    for i := 1; i < tupleSize; i++ {
        if t[i] > result {
            result = t[i]
        }
    }
    
    return
}

func (t tuple) min() (result int) {
    result = t[0]
    
    for i := 1; i < tupleSize; i++ {
        if t[i] < result {
            result = t[i]
        }
    }
    
    return
}