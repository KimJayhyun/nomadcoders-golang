## Go Routine
> See `check_url/check_url.go`

Go 언어의 'goroutine'(고루틴)에 대해 설명해드리겠습니다.

Goroutine은 Go 언어에서 제공하는 경량 스레드로, 동시성 프로그래밍을 위한 핵심 기능입니다. 일반적인 스레드와 비교했을 때 다음과 같은 특징이 있습니다:

1. 경량성
- 일반 스레드는 보통 1-2MB의 스택 크기를 가지지만, goroutine은 초기에 약 2KB로 시작합니다
- 필요에 따라 동적으로 스택 크기가 조절됩니다
- 시스템 자원을 적게 사용하므로 수천, 수만 개의 goroutine을 동시에 실행할 수 있습니다

2. 주요 특징
- 멀티코어 활용: Go 런타임이 자동으로 goroutine들을 여러 OS 스레드에 분배합니다
- 동시성 관리: 채널과 select 문을 통해 goroutine 간 통신과 동기화를 쉽게 할 수 있습니다
- 에러 처리: panic이 발생해도 다른 goroutine에는 영향을 주지 않습니다

Goroutine은 Go 언어의 "동시성이 쉬워야 한다"는 철학을 잘 보여주는 기능입니다. 일반적인 스레드보다 사용하기 쉽고, 리소스 효율성도 뛰어나며, 채널을 통한 통신 방식으로 동시성 프로그래밍에서 발생할 수 있는 많은 문제들을 예방할 수 있습니다.

---

## 예제 Code

1. `main` 함수는 `goroutine`이 끝날 때까지 대기하지 않음

```go
package main

import (
	"fmt"
	"time"
)

func main() {
    go sexyCount("nico")
    // 정상 동작
    sexyCount("jhkim")
    // 아래 코드를 사용할 경우, main 함수가 `goroutine`을 기달리지 않아
    // 바로 종료됨
    // go sexyCount("jhkim") 
}

func sexyCount(person string) {
	for i := 0; i < 10; i++ {
		fmt.Println(person, "is sexy", i)
		time.Sleep(time.Second)
	}
```

2. `channel`을 활용하여 `goroutine` 함수 결과 받기
    - return할 chanel의 수보다 많이 `<-`로 받을려고 하면 `dead lock` 에러 발생
    - channel에 데이터를 할당하는 것은 `blocking operation`
```go
package main

import (
	"fmt"
	"time"
)

func main() {
    c := make(chan bool)
	people := [2]string{"nico", "flynn"}
	for _, person := range people {
		go isSexy(person, c)
	}

    fmt.Println(<-c)
    fmt.Println(<-c)
}

func isSexy(person string, c chan bool) {
	time.Sleep(time.Second * 5)
	fmt.Println(person)
	c <- true
}
```