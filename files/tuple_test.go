package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

func TestTupleMax(t *testing.T) {
    assert.Equal(t, 5, tuple{1,2,5,3,2}.max())
}

func TestTupleMin(t *testing.T) {
    assert.Equal(t, 2, tuple{10,2,5,3,2}.min())
}