package v1

import (
	"apiproject/model"
	"apiproject/pkg"
	"apiproject/pkg/page"
	"apiproject/service"
	"context"
	"net/http"
	"net/url"

	"apiproject/repository"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/tencentyun/cos-go-sdk-v5"
)

var orderService = service.NewOrderService()

type OrderTypeReq struct {
	OrderType string `json:"sport_type"`
}
type OrderLocationReq struct {
	OrderLocation string `json:"location"`
}
type OrderLocatonReq struct {
	Longitude *float64 `json:"longitude"`
	Latitude  *float64 `json:"latitude"`
}

type OrderID struct {
	OrderId string `json:"order_id"`
}
type ResponseRep struct {
	page.Page
	Value interface{}
}

func ListOrders(c *gin.Context) {
	// 序列化 请求方http body
	var req model.OrderReq
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	//
	var orders []model.Order
	if req.Latitude != nil || req.Longitude != nil {
		orders, err = repository.ListOrdersByLongitudeAndLatitud(*req.Longitude, *req.Latitude)
	} else {
		orders, err = repository.ListOrders()

	}
	if err != nil {
		pkg.Response(c, pkg.ErrorListFail, orders)
		return
	}
	flag, start, end := page.PageTool(len(orders), req.PageSize, req.PageNo)
	if !flag {
		logrus.Println("分页失败")
		pkg.Response(c, pkg.ErrorPageFail, orders)
		return
	}
	req.Page.Total = len(orders)
	pkg.Response(c, pkg.Success, ResponseRep{req.Page, orders[start:end]})
	return

}
func GetOrderById(c *gin.Context) {
	// 反序列化 请求方http body

	var orderId OrderID
	err := c.ShouldBind(&orderId)
	fmt.Println("req", orderId)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return

	}
	order := orderService.GetOrderById(orderId.OrderId)
	pkg.Response(c, pkg.Success, order)
	return

}

//根据地区以及喜欢的类型进行定时推送
func GetOrdersByLoacation(c *gin.Context) {
	var req model.OrderReq
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}
	//data := repository.GetOrderByTypeAndLocation(req.OrderType)

}

func UpdateOrder(c *gin.Context) {
	//这是json
	var order model.OrderReq
	err := c.ShouldBind(&order)
	logrus.Println(order)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	urls := UploadPicture(c)
	if urls != nil {
		order.PicUrls = urls[0]
		fmt.Println("上传成功url", order.PicUrls)
	}

	orderService.CreateOrder(order)

	pkg.Response(c, pkg.Success, urls)
	return

}

type GrabOrder struct {
	OrderId string
	UserId  string
}

//抢单
func GrabOrder1(c *gin.Context) {
	// grabOrder := GrabOrder{}
	// fmt.Println(o.Ctx.Input.RequestBody)
	// //这是json
	// json.Unmarshal(o.Ctx.Input.RequestBody, &grabOrder)
	// fmt.Println("g", grabOrder)
	// if (grabOrder == GrabOrder{}) {
	// 	o.Data["json"] = nil
	// 	o.ServeJSON()

	// }
	// //createrId string,orderLocation string,sportType string,description string,peopleNumber int,longitude,latitude string ,endTime time.Time)
	// orderService.GrabOrder(grabOrder.UserId, grabOrder.OrderId)
	// o.Data["json"] = nil
	// o.ServeJSON()
}

//删除订单
func DeleteOrder(c *gin.Context) {
	var order OrderID
	err := c.ShouldBind(&order)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})

		return
	}
	err = repository.DeleteOrder(order.OrderId)
	if err != nil {
		pkg.Response(c, pkg.ErrorDeleteFail, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})
		return
	}

	pkg.Response(c, pkg.Success, map[string]string{
		"ok": fmt.Sprintf("%v", err),
	})

}
func UploadPicture(c *gin.Context) []string {

	multiForm, err := c.MultipartForm()

	if err != nil {
		pkg.Response(c, pkg.ErrorSaveImageFail, err)
		return nil
	}
	files := multiForm.File["file"]
	urls := make([]string, 0)
	for _, file := range files {
		filePath := "/data/SuperSportParty/temp/" + file.Filename
		if err = c.SaveUploadedFile(file, filePath); err != nil {
			if err != nil {
				pkg.Response(c, pkg.ErrorSaveImageFail, map[string]string{
					"error": fmt.Sprintf("%v", err),
				})
				return nil

			}
		}
		url := UploadTencentCloud(filePath, file.Filename)
		urls = append(urls, url)

	}

	//fmt.Println(filePath)
	//

	return urls
}

func UploadTencentCloud(filesPath string, name string) string {
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

	_, err := c.Object.PutFromFile(context.Background(), name, filesPath, nil)
	if err != nil {
		panic(err)
	}
	key := name
	url := c.Object.GetObjectURL(key)
	myurl := url.Scheme + "://" + url.Host + url.Path
	return myurl
}

func ListUserSportRecord(c *gin.Context) {
	var req UserId
	err := c.ShouldBind(&req)
	if err != nil {
		logrus.Errorf("json bind error!")
		pkg.Response(c, pkg.ErrorParams, map[string]string{
			"error": fmt.Sprintf("%v", err),
		})

		return
	}

}
