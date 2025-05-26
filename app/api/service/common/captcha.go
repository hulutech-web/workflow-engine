package service

import (
	"encoding/json"
	"fmt"
	"github.com/wenlng/go-captcha-assets/helper"
	"github.com/wenlng/go-captcha-assets/resources/images"
	"github.com/wenlng/go-captcha/v2/base/option"
	"github.com/wenlng/go-captcha/v2/rotate"
	"log"
	"strconv"
	"sync"
	"time"
)

type CaptchaService interface {
	Generate() (string, int, string, string, error)
	CheckAngle(angle string, key string) (int, bool)
}

var rotateCapt rotate.Captcha

func init() {
	builder := rotate.NewBuilder(rotate.WithRangeAnglePos([]option.RangeVal{
		{Min: 20, Max: 330},
	}))

	// background images
	imgs, err := images.GetImages()
	if err != nil {
		log.Fatalln(err)
	}

	// set resources
	builder.SetResources(
		rotate.WithImages(imgs),
	)

	rotateCapt = builder.Make()
}

type captchaService struct {
}

func (c *captchaService) Generate() (string, int, string, string, error) {
	code := 0
	image_base64 := ""
	thumb_base64 := ""
	captchaData, err := rotateCapt.Generate()
	dotsByte, _ := json.Marshal(captchaData.GetData())
	captcha_key := helper.StringToMD5(string(dotsByte))
	//加入缓存
	WriteCache(captcha_key, dotsByte)
	if err != nil {
		return captcha_key, code, image_base64, thumb_base64, err
	}
	image_base64, _ = captchaData.GetMasterImage().ToBase64()
	thumb_base64, _ = captchaData.GetThumbImage().ToBase64()

	return captcha_key, code, image_base64, thumb_base64, nil
}
func (c *captchaService) CheckAngle(angle string, key string) (int, bool) {
	code := 1
	if angle == "" || key == "" {
		return code, false
	}

	cacheDataByte := ReadCache(key)
	if len(cacheDataByte) == 0 {
		return code, false
	}

	var dct *rotate.Block
	if err := json.Unmarshal(cacheDataByte, &dct); err != nil {
		return code, false
	}

	sAngle, _ := strconv.ParseFloat(fmt.Sprintf("%v", angle), 64)
	chkRet := rotate.CheckAngle(int64(sAngle), int64(dct.Angle), 2)

	if chkRet {
		code = 0
		return code, true
	} else {
		return code, false
	}
}
func NewCaptchaService() CaptchaService {
	return &captchaService{}
}

type cachedata = struct {
	data     []byte
	createAt time.Time
}

var mux sync.Mutex

var cachemaps = make(map[string]*cachedata)

// WriteCache .
func WriteCache(key string, data []byte) {
	mux.Lock()
	defer mux.Unlock()
	cachemaps[key] = &cachedata{
		createAt: time.Now(),
		data:     data,
	}
}

// ReadCache .
func ReadCache(key string) []byte {
	mux.Lock()
	defer mux.Unlock()
	if cd, ok := cachemaps[key]; ok {
		return cd.data
	}

	return []byte{}
}

// ClearCache .
func ClearCache(key string) {
	mux.Lock()
	defer mux.Unlock()
	delete(cachemaps, key)
}

// RunTimedTask .
func RunTimedTask() {
	ticker := time.NewTicker(time.Minute * 5)
	go func() {
		for range ticker.C {
			checkCacheOvertimeFile()
		}
	}()
}

func checkCacheOvertimeFile() {
	var keys = make([]string, 0)
	for key, data := range cachemaps {
		ex := time.Now().Unix() - data.createAt.Unix()
		if ex > (60 * 30) {
			keys = append(keys, key)
		}
	}

	for _, key := range keys {
		ClearCache(key)
	}
}
