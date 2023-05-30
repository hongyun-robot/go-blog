package router

import (
	"front-gin/router/article"
	"front-gin/router/home"
	"front-gin/rpc/server/svc"
	"github.com/gin-gonic/gin"
	"github.com/microcosm-cc/bluemonday"
	"github.com/russross/blackfriday"
	"html/template"
	"time"
)

// html 字符串转为 html
// "<h1>111</h1>" => <h1>111</h1>
func transition(str string) template.HTML {
	return template.HTML(str)
}

// markdown 转为 html
// # 111 => <h1>111</h1>
func markdownToHtml(str string) template.HTML {
	unsafe := blackfriday.MarkdownCommon([]byte(str))
	htmlStr := string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
	return transition(htmlStr)
}

// 时间戳转为字符串 YYYY-MM-DD
func timeStampFormat(stamp int64) string {
	t := time.Unix(int64(stamp/1000), 0)
	return t.Format("2006-01-02")
}

func InitRouter(serverCtx *svc.ServiceContext) *gin.Engine {
	r := gin.New()
	r.SetFuncMap(template.FuncMap{
		"transition":      transition,
		"markdownToHtml":  markdownToHtml,
		"timeStampFormat": timeStampFormat,
	})
	r.Use(gin.Logger())
	r.Use(gin.Recovery())
	home.InitHome(r)
	article.InitArticle(r, serverCtx)
	return r
}