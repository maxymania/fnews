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
import "os"
import "flag"
import "path/filepath"
import "golang.org/x/sys/windows/svc/mgr"

// go build github.com/maxymania/fnews/win.tools/fnews_svc_install

var folder = flag.String("path",".","file path to the fnews binaries")

type Service struct{
	Name,File,Descr string
}

type ErrorHaving struct{
	Service
	Err error
}

func main() {
	flag.Parse()
	elems := []Service{
		{"fnews_bktstor_service","","FNews Bucketstore"},
		{"fnews_service","","FNews NNTP Server"},
	}
	for i := range elems {
		subelem := elems[i].Name
		fn,e := filepath.Abs(filepath.Join(*folder,subelem+".exe"))
		if e!=nil { fmt.Println(e); flag.PrintDefaults(); return }
		f,e := os.Open(fn)
		if e!=nil { fmt.Println(e); flag.PrintDefaults(); return }
		f.Close()
		elems[i].File = fn
	}
	mconn,err := mgr.Connect()
	if err!=nil { fmt.Println(err); return }
	
	var eh []ErrorHaving
	for _,elem := range elems {
		s,e := mconn.CreateService(elem.Name, elem.File, mgr.Config{DisplayName:elem.Descr})
		if e!=nil {
			eh = append(eh,ErrorHaving{elem,e})
		}else{
			s.Close()
		}
	}
	for _,ee := range eh {
		fmt.Println(ee.Name,ee.File,ee.Descr,ee.Err)
	}
}

