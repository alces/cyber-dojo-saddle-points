package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
    flat := Landscape([25]int{})
    
    assert.Len(t, Find(flat), 25)
}
