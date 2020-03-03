package dbgr

import "fmt"
import "github.com/maxymania/fastnntp-polyglot/gold"
import "github.com/maxymania/fastnntp-polyglot/postauth"
import "github.com/maxymania/fastnntp-polyglot"

type PostingDebuggerDirect struct {
	gold.ArticleDirectEX
}
func (pdd PostingDebuggerDirect) ArticleDirectStore(exp uint64, ov *newspolyglot.ArticleOverview, obj *newspolyglot.ArticleObject) (err error) {
	err = pdd.ArticleDirectEX.ArticleDirectStore(exp,ov,obj)
	fmt.Println("ArticleDirectStore",err)
	return
}

type PostingDebuggerGroup struct {
	gold.ArticleGroupEX
}
func (pdg PostingDebuggerGroup) StoreArticleInfos(groups [][]byte, nums []int64, exp uint64, ov *newspolyglot.ArticleOverview) (err error) {
	err = pdg.ArticleGroupEX.StoreArticleInfos(groups,nums,exp,ov)
	fmt.Println("StoreArticleInfos",err)
	return
}

type GroupDebugger struct {
	gold.GroupListDB
}
func (gd GroupDebugger) GroupHeadFilterWithAuth(rank postauth.AuthRank, groups [][]byte) (res [][]byte, err error) {
	res,err = gd.GroupListDB.GroupHeadFilterWithAuth(rank,groups)
	fmt.Printf("GroupHeadFilterWithAuth (%v,%q) -> %q %v\n",rank,groups,res,err)
	return
}

