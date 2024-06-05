package saddle_points

import (
    "testing"
    "github.com/stretchr/testify/assert"
)

var testLandscapeData = [25]int{1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25}

func TestLandscape(t *testing.T) {
    l := MakeLandscape(testLandscapeData)
    
    assert.Equal(t, 1, l[0][0])
    assert.Equal(t, 25, l[4][4])
    assert.Equal(t, 6, l[0][1])
}

func TestLandscapeColumn(t *testing.T) {
    l := MakeLandscape(testLandscapeData)
    r := l.column(1)
    
    assert.Equal(t, 22, r[4])
}

func TestLandscapeRow(t *testing.T) {
    l := MakeLandscape(testLandscapeData)
    r := l.row(2)
    
    assert.Equal(t, 12, r[1])
}

func TestLandscapeRowHighestPoints(t *testing.T) {
    l := MakeLandscape(testLandscapeData)
    rh := l.rowHighestPoints()
    
    assert.Len(t, rh, 5)
    assert.Contains(t, rh, point(4, 0))
}
