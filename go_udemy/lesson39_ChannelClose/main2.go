package main

func doubleNumber(number int, resultChan chan<- int) {
	resultChan <- number * 2
}

//
