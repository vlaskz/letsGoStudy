package decorator

import "fmt"

type IProcess interface {
	process()
}

type ProcessClass struct{}

func (process *ProcessClass) Process() {
	fmt.Println("ProcessClass process")
}

type ProcessDecorator struct {
	processInstance *ProcessClass
}

func (decorator *ProcessDecorator) process() {
	if decorator.processInstance == nil {
		fmt.Println("ProcessDecorator process")
	} else {
		fmt.Println("ProcessDecorator process and ")
		decorator.processInstance.Process()
	}

}

func Example() {
	var process = &ProcessClass{}
	var decorator = &ProcessDecorator{}
	decorator.process()
	decorator.processInstance = process
	decorator.process()
}
