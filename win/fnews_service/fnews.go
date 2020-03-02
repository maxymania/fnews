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


package main

import "fmt"
import "github.com/maxymania/fnews/common/fnewsd"
import "golang.org/x/sys/windows/svc"

// go build github.com/maxymania/fnews/win/fnews_service

type service struct {}

func doserve() {
	lc := fnewsd.NewLifecycle()
	e := lc.Load()
	if e!=nil { fmt.Println(e); return }
	lc.Serve()
}

func (_ service) Execute(args []string, r <-chan svc.ChangeRequest, changes chan<- svc.Status) (svcSpecificEC bool, exitCode uint32) {
	const cmdsAccepted = svc.AcceptStop | svc.AcceptShutdown
	changes <- svc.Status{State: svc.Running, Accepts: cmdsAccepted}
	defer func(){ changes <- svc.Status{State: svc.StopPending} }()
	go doserve()
	for {
		c := <- r
		switch c.Cmd {
		case svc.Stop, svc.Shutdown: return
		}
	}
	return
}

func main() {
	isOK,err := svc.IsAnInteractiveSession()
	if err!=nil || isOK {
		fmt.Println("Running in interactive mode")
		doserve()
		return
	}
	svc.Run("fnews_service",service{})
}


