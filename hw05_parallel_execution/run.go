package hw05parallelexecution

import (
	"errors"
	"sync"
)

var ErrErrorsLimitExceeded = errors.New("errors limit exceeded")

type Task func() error

// Run starts tasks in n goroutines and stops its work when receiving m errors from tasks.
func Run(tasks []Task, countOfChannels, maxErrorCount int) error {
	if maxErrorCount <= 0 {
		return ErrErrorsLimitExceeded
	}

	wGroup := sync.WaitGroup{}
	wGroup.Add(len(tasks))

	channels := createChannels(countOfChannels)
	countOfErrors := 0
	startGoroutines(channels, &wGroup, &countOfErrors, maxErrorCount)
	pushTasksToChannels(tasks, channels)

	wGroup.Wait()

	if countOfErrors >= maxErrorCount {
		return ErrErrorsLimitExceeded
	}
	return nil
}

func createChannels(countOfChannels int) []chan Task {
	channels := make([]chan Task, countOfChannels)
	for i := 0; i < countOfChannels; i++ {
		channels[i] = make(chan Task)
	}
	return channels
}

func startGoroutines(channels []chan Task, wGroup *sync.WaitGroup, countOfErrors *int, maxErrorCount int) {
	mutex := &sync.Mutex{}
	for _, channel := range channels {
		go func(channel chan Task, maxErrorCount int) {
			for {
				task, ok := <-channel
				if !ok {
					break
				}
				taskExecute(task, wGroup, mutex, countOfErrors, maxErrorCount)
			}
		}(channel, maxErrorCount)
	}
}

func taskExecute(task Task, wGroup *sync.WaitGroup, mutex *sync.Mutex, countOfErrors *int, maxErrorCount int) {
	defer wGroup.Done()

	mutex.Lock()
	if *countOfErrors >= maxErrorCount {
		mutex.Unlock()
		return
	}
	mutex.Unlock()

	if err := task(); err != nil {
		mutex.Lock()
		*countOfErrors++
		mutex.Unlock()
		return
	}
}

func pushTasksToChannels(tasks []Task, channels []chan Task) {
	channelIndex := 0
	for _, task := range tasks {
		channel := channels[channelIndex]
		channel <- task

		channelIndex++
		if channelIndex == len(channels) {
			channelIndex = 0
		}
	}
	closeChannels(channels)
}

func closeChannels(channels []chan Task) {
	for i := 0; i < len(channels); i++ {
		close(channels[i])
	}
}
