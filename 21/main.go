package main

import "fmt"

type AmericanSocket interface {
	PlugInUS() string
}

type USARose struct{}

func (s *USARose) PlugInUS() string {
	return "Device is plugged into American socket"
}

type EuroPlug interface {
	PlugInEU() string
}

type EuropeanDevice struct{}

func (d *EuropeanDevice) PlugInEU() string {
	return "Device is plugged into European socket"
}

type Adapter struct {
	device EuroPlug
}

func (a *Adapter) PlugInUS() string {
	return a.device.PlugInEU() + " using an adapter"
}

func ClientCode(socket AmericanSocket) {
	fmt.Println(socket.PlugInUS())
}

func main() {
	euroDevice := &EuropeanDevice{}

	adapter := &Adapter{device: euroDevice}

	americanSocket := &USARose{}

	fmt.Println("Using adapter to connect European device:")
	ClientCode(adapter)

	fmt.Println("Connecting directly to American socket:")
	ClientCode(americanSocket)
}
