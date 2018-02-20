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

import "github.com/maxymania/fnews/common/config"
import "github.com/maxymania/fastnntp-polyglot-labs/bucketstore/remote"
import "github.com/maxymania/fastnntp-polyglot-labs/articlewrap"
import "database/sql"
import "github.com/maxymania/fastnntp-polyglot-labs/articlewrap/sqldb"
import "github.com/maxymania/fastnntp-polyglot-labs/groupdb/semigroupdb"
import "github.com/maxymania/fastnntp-polyglot/caps"

import "github.com/valyala/fasthttp"
import "net"

type ClientLoader func(bs *config.BucketServer) remote.HttpClient

var ClientLoaders = map[string]ClientLoader{
	"http":ccHttp,
}

func dial(s string) (net.Conn,error) { return net.Dial("tcp",s) }

func ccHttp(bs *config.BucketServer) remote.HttpClient {
	hc := &fasthttp.HostClient{Addr: bs.Address}
	
	hc.Dial = dial
	return hc
}

func Create(a *config.ArticleBackendCfg, c *caps.Caps) {
	awrap := new(articlewrap.ArticleDirectBackend)
	proto := a.Bucket.Protocol
	lfu,ok := ClientLoaders[proto]
	if !ok { panic("unknown Client Loader: '"+proto+"'") }
	hc := lfu(a.Bucket)
	awrap.Store = remote.NewMultiClient(hc)
	db,err := sql.Open(a.Database.Driver, a.Database.DataSource)
	if err!=nil { panic(err) }
	awrap.Bdb = &sqldb.Base{db}
	c.ArticlePostingDB = awrap
	c.ArticleDirectDB  = awrap
	c.ArticleGroupDB   = awrap
	sdb := &semigroupdb.Base{db}
	c.GroupHeadCache   = &semigroupdb.AuthBase{*sdb,semigroupdb.ARUser}
	if a.Database.Driver=="postgresql" {
		pdb := &semigroupdb.PgBase{*sdb}
		c.GroupHeadDB      = pdb
		c.GroupRealtimeDB  = pdb
		c.GroupStaticDB    = pdb
	} else {
		c.GroupHeadDB      = sdb
		c.GroupRealtimeDB  = sdb
		c.GroupStaticDB    = sdb
	}
}

func SemiCreate(a *config.ArticleBackendCfg) (*sqldb.Base,*semigroupdb.Base,error) {
	db,err := sql.Open(a.Database.Driver, a.Database.DataSource)
	if err!=nil { return nil,nil,err }
	return &sqldb.Base{db},&semigroupdb.Base{db},nil
}


