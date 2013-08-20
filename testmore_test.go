package testmore

import (
	"fmt"
	"testing"
)

func TestOk(t *testing.T) {
	Ok(t, 1)
	Ok(t, 0 == 0, "message")
	Ok(t, 1+2 == 3, "message")
	Ok(t, "a", "message")
	Ok(t, "string")
	Ok(t, " ")
	Ok(t, []byte("string"))
	Ok(t, []rune("строка"))
	Ok(t, []rune("ライン"))
	Ok(t, []byte("線"))
	Ok(t, []byte(" "))
	m1 := map[int]string{1: "a", 2: "c"}
	m2 := map[int]string{2: "c", 1: "a"}
	Is(t, m1, m2, "Maps")

}

func TestIs(t *testing.T) {
	Is(t, 1, 1)
	Is(t, 0, 0, "message")
	Is(t, 1, 2-1)
	Is(t, 100, 10*10)
	Is(t, fmt.Sprintf("%s", "string"), "string")
	Is(t, "string", fmt.Sprintf("%s", "string"))
	Is(t, "string", "string")
	Is(t, "string", "str"+"ing")
	Is(t, 1.5, 0.5*3)
	Is(t, 0.5, 1.5/3)
	a := []byte("string")
	b := []byte("string")
	Is(t, a, b)
	Is(t, []byte("string"), []byte("string"))
	Is(t, t, t)
	t1 := t
	Is(t, t, t1)
	Is(t, *t, *t)
	Is(t, *t, *t1)
}

func TestIsnt(t *testing.T) {
	Isnt(t, 1, 1.1)
	Isnt(t, 0, "0", "message")
	Isnt(t, 1, 2-1.0000001)
	Isnt(t, 100, 10^10)
	Isnt(t, fmt.Sprintf("%v", "string1"), []byte("string1"))
	Isnt(t, []rune("string2"), fmt.Sprintf("%s", "string2"))
	Isnt(t, []byte("string3"), []rune("string3"))
	Isnt(t, "string", "str"+"in")
	Isnt(t, 1.6, 0.5333333333333333*3)
	Isnt(t, 0.5333333333333, 1.6/3)
	a := []byte("string1")
	b := []byte("string2")
	Isnt(t, a, b)
	Isnt(t, []byte("string3"), []byte("string4"))
	Isnt(t, *t, t)
	Isnt(t, t, *t)
	m1 := map[int]string{1: "a", 2: "c"}
	m2 := map[int]string{2: "c", 1: "b"}
	Isnt(t, m1, m2, "Maps")
}

func TestLike(t *testing.T) {
	Like(t, "qwerty", "w")
	Like(t, "qwerty", "we")
	Like(t, "qwerty", "^q")
	Like(t, "qwerty", ".")
	Like(t, "qwerty", ".{5}")
	Like(t, "qwerty", "[w]")
	Like(t, "qwerty", "[w]a?e", "Message")
}

func TestUnlike(t *testing.T) {
	Unlike(t, "qwerty", "[a]")
	Unlike(t, "qwerty", "[w]e+r$", "Message")
	Unlike(t, "qwerty", ".{9}")
}

func TestDiag(t *testing.T) {
	Diag(t, "My super message")
	Diag(t, "我的超級消息")
	Diag(t, "私のスーパーメッセージ")
	Diag(t, "Моё важное сообщение")
}
