### 윈도우에서 VS Code로 Go 실행 요약

1. **Go 설치 확인**:  
   - Go 설치 후 `C:\Program Files\go` (GOROOT) 폴더가 생성되었는지 확인.  
   - `cmd`에서 `go env` 실행해 환경변수 확인.

2. **VS Code 설정**:  
   - VS Code 마켓플레이스에서 Go 익스텐션 설치.

3. **GOPATH 설정**:  
   - 기본 경로는 `C:\Users\{사용자이름}\go`.  
   - 아래 폴더 구조 생성:  
     ```
     C:\Users\{사용자이름}\go\bin  
     C:\Users\{사용자이름}\go\pkg  
     C:\Users\{사용자이름}\go\src  
     ```

4. **프로젝트 생성**:  
    > `learngo` 프로젝트라고 가정
   - 프로젝트 폴더 예:  
     ```
     C:\Users\{사용자이름}\go\src\github.com\{GitHub 사용자명}\learngo
     ```

5. **코드 작성**:  
   - `learngo` 폴더에 `main.go` 파일 생성 후 아래 코드 작성:  
     ```go
     package main

     import "fmt"

     func main() {
         fmt.Println("Hello world!")
     }
     ```

6. **필요한 패키지 설치**:  
   - VS Code 하단 메시지에서 "Install All" 버튼 클릭.

7. **Go 모듈 초기화**:  
   - `go.mod file not found` 오류 발생 시, VS Code 터미널에서 `go mod init learngo` 실행.

8. **코드 실행**:  
   - F5 키로 실행.  

**참고**: GOROOT는 Go가 설치된 경로(`C:\Program Files\go`), GOPATH는 작업 공간 경로(`C:\Users\{사용자이름}\go`).