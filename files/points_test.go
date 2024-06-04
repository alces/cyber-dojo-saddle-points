package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
    p := Point(1, 2)
    
    assert.Equal(t, 1, p.x, "unexpected X")
    assert.Equal(t, 2, p.y, "unexpected Y")
}

func TestPointSetAdd(t *testing.T) {
    s := pointSet{}
    
    assert.Len(t, s, 0)
    
    s.Add(Point(1, 2))

    assert.Len(t, s, 1)
}

func TestPointSetIn(t *testing.T) {
    s := pointSet{}
    
    s.Add(Point(1,1))
    
    assert.True(t, s.In(Point(1,1)), "should countain Point(1,1)")
    assert.False(t, s.In(Point(1,2)), "should not countain Point(1,2)")
}
    
func TestPointSetIntersect(t *testing.T) {
    s1 := pointSet{}
    s2 := pointSet{}
    
    s1.Add(Point(1,1))
    s1.Add(Point(1,5))
    s1.Add(Point(2,2))
    s1.Add(Point(3,3))
    
    s2.Add(Point(1,1))
    s2.Add(Point(2,2))
    s2.Add(Point(3,1))
    
    r := s1.Intersect(s2)
    
    assert.Len(t, r, 2)
    assert.Contains(t, r, Point(1,1))
    assert.Contains(t, r, Point(2,2))
}