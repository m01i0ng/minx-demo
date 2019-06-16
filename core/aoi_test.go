package core

import "testing"

func TestNewAoi(t *testing.T) {
    aoi := NewAoi(0, 250, 0, 250, 5, 5)
    t.Log(aoi)
}
