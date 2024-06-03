package saddle_points

type tuple [5]int

type landscape [5]tuple

type point struct {
    x, y int
}

func Landscape(data [25]int) (result landscape) {    
    for i := 0; i < len(data); i++ {
        c := i % len(tuple)
        r := i / len(tuple)
        
        result[c][r] = data[i]
    }
        
    return
}

func (l landscape) Column(num int) tuple {
    return l[num]
}

func (l landscape) Row(num int) (result tuple) {
    for i := 0; i < len(tuple); i++ {
        result[i] = l[i][num]
    }
    
    return
}

func Point(x, y int) point {
    return point{
        x: x,
        y: y,
    }
}