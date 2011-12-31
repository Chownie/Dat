// controller.go
package controller

import (
	"net"
	"parsing"
	"fmt"
)

type ChanPair struct {
	in, out chan string
}

//Creates a new routine
func NewRoutine(clist *[]ChanPair, controlc chan chan string, conn *net.TCPConn) {
	go readloop(conn, clist, controlc)
	var l ChanPair
	l.out = <- controlc
	l.in = <- controlc
	*clist = append(*clist, l)
}

//The "mainloop" of the server
func ControlLoop(clist *[]ChanPair, controlc chan chan string) {
	fmt.Println("Control Loop beginning")
	limit := 0
	for {
		limit = len(*clist)
		for i := 0; i < limit; i++ {
			l := *clist
			select{
				case outbound := <- l[i].out:
					if len(outbound) > 0{fmt.Println(outbound)}
					l[i].in <- outbound
				default:
			}
		}
	}
}

func readloop(conn *net.TCPConn, clist *[]ChanPair, controlc chan chan string) {
	output := make(chan string, 2048)
	input := make(chan string, 2048)
	controlc <- output
	controlc <- input
	address := conn.RemoteAddr()
	player := parsing.NewPlayer(address.String())
	for {
		b := make([]byte, 4096)
		n, err := conn.Read(b[:])
		data := b[:n]
		if err != nil {
			fmt.Println(err)
		}

		select {
		case str := <-input:
			conn.Write([]uint8(str))
		default:
		}

		if len(string(data)) == 0 {
			fmt.Println("PARTING:", address.String())
			conn.Close()
			return
		}
		conn.Write(parsing.Parse(player, string(data), output))
	}
}
