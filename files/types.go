package saddle_points

type landscape [5][5]int

type point struct {
    x, y int
}

func Landscape(data [25]int) landscape {
    return landscape{}
}

func Point(x, y int) point {
    return point{
        x: x,
        y: y,
    }
}