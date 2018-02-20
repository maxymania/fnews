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


package fnewsd

import "net"
import "time"

import "github.com/maxymania/fnews/common/win"
import "github.com/maxymania/fnews/common/loader"
import "github.com/byte-mug/fastnntp"
import "github.com/byte-mug/fastnntp/posting"
import "github.com/maxymania/fastnntp-polyglot/caps"

import _ "github.com/lib/pq"
import _ "github.com/maxymania/fnews/common/loader/oohttp"
import _ "github.com/maxymania/fnews/common/loader/kcphttp"

// go build github.com/maxymania/fnews/win/fnews_run

type Lifecycle struct{
	cfgf string
	h *fastnntp.Handler
}
func NewLifecycle() (l *Lifecycle) {
	l = new(Lifecycle)
	l.cfgf = win.FindEtcConfig()
	return
}

func (l *Lifecycle) Load() error {
	c := new(caps.Caps)
	c.Stamper = posting.HostName("localhost")
	e := loader.LoadConfig(c,l.cfgf)
	if e!=nil { return e }
	l.h = &fastnntp.Handler{
		GroupCaps:c,
		ArticleCaps:c,
		PostingCaps:c,
		GroupListingCaps:c,
		//LoginCaps:adb,
	}
	return nil
}


func mainBolt(h *fastnntp.Handler, n net.Listener) {
	for {
		conn,err := n.Accept()
		if err!=nil { time.Sleep(time.Second) }
		go h.ServeConn(conn)
	}
}
func (l *Lifecycle) Serve() error {
	lis,err := net.Listen("tcp",":63119")
	if err!=nil { return err }
	mainBolt(l.h,lis)
	return nil
}


