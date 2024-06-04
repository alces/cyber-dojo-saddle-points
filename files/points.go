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

func (s pointSet) In(p point) bool {
    _, ok := s[p]
    
    return ok
}

func (s pointSet) Intersect(f pointSet) pointSet {
    result := pointSet{}
    
    for k := range s {
        if f.In(k) {
            result.Add(k)
        }
    }
    
    return result
}