package foo

type foo struct {
	a [1024]int8
	b [1024 * 63]int8
}

func getFoo() foo {
	var a [1024]int8
	var b [1024 * 63]int8
	return foo{a, b}
}

func getFooPointer() *foo {
	var a [1024]int8
	var b [1024 * 63]int8
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
