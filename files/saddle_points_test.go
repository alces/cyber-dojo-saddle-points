package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
    assert.Len(t, Find(Landscape([25]int{})), 0)
}
