package main

import (
	"fmt"
)

// 인사말 문자열을 반환하는 인터페이스
type Greeting interface {
	Greeting() string
}

// Greeting인터페이스와 fmt.Stringer 인터페이스를 임베딩해서 새로운 인터페이스 생성
type SelfIntroducing interface {
	Greeting
	fmt.Stringer
}

// 이름을 갖는 Human 구조체
type Human struct {
	name string
}

// Human 타입과 Greeting 메서드 연결
func (r Human) Greeting() string {
	return fmt.Sprintln("Hello! I'm", r.name)
}

// Human 타입과 String() 메서드 연결
func (r Human) String() string {
	return r.Greeting()
}

func main() {
	mike := Human{name: "Mike"}
	greeting(mike)
}

// SelfIntroducing를 받아서 인사말을 출력한다.
func greeting(introduce SelfIntroducing) {
	fmt.Println(introduce)
}
