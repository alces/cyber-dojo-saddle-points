package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
    flat := Landscape([25]int{})
    
    assert.Len(t, Find(flat), 25)
    
    hilly := Landscape([25]int{
        10, 10, 50, 10, 10,
        10, 30, 60, 30, 10,
        10, 20, 40, 20, 10,
        10, 30, 60, 30, 10,
        10, 10, 50, 10, 10,
    })
    
    assert.Len(t, Find(hilly), 1)
}
