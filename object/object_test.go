package object

import "testing"

func TestStreingHashKey(t *testing.T) {
	hello1 := &String{Value: "Hello world!"}
	hello2 := &String{Value: "Hello world!"}
	diff1 := &String{Value: "My Name is Tom"}
	diff2 := &String{Value: "My Name is Tom"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("same string but wrong key")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("same string but wrong key")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("wrong string but same key")
	}
}

func TestBooleanHashKey(t *testing.T) {
	hello1 := &Boolean{Value: true}
	hello2 := &Boolean{Value: true}
	diff1 := &Boolean{Value: false}
	diff2 := &Boolean{Value: false}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("same string but wrong key")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("same string but wrong key")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("wrong string but same key")
	}
}

func TestIntegerHashKey(t *testing.T) {
	hello1 := &Integer{Value: 1}
	hello2 := &Integer{Value: 1}
	diff1 := &Integer{Value: 10}
	diff2 := &Integer{Value: 10}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("same string but wrong key")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("same string but wrong key")
	}

	if hello1.HashKey() == diff1.HashKey() {
		t.Errorf("wrong string but same key")
	}
}
