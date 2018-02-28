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


package kcphttp

import "net"
import "github.com/maxymania/fnews/common/loader"
import "github.com/maxymania/fnews/common/config"
import "github.com/maxymania/fastnntp-polyglot-labs/kcphttp"
import "github.com/xtaci/kcp-go"
import "github.com/maxymania/fastnntp-polyglot-labs/bucketstore/remote"

func kcpDialer(bs *config.BucketServer) func() (net.Conn,error) {
	addr := bs.Address
	return func() (net.Conn,error) {
		var cipher kcp.BlockCrypt
		cipher = nil
		ce,err := kcp.DialWithOptions(addr, cipher, bs.DataShards, bs.ParityShards)
		if ce!=nil && bs.TurboMode { ce.SetNoDelay(1,40,1,1) }
		return ce,err
	}
}

func ccKcp(bs *config.BucketServer) remote.HttpClient {
	kc := kcphttp.NewClient(kcpDialer(bs))
	return kc
}

func init() {
	loader.ClientLoaders["kcphttp"] = ccKcp
}
