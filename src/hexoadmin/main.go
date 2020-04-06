package main

import (
	"github.com/gogf/gf/frame/g"
	"github.com/gogf/gf/net/ghttp"
	"gowww/hexoadmin/articleutil"
	"gowww/hexoadmin/param"
)





func main() {

	s := g.Server()
	s.BindHandler("/createArticle", func(r *ghttp.Request) {
		params := new(param.ReqUrlParam)
		r.GetToStruct(params)

		//发布一篇文章
		articleutil.PostArticle(params)

	})
	s.SetPort(5050)
	s.Run()
}

