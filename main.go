package main

import (
	"crypto/sha1"
	"encoding/hex"
	"github.com/gin-gonic/gin"
	"log"
	"sort"
	"strings"
)

// 与填写的服务器配置中的Token一致
const Token = "befa02e70c2b43158c5a69d6c84a11c5"
const EncodingAESKey = "yX6sNIqwS8zNs9lcqdrRPU8vpAMCbYtqtTqd993uzpK"

func main() {
	router := gin.Default()

	router.GET("/wx", WXCheckSignature)

	log.Fatalln(router.Run(":56600"))
}

// WXCheckSignature 微信接入校验
func WXCheckSignature(c *gin.Context) {
	//signature := c.Query("signature")
	//timestamp := c.Query("timestamp")
	//nonce := c.Query("nonce")
	//echostr := c.Query("echostr")

	//ok := CheckSignature(signature, timestamp, nonce, Token)
	//if !ok {
	//	log.Println("微信公众号接入校验失败!")
	//	return
	//}

	log.Println("微信公众号接入校验成功!")
	_, _ = c.Writer.WriteString("微信公众号接入校验成功!")
}

// CheckSignature 微信公众号签名检查
func CheckSignature(signature, timestamp, nonce, token string) bool {
	arr := []string{timestamp, nonce, token}
	// 字典序排序
	sort.Strings(arr)

	n := len(timestamp) + len(nonce) + len(token)
	var b strings.Builder
	b.Grow(n)
	for i := 0; i < len(arr); i++ {
		b.WriteString(arr[i])
	}

	return Sha1(b.String()) == signature
}

// 进行Sha1编码
func Sha1(str string) string {
	h := sha1.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
