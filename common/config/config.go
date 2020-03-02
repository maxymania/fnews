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


package config

type Keyspace struct{
	Keyspace string `inn:"$keyspace"`
	Cluster  string `inn:"$cluster"`
	Hosts  []string `inn:"@host"`
	Username string `inn:"$user"`
	Password string `inn:"$pass"`
	
	WriteCs1 string `inn:"$write-consistency"`
	WriteCs2 string `inn:"$atomic-consistency"`
}
func (ks *Keyspace) KSKey() string { return ks.Keyspace+ks.Cluster }

type ArticleDirect struct{
	Method string `inn:"$store"`
	Keyspace *Keyspace `inn:"$keyspace"`
	Dburl  string `inn:"$dburl"`
}
type ArticleGroup struct{
	Method string `inn:"$groupindex"`
	Keyspace *Keyspace `inn:"$keyspace"`
	Dburl  string `inn:"$dburl"`
}
type GroupList struct {
	Method string `inn:"$grouplist"`
	Keyspace *Keyspace `inn:"$keyspace"`
	Dburl  string `inn:"$dburl"`
}
type GroupHead struct {
	Method string `inn:"$grouphead"`
	Keyspace *Keyspace `inn:"$keyspace"`
	Dburl  string `inn:"$dburl"`
}

type ArticleBackendCfg struct{
	Store     *ArticleDirect `inn:"$store"`
	GroupIdx  *ArticleGroup  `inn:"$groupindex"`
	GroupLst  *GroupList     `inn:"$grouplist"`
	GroupHead *GroupHead     `inn:"$grouphead"`
}

// #
