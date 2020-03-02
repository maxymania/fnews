/*
Copyright (c) 2020 Simon Schmidt

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


package loaderauth

import "io/ioutil"
import "path/filepath"
import "fmt"
import "github.com/byte-mug/fastnntp"
import "github.com/byte-mug/goconfig"
import "github.com/maxymania/fnews/common/config"
import "github.com/maxymania/fastnntp-polyglot/caps"
import "os"

func loadConfig(cfgf string) (a *config.AuthCfg,e error) {
	data,err := ioutil.ReadFile(filepath.Join(cfgf,"auth.conf"))
	if err!=nil { e = err; return }
	a = new(config.AuthCfg)
	e = goconfig.Parse(data,goconfig.CreateReflectHandler(a))
	return
}

func LoadConfig(c *caps.Caps, h *fastnntp.Handler,cfgf string) (e error) {
	defer func(){
		p := recover()
		if p==nil { return }
		if f,ok := p.(error) ; ok { e = f; return }
		e = fmt.Errorf("%v",p)
	}()
	a,err := loadConfig(cfgf)
	if err!=nil {
		if _,ok := err.(*os.PathError); ok { return } /* If file was not found: Ignore it. */
		e = err
		return
	}
	Create(a,c,h)
	return
}

func LoadConfigSemi(cfgf string) (s1 []FakeSqlModel,s2 LoginAdm,e error) {
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


