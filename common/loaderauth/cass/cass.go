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


package cass

import (
	"github.com/maxymania/fnews/common/config"
	"github.com/gocql/gocql"
	
	"github.com/maxymania/fastnntp-polyglot/postauth/advauth"
	"github.com/maxymania/fastnntp-polyglot/auth/cassauth"
	
	"github.com/maxymania/fnews/common/loaderauth"
)

func loadCass(a *config.AuthCfg) (advauth.LoginHook,error) {
	cluster := gocql.NewCluster(a.Method.Hosts...)
	cluster.Keyspace = a.Method.Dbname
	if a.Method.Username!="" && a.Method.Password!="" {
		cluster.Authenticator = gocql.PasswordAuthenticator{a.Method.Username,a.Method.Password}
	}
	session,err := cluster.CreateSession()
	if err!=nil { return nil,err }
	cassauth.Initialize(session)
	lh := &cassauth.CassLoginHook{DB:session}
	err = lh.InitHash(a.Method.Pwhash)
	if err!=nil { return nil,err }
	return lh,nil
}

func init() {
	loaderauth.LoginHookLoaders["cass"] = loadCass
}
