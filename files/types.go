package saddle_points

type landscape [5][5]int

type point struct {
    x, y int
}

func Landscape(data [25]int) (result landscape) {    
    for i := 0; i < len(data); i++ {
        c := i % 5
        r := i / 5
        
        result[c][r] = data[i]
    }
        
    return
}

func Point(x, y int) point {
    return point{
        x: x,
        y: y,
    }
}