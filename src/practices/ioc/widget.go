package ioc

type Widget struct {
	X, Y int
}

type Label struct {
	Widget
	Text string
}

type Button struct {
	Label
}

type ListBox struct {
	Widget
	Texts []string
	Index int
}

func Abc() {
	label := Label{Widget{10, 10}, "State:"}
	label.X = 11
	label.Y = 12
	print(label)
}
