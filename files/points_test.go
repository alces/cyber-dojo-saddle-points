package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestPoint(t *testing.T) {
    p := point(1, 2)
    
    assert.Equal(t, 1, p.X, "unexpected X")
    assert.Equal(t, 2, p.Y, "unexpected Y")
}

func TestPointString(t *testing.T) {
    assert.Equal(t, "(x:1, y:2)", point(1, 2).String())
}

func TestPointSetAdd(t *testing.T) {
    s := pointSet{}
    
    assert.Len(t, s, 0)
    
    s.add(point(1, 2))

    assert.Len(t, s, 1)
}

func TestPointSetIn(t *testing.T) {
    s := pointSet{}
    
    s.add(point(1,1))
    
    assert.True(t, s.In(point(1,1)), "should countain Point(1,1)")
    assert.False(t, s.In(point(1,2)), "should not countain Point(1,2)")
}
    
func TestPointSetIntersect(t *testing.T) {
    s1 := pointSet{}
    s2 := pointSet{}
    
    s1.add(point(1,1))
    s1.add(point(1,5))
    s1.add(point(2,2))
    s1.add(point(3,3))
    
    s2.add(point(1,1))
    s2.add(point(2,2))
    s2.add(point(3,1))
    
    r := s1.intersect(s2)
    
    assert.Len(t, r, 2)
    assert.Contains(t, r, point(1,1))
    assert.Contains(t, r, point(2,2))
}

func TestPointSetToSlice(t *testing.T) {
    s := pointSet{}
    
    s.Add(point(1,1))
    s.Add(point(3,3))
    
    r := s.toSlice()
    
    assert.Len(t, r, 2)
    assert.Contains(t, r, point(1,1))
    assert.Contains(t, r, point(3,3))
}    