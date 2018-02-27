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


package loaderlst

import "github.com/maxymania/fnews/common/config"
//import "net"
import "crypto/tls"
import "path/filepath"
import "strings"

var seperator = string([]byte{filepath.Separator})

func GetTLS(cfg *config.NntpListener) (*tls.Config,error) {
	tc := cfg.TlsConfig
	if tc==nil { return nil,nil }
	
	tcf := &tls.Config{}
	
	tcf.Certificates = make([]tls.Certificate,len(tc.Crts))
	for i,crcf := range tc.Crts {
		crt := strings.Replace(crcf.Crt,"/",seperator,-1)
		key := strings.Replace(crcf.Key,"/",seperator,-1)
		cert,err := tls.LoadX509KeyPair(crt,key)
		if err!=nil { return nil,err }
		tcf.Certificates[i] = cert
	}
	
	return tcf,nil
}