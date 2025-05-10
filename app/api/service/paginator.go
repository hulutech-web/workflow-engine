package service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/spf13/cast"
	"gorm.io/gorm"
	"math"
	"net/http"
	"strconv"
	"strings"
)

type Meta struct {
	TotalPage   int   `json:"total_page"`
	CurrentPage int   `json:"current_page"`
	PerPage     int   `json:"per_page"`
	Total       int64 `json:"total"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type PageResult struct {
	Data  any   `json:"data"` // List of data
	Total int64 `json:"total"`
	Links Links `json:"links"`
	Meta  Meta  `json:"meta"`
}

type paginateServiceImpl interface {
	SearchByParams(params map[string]string, conditionMap map[string]interface{}, excepts ...string) paginateServiceImpl
	ResultPagination(dest any, withes ...string) (error, *PageResult)
}

type PaginateService struct {
	Query *gorm.DB
	Ctx   *gin.Context
}

func NewPaginatorServiceImpl(db *gorm.DB, ctx *gin.Context) paginateServiceImpl {
	return &PaginateService{Query: db, Ctx: ctx}
}

// SearchByParams
// example SearchByParams(map[string]{}{"name":"user"}, map[string]interface{}{"state",1}, []string{"age"}...)
// ?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx
func (h *PaginateService) SearchByParams(params map[string]string, conditionMap map[string]interface{}, excepts ...string) paginateServiceImpl {
	for _, except := range excepts {
		delete(params, except)
	}

	if h.Query == nil {
		return h
	}

	query := h.Query

	// 处理过滤条件
	for key, val := range conditionMap {
		query = query.Where(fmt.Sprintf("%s = ?", key), val)
	}

	// 处理URL查询参数
	for key, value := range params {
		if strings.Contains(key, "[]") || value == "" ||
			key == "pageSize" || key == "total" ||
			key == "currentPage" || key == "sort" || key == "order" {
			continue
		}

		if strings.Contains(key, "[]") {
			key = strings.Replace(key, "[]", "", -1)
			if value == "" {
				continue
			}
			ranges := strings.Split(value, ",")
			if len(ranges) == 2 {
				query = query.Where(key+" BETWEEN ? AND ?", ranges[0], ranges[1])
			}
		} else {
			query = query.Where(gorm.Expr(key+" LIKE ?", "%"+value+"%"))
		}
	}

	h.Query = query
	return h
}
func (r *PaginateService) ResultPagination(dest any, withes ...string) (error, *PageResult) {
	pageSize := r.Ctx.DefaultQuery("pageSize", "10")
	pageSizeInt := cast.ToInt(pageSize)
	currentPage := r.Ctx.DefaultQuery("currentPage", "1")
	currentPageInt := cast.ToInt(currentPage)

	for _, with := range withes {
		r.Query = r.Query.Preload(with)
	}
	r.Query = r.Query.Order("id")

	var total int64
	offset := (currentPageInt - 1) * pageSizeInt

	if err := r.Query.Model(dest).Count(&total).Error; err != nil {
		r.Ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "count failed",
		})
		return err, nil
	}

	if err := r.Query.Offset(offset).Limit(pageSizeInt).Find(dest).Error; err != nil {
		r.Ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "query failed",
		})
		return err, nil
	}

	URL_PATH := r.Ctx.Request.URL.Path
	proto := "http://"
	if r.Ctx.Request.TLS != nil {
		proto = "https://"
	}
	host := r.Ctx.Request.Host

	baseURL := proto + host + URL_PATH

	// 使用 url.Values 构建分页链接
	query := func(page int) string {
		v := make(map[string]string)
		v["pageSize"] = strconv.Itoa(pageSizeInt)
		v["currentPage"] = strconv.Itoa(page)
		var q []string
		for k, v := range v {
			q = append(q, fmt.Sprintf("%s=%s", k, v))
		}
		return baseURL + "?" + strings.Join(q, "&")
	}

	totalPage := int(math.Ceil(float64(total) / float64(pageSizeInt)))

	links := Links{
		First: query(1),
		Last:  query(totalPage),
		Prev: func() string {
			if currentPageInt > 1 {
				return query(currentPageInt - 1)
			}
			return ""
		}(),
		Next: func() string {
			if currentPageInt < totalPage {
				return query(currentPageInt + 1)
			}
			return ""
		}(),
	}

	meta := Meta{
		TotalPage:   totalPage,
		CurrentPage: currentPageInt,
		PerPage:     pageSizeInt,
		Total:       total,
	}

	pageResult := &PageResult{
		Data:  dest,
		Total: total,
		Links: links,
		Meta:  meta,
	}
	return nil, pageResult
}
func NewPaginatorService(db *gorm.DB, ctx *gin.Context) paginateServiceImpl {
	return &PaginateService{Query: db, Ctx: ctx}
}
