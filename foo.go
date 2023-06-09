package foo

type foo struct {
	a [64]int64
	b [64]int64
}

func getFoo() foo {
	var a [64]int64
	var b [64]int64
	return foo{a, b}
}

func getFooPointer() *foo {
	var a [64]int64
	var b [64]int64
	return &foo{a, b}
}

func processFoo() {
	foo := getFoo()
	for i := 0; i < len(foo.a); i++ {
		_ = foo.a[i] * 2
	}
	for i := 0; i < len(foo.b); i++ {
		_ = foo.b[i] * 2
	}
}

func processFooPointer() {
	foo := *getFooPointer()
	for i := 0; i < len(foo.a); i++ {
		_ = foo.a[i] * 2
	}
	for i := 0; i < len(foo.b); i++ {
		_ = foo.b[i] * 2
	}
}
