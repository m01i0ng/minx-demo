package core

import (
    "fmt"
    "sync"

    "github.com/kataras/golog"
)

//AOI 地图网格

type Grid struct {
    Id int
    //左边界坐标
    MinX int
    //右边界坐标
    MaxX int
    //上边界坐标
    MinY int
    //下边界坐标
    MaxY int
    //玩家集合，id 为键
    players     map[int]bool
    playersLock sync.RWMutex
}

func NewGrid(id, minX, maxX, minY, maxY int) *Grid {
    return &Grid{
        Id:      id,
        MinX:    minX,
        MaxX:    maxX,
        MinY:    minY,
        MaxY:    maxY,
        players: make(map[int]bool),
    }
}

func (g *Grid) Add(playerId int) {
    g.playersLock.Lock()
    defer g.playersLock.Unlock()

    g.players[playerId] = true
    golog.Infof("player %d join", playerId)
}

func (g *Grid) Remove(playerId int) {
    g.playersLock.Lock()
    defer g.playersLock.Unlock()

    delete(g.players, playerId)
    golog.Infof("player %d left", playerId)
}

func (g *Grid) GetPlayers() []int {
    g.playersLock.RLock()
    defer g.playersLock.RUnlock()

    var playerIds []int
    for id := range g.players {
        playerIds = append(playerIds, id)
    }

    return playerIds
}

func (g *Grid) String() string {
    return fmt.Sprintf("Grid id:%d, minX:%d, maxX:%d, minY:%d, maxY:%d, players:%v", g.Id, g.MinX, g.MaxX, g.MinX, g.MaxY, g.players)
}
