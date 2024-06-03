package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestFind(t *testing.T) {
    assert.Size(t, 1, Find([25]int{}))
}
