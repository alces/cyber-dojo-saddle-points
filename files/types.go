package saddle_points

type point struct {
    x, y int
}

func Point(x, y int) point {
    return point{
        x: x,
        y: y,
    }
}