package saddle_points

import (
    "fmt"
)

type Point struct {
    X, Y int
}

func point(x, y int) Point {
    return Point{
        X: x,
        Y: y,
    }
}

func (p Point) String() string {
    return fmt.Sprintf("(x:%d, y:%d)", p.X, p.Y)
}

type pointSet map[Point]struct{}

func (s pointSet) add(p Point) {
    var value struct{}
    
    s[p] = value
}

func (s pointSet) in(p Point) bool {
    _, ok := s[p]
    
    return ok
}

func (s pointSet) intersect(f pointSet) pointSet {
    result := pointSet{}
    
    for k := range s {
        if f.in(k) {
            result.add(k)
        }
    }
    
    return result
}

func (s pointSet) toSlice() []Point {
    result := make([]Point, 0, len(s))
    
    for k := range s {
        result = append(result, k)
    }
    
    return result
}