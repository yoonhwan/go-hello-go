package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"

	"github.com/yoonhwan/go-awssample"
)

func say(s string) {
	for i := 0; i < 10; i++ {
		fmt.Println(s, "***", i)
	}
}

func main() {
	// 4개의 CPU 사용
	fmt.Println(runtime.NumCPU())
	runtime.GOMAXPROCS(8)

	// WaitGroup 생성. 2개의 Go루틴을 기다림.
	var wait sync.WaitGroup
	wait.Add(2)

	// 익명함수를 사용한 goroutine
	go func() {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println("Hello")
	}()

	// 익명함수에 파라미터 전달
	go func(msg string) {
		defer wait.Done() //끝나면 .Done() 호출
		fmt.Println(msg)
	}("Hi")

	wait.Wait() //Go루틴 모두 끝날 때까지 대기

	var sum func(n ...int) int
	sum = func(n ...int) int { //익명함수 정의
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5) //익명함수 호출
	println(result)

	s := []int{0, 1, 2, 3, 4, 5}
	s = s[2:5]     // 2, 3, 4
	s = s[1:]      // 3, 4
	fmt.Println(s) // 3, 4 출력

	// 함수를 동기적으로 실행
	say("Sync")

	// 함수를 비동기적으로 실행
	go say("Async1")
	go say("Async2")
	go say("Async3")

	// 3초 대기
	time.Sleep(time.Second * 3)

	// 정수형 채널을 생성한다
	ch := make(chan int, 10)

	go func() {
		ch <- 123 //채널에 123을 보낸다
	}()
	select {
		case v := <-ch:
			fmt.Println(v)
	}

	ch = make(chan int, 1)

	//수신자가 없더라도 보낼 수 있다.
	ch <- 101

	fmt.Println(<-ch)

	ch = make(chan int, 2)

	// 채널에 송신
	ch <- 1
	ch <- 2

	// 채널을 닫는다
	close(ch)

	// 채널 수신
	println(<-ch)
	println(<-ch)

	if _, success := <-ch; !success {
		println("더이상 데이타 없음.")
	}


	awssample.StartSample()
}
