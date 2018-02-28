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

import "github.com/maxymania/fnews/common/win"
import "github.com/maxymania/fnews/common/loader"
import "github.com/maxymania/fnews/common/loaderauth"

import "github.com/maxymania/fastnntp-polyglot-labs/util/sqlutil"
import "fmt"

// go build github.com/maxymania/fnews/win.tools/fnews_prepare

func main() {
	cfgf := win.FindEtcConfig()
	s1,_,e := loader.LoadConfigSemi(cfgf)
	if e!=nil { fmt.Println(e); goto done1 }
	for _,sm := range s1 {
		e = sm.CreateSqlModel(sqlutil.PgDialect)
		if e!=nil { fmt.Println(e) }
	}
	
	done1:
	
	s1,_,e = loaderauth.LoadConfigSemi(cfgf)
	if e!=nil { fmt.Println(e); return }
	for _,sm := range s1 {
		e = sm.CreateSqlModel(sqlutil.PgDialect)
		if e!=nil { fmt.Println(e) }
	}
	
}


