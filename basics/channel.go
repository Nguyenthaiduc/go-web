package main

import "log"

var number int

func CaculateNumber(inChan chan int) {
	number = 10
	numberRandom := RandomNumber(number)

	inChan <- numberRandom

}

func Channel() {
	inChan := make(chan int)
	defer close(inChan)

	go CaculateNumber(inChan)

	num := <-inChan

	log.Println(num)
}
