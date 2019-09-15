package main

import (
	"github.com/crawler/danke/parser"
	"github.com/crawler/engine"
	"github.com/crawler/persist"
	"github.com/crawler/scheduler"
)

func main() {
	e:=engine.ConcurrentEngine{
		Scheduler:   &scheduler.QueuedScheduler{},
		WorkerCount: 10,
		ItemChan:    persist.ItemSaverSql(),
	}
	e.Run(engine.Request{
		Url:        "https://www.danke.com/room/sz/d.html",
		ParserFunc: parser.ParseAreaList,
	})
}
