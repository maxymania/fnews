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
	"github.com/maxymania/fnews/common/loader"
	"github.com/maxymania/fnews/common/config"
	"github.com/maxymania/fastnntp-polyglot/gold"
	
	"github.com/gocql/gocql"
	"errors"
)

import (
	adcass "github.com/maxymania/fastnntp-polyglot/gold/ad.cass"
	cassm  "github.com/maxymania/fastnntp-polyglot/gold/ag.cassm"
	glcass "github.com/maxymania/fastnntp-polyglot/gold/gl.cass"
)

var ErrNoKeyspace = errors.New("Keyspace not defined")

var keyspaces = make(map[string]*gocql.Session)

func openSession(ks *config.Keyspace) (*gocql.Session,error) {
	if ks==nil { return nil,ErrNoKeyspace }
	sess,ok := keyspaces[ks.KSKey()]
	if ok { return sess,nil }
	cluster := gocql.NewCluster(ks.Hosts...)
	cluster.Keyspace = ks.Keyspace
	if ks.Username!="" && ks.Password!="" {
		cluster.Authenticator = gocql.PasswordAuthenticator{ks.Username,ks.Password}
	}
	
	sess,err := cluster.CreateSession()
	if err!=nil { return nil,err }
	keyspaces[ks.KSKey()] = sess
	return sess,nil
}

func loadAD(bs *config.ArticleDirect) (gold.ArticleDirectEX,error) {
	sess,err := openSession(bs.Keyspace)
	if err!=nil { return nil,err }
	
	update,err := gocql.ParseConsistencyWrapper(bs.Keyspace.WriteCs1)
	if err!=nil { return nil,err }
	
	adcass.Initialize(sess)
	
	return &adcass.Storage{sess,update},nil
}

func loadAG(bs *config.ArticleGroup) (gold.ArticleGroupEX,error) {
	sess,err := openSession(bs.Keyspace)
	if err!=nil { return nil,err }
	
	assign,err := gocql.ParseConsistencyWrapper(bs.Keyspace.WriteCs1)
	if err!=nil { return nil,err }
	increm,err := gocql.ParseConsistencyWrapper(bs.Keyspace.WriteCs2)
	if err!=nil { return nil,err }
	
	cassm.InitializeN2Layer(sess)
	
	n2l := new(cassm.N2LayerGroupDB)
	n2l.Session     = sess
	n2l.OnAssign    = assign
	n2l.OnIncrement = increm
	return n2l,nil
}

func loadGL(bs *config.GroupList) (gold.GroupListDB,error) {
	sess,err := openSession(bs.Keyspace)
	if err!=nil { return nil,err }
	
	update,err := gocql.ParseConsistencyWrapper(bs.Keyspace.WriteCs1)
	if err!=nil { return nil,err }
	
	glcass.Initialize(sess)
	
	return &glcass.Database{sess,update},nil
}

func init() {
	loader.ArticleDirectLoaders["cass"] = loadAD
	loader.ArticleGroupLoaders ["cass"] = loadAG
	loader.GroupListLoaders    ["cass"] = loadGL
}


// #
