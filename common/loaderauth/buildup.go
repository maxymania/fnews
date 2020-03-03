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

import (
	"fmt"
	
	"github.com/byte-mug/fastnntp"
	"github.com/maxymania/fnews/common/config"
	
	"github.com/gocql/gocql"
	
	"github.com/maxymania/fastnntp-polyglot/caps"
	"github.com/maxymania/fastnntp-polyglot/postauth"
	"github.com/maxymania/fastnntp-polyglot/postauth/advauth"
	"github.com/maxymania/fastnntp-polyglot/postauth/advauthif"
	"github.com/maxymania/fastnntp-polyglot/auth/cassauth"
)

var NoAuth = fmt.Errorf("No auth.")

type LoginAdm interface {
	InsertUser(user, password []byte,rank postauth.AuthRank) error
	UpdateUserPassword(user, password []byte) error
	UpdateUserRank(user []byte, rank postauth.AuthRank) error
}

func Open(a *config.AuthCfg) (advauth.LoginHook,error) {
	if a==nil { return nil,NoAuth }
	if !a.Enable { return nil,NoAuth }
	if a.Method==nil { return nil,NoAuth }
	
	switch a.Method.Method {
	case "cass","cassandra","cql": {
			cluster := gocql.NewCluster(a.Method.Hosts...)
			cluster.Keyspace = a.Method.Dbname
			session,err := cluster.CreateSession()
			if err!=nil { return nil,err }
			lh := &cassauth.CassLoginHook{DB:session}
			err = lh.InitHash(a.Method.Pwhash)
			if err!=nil { return nil,err }
			return lh,nil
		}
	}
	return nil,fmt.Errorf("unknown METHOD: %s",a.Method.Method)
}

type loginDbg struct{
	fastnntp.LoginCaps
}
func (ld loginDbg) AuthinfoUserPass(user, password []byte, oldh *fastnntp.Handler) (bool, *fastnntp.Handler) {
	fmt.Printf("LOGIN %q %q\n",user,password)
	return ld.LoginCaps.AuthinfoUserPass(user,password,oldh)
}

type writable struct{
	fastnntp.LoginCaps
}
func (w writable) AuthinfoCheckPrivilege(p fastnntp.LoginPriv, h *fastnntp.Handler) bool {
	if p==fastnntp.LoginPriv_Post { return true }
	return w.LoginCaps.AuthinfoCheckPrivilege(p,h)
}

func Create(a *config.AuthCfg, c *caps.Caps, h *fastnntp.Handler) {
	var err error
	if a==nil { return }
	if !a.Enable { return }
	if a.Method==nil { return }
	
	auther := new(advauth.Auther)
	auther.Base = h
	
	auther.Hook,err = Open(a)
	if err!=nil { panic(err) }
	
	gch := advauthif.ExtractGhead(c)
	auther.Deriv = &advauthif.Deriver{Caps:c,Ghead:gch }
	auther.Backlink()
	
	if a.NoAuth.NoReading {
		*h = fastnntp.Handler{} // Clear handler to prevent reading.
	} else if a.NoAuth.NoPosting {
		c.GroupHeadCache = &postauth.GroupHeadCacheAuthed{gch,postauth.ARReader}
	} else {
		c.GroupHeadCache = &postauth.GroupHeadCacheAuthed{gch,postauth.ARUser}
	}
	
	if a.NoAuth.NoPosting||a.NoAuth.NoReading {
		// in case of NoPosting, posting is disallowed.
		// in case of NoReading, ANY access is disallowed.
		h.LoginCaps = auther
	} else {
		// IF posting and reading is allowed for anonymous users, enable write.
		h.LoginCaps = writable{auther}
	}
}

type FakeSqlModel interface {
    CreateSqlModel(d interface{}) error
}
func SemiCreate(a *config.AuthCfg) ([]FakeSqlModel,LoginAdm,error) {
	adm,err := Open(a)
	return nil,adm,err
	/*
	switch a.Method.Method {
	case "db":{
			db,err := sql.Open(a.Method.Driver, a.Method.Db)
			if err!=nil { panic(err) }
			sqla := &sqlauth.LoginDB{db}
			return []sqlutil.SqlModel{sqla},sqla,nil
		}
	}
	return nil,nil,fmt.Errorf("unknown METHOD: %s",a.Method.Method)
	*/
}


