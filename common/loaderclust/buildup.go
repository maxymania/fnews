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


package loaderclust

import "github.com/maxymania/fnews/common/config"
import "github.com/maxymania/fastnntp-polyglot-labs/bucketstore"
import "github.com/maxymania/fastnntp-polyglot-labs/bucketstore/cluster"
import "github.com/hashicorp/memberlist"
import "net"
import "path/filepath"
import "strings"
import "fmt"

var ProtocolMap = map[string]uint{
	"" : cluster.Proto_HTTP,
	"http" : cluster.Proto_HTTP,
	"oohttp" : cluster.Proto_OO,
	"kcphttp" : cluster.Proto_KCP,
}

func configureMetadata(l *config.BucketServiceListener,m *cluster.MetaData) {
	var ok bool
	m.Proto,ok = ProtocolMap[l.Protocol]
	if !ok { panic("no such protocol: "+l.Protocol) }
	m.Port = l.Port
	m.KCP.DataShards   = l.DataShards
	m.KCP.ParityShards = l.ParityShards
	m.KCP.Salsa20Key   = nil // Still missing.
	m.KCP.TurboMode    = l.TurboMode
}

var fsSep = string([]byte{filepath.Separator})

func Create(cfg *config.BucketServiceCfg) *cluster.Membered {
	if cfg.Listener==nil { panic("not ready") }
	
	member := cluster.NewMembered()
	
	configureMetadata(cfg.Listener,&(member.Meta))
	
	for _,bak := range cfg.Backends {
		if bak.Config==nil { panic("No config @ "+bak.Type+" "+bak.Path) }
		pth := bak.Path
		pth = strings.Replace(pth,"/",fsSep,-1)
		buk,err := bucketstore.OpenStore(bak.Type,pth,bak.Config)
		if err!=nil { panic(err) } /* TODO: Reliability!? */
		member.Router.AddLocal(buk)
		member.Meta.Buckets = append(member.Meta.Buckets,buk.Uuid)
	}
	
	return member
}

func CreateClusterConfig(clust *config.BucketCluster,memb *cluster.Membered) (mc *memberlist.Config,e error) {
	switch clust.NodeSpecs {
	case "LAN","lan": mc = memberlist.DefaultLANConfig()
	case "WAN","wan": mc = memberlist.DefaultWANConfig()
	case "local","Local": mc = memberlist.DefaultLocalConfig()
	default: e = fmt.Errorf("unknown specs %v",clust.NodeSpecs); return
	}
	mc.Delegate = memb
	mc.Events   = memb
	/* mc.Conflict = memb */
	mc.Merge    = memb
	/* mc.Ping     = memb */
	mc.Alive    = memb
	
	if clust.NodeName!="" { mc.Name = clust.NodeName }
	if clust.BindAddr!="" { mc.BindAddr = clust.BindAddr }
	if clust.BindPort!=0  { mc.BindPort = clust.BindPort }
	if clust.AdvertiseAddr!="" { mc.AdvertiseAddr = clust.AdvertiseAddr }
	if clust.AdvertisePort!=0  { mc.AdvertisePort = clust.AdvertisePort }
	
	return
}
func CreateCluster(clust *config.BucketCluster, mlist *memberlist.Memberlist) (int, error) {
	ipv := ""
	switch clust.JoinPreferVer {
	case 4: ipv = "4"
	case 6: ipv = "6"
	}
	
	jnc := make([]string,0,len(clust.JoinCluster))
	for _,s := range clust.JoinCluster {
		addr,err := net.ResolveTCPAddr("tcp"+ipv,s)
		if err==nil { jnc = append(jnc,addr.String()) ; continue }
		addr,err  = net.ResolveTCPAddr("tcp",s)
		if err==nil { jnc = append(jnc,addr.String()) }
		/* The port is required! */
	}
	return mlist.Join(jnc)
}


func listenAndServcePrimary(memb *cluster.Membered, res chan <- int) {
	defer func(){res<-1}()
	memb.ListenAndServe()
}

func listenAndServceSecondary(memb *cluster.Membered, meta *cluster.MetaData, res chan <- int) {
	defer func(){res<-1}()
	las := cluster.ServerPlugins[meta.Proto]
	if las!=nil { las(meta,memb.Router.Handler) }
}

func ListenAndServe(cfg *config.BucketServiceCfg, memb *cluster.Membered) {
	targ := make(chan int,16)
	n := 0
	for i := range cfg.OuterSvc {
		m := new(cluster.MetaData)
		configureMetadata(&(cfg.OuterSvc[i]),m)
		n++
		go listenAndServceSecondary(memb,m,targ)
	}
	n++
	go listenAndServcePrimary(memb,targ)
	for ;n>0;n-- {
		<-targ
	}
}

