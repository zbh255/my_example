package main

import (
	"bytes"
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
	"strings"
)



func GinFileUpload() {
	r := gin.Default()
	r.POST("/file",click)
	_ = r.Run("127.0.0.1:8080")
}

func click(ctx *gin.Context) {
	// gin将上传的文件流式储存
	file, _ := ctx.FormFile("file")
	println(file.Filename)
	if file.Size > 0 && file.Filename != "" {
		ctx.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": file.Filename,
		})
	}
	src,_ := file.Open()
	println(src)
	buf := bytes.NewBuffer(nil)
	_, err := io.Copy(buf, src)
	if err != nil {
		panic(err)
	}
	//fmt.Println(buf.String())
	bufStr := strings.NewReader(buf.String())
	u,_ := url.Parse("https://examplebucket-1250000000.cos.COS_REGION.myqcloud.com")
	// 用于Get Service 查询，默认全地域 service.cos.myqcloud.com
	su, _ := url.Parse("https://cos.COS_REGION.myqcloud.com")

	b := cos.BaseURL{BucketURL: u,ServiceURL: su}

	// 永久密钥
	client := cos.NewClient(&b,&http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID: "",
			SecretKey: "",
		},
	})

	if client != nil {
		// 调用cos请求
		s,_,err := client.Service.Get(context.Background())
		if err != nil {
			panic(err)
		}
		for _,b := range s.Buckets {
			fmt.Printf("%#v\n", b)
		}

		// 将文件上传到桶中
		_,err2 := client.Object.Put(context.Background(),file.Filename, bufStr,nil)
		if err2 != nil {
			panic(err)
		}
	}
}
func main() {
	GinFileUpload()
}