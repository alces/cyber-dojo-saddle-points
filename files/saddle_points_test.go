package saddle_points

import (
    "fmt"
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
    
    fmt.Println(Find(hilly))
    assert.Len(t, Find(hilly), 1)
}
