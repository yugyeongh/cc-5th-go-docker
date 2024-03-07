package main

import (
	"fmt"
	"sync"
)

func main() {
	// sync 패키지의 WaitGroup 구조체를 생성하는 코드
	wg := sync.WaitGroup{} // sync.WaitGroup 구조체의 인스턴스를 생성하고 이를 wg에 할당, {}를 통해 생성자 함수 호출하고 구조체의 기본값으로 초기화
	wg.Add(1)              // 대기해야 할 고루틴의 수를 추가
	go func() {
		defer wg.Done() // 각각의 고루틴이 작업을 완료할 때마다 done 메서드 호출
		fmt.Println("Let's Go")
	}()
	wg.Wait() // 모든 고루틴이 종료될 때까지 대기
}
