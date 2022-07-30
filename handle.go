package main

import (
	"context"
	"sync"
	"time"

	write_proc "github.com/xinzezhu/barrage_protocol/write_proc" // 引入编译生成的包
)

type handleWriteProc struct{}

var HandleWriteProc handleWriteProc
var BarrageSecMap sync.Map

func (h handleWriteProc) SendBarrage(ctx context.Context, req *write_proc.Request) (*write_proc.Response, error) {
	room_id := req.Header.RoomId
	content := req.Body.Content
	now_time := time.Now().Unix()
	var roomIdMap map[uint32][]string
	if v, ok := BarrageSecMap.Load(now_time); ok {
		roomIdMap = v.(map[uint32][]string)
	} else {
		roomIdMap = make(map[uint32][]string)
	}

	roomIdMap[room_id] = append(roomIdMap[room_id], content)

	BarrageSecMap.Store(now_time, roomIdMap)

	res := new(write_proc.Response)
	res.Header = new(write_proc.Header)
	res.Header.RoomId = room_id
	res.Header.RetCode = 0
	res.Header.RetMsg = "success"
	return res, nil
}
