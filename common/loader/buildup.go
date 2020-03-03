/*
Copyright (c) 2018,2020 Simon Schmidt

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

import "github.com/maxymania/fnews/common/config"
import "github.com/maxymania/fastnntp-polyglot/caps"

import "fmt"

import "github.com/maxymania/fastnntp-polyglot"
import "github.com/maxymania/fastnntp-polyglot/gold"
import "github.com/maxymania/fastnntp-polyglot/gold/setup"
import "github.com/maxymania/fastnntp-polyglot/postauth"

import "time"

type ArticleDirectLoader func(bs *config.ArticleDirect) (gold.ArticleDirectEX,error)
var ArticleDirectLoaders = make(map[string]ArticleDirectLoader)

type ArticleGroupLoader func(bs *config.ArticleGroup) (gold.ArticleGroupEX,error)
var ArticleGroupLoaders = make(map[string]ArticleGroupLoader)

type GroupListLoader func(bs *config.GroupList) (gold.GroupListDB,error)
var GroupListLoaders = make(map[string]GroupListLoader)

type GroupHeadDBLoader func(bs *config.GroupHead) (newspolyglot.GroupHeadDB,error)
var GroupHeadDBLoaders = make(map[string]GroupHeadDBLoader)

func assert2(ok bool,err string) {
	if !ok { panic(err) }
}
func failon(err error) {
	if err!=nil { panic(err) }
}

type days int
func (i days) DecideLite(groups [][]byte, lines, length int64) gold.PostingDecisionLite {
	return gold.PostingDecisionLite{ExpireAt: time.Now().UTC().AddDate(0,0,int(i))}
}

func Create(a *config.ArticleBackendCfg, c *caps.Caps) {
	
	adl,ok := ArticleDirectLoaders[a.Store.Method]
	assert2(ok,"unknown Article-Direct Backend: '"+a.Store.Method+"'")
	agl,ok := ArticleGroupLoaders[a.GroupIdx.Method]
	assert2(ok,"unknown Article-Group Backend: '"+a.GroupIdx.Method+"'")
	gll,ok := GroupListLoaders[a.GroupLst.Method]
	assert2(ok,"unknown Group-List Backend: '"+a.GroupLst.Method+"'")
	ghl,ok := GroupHeadDBLoaders[a.GroupHead.Method]
	assert2(ok,"unknown Group-Head-DB Backend: '"+a.GroupHead.Method+"'")
	
	ad,err := adl(a.Store)
	failon(err)
	ag,err := agl(a.GroupIdx)
	failon(err)
	gl,err := gll(a.GroupLst)
	failon(err)
	var pp gold.PostingPolicyLite = days(30)
	
	setup.Setup(c,ad,ag,gl,pp)
	c.GroupHeadCache = &postauth.GroupHeadCacheAuthed{gl,postauth.ARUser}
	
	gh,err := ghl(a.GroupHead)
	failon(err)
	c.GroupHeadDB = gh
}
type FakeSqlModel interface {
    CreateSqlModel(d interface{}) error
}
type Backends struct {
	gold.ArticleDirectEX
	gold.ArticleGroupEX
	gold.GroupListDB
	newspolyglot.GroupHeadDB
}
func SemiCreate(a *config.ArticleBackendCfg) (r1 []FakeSqlModel,r2 *Backends,e error) {
	defer func(){
		p := recover()
		if p==nil { return }
		if f,ok := p.(error) ; ok { e = f; return }
		e = fmt.Errorf("%v",p)
	}()
	
	adl,ok := ArticleDirectLoaders[a.Store.Method]
	assert2(ok,"unknown Article-Direct Backend: '"+a.Store.Method+"'")
	agl,ok := ArticleGroupLoaders[a.GroupIdx.Method]
	assert2(ok,"unknown Article-Group Backend: '"+a.GroupIdx.Method+"'")
	gll,ok := GroupListLoaders[a.GroupLst.Method]
	assert2(ok,"unknown Group-List Backend: '"+a.GroupLst.Method+"'")
	ghl,ok := GroupHeadDBLoaders[a.GroupHead.Method]
	assert2(ok,"unknown Group-Head-DB Backend: '"+a.GroupHead.Method+"'")
	
	ad,err := adl(a.Store)
	failon(err)
	ag,err := agl(a.GroupIdx)
	failon(err)
	gl,err := gll(a.GroupLst)
	failon(err)
	gh,err := ghl(a.GroupHead)
	failon(err)
	
	return nil,&Backends{ad,ag,gl,gh},nil
}


