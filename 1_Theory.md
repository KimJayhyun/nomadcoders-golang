# 1. Theory
## 0. Main Package
- 일반적으로 패키지는 라이브러리로서 사용되지만, `"main"`이라고 명명된 패키지는 Go Compiler에 의해 특별하게 인식된다.
- `package main`인 경우, 컴파일러는 해당 패키지를 공유 라이브러리가 아닌 실행(executable) 프로그램으로 만든다.
- 그리고 이 `package main` 안의 `main()` 함수가 프로그램의 시작점, 즉 Entry Point가 된다.
- 패키지를 공유 라이브러리로 만들 때에는, `package main`나 `main()` 함수를 사용해서는 안 된다.

## 1. Packages and Imports
- 대문자로 시작하는 함수는 `Public`
- 소문자로 시작하는 함수는 `Private`
- ["something.go"로 확인해보기](/something/something.go)

## 2. Variables and Constants

```go
const name = "jhkim" // 자동으로 Type을 정함
const name string = "jhkim" // Type을 지정하는 방법

// 자동으로 Type을 정함
var name = "jhkim"

// 함수 내에서만 사용 가능한 문법
// 지정된 Type을 변경할 수 없음
name := "jhkim" 

// Type을 지정
var name string = "jhkim"
```

## 3. Functions part One
### Type 지정

함수의 Parameter와 Return 타입을 모두 지정해야 한다.

```go 
func multiply(a int, b int) int {
    return a * b
}
```

`a, b`의 Type이 같다면 아래와 같이 작성할 수 있다.

```go
func multiply(a, b int) int {
    return a * b
}
```

### 여러 값 반환

Golang에서는 함수가 여러 값을 반환할 수 있다.
Python과 같이 리턴값을 무시할 수 있다.

```go
func lenAndUpper(name string) (int, string) {
	return len(name), strings.ToUpper(name)
}

func main() {
    // totalLength, upperName := lenAndUpper("jhkim")
	// fmt.Println(totalLength, upperName)

    totalLength, _ := lenAndUpper("jhkim")
	fmt.Println(totalLength)

}
```

### 많은 Param 받기

점(`.`)을 3 개 사용하면 함수에서 Param을 여러 개 받을 수 있다.

```go
func repeatMe(words ...string) {
	fmt.Println(words)
}

func main() {
	repeatMe("jhkim", "nico", "dal")
}
```
```bash
$ go run main.go
[jhkim nico dal] # array Type
```

## 4. Functions part Two
### Naked return

리턴값 정의를 함수 선언에 하여 값만 업데이트하여 return할 수 있다.

```go
func lenAndUpper(name string) (length int, uppercase string) {
	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func main() {
	fmt.Println(lenAndUpper("jhkim"))
}
```

### Defer

함수가 종료된 후의 행동을 작성할 수 있다.

```go
func lenAndUpper(name string) (length int, uppercase string) {
	defer fmt.Println("I'm done")

	length = len(name)
	uppercase = strings.ToUpper(name)

	return
}

func main() {
	fmt.Println(lenAndUpper("jhkim"))
}
```

## 5. for, range, ...args

Golang에는 오직 for만 있다.

```go
func superAdd(numbers ...int) int {
	total := 0

	// for index := range numbers { // 하나의 변수만 선언하면 index만 준다.
	// for index, number := range numbers {
    // for i := 0; i < len(numbers); i++ {
	for _, number := range numbers {
		total += number
	}

	return total
}

func main() {
	result := superAdd(1, 2, 3, 4, 5, 6, 7)
	fmt.Println(result)
}
```

## 6. If with a Twist

```go
func canIDrink(age int) bool {
	if koreanAge := age + 2; koreanAge < 18 { // if 문에 변수를 선언할 수 있음
		return false
	} else {
		return true
	}
}
```

## 7. Switch

```go
func canIDrink(age int) bool {
	switch {
	case age < 18:
		return false
	case age == 18:
		return true
	case age > 50:
		return false
	}

	return false
}
```

## 8. Pointers!

```go
func main() {
	a := 2
	b := &a // b에 a의 주소를 할당
	fmt.Println(*b) // a의 값

    *b = 20
    fmt.Println(a) // a의 값 update
}
```

## 9. Arrays and Slices

Array는 길이를 선언해야하지만, Slices는 길이를 선언하지 않아도 된다.

```go
func main() {
	names := [5]string{"jhkim", "nico", "dal"}

	names[3] = "alala"
	names[4] = "alala"
	// names[5] = "alala" 0-based 
	fmt.Println(names)

	slice_names := []string{"jhkim", "nico", "dal"}
	slice_names = append(slice_names, "alala") // 기존 값을 수정하지 않고 새로운 값을 만듬

	fmt.Println(slice_names)
}
```

## 10. Maps

```go
func main() {
	nico := map[string]string{"name": "nico", "age": "12"}

	// for key, value := nico {
	for _, value := range nico {
		fmt.Println(value)
	}

	fmt.Println(nico)
}
```

## 11. Structs

```go
func main() {
	favFood := []string{"a", "b"}
	nico := person{name: "nico", age: 12, favFood: favFood}

	fmt.Println(nico.name)
}
```