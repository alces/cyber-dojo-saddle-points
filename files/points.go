package saddle_points

type Point struct {
    X, Y int
}

func point(x, y int) Point {
    return Point{
        X: x,
        Y: y,
    }
}

type pointSet map[Point]struct{}

func (s pointSet) Add(p Point) {
    var value struct{}
    
    s[p] = value
}

func (s pointSet) In(p Point) bool {
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

func (s pointSet) ToSlice() []Point {
    result := make([]Point, 0, len(s))
    
    for k := range s {
        result = append(result, k)
    }
    
    return result
}