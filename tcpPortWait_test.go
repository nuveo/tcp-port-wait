package tcpPortWait

import (
	"fmt"
	"net"
	"os"
	"testing"
	"time"
)

func simpleServer() {
	rAddr, err := net.ResolveTCPAddr("tcp", ":9999")
	if err != nil {
		panic(err)
	}
	l, err := net.ListenTCP("tcp", rAddr)
	if err != nil {
		fmt.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	fmt.Println("Server listen at ", rAddr)
	// Close the listener when the application closes.
	defer func() {
		cErr := l.Close()
		if cErr != nil {
			fmt.Println(cErr)
		}
	}()
	for {
		// Listen for an incoming connection.
		conn, err := l.AcceptTCP()
		if err != nil {
			fmt.Println("Error accepting: ", err.Error())
			continue
		}
		fmt.Println("client at", conn.RemoteAddr())
	}
}

func TestCheck(t *testing.T) {
	var p Port
	p.Timeout = time.Duration(1) * time.Second
	timeout, err := p.Check(":9999")
	if err != nil {
		t.Fatal(err)
	}
	if !timeout {
		t.Fatal("Expected timeout true")
	}
	go simpleServer()
	p.Timeout = 0
	<-time.After(time.Duration(1) * time.Second)
	timeout, err = p.Check(":9999")
	if err != nil {
		t.Fatal(err)
	}
	if timeout {
		t.Fatal("Expected timeout false")
	}

}
