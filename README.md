# 产品简介
<p align="center">
<img align="center" width="150px" src="https://img-blog.csdnimg.cn/img_convert/6e6a82f5804d5c06542c925ee99c1d78.png#pic_center">
</p>

Crypto-m is an easy-used hybrid-encryption middleware manager which used for go-zero framework。

Crypto-m是一个基于go-zero框架的通信加、解密中间件管理工具，使用不对称加密和对称加密的混合加密策略（hybrid-encryption）。本包设计清晰简洁，使用方便，功能完善，完美契合go-zero原生体系，欢迎广大新老同行选用。
（名称中的后缀-m，既代表中间件middleware，也代表manager，还代表可手动manual。）

# 相关链接
源码地址：[https://github.com/chenkaiwei/crypto-m](https://github.com/chenkaiwei/crypto-m)

教程示例：[https://github.com/chenkaiwei/crypto-m-samples](https://github.com/chenkaiwei/crypto-m-samples)

# 策略简述
**混合加密策略（hybrid-encryption）**，是当今非常主流的通信加密策略。即由客户端生成一个随机密钥(cek)并以该密钥对称加密消息内容(content)，继而以不对称加密方式加密cek，并将加密后的cek与content一并发往服务端。该加密方式博采众长，兼具不对称加密的安全性和对称加密的高效性，只须保证客户端持有的公钥不被替换（公钥不怕泄露）以及服务端持有的对应私钥没有泄漏，即可非常可靠且高效地保障通信安全。

本产品提供一套非常便捷的工具帮助您在go-zero项目中引入对该策略的支持，仅需几步简单配置即可完成全套通信加解密体系的搭建

# quick start
### 服务端：
1. 拉取包到本地
   > go get github.com/chenkaiwei/crypto-m
2. 在api文件中按go-zero的规则加入中间件的定义（全套代码详见simpleDemo目录）
   ```api
   type (
    SimpleMsg {
        Message string `json:"msg,optional"`
    }
    )
    
    //⬇️中间件名称可自行定义，初次使用建议和demo保持一致
    @server(
    middleware: CryptionRequest,CryptionResponse
    )
    service simpleDemo-api {
    
        @handler CryptionTest
        post /cryptionTest (SimpleMsg) returns (SimpleMsg)
    }
   ```
3. 删除自动生成的空白中间件

   ![img.png](https://img-blog.csdnimg.cn/img_convert/2c79375a4432e10e196e81c3e9d1a8e6.png)
4. 关键步骤：在serviceContext.go文件中配置cryptom.cryptomManager

   ```go
   package svc
   
   import (
   "github.com/chenkaiwei/crypto-m/cryptom"
   "simpleDemo/internal/config"
   "github.com/zeromicro/go-zero/rest"
   )
   
   const PRIVATE_KEY = "MIICWwIBAAKBgQC9G7U67pfDHfAfUvMkZ2uBPyEbvqYrlr5xEr2zvcoKxXfdfh4+n//ycD2wTE4EkwwAJiVAqaT3s1KilhkMf4RnB5sE1Fj0Aq7n+8tYsmTdK95BIEeSGuO2qZIni5S7EaAZVSQiv8HGbedlteAv2Ja1XRZsnIoe0E9qQBOCAhpB6QIDAQABAoGAV/Kd62VxISY4OWkreP+8CKTicfPNdjIqKY4suX4Hi9DgeRshV8CzmP3IQsiJ9CirCRq0cokzFpvIT6L8zUo0uaRJdaA7cpo2iJLUtoOHc6zdUHdBHBaVEa9ymjEqx2wrIw5kkcy4SBe472DWJDtphgig9S3THww9k4jzn0L6DL0CQQDxsfPRm3nwOj8qaVHBJCYcJ0+rdXL9a7UIYomUBEF6Z+DRlG6iYI5ZXSJZ+Sxhp20QawVk47JlMOGGRM7o2ApbAkEAyEz7/5kkp6IZ3lC4yRrTD86eKjJ9yzSXiOBQiVmW54lluxvK6o4sXP/+hVb9zqiD5AYOwjFbLiCTtSLZgv5wCwJAZSuxPPdQ5p7rG+y0HR3tmfFWpxXlyXDRea4NmtjhM8TR1cjFOtEiJQQYQgNMcaAsxieWPXIWlccNUC/zUIJGawJAKpudu3hjQLmN0SnQtQ7cuO8V3BoTgkd0uKwm1aDWJfinSE8YMh7+NuZJySmBIhXcwIO9XffL0pshcJWyOVhQkwJANTQejTHOgu7bPkGe5Zg+HfHQGMrimffbWckXMTDQpA1y5l5fq7+B7cauSYoNcQZ41Bp38USCzsxkwTn5659iFA=="
   
   type ServiceContext struct {
   Config           config.Config
   CryptionRequest  rest.Middleware
   CryptionResponse rest.Middleware
   }
   
   func NewServiceContext(c config.Config) *ServiceContext {
   

        //针对cek的不对称加密策略
        cekAlgo := algom.NewCekAlgoRsaBase64(PRIVATE_KEY)
        //针对消息内容的对称加密策略
        contentAlgo := algom.NewContentAlgoAesCbcHex([]byte("1111222233334444"))
        //组装成manager
        cryptomManager := cryptom.NewStandardCryptomManager(cekAlgo, contentAlgo)
   
       return &ServiceContext{
           Config:           c,
           CryptionRequest:  cryptomManager.RequestHandle,
           CryptionResponse: cryptomManager.ResponseHandle,
       }
   }

   ```
   其中cryptomManager.RequestHandle功能为对所有应用该中间件的接口的请求消息体进行解密；cryptomManager.ResponseHandle的功能为对所有应用该中间件的接口的响应消息体进行加密。

   > *至此crypto-m的配置和使用已经完成

5. 按照业务需要正常完成逻辑代码功能等
> (e.g.⬇️simpleDemo/internal/logic/cryptionTestLogic.go)
   ```go
   func (l *CryptionTestLogic) CryptionTest(req *types.SimpleMsg) (resp *types.SimpleMsg, err error) {
resp = &types.SimpleMsg{
Message: "加密信息已收到，内容为--" + req.Message,
}
return
}
   ```
> ⬆️由于中间件的配置，在logic代码中的req、resp即为请求解密后、响应加密前的内容，直接按明文使用即可。

至此服务端的配置就已全部完成，以下是对客户端的开发要求
### 客户端：
客户端的示例代码详见samples/client_demo/simple_demo_test.go
这里用go程序作为演示
```go
package client_demo

import (
   "bytes"
   "fmt"
   "github.com/wumansgy/goEncrypt/aes"
   "github.com/wumansgy/goEncrypt/rsa"
   "io/ioutil"
   "net/http"
   "testing"
)
func TestSimpleDemo(t *testing.T) {

   aesIv := []byte("1111222233334444")

   cek := newRandomCek(16)
   fmt.Println("新生成随机cek--", cek)

   //==请求服务器阶段：
   client := &http.Client{}
   url := "http://localhost:11111/cryptionTest"

   //组装消息体
   s := "{\"msg\":\"来自TestSimpleDemo的加密测试数据\"}"
   eContent, err := aes.AesCbcEncryptHex([]byte(s), []byte(cek), aesIv)
   if err != nil {
      return
   }
   request, err := http.NewRequest("POST", url, bytes.NewBuffer([]byte(eContent)))
   if err != nil {
      t.Fatal(err)
   }
   //组装消息头
   eCek, err := rsa.RsaEncryptToBase64([]byte(cek), PUB_KEY)
   if err != nil {
      t.Fatal(err)
   }
   request.Header.Set("ECEK", eCek)

   request.Header.Set("Content-Type", "application/json")

   response, _ := client.Do(request)

   //处理响应阶段：
   defer response.Body.Close() //

   status := response.StatusCode
   fmt.Printf("响应状态码--%v\n", status)
   body, err := ioutil.ReadAll(response.Body)
   if err != nil {
      t.Fatal(err)
   }

   fmt.Println("响应消息体--", string(body))
   //解密响应消息体（服务端采用ResponseHandle中间件策略时）
   decryptedRespMsg, err := aes.AesCbcDecryptByHex(string(body), []byte(cek), aesIv)
   if err != nil {
      return
   }
   fmt.Println("响应消息体（解密后）--", string(decryptedRespMsg))

}

```
具体策略简述如下：

1. 准备好与服务端事先约定的不对成加密的公钥和对称加密的vi，以及待发送的消息内容content
2. 随机生成一个仅供当次使用的CEK(内容密钥｜content encryption key)，注意应符合对应内容加密策略的长度要求。
3. 以cek和与服务端约定的内容加密（对称加密）策略加密正文（content）
4. 以不对称加密策略的公钥加密cek，加密后的结果为ecek
5. 将ecek以ECEK为名存入消息头中，并将头信息Content-Type更改为application/json
6. 发送请求
7. 获得响应后同样按照内容对称加密的策略、先前生成的cek，解密所返回的响应消息体。至此加解密通信环节完成。
8. 对解密后的消息体进行正常的业务逻辑处理。

# 进阶用法
在 samples/standardDemo项目中你可以找到本工具的几项进阶用法。具体方向包括：
### 1. 自选其他加密策略
e.g.(svc/serviceContext.go)
   ```go
   cryptom.NewStandardCryptomManager(algom.NewCekAlgoRsaBase64(PRIVATE_KEY), algom.NewContentAlgoTripleDesBase64(nil))
   ```
其中NewContentAlgoTripleDesBase64即以三次des的加解密策略来处理正文（content）数据。
### 2. 多套加解密策略共用
e.g.(svc/serviceContext.go)
   ```go
   func NewServiceContext(c config.Config) *ServiceContext {

multi1CryptomManager := cryptom.NewStandardCryptomManager(algom.NewCekAlgoRsaBase64(PRIVATE_KEY), algom.NewContentAlgoTripleDesBase64(nil))
multi2CryptomManager := cryptom.NewStandardCryptomManager(algom.NewCekAlgoRsaBase64(PRIVATE_KEY), algom.NewContentAlgoAesCtrHex(nil))

return &ServiceContext{
Config: c,

//middleware
Multi1Request:  multi1CryptomManager.RequestHandle,
Multi1Response: multi1CryptomManager.ResponseHandle,
Multi2Request:  multi2CryptomManager.RequestHandle,
}
}
   ```
示例中配置了multi1和multi2两套加解密策略，而其各自对应的接口分配则依然在api中设置
(samples/standardDemo/standardDemo.api)
   ```api
   @server(
   middleware: Multi1Request,Multi1Response
)
service standardDemo-api {

   @handler MultiDemo1
   post /multiDemo1 (SimpleMsg) returns (SimpleMsg)
}

@server(
   //自由组合，仅发送时加密
   middleware: Multi2Request
)
service standardDemo-api {

   @handler MultiDemo2
   post /multiDemo2 (SimpleMsg) returns (SimpleMsg)
}  
   ```
以上MultiDemo2的中间件配置同时也是令中间件仅对请求加密/仅对响应解密的配置演示
> *对应客户端示例见samples/client_demo/standard_demo_test.go，和simple差别不大故此不加赘述。
### 3. 手动选择部分字段或内容加解密
   ```go
   func NewServiceContext(c config.Config) *ServiceContext {

multi1CryptomManager := cryptom.NewStandardCryptomManager(algom.NewCekAlgoRsaBase64(PRIVATE_KEY), algom.NewContentAlgoTripleDesBase64(nil))

return &ServiceContext{
Config: c,

//manager|client
Multi1CryptomManager: multi1CryptomManager, //供手动加解密（Manual）模式使用

//middleware
Muti1Manual:    multi1CryptomManager.ManualHandle,
}
}
   ```
手动加密时，须将对应cryptomManager对象也添加进ServiceContext（即上述代码中的Multi1CryptomManager），供logic部分手动加解密使用。示例如下
(samples/standardDemo/internal/logic/manualDemoLogic.go)：
   ```go
   func (l *ManualDemoLogic) ManualDemo(req *types.StandardMsg) (resp *types.StandardMsg, err error) {

cryptomManager := l.svcCtx.Multi1CryptomManager
//手动解密
decryptCustomMessage, err := cryptomManager.ContentDecrypt(l.ctx, req.EncryptedMessage)
if err != nil {

//复合型错误的判断示例
var ce *cryptom.CryptomError
if errors.As(err, &ce) && ce.ErrType() == cryptom.ErrTypeContentDecryptFailure {
logx.Error("ErrType() == cryptom.ErrTypeContentDecryptFailure")
}

return nil, err
}

//手动加密
encryptMsg, err := l.svcCtx.Multi1CryptomManager.ContentEncrypt(l.ctx, "来自服务端的加密消息4321")
if err != nil {
return nil, err
}

resp = &types.StandardMsg{
Message:          "your message is revieved --" + req.Message,
EncryptedMessage: encryptMsg,
DecryptedMessage: "您发送的加密数据为：" + decryptCustomMessage,
}

return
}
   ```
> *对应客户端演示代码亦见samples/client_demo/standard_demo_test.go

### 4. 自定义加解密策略

这个用法的示例写在 samples/customDemo 里，核心代码参考samples/customDemo/internal/cryptomx/myContentAlgo.go，也即自己实现一个ContentAlgo接口。
   ```go
   package cryptomx

import (
   "github.com/pkg/errors"
)

type MyCaesarAlgo struct {
}

/*
   本演示用的自定义加密策略，古老的凯撒加密法：
   cek为偏移量，加密过程为每个字节byte在原数字的基础上加上cek，超出2^8时循环回1计数
*/

func (ca *MyCaesarAlgo) Encrypt(data []byte, cek []byte) (res string, err error) {
   offset := int(cek[0])
   if err != nil || offset < 0 || offset > 255 {
      return "", errors.WithMessage(err, "cek不合法，请确保其为一个0~255范围内的数字")
   }

   for i, datum := range data {
      sum := int(datum) + offset
      if sum >= 256 {
         sum = sum - 256
      }
      data[i] = byte(sum)
   }
   res = string(data)
   return
}

func (ca *MyCaesarAlgo) Decrypt(s string, cek []byte) (res []byte, err error) {

   offset := int(cek[0])
   if err != nil || offset < 0 || offset > 255 {
      return nil, errors.WithMessage(err, "cek不合法，请确保其为一个0~255范围内的数字")
   }
   data := []byte(s)
   for i, datum := range data {
      sum := int(datum) - offset
      if sum < 0 {
         sum = sum + 256
      }
      data[i] = byte(sum)
   }
   res = data
   return
}

   ```
这里我手搓了一套凯撒加密法（的变体）作为演示，具体做法为将每一个字节作为数字进行一定量的偏移（原版是26个字母按照次序偏移）。
由于不对称加密过于高端我辈手搓不能，就依然沿用一家独大的Rsa策略了，然后原本用作密钥的cek就由凯撒加密法的偏移量代替，反正从功能上看也算是一种密钥了，二者结合依然是个足够凑合的混合加密策略。
> *对应客户端示例 samples/client_demo/custom_demo_test.go

# 交流建议
CSDN的介绍贴的评论、私信，或者github的issue都可以。
QQ群：784219974
