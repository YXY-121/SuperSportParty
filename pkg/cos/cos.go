package main

import (
	"context"
	"net/http"
	"net/url"
	"os"

	"github.com/tencentyun/cos-go-sdk-v5"
)

var CosClient *cos.Client

// SetUpCosClient 初始化cos客户端
func SetUpCosClient() {
	// 将 examplebucket-1250000000 和 COS_REGION 修改为真实的信息
	// 存储桶名称，由bucketname-appid 组成，appid必须填入，可以在COS控制台查看存储桶名称。https://console.cloud.tencent.com/cos5/bucket
	// COS_REGION 可以在控制台查看，https://console.cloud.tencent.com/cos5/bucket, 关于地域的详情见 https://cloud.tencent.com/document/product/436/6224
	u, _ := url.Parse("https://sport-1305086242.cos.ap-guangzhou.myqcloud.com")
	b := &cos.BaseURL{BucketURL: u}
	c := cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  "AKID98tPghTzGH3Zpqb7NO9fvvPedNr53mFS", // 替换为用户的 SecretId，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
			SecretKey: "0kESRAesra4IH0ThfNq30h9LI0Sfcknd",     // 替换为用户的 SecretKey，请登录访问管理控制台进行查看和管理，https://console.cloud.tencent.com/cam/capi
		},
	})
	CosClient = c
	// 对象键（Key）是对象在存储桶中的唯一标识。
	// 例如，在对象的访问域名 `examplebucket-1250000000.cos.COS_REGION.myqcloud.com/test/objectPut.go` 中，对象键为 test/objectPut.go

}
func UploadPictureToCos(fileName string) {
	name := "cos.go"

	// 2.通过本地文件上传对象
	_, err := CosClient.Object.PutFromFile(context.Background(), name, "cos.go", nil)
	if err != nil {
		panic(err)
	}
	// 3.通过文件流上传对象
	fd, err := os.Open("cos.go")
	if err != nil {
		panic(err)
	}
	defer fd.Close()
	_, err = CosClient.Object.Put(context.Background(), name, fd, nil)
	if err != nil {
		panic(err)
	}
}
