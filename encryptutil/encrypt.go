package encryptutil

import (
    "time"
    "math/rand"
    "crypto/md5"
    "encoding/hex"
)

/**
生成盐
 */
func GenerateSalt() string {
    salt := string(time.Now().UnixNano()) + string(rand.Intn(99999999))
    m5 := md5.New()
    m5.Write([]byte(salt))
    st := m5.Sum(nil)
    return hex.EncodeToString(st)
}

/**
加盐加密
 */
func EncryptWithSalt(plainStr, salt string) (strEncrypt string)  {
    m5 := md5.New()
    m5.Write([]byte(plainStr + salt))
    st := m5.Sum(nil)
    strEncrypt = hex.EncodeToString(st)
    return
}

/**
md5编码
 */
func Md5(plainStr string) (strEncrypt string) {
    m5 := md5.New()
    m5.Write([]byte(plainStr))
    st := m5.Sum(nil)
    strEncrypt = hex.EncodeToString(st)
    return
}

