package resize

import (
	"fmt"
	"testing"
)

func TestResize(t *testing.T) {
	err := Resize("txx.png", "s.png", 100, 100)
	if err == nil {
		t.Error("not right")
		return
	}
	err = Resize("t.png", "s.png", 100, 100)
	if err != nil {
		t.Error("not right")
		return
	}
	err = Resize("t.png", "s.png", 100, 100)
	if err != nil {
		t.Error(err.Error())
		return
	}
	err = Resize("t.jpg", "s.jpg", 100, 100)
	if err != nil {
		t.Error(err.Error())
		return
	}
	fmt.Println("all...")
}
