package main

import (
	"fmt"
	"net"
	"controller"
)

func main() {
	i, _adderr := net.ResolveTCPAddr("tcp", "0.0.0.0:6580")
	l, _liserr := net.ListenTCP("tcp", i)
	controlchan := make(chan chan string)
	clist := make([]controller.ChanPair,1)
	go controller.ControlLoop(&clist, controlchan)

	for {
		conn, _accerr := l.AcceptTCP()

		address := conn.RemoteAddr()
		fmt.Println("INCOMING:", address.String())

		controller.NewRoutine(&clist, controlchan, conn)

		if _accerr != nil {
			fmt.Println(_accerr)
		}
	}
	l.Close()
	fmt.Println(_adderr)
	fmt.Println(_liserr)
}
