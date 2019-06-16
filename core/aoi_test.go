package core

import "testing"

func TestNewAoi(t *testing.T) {
    aoi := NewAoi(0, 250, 0, 250, 5, 5)
    t.Log(aoi)
}

func TestAoi_GetSurroundGrids(t *testing.T) {
    aoi := NewAoi(0, 250, 0, 250, 5, 5)

    for gridId := range aoi.grids {
        grids := aoi.GetSurroundGrids(gridId)
        t.Logf("grid: %d, len: %d\n", gridId, len(grids))

        gridIds := make([]int, 0, len(grids))
        for _, grid := range grids {
            gridIds = append(gridIds, grid.Id)
        }
        t.Logf("surroundGrids are: %v", gridIds)
    }
}
