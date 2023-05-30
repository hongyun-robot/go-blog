package article

import (
	"fmt"
	"front-gin/rpc/server/svc"
	article "front-gin/src/server/article/detail"
	classify "front-gin/src/server/classify/list"
	"front-gin/src/types"
	"github.com/gin-gonic/gin"
	"github.com/zeromicro/go-zero/core/timex"
	"net/http"
	"strconv"
)

func InitArticle(r *gin.Engine, serverCtx *svc.ServiceContext) {
	r.LoadHTMLGlob("templates/view/**/*")
	r.StaticFS("/static", http.Dir("templates/static"))
	router := r.Group("/article")

	list(router, serverCtx)
	detail(router, serverCtx)
}

func list(r *gin.RouterGroup, serverCtx *svc.ServiceContext) {
	ctx := &http.Request{}
	l := classify.NewGetClassifyLogic(ctx.Context(), serverCtx)
	req := &types.GetClassifyReq{
		ID: -1,
	}
	data, err := l.GetClassify(req)
	if err != nil {
		fmt.Println("请求rpc错误", err)
		return
	}

	//for _, datum := range data.Data {
	//	unsafe := blackfriday.MarkdownCommon([]byte(datum.Content))
	//	datum.Content = string(bluemonday.UGCPolicy().SanitizeBytes(unsafe))
	//}

	r.GET("/list", func(c *gin.Context) {
		c.HTML(http.StatusOK, "article/list", gin.H{
			"code":  data.Code,
			"data":  data.Data,
			"title": "文章列表",
			"css":   [1]string{"/static/css/list/index.css?timestamp=" + strconv.Itoa(int(timex.Now().Microseconds()))},
			"js":    [1]string{"/static/js/list/index.js?timestamp=" + strconv.Itoa(int(timex.Now().Microseconds()))},
		})
	})
}

func detail(r *gin.RouterGroup, serverCtx *svc.ServiceContext) {
	r.GET("/detail/:id", func(c *gin.Context) {
		id := c.Params[0].Value
		ctx := &http.Request{}
		l := article.NewGetArticleLogic(ctx.Context(), serverCtx)
		atoi, err := strconv.Atoi(id)
		if err != nil {
			return
		}
		req := &types.GetArticleReq{ID: uint64(atoi)}
		data, err := l.GetArticle(req)
		if err != nil {
			return
		}
		if data.Message != "SUCCESS" {
			temp := &types.ArticleData{
				ID:           0,
				Title:        data.Message,
				Content:      "> " + data.Message,
				CreateAt:     0,
				UpdateAt:     0,
				DeleteAt:     0,
				Status:       0,
				Draft:        false,
				ClassifyData: nil,
			}
			data.Data = append(data.Data, temp)
		}
		c.HTML(http.StatusOK, "article/detail", gin.H{
			"title": data.Data[0].Title,
			"id":    id,
			"data":  data.Data[0],
		})
	})
}