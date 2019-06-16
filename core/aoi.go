package core

import "fmt"

//AOI 地图网格管理
//alias grid_manager

type Aoi struct {
    //区域坐标
    //左边界
    MinX int
    //右边界
    MaxX int
    //上边界
    MinY int
    //下边界
    MaxY int
    //X 方向网格数量
    SumX int
    //Y 方向网格数量
    SumY int
    //包含网格
    grids map[int]*Grid
}

func NewAoi(minX, maxX, minY, maxY, sumX, sumY int) *Aoi {
    aoi := &Aoi{MinX: minX, MaxX: maxX, MinY: minY, MaxY: maxY, SumX: sumX, SumY: sumY, grids: make(map[int]*Grid)}

    //初始化区域网格
    for y := 0; y < sumY; y++ {
        for x := 0; x < sumX; x++ {
            //计算网格编号
            gridId := y*sumX + x
            aoi.grids[gridId] = NewGrid(
                gridId,
                aoi.MinX+x*aoi.gridWidth(),
                aoi.MinX+(x+1)*aoi.gridWidth(),
                aoi.MinY+y*aoi.gridLength(),
                aoi.MinY+(y+1)*aoi.gridLength(),
            )
        }
    }

    return aoi
}

//每个网格 X 方向宽度
func (a *Aoi) gridWidth() int {
    return (a.MaxX - a.MinX) / a.SumX
}

//每个网格 Y 方向长度
func (a *Aoi) gridLength() int {
    return (a.MaxY - a.MinY) / a.SumY
}

func (a *Aoi) String() string {
    s := fmt.Sprintf("Aoi:\n minX:%d, maxX:%d, minY:%d, maxY:%d, sumX:%d, sumY:%d\nGrids:\n", a.MinX, a.MaxX, a.MinY, a.MaxY, a.SumX, a.SumY)

    for _, grid := range a.grids {
        s += " " + fmt.Sprintln(grid)
    }

    return s
}

//GetSurroundGrids 根据网格 ID 得到周边九宫格集合
func (a *Aoi) GetSurroundGrids(gridId int) []*Grid {
    //判断 gridId 是否在 aoi 中
    if _, ok := a.grids[gridId]; !ok {
        return nil
    }

    var grids []*Grid
    grids = append(grids, a.grids[gridId])

    //获取 X 方向所有网格
    idx := gridId % a.SumX
    if idx > 0 {
        grids = append(grids, a.grids[gridId-1])
    }

    if idx < a.SumX-1 {
        grids = append(grids, a.grids[gridId+1])
    }

    gridIdsX := make([]int, 0, len(grids))
    for _, grid := range grids {
        gridIdsX = append(gridIdsX, grid.Id)
    }

    for _, v := range gridIdsX {
        idy := v / a.SumY

        if idy > 0 {
            grids = append(grids, a.grids[v-a.SumX])
        }

        if idy < a.SumY-1 {
            grids = append(grids, a.grids[v+a.SumX])
        }
    }

    return grids
}

//GetGridId 通过坐标获取网格 ID
func (a *Aoi) GetGridId(x, y float32) int {
    idx := (int(x) - a.MinX) / a.gridWidth()
    idy := (int(y) - a.MinY) / a.gridLength()

    return idy*a.SumX + idx
}

//GetSurroundPlayerIds 通过坐标获取周围所有玩家 ID
func (a *Aoi) GetSurroundPlayerIds(x, y float32) []int {
    gridId := a.GetGridId(x, y)
    grids := a.GetSurroundGrids(gridId)

    var playerIds []int

    for _, grid := range grids {
        playerIds = append(playerIds, grid.GetPlayerIds()...)
    }

    return playerIds
}
