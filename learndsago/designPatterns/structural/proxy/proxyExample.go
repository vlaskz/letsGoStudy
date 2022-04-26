package proxy

import "fmt"

type IRealObject interface {
	performAction()
}
type RealObject struct{}

func (realObject *RealObject) performAction() {
	fmt.Println("RealObject performAction()")
}

type VirtualProxy struct {
	realObject *RealObject
}

func (virtualProxy *VirtualProxy) performAction() {
	if virtualProxy.realObject == nil {
		virtualProxy.realObject = &RealObject{}
	}
	fmt.Println("VirtualProxy performAction()")
	virtualProxy.realObject.performAction()
}
func Example() {
	var object VirtualProxy = VirtualProxy{}
	object.performAction()
}
