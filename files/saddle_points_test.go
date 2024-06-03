package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
    assert.Len(t, Find([25]int{}), 1)
}
