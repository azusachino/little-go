package ioc

import "errors"

type Undo []func()

func (undo *Undo) Add(function func()) {
	*undo = append(*undo, function)
}

func (undo *Undo) Undo() error {
	functions := *undo
	if len(functions) == 0 {
		return errors.New("No functions to undo")
	}
	index := len(functions) - 1
	if function := functions[index]; function != nil {
		function()
		functions[index] = nil // For garbage collection
	}
	*undo = functions[:index]
	return nil
}

type IntegerSet struct {
	data map[int]bool
	undo Undo
}

func NewIntegerSet() IntegerSet {
	return IntegerSet{data: make(map[int]bool)}
}

func (set *IntegerSet) Undo() error {
	return set.undo.Undo()
}

func (set *IntegerSet) Contains(x int) bool {
	return set.data[x]
}

func (set *IntegerSet) Add(x int) {
	if !set.Contains(x) {
		set.data[x] = true
		set.undo.Add(func() { set.Delete(x) })
	} else {
		set.undo.Add(nil)
	}
}

func (set *IntegerSet) Delete(x int) {
	if set.Contains(x) {
		delete(set.data, x)
		set.undo.Add(func() { set.Add(x) })
	} else {
		set.undo.Add(nil)
	}
}
