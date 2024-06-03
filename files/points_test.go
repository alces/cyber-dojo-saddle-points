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