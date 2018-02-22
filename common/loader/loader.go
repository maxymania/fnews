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


package loader

import "io/ioutil"
import "path/filepath"
import "fmt"
import "github.com/byte-mug/goconfig"
import "github.com/maxymania/fnews/common/config"
import "github.com/maxymania/fastnntp-polyglot-labs/util/sqlutil"
import "github.com/maxymania/fastnntp-polyglot-labs/util/groupadm"
import "github.com/maxymania/fastnntp-polyglot/caps"

func loadConfig(cfgf string) (a *config.ArticleBackendCfg,e error) {
	data,err := ioutil.ReadFile(filepath.Join(cfgf,"backend.conf"))
	if err!=nil { e = err; return }
	a = new(config.ArticleBackendCfg)
	e = goconfig.Parse(data,goconfig.CreateReflectHandler(a))
	return
}

func LoadConfig(c *caps.Caps,cfgf string) (e error) {
	defer func(){
		p := recover()
		if p==nil { return }
		if f,ok := p.(error) ; ok { e = f; return }
		e = fmt.Errorf("%v",p)
	}()
	a,err := loadConfig(cfgf)
	if err!=nil { e = err; return }
	Create(a,c)
	return
}

func LoadConfigSemi(cfgf string) (s1 []sqlutil.SqlModel,s2 groupadm.GroupAdm,e error) {
	defer func(){
		p := recover()
		if p==nil { return }
		if f,ok := p.(error) ; ok { e = f; return }
		e = fmt.Errorf("%v",p)
	}()
	a,err := loadConfig(cfgf)
	if err!=nil { e = err; return }
	s1,s2,e = SemiCreate(a)
	return
}


