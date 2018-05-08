/*
*   @author SJC
*   @工具集，包含md5，sha1，AES加密解密算法，以及签名和校验签名工具
*   @18.4.28
*/

package main

import (
	"fmt"
	"crypto/sha1"
	"crypto/md5"
    "regexp"
    "crypto/aes"
    "crypto/cipher"
//    "encoding/base64"
//    "strings"
    "bytes"
	"time"
    "math/rand"
)


/*
sha1_encode
产生哈希散列
*/
func sha1_encode (s string) string {
    h := sha1.New()
	h.Write([]byte(s))
	bs := h.Sum(nil)
    return fmt.Sprintf("%x",bs)
}

/*
md5_encode
产生md5散列
*/
func md5_encode (s string) string {
    md5Ctx := md5.New()
    md5Ctx.Write([]byte(s))
	cipherStr := md5Ctx.Sum(nil)
	ret := fmt.Sprintf("%x",cipherStr)
	return ret
}

func pregCheck(p1 string,ept bool)error{
    //入口正则
    //"^[A-Za-z0-9]+$"英文字母数字//ept参数设置是否允许空值
    m1, err00 := regexp.MatchString("^[A-Za-z0-9%.]+$", p1)
    if (err00!=nil){
        return fmt.Errorf("pregErr:%s",err00)
    }
    //空串或无效则返回错误
    if !(m1){
        return fmt.Errorf("invaild params")
    }
    if(ept){
        return nil
    }else{
        if(p1!=""){
            return nil
        }else{
            return fmt.Errorf("invaild empty params")
        }
    }
}

func pregCheck2(p1 string,p2 string,ept bool)error{
    //入口正则
    //"^[A-Za-z0-9]+$"英文字母数字//ept参数设置是否允许空值
    m1, err00 := regexp.MatchString("^[A-Za-z0-9%.]+$", p1)
    m2, err01 := regexp.MatchString("^[A-Za-z0-9%.]+$", p2)
    if (err00!=nil){
        return fmt.Errorf("pregErr:%s",err00)
    }
    if (err01!=nil){
        return fmt.Errorf("pregErr:%s",err01)
    }
    //空串或无效则返回错误
    if !(m1&&m2){
        return fmt.Errorf("invaild params")
    }
    if(ept){
        return nil
    }else{
        if(p1!=""&&p2!=""){
            return nil
        }else{
            return fmt.Errorf("invaild empty params")
        }
    }
}


func pregCheck3(p1 string,p2 string,p3 string,ept bool)error{
    //入口正则
    //"^[A-Za-z0-9]+$"英文字母数字
    m1, err00 := regexp.MatchString("^[A-Za-z0-9%.]+$", p1)
    m2, err01 := regexp.MatchString("^[A-Za-z0-9%.]+$", p2)
    m3, err02 := regexp.MatchString("^[A-Za-z0-9%.]+$", p3)
    if (err00!=nil){
        return fmt.Errorf("pregErr:%s",err00)
    }
    if (err01!=nil){
        return fmt.Errorf("pregErr:%s",err01)
    }
    if (err02!=nil){
        return fmt.Errorf("pregErr:%s",err02)
    }
    //空串或无效则返回错误
    if !(m1&&m2&&m3){
        return fmt.Errorf("invaild params")
    }
    if(ept){
        return nil
    }else{
        if(p1!=""&&p2!=""&&p3!=""){
            return nil
        }else{
            return fmt.Errorf("invaild empty params")
        }
    }
}
func pregCheck4(p1 string,p2 string,p3 string,p4 string,ept bool)error{
    //入口正则
    //"^[A-Za-z0-9]+$"英文字母数字
    m1, err00 := regexp.MatchString("^[A-Za-z0-9%.]+$", p1)
    m2, err01 := regexp.MatchString("^[A-Za-z0-9%.]+$", p2)
    m3, err02 := regexp.MatchString("^[A-Za-z0-9%.]+$", p3)
    m4, err03 := regexp.MatchString("^[A-Za-z0-9%.]+$", p4)
    if (err00!=nil){
        return fmt.Errorf("pregErr:%s",err00)
    }
    if (err01!=nil){
        return fmt.Errorf("pregErr:%s",err01)
    }
    if (err02!=nil){
        return fmt.Errorf("pregErr:%s",err02)
    }
    if (err03!=nil){
        return fmt.Errorf("pregErr:%s",err03)
    }
    //空串或无效则返回错误
    if !(m1&&m2&&m3&&m4){
        return fmt.Errorf("invaild params")
    }
    if(ept){
        return nil
    }else{
        if(p1!=""&&p2!=""&&p3!=""&&p4!=""){
            return nil
        }else{
            return fmt.Errorf("invaild empty params")
        }
    }
}

//app_key   nonce   school_code timestamp
func wx_sign(app_key string,nonce string,school_code string,timestamp string)string{
    key:=""
    str := "app_key="+app_key+"&nonce="+nonce+"&school_code="+school_code+"&timestamp="+timestamp+"&key="+key;
    return md5_encode(str)
}

func wx_sign_check(app_key string,nonce string,school_code string,timestamp string,sign string)bool{
    key:=""
    str := "app_key="+app_key+"&nonce="+nonce+"&school_code="+school_code+"&timestamp="+timestamp+"&key="+key;
    sign2 := md5_encode(str)
    if(sign==sign2){
        return true
    }else{
        return false
    }
}

//AES加密算法
//from: http://www.baike.com/wiki/AES%E5%8A%A0%E5%AF%86%E7%AE%97%E6%B3%95
func AesEncrypt(origData, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
         return nil, err
    }
    blockSize := block.BlockSize()
    origData = PKCS5Padding(origData)
    // origData = ZeroPadding(origData, block.BlockSize())
    blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
    crypted := make([]byte, len(origData))
    // 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
    // crypted := origData
    blockMode.CryptBlocks(crypted, origData)
    return crypted, nil
}

func AesDecrypt(crypted, key []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
         return nil, err
    }
    blockSize := block.BlockSize()
    blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
    origData := make([]byte, len(crypted))
    // origData := crypted
    blockMode.CryptBlocks(origData, crypted)
    origData = PKCS5Padding(origData)
    // origData = ZeroUnPadding(origData)
    return origData, nil
}

func PKCS5Padding(ciphertext []byte) []byte {
    blockSize := 128
    padding := blockSize - len(ciphertext)%blockSize
    padtext := bytes.Repeat([]byte{byte(padding)}, padding)
    return append(ciphertext, padtext...)
}



//生成指定长度随机字符串
func  GetRandomString(l int) string {  
    str := "0123456789abcdefghijklmnopqrstuvwxyz"  
    bytes := []byte(str)  
    result := []byte{}  
    r := rand.New(rand.NewSource(time.Now().UnixNano()))  
    for i := 0; i < l; i++ {  
        result = append(result, bytes[r.Intn(len(bytes))])  
    }  
    return string(result)  
}  

func makeErrJson(errcode int,errdata interface{})(int , map[string]interface{}){
    var m map[string]interface{}
    m = make(map[string]interface{})
    m["error"] = errcode
    m["msg"] = fmt.Sprint(errdata)
    return 200,m
}

func makeErrJson401(errcode int,errdata interface{})(int , map[string]interface{}){
    var m map[string]interface{}
    m = make(map[string]interface{})
    m["error"] = errcode
    m["msg"] = fmt.Sprint(errdata)
    return 200,m
}

func makeSuccessJson(data string)(int , map[string]interface{}){
    var m map[string]interface{}
    m = make(map[string]interface{})
    m["error"] = 0
    m["msg"] = data
    return 200,m
}