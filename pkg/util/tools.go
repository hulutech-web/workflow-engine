package util

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"github.com/google/uuid"
	"github.com/jinzhu/copier"
	"go.uber.org/zap"
	"math"
	"math/rand"
	"reflect"
	"strconv"
	"strings"
	"time"
)

type toolsUtil struct{}

var ToolsUtil = toolsUtil{}

// Copy
func (t toolsUtil) Copy(toValue interface{}, fromValue interface{}) interface{} {
	if err := copier.Copy(toValue, fromValue); err != nil {
		zap.S().Panic("copy error", zap.Error(err))
		return nil
	}
	return toValue
}

// RandomString 返回随机字符串
func (tu toolsUtil) RandomString(length int) string {
	allRandomStr := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
	byteList := make([]byte, length)
	for i := 0; i < length; i++ {
		byteList[i] = allRandomStr[rand.Intn(62)]
	}
	return string(byteList)
}

// MakeUuid 制作UUID
func (tu toolsUtil) MakeUuid() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}

// MakeMd5 制作MD5
func (tu toolsUtil) MakeMd5(data string) string {
	sum := md5.Sum([]byte(data))
	return hex.EncodeToString(sum[:])
}

// Round float四舍五入
func (tu toolsUtil) Round(val float64, n int) float64 {
	base := math.Pow(10, float64(n))
	return math.Round(base*val) / base
}

// JsonToObj JSON转Obj
func (tu toolsUtil) JsonToObj(jsonStr string, toVal interface{}) (err error) {
	return json.Unmarshal([]byte(jsonStr), &toVal)
}

// ObjToJson Obj转JSON
func (tu toolsUtil) ObjToJson(data interface{}) (res string, err error) {
	b, err := json.Marshal(data)
	if err != nil {
		return res, err
	}
	res = string(b)
	return res, nil
}

// ObjsToMaps
func (tu toolsUtil) ObjsToMaps(objs interface{}) []map[string]interface{} {
	var maps []map[string]interface{}
	val := reflect.ValueOf(objs)
	if val.Kind() == reflect.Slice {
		for i := 0; i < val.Len(); i++ {
			m := make(map[string]interface{})
			v := val.Index(i).Interface()
			t := v.(reflect.Value).Type()
			for j := 0; j < t.NumField(); j++ {
				m[t.Field(j).Name] = v.(reflect.Value).Field(j).Interface()
			}
			maps = append(maps, m)
		}
	}
	return maps
}

// MakeToken 生成唯一Token
func (tu toolsUtil) MakeToken() string {
	ms := time.Now().UnixMilli()
	token := tu.MakeMd5(tu.MakeUuid() + strconv.FormatInt(ms, 10) + tu.RandomString(8))
	tokenSecret := token + "zhoudm1743"
	return tu.MakeMd5(tokenSecret) + tu.RandomString(6)
}

// ParseDuration 解析时间字符串
func (tu toolsUtil) ParseDuration(durationStr string) (time.Duration, error) {
	return time.ParseDuration(durationStr)
}

// Contains 判断src是否包含elem元素
func (tu toolsUtil) Contains(src interface{}, elem interface{}) bool {
	srcArr := reflect.ValueOf(src)
	if srcArr.Kind() == reflect.Ptr {
		srcArr = srcArr.Elem()
	}
	if srcArr.Kind() == reflect.Slice {
		for i := 0; i < srcArr.Len(); i++ {
			if srcArr.Index(i).Interface() == elem {
				return true
			}
		}
	}
	return false
}
