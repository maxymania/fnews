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


package loaderauth

import (
	"database/sql"
	"fmt"
	
	"github.com/byte-mug/fastnntp"
	"github.com/maxymania/fnews/common/config"
	
	"github.com/maxymania/fastnntp-polyglot/caps"
	"github.com/maxymania/fastnntp-polyglot/postauth"
	"github.com/maxymania/fastnntp-polyglot-labs/util/sqlutil"
	"github.com/maxymania/fastnntp-polyglot-labs/auth/sqlauth"
	"github.com/maxymania/fastnntp-polyglot-labs/auth/advauth"
	"github.com/maxymania/fastnntp-polyglot-labs/auth/advauthif"
)

type LoginAdm interface {
	InsertUser(user, password []byte,rank postauth.AuthRank) error
	UpdateUserPassword(user, password []byte) error
	UpdateUserRank(user []byte, rank postauth.AuthRank) error
}

func Create(a *config.AuthCfg, c *caps.Caps, h *fastnntp.Handler) {
	if a==nil { return }
	if !a.Enable { return }
	if a.Method==nil { return }
	
	auther := new(advauth.Auther)
	auther.Base = h
	
	switch a.Method.Method {
	case "db":{
			db,err := sql.Open(a.Method.Driver, a.Method.Db)
			if err!=nil { panic(err) }
			auther.Hook = &sqlauth.LoginDB{db}
		}
	default:
		panic(fmt.Errorf("unknown METHOD: %s",a.Method.Method))
	}
	
	auther.Deriv = &advauthif.Deriver{Caps:c,Ghead: advauthif.ExtractGhead(c)}
	
	if a.NoAuth.NoReading {
		*h = fastnntp.Handler{} // Clear handler to prevent reading.
	} else if !a.NoAuth.NoPosting { /* Unauthenticated Posting is disabled by default. */
		/* TODO: implement, enable. */
		
	}
	
	h.LoginCaps = auther
}

func SemiCreate(a *config.AuthCfg) ([]sqlutil.SqlModel,LoginAdm,error) {
	switch a.Method.Method {
	case "db":{
			db,err := sql.Open(a.Method.Driver, a.Method.Db)
			if err!=nil { panic(err) }
			sqla := &sqlauth.LoginDB{db}
			return []sqlutil.SqlModel{sqla},sqla,nil
		}
	}
	return nil,nil,fmt.Errorf("unknown METHOD: %s",a.Method.Method)
}


