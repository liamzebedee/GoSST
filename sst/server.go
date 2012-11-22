/*
   This file is part of GoSST.

   GoSST is free software: you can redistribute it and/or modify
   it under the terms of the GNU General Public License as published by
   the Free Software Foundation, either version 3 of the License, or
   (at your option) any later version.

   GoSST is distributed in the hope that it will be useful,
   but WITHOUT ANY WARRANTY; without even the implied warranty of
   MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
   GNU General Public License for more details.

   You should have received a copy of the GNU General Public License
   along with GoSST.  If not, see <http://www.gnu.org/licenses/>.
*/

package sst

// Contains code for serving connections to other nodes.

import (
	"net"
)

// A listener to accept incoming streams.
type Listener struct {
	udpConn      net.Listener // Underlying network connection
	closeChannel chan bool
}

// Listen announces on the address (using UDP) and returns a listener that listens for streams.
func Listen(address string) (ln *Listener, err error) {
	ln = &Listener{}

	// Construct underlying network connection
	ln.udpConn, err = net.Listen("udp", address)
	if err != nil {
		logger.Printf("couldn't create listener - %s\n", err.Error())
		return nil, err
	}

	ln.closeChannel = make(chan bool)

	go ln.listen()
	return ln, nil
}

func (ln *Listener) listen() {
	defer ln.udpConn.Close()

	for {
		select {
		case <-ln.closeChannel:
			return
		default:
			// Listen for connections
			return
		}
	}
}

// Accept waits for and returns the next connection (a stream) to the listener.
func (ln *Listener) Accept() (*BaseStream, error) {
	// TODO
	return nil, nil
}
