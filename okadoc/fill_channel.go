package okadoc

import "sync"

func FillChannel(functions ...func() int) chan int {
	var wg sync.WaitGroup
	var intChan = make(chan int)
	var intFunc func() int
	for _, function := range functions {
		wg.Add(1)
		go func(function func() int) {
			defer wg.Done()
			intFunc = function
			intChan <- intFunc()
		}(function)
	}

	go func() {
		defer close(intChan)
		wg.Wait()
	}()

	return intChan
}
