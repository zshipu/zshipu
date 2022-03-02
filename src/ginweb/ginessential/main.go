package main

import (
	"github.com/gin-gonic/gin"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	r := gin.Default()
	r.POST("/api/auth/register", func(ctx *gin.Context) {
		// 获取参数
		name := ctx.PostForm("name")
		mobile := ctx.PostForm("mobile")
		passwd := ctx.PostForm("passwd")

		//数据验证
		if len(mobile) != 11 {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"手机号码必须11位"})
			return
		}
		if len(passwd) < 6 {
			ctx.JSON(http.StatusUnprocessableEntity,gin.H{"code":422,"msg":"密码不能少于6位"})
			return
		}
		if len(name) ==0 {
			name = RandomString(10)
		}

		log.Println(name,mobile,passwd)
		//判断手机号码


		//创建用户


		//返回结果
		ctx.JSON(200, gin.H{
			"msg": "注册成功",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}

func RandomString(n int) string {
	var letters = []byte("sjgaf4534giuyuyuagfsdfsffgjhljhilkjhoiu46436542342136890987jskhjdfster65w5dhruip3454jdsrekfhjrueg")
	resutl := make([]byte,n)
	rand.Seed(time.Now().Unix())
	for i:= range resutl{
		resutl[i] = letters[rand.Intn(len(letters))]
	}
	return string(resutl)
}