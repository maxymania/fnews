/*
Copyright (c) 2018 Simon Schmidt

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*/


package loaderlst

import "github.com/maxymania/fnews/common/config"
import "net"
import "strconv"
import "golang.org/x/net/ipv4"
import "golang.org/x/net/ipv6"

type Modifier interface{
	ModifySocket(sck net.Conn)
}

type nonesense struct{}
func (_ nonesense) ModifySocket(sck net.Conn) {}

var defModifier Modifier = nonesense{}

type modifier struct{
	hopLimit,dscp int
}
var _ Modifier = modifier{}

func (m modifier) ModifySocket(sck net.Conn) {
	var i net.IP
	if ipa,ok := sck.RemoteAddr().(*net.IPAddr) ; ok { i = ipa.IP }
	if i.To4()==nil {
		c := ipv6.NewConn(sck)
		if m.hopLimit!=0 { c.SetHopLimit(m.hopLimit) }
		if m.dscp!=0 { c.SetTrafficClass(m.dscp<<2) }
	} else {
		c := ipv4.NewConn(sck)
		if m.hopLimit!=0 { c.SetTTL(m.hopLimit) }
		if m.dscp!=0 { c.SetTOS(m.dscp<<2) }
	}
}

func GetModifier(cfg *config.NntpListener) Modifier {
	var m modifier
	m.hopLimit = cfg.HopLimit
	m.dscp = dscp2int(cfg.DscpValue)
	if m.hopLimit!=0 || m.dscp!=0 { return m }
	return defModifier
}

func dscp2int(s string) int{
	switch s {
	case "CS0" : return 0
	case "CS1" : return 0x08
	case "CS2" : return 0x10
	case "CS3" : return 0x18
	case "CS4" : return 0x20
	case "CS5" : return 0x28
	case "CS6" : return 0x30
	case "CS7" : return 0x38
	case "AF11": return 0x0a
	case "AF12": return 0x0c
	case "AF13": return 0x0e
	case "AF21": return 0x12
	case "AF22": return 0x14
	case "AF23": return 0x16
	case "AF31": return 0x1a
	case "AF32": return 0x1c
	case "AF33": return 0x1e
	case "AF41": return 0x22
	case "AF42": return 0x24
	case "AF43": return 0x26
	case "EF"  : return 0x2e
	}
	u,err := strconv.ParseUint(s,0,32)
	if err!=nil { u = 0 }
	return int(u)
}


