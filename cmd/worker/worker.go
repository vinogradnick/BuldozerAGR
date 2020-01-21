package worker

import (
	"fmt"
	"github.com/bradhe/stopwatch"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
)

type Metrics struct {
	ResponseTime time.Duration
}

type WorkerFactory interface {
	createWorker()
}
type ConstWorkerFactory struct {
	wg         sync.WaitGroup
	metricChan chan Metrics
	counter    int32
}

func NewConstWorkerFactory(wg sync.WaitGroup) *ConstWorkerFactory {
	return &ConstWorkerFactory{wg: wg, metricChan: make(chan Metrics), counter: 0}
}

type LinearWorkerFactory struct {
	wg         sync.WaitGroup
	metricChan chan Metrics
	counter    uint32
}

func (c ConstWorkerFactory) createWorker() {
	c.wg.Add(1)
	go InitWorker(&c.wg, c.metricChan)
	c.counter++
}

func InitWorker(wg *sync.WaitGroup, metricChan chan Metrics) {
	metricChan <- httpRequest()
	wg.Done()

}
func httpRequest() Metrics {
	watch := stopwatch.Start()
	resp, err := http.Get("http://localhost:8000")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	watch.Stop()
	fmt.Println(body)
	fmt.Printf("Milliseconds elapsed: %v\n", watch.Milliseconds())
	return Metrics{ResponseTime: watch.Milliseconds()}
}
