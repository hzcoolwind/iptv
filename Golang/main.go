// Package Golang
// @Time:2023/02/03 02:28
// @File:main.go
// @SoftWare:Goland
// @Author:feiyang
// @Contact:TG@feiyangdigital

package main

import (
	"Golang/liveurls"
	"github.com/gin-gonic/gin"
	"log"
//	"log/syslog"
	"net/http"
)

var logger *log.Logger = log.Default()

func duanyan(adurl string, realurl any) string {
	var liveurl string
	if str, ok := realurl.(string); ok {
		liveurl = str
	} else {
		liveurl = adurl
	}
	return liveurl
}

func setupRouter(adurl string) *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.GET("/douyin", func(c *gin.Context) {
		url := c.Query("url")
		douyinobj := &liveurls.Douyin{}
		douyinobj.Shorturl = url
		c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyinobj.GetRealurl()))
	})

	r.GET("/:path/:rid", func(c *gin.Context) {
		path := c.Param("path")
		rid := c.Param("rid")
		switch path {
		case "douyin":
			douyinobj := &liveurls.Douyin{}
			douyinobj.Rid = rid
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyinobj.GetDouYinUrl()))
		case "douyu":
			douyuobj := &liveurls.Douyu{}
			douyuobj.Rid = rid
			douyuobj.Stream_type = c.DefaultQuery("stream", "hls")
			douyuobj.Cdn_type = c.DefaultQuery("cdn", "akm-tct")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, douyuobj.GetRealUrl()))
		case "huya":
			huyaobj := &liveurls.Huya{}
			huyaobj.Rid = rid
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, huyaobj.GetLiveUrl()))
		case "bilibili":
			biliobj := &liveurls.BiliBili{}
			biliobj.Rid = rid
			biliobj.Platform = c.DefaultQuery("platform", "web")
			biliobj.Quality = c.DefaultQuery("quality", "10000")
			biliobj.Line = c.DefaultQuery("line", "second")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, biliobj.GetPlayUrl()))
		case "youtube":
			ytbObj := &liveurls.Youtube{}
			ytbObj.Rid = rid
			ytbObj.Quality = c.DefaultQuery("quality", "1080")
			c.Redirect(http.StatusMovedPermanently, duanyan(adurl, ytbObj.GetLiveUrl()))
		}
	})
	return r
}

func main() {
	//key := []byte("6354127897263145")
	//defstr, _ := base64.StdEncoding.DecodeString("Mf5ZVkSUHH5xC9fH2Sao+2LgjRfydmzMgHNrVYX4AcSoI0nktkV7z1jSU6nSihf7ny+PexV73YjDoEtG7qu+Cw==")
	//defurl, _ := openssl.AesECBDecrypt(defstr, key, openssl.PKCS7_PADDING)
	//sysLog, err := syslog.Dial("tcp", "localhost:1234",
	//	syslog.LOG_WARNING|syslog.LOG_DAEMON, "gom3u8")
	//sysLog, err := syslog.New(syslog.LOG_SYSLOG|syslog.LOG_WARNING, "gom3u8")
    //	if err != nil {
	//	logger.Print("syslog is error.")
	//}
	r := setupRouter(string("http://10.10.10.207:8123/local/test.m3u"))
	logger.Print("run on 35455.")
	//sysLog.Info("listen on 35455...")
	r.Run(":35455")
}
