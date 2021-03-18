package models

import (
	"bytes"
	"crypto/des"
	"encoding/hex"
	"encoding/json"
	"errors"
	"github.com/minio/minio-go/v6"
	"math/rand"
	"time"
  "io"
	"fmt"
	"net/url"
)

func CurTime() int64 {
	return time.Now().Unix()
}

func UploadImgToMinio(name string,file io.Reader) string {
	minioClient, err := minio.New("minio.soca.ink", "AKIAIOSFODCC7EXAMPLE", "wJalrXUtnFEMI/K88DENG/bPxRfiCYEXAMPLELOY", false)
	if err != nil {
		fmt.Println(err)
		return ""
	}
  _, err = minioClient.PutObject("zcxt", name, file, -1,minio.PutObjectOptions{ContentType:"image/" + name[len(name)-3:],})
  if err != nil {
		fmt.Print("49494")
		fmt.Println(err)
    return ""
  }
	fmt.Print("494")
  return name
}

func DelImgInMinio(name string) {
	minioClient, _ := minio.New("minio.soca.ink", "AKIAIOSFODCC7EXAMPLE", "wJalrXUtnFEMI/K88DENG/bPxRfiCYEXAMPLELOY", false)
	minioClient.RemoveObject("zcxt", name)
}

func GetImgUrlInMinio(name string) string {
	minioClient, _ := minio.New("minio.soca.ink", "AKIAIOSFODCC7EXAMPLE", "wJalrXUtnFEMI/K88DENG/bPxRfiCYEXAMPLELOY", false)
	url, _ := minioClient.PresignedGetObject("zcxt", name, time.Second * 24 * 60 * 60, make(url.Values))
	return url.String()
}

/*
 *   获取cookie,使用DES解密
 *   用法:id := models.GetCookie(this.Ctx.GetCookie("token"),"id")
 */
func GetCookie(token, key string) string {
	keys := []byte("12345678")
	de, _ := DesDecrypt(token, keys)
	var mapResult map[string]string
	json.Unmarshal([]byte(de), &mapResult)
	return mapResult[key]
}

/*
 *   设置cookie,使用DES加密  604800为秒,cookie存活时间
 *   用法:this.Ctx.SetCookie("token", models.SetCookie(map[string]string{"id": "123"}), 604800)
 */
func SetCookie(ctx map[string]string) string {
	b, _ := json.Marshal(ctx)
	keys := []byte("12345678")
	en, _ := DesEncrypt(string(b), keys)
	return string(en)
}

//随机创建30位字符
func RandomCreater() string {
	bytes := []byte("0123456789abcdefghijklmnopqrstuvwxyz")
	result := []byte{}
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < 30; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}
	return string(result)
}

//状态码函数
func Status(stateCode int, resData ...interface{}) map[string]interface{} {
	StatusToMessage := make(map[int]string)
	StatusToMessage[200] = "成功"
	StatusToMessage[141] = "查询数据库出错"
	StatusToMessage[142] = "结果不唯一"
	StatusToMessage[131] = "非法session"
	StatusToMessage[132] = "参数异常132"
	StatusToMessage[133] = "参数异常133"
	StatusToMessage[134] = "参数异常134"
	StatusToMessage[135] = "未登录"

	output := make(map[string]interface{})
	output["status"] = stateCode
	output["message"] = StatusToMessage[stateCode]

	for _, item := range resData {
		output["data"] = item
		return output
	}

	output["data"] = nil
	return output
}

//DES加密解密函数
func ZeroPadding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{0}, padding)
	return append(ciphertext, padtext...)
}

func ZeroUnPadding(origData []byte) []byte {
	return bytes.TrimFunc(origData,
		func(r rune) bool {
			return r == rune(0)
		})
}

func DesEncrypt(text string, key []byte) (string, error) {
	src := []byte(text)
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	bs := block.BlockSize()
	src = ZeroPadding(src, bs)
	if len(src)%bs != 0 {
		return "", errors.New("Need a multiple of the blocksize")
	}
	out := make([]byte, len(src))
	dst := out
	for len(src) > 0 {
		block.Encrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	return hex.EncodeToString(out), nil
}

func DesDecrypt(decrypted string, key []byte) (string, error) {
	src, err := hex.DecodeString(decrypted)
	if err != nil {
		return "", err
	}
	block, err := des.NewCipher(key)
	if err != nil {
		return "", err
	}
	out := make([]byte, len(src))
	dst := out
	bs := block.BlockSize()
	if len(src)%bs != 0 {
		return "", errors.New("crypto/cipher: input not full blocks")
	}
	for len(src) > 0 {
		block.Decrypt(dst, src[:bs])
		src = src[bs:]
		dst = dst[bs:]
	}
	out = ZeroUnPadding(out)
	return string(out), nil
}
