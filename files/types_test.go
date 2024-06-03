package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestLandscape(t *testing.T) {
    l := Landscape([25]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25})
    
    assert.Equal(t, 1, l[0][0])
}


func TestPoint(t *testing.T) {
    p := Point(1, 2)
    
    assert.Equal(t, 1, p.x, "unexpected X")
    assert.Equal(t, 2, p.y, "unexpected Y")
}