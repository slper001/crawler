package engine

import (
	"github.com/crawler/danke/fetcher"
	"github.com/uber/jaeger-client-go/crossdock/log"
)

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkerCount int
	ItemChan chan interface{}
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}
func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i:=0; i<e.WorkerCount; i++{
		createWorker(e.Scheduler.WorkerChan(),out, e.Scheduler)
	}

	for _, r := range seeds{
		e.Scheduler.Submit(r)
	}

	for  {
		result := <-out
		for _, item := range result.Items{
			go func() {e.ItemChan <- item}()
		}

		for _,request := range result.Requests{
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadyNotifier)  {
	go func() {
		for  {
			ready.WorkerReady(in)
			request := <- in
			result, err := worker(request)
			if err != nil{
				continue
			}
			out <- result
		}
	}()
}
func worker(r Request)  (ParseResult, error){
	log.Printf("Fetching %s",r.Url)
	body, err := fetcher.Fetch(r.Url)
	if err != nil{
		log.Printf("Fetcher: error" + "fetching url %s: %v", r.Url, err)
		return ParseResult{},err
	}

	return r.ParserFunc(body), nil
}