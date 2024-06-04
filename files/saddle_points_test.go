package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
    flat := Landscape([25]int{})
    
    assert.Len(t, Find(flat), 25)
    
    hilly := Landscape([25]int{
        10, 10, 50, 80, 10,
        10, 30, 60, 70, 10,
        10, 20, 40, 90, 10,
        10, 30, 60, 70, 10,
        10, 10, 50, 80, 10,
    })
    
    r := Find(hilly)
    
    assert.Len(t, r, 2)
    assert.Contains(t, r, Point(3, 1))
    assert.Contains(t, r, Point(3, 3))
}
