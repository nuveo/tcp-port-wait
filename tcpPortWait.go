package tcpPortWait

import (
	"net"
	"time"
)

// Port contains tcpPortWait config parameters
type Port struct {
	Timeout time.Duration
}

// Check if tcp port is open
func (p *Port) Check(host string) (ret bool, err error) {
	for {
		var timeout time.Duration
		if p.Timeout == 0 {
			timeout = time.Duration(10) * time.Minute
		}
		var conn net.Conn
		conn, err = net.DialTimeout("tcp", host, timeout)
		if err != nil {
			ret = true
			err = nil
			return
		}
		if conn != nil {
			err = conn.Close()
			return
		}
	}
}
