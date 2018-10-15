package main

import (
	"fmt"
	"sort"
	"strconv"
	"sync"
)

func ExecutePipeline(jobs ...job) {
	wg := &sync.WaitGroup{}
	var pipeChan chan interface{}

	for _, currentJob := range jobs {
		wg.Add(1)
		outChan := make(chan interface{})

		go func(jobFunc job, wg *sync.WaitGroup, in, out chan interface{}) {
			defer wg.Done()
			defer close(out)
			jobFunc(in, out)
		}(currentJob, wg, pipeChan, outChan)

		pipeChan = outChan
	}

	wg.Wait()
}

func SingleHash(in, out chan interface{}) {
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	dataIndex := 0
	hashResults := make(map[int]string)

	for value := range in {
		wg.Add(2)
		data := strconv.Itoa(value.(int))
		newData := DataSignerMd5(data)

		go func(hashResults map[int]string, inputData string, dataIndex int, wg *sync.WaitGroup, mutex *sync.Mutex) {
			defer wg.Done()
			crc := DataSignerCrc32(inputData)
			mutex.Lock()
			hashResults[dataIndex*2] = crc
			mutex.Unlock()
		}(hashResults, data, dataIndex, wg, mutex)

		go func(hashResults map[int]string, inputData string, dataIndex int, wg *sync.WaitGroup, mutex *sync.Mutex) {
			defer wg.Done()
			crc := DataSignerCrc32(inputData)
			mutex.Lock()
			hashResults[dataIndex*2+1] = crc
			mutex.Unlock()
		}(hashResults, newData, dataIndex, wg, mutex)

		dataIndex++
	}
	wg.Wait()

	for i := 0; i < len(hashResults); i += 2 {
		out <- (hashResults[i] + "~" + hashResults[i+1])
	}
}

func MultiHash(in, out chan interface{}) {
	const hashCount = 6
	wg := &sync.WaitGroup{}
	mutex := &sync.Mutex{}
	hashResults := make(map[int]string)
	dataIndex := 0

	for value := range in {
		wg.Add(hashCount)
		for i := 0; i < hashCount; i++ {
			go func(hashResults map[int]string, data string, dataIndex int, i int, wg *sync.WaitGroup, mutex *sync.Mutex) {
				defer wg.Done()
				data = DataSignerCrc32(strconv.Itoa(i) + data)
				mutex.Lock()
				hashResults[dataIndex*hashCount+i] = data
				mutex.Unlock()
			}(hashResults, value.(string), dataIndex, i, wg, mutex)
		}
		dataIndex++
	}
	wg.Wait()

	for i := 0; i < dataIndex; i++ {
		var result string
		for j := 0; j < hashCount; j++ {
			result += hashResults[i*hashCount+j]
		}
		out <- result
	}
}

func CombineResults(in, out chan interface{}) {
	resultContainer := make([]string, 0)

	for data := range in {
		resultContainer = append(resultContainer, data.(string))
	}
	sort.Strings(resultContainer)

	var result = resultContainer[0]
	for i := 1; i < len(resultContainer); i++ {
		result += (fmt.Sprintf("_%s", resultContainer[i]))
	}
	out <- result
}
