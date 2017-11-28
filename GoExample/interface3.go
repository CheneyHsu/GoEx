package main

import "fmt"

type USB interface {
	Name() string
	Connecter
}

type Connecter interface {
	Connect()
}

type Phoneconnecter struct {
	name string
}

func (pc Phoneconnecter) Name() string {
	return pc.Name()
}

func (pc Phoneconnecter) Connect() {
	fmt.Println("connected", pc.name)
}

func main() {

	a := Phoneconnecter{"Phoneconnecter"}
	a.Connect()
	Disconnected(a)

}
func Disconnected(usb USB) {
	fmt.Println("Disconnected.")
}
