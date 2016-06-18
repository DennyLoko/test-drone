package main

type Foo struct {
	bar string
}

func (f *Foo) Bar() string {
	return f.bar
}

func (f *Foo) SetBar(bar string) {
	f.bar = bar
}
