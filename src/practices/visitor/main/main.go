package main

import . "practices/visitor"

func main() {
	info := Info{}
	var v Visitor = &info
	//v = LogVisitor{v}
	//v = NameVisitor{v}
	//v = OtherThingsVisitor{v}
	loadFile := func(info *Info, err error) error {
		info.Name = "Hao Chen"
		info.Namespace = "MegaEase"
		info.OtherThings = "We are running as remote team."
		return nil
	}

	_ = v.Visit(loadFile)
}
