package saddle_points

type tuple [5]int

func (t tuple) Max() (result int) {
    result = t[0]
    
    for i = 1; i < len(t); i++ {
        if t[i] > result {
            result = t[i]
        }
    }
    
    return
}