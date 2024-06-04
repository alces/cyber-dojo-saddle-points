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

type pointSet map[point]struct{}

func (s pointSet) Add(p point) {
    var value struct{}
    
    s[p] = value
}

func (s pointSet) Intersect(f pointSet) pointSet {
    result := pointSet{}
    
    for k := range s {
        if _, ok := f[k]; ok {
            result.Add(k)
        }
    }
    
    return result
}