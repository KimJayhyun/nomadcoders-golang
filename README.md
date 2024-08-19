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