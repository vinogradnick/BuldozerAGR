package generator

import (
	"container/list"
	"fmt"
	"github.com/BuldozerAGR/cmd/worker"
	"github.com/BuldozerAGR/configs"
	"strconv"
	"strings"
	"sync"
	"time"
)

func durationConvertation(duration string) int64 {
	var str string
	var data int64
	var err error
	if strings.Contains(duration, "min") {
		str = duration[0:strings.Index(duration, "m")]
		data, err = strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}
		return data * 60000
	}

	if strings.Contains(duration, "sec") {
		str = duration[0:strings.Index(duration, "s")]
		data, err = strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}
		return data * 1000
	}
	if strings.Contains(duration, "ms") {
		str = duration[0:strings.Index(duration, "m")]
		data, err = strconv.ParseInt(str, 10, 32)
		if err != nil {
			panic(err)
		}
		return data
	}
	return 0
}

func StepAlgorithm(config *configs.StepConf) {
	sFactory := worker.NewConstWorkerFactory(sync.WaitGroup{})
	var i int32
	startWorkers := config.Start
	timerData := durationConvertation(config.Duration)
	tickerDuration := time.Duration(timerData)
	for i = 0; i < startWorkers; i++ {
		sFactory.createWorker()
	}

	ticker := time.NewTicker(tickerDuration * time.Millisecond)

	for range ticker.C {
		if sFactory.counter < config.End {
			for i = 0; i < config.Step && sFactory.counter < config.End; i++ {
				sFactory.createWorker()
			}
		} else {
			sFactory.wg.Wait()
			ticker.Stop()
			break
		}

	}
	arr := list.List{}
	for i = 0; i < config.End; i++ {
		select {
		case res := <-sFactory.metricChan:
			arr.PushBack(res)
			fmt.Println(res.ResponseTime)
		}
	}

}

