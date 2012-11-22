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

import (
	"net"
)

type stream interface {
	// Writes data to the stream.
	WriteTo(data []byte)

	// Writes data to the parent stream. 
	// Called only by substreams as a method of facilitating hereditary structure.
	writeTo(from stream, data []byte)

	// Indicate a willingness to accept substreams on this stream.
	ListenForSubstreams()
	AcceptSubstream() (*SubStream, error)

	SendDatagram()
	ReceiveDatagram()

	LocalAddr() net.Addr
	RemoteAddr() net.Addr
}

// Top-level application stream.
type BaseStream struct {
}
type Conn BaseStream // To align with naming conventions in go/net

func (baseStream *BaseStream) WriteTo(b []byte) {}

// Writes data to the underlying channel.
func (baseStream *BaseStream) writeTo(from *stream, data []byte) {}
func (baseStream *BaseStream) ListenForSubstreams()              {}
func (baseStream *BaseStream) AcceptSubstream() (*SubStream, error) {
	return nil, nil
}
func (baseStream *BaseStream) SendDatagram()    {}
func (baseStream *BaseStream) ReceiveDatagram() {}
func (baseStream *BaseStream) LocalAddr() net.Addr {
	return nil
}
func (baseStream *BaseStream) RemoteAddr() net.Addr {
	return nil
}

// A child stream of a BaseStream. 
type SubStream struct {
	parent stream
}

func (this *SubStream) WriteTo(data []byte) {
	this.writeTo(stream(this), data)
}

func (this *SubStream) writeTo(from stream, data []byte) {
	// Write to parent stream
	this.parent.writeTo(stream(this), data)
}

func (this *SubStream) ListenForSubstreams() {}
func (this *SubStream) AcceptSubstream() (*SubStream, error) {
	return nil, nil
}
func (this *SubStream) SendDatagram()    {}
func (this *SubStream) ReceiveDatagram() {}
func (this *SubStream) LocalAddr() net.Addr {
	return nil
}
func (this *SubStream) RemoteAddr() net.Addr {
	return nil
}
