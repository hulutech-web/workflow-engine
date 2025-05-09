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
	SearchByParams(params map[string]string, conditionMap map[string]interface{}, excepts ...string) *PaginateService
	ResultPagination(dest any, withes ...string) error
}

type PaginateService struct {
	Query *gorm.DB
	Ctx   *gin.Context
}

// SearchByParams
// example SearchByParams(map[string]{}{"name":"user"}, map[string]interface{}{"state",1}, []string{"age"}...)
// ?name=xxx&pageSize=1&currentPage=1&sort=xxx&order=xxx
func (h *PaginateService) SearchByParams(params map[string]string, conditionMap map[string]interface{}, excepts ...string) *PaginateService {
	for _, except := range excepts {
		delete(params, except)
	}
	if h.Query != nil {
		query := h.Query
		// 再处理url查询

		h.Query = func(q *gorm.DB) *gorm.DB {
			//处理日期时间
			// 先处理过滤条件
			for key, val := range conditionMap {
				q.Where(fmt.Sprintf("%s=?%v", key, val))
			}
			for key, value := range params {
				//如果key包含了[]符号

				if strings.Contains(key, "[]") || value == "" || key == "pageSize" || key == "total" || key == "currentPage" || key == "sort" || key == "order" {
					continue
				} else {
					q = q.Where(gorm.Expr(key+" LIKE ?", "%"+value+"%"))
				}
				//则表示是日期时间范围
				/**
				created_at[]: 2024-10-21 00:00:00
				created_at[]: 2024-10-21 23:59:59
				*/
				if strings.Contains(key, "[]") {
					key = strings.Replace(key, "[]", "", -1)
					if value == "" {
						continue
					}
					//按照，拆分value
					ranges := strings.Split(value, ",")
					if len(ranges) == 2 {
						q = q.Where(key+" BETWEEN ? AND ?", ranges[0], ranges[1])
					} else {
						continue
					}
				}
			}

			return q
		}(query)
	} else {
		h.Query = func(q *gorm.DB) *gorm.DB {
			//处理日期时间
			// 先处理过滤条件
			for key, val := range conditionMap {
				q.Where(fmt.Sprintf("%s=?%v", key, val))
			}
			for key, value := range params {
				//如果key包含了[]符号

				if strings.Contains(key, "[]") || value == "" || key == "pageSize" || key == "total" || key == "currentPage" || key == "sort" || key == "order" {
					continue
				} else {
					q = q.Where(gorm.Expr(key+" LIKE ?", "%"+value+"%"))
				}
				//则表示是日期时间范围
				/**
				created_at[]: 2024-10-21 00:00:00
				created_at[]: 2024-10-21 23:59:59
				*/
				if strings.Contains(key, "[]") {
					key = strings.Replace(key, "[]", "", -1)
					if value == "" {
						continue
					}
					//按照，拆分value
					ranges := strings.Split(value, ",")
					if len(ranges) == 2 {
						q = q.Where(key+" BETWEEN ? AND ?", ranges[0], ranges[1])
					} else {
						continue
					}
				}
			}

			return q
		}(h.Query)
	}

	return h
}

func (r *PaginateService) ResultPagination(dest any, withes ...string) error {
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
		return err
	}

	if err := r.Query.Offset(offset).Limit(pageSizeInt).Find(dest).Error; err != nil {
		r.Ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
			"error": "query failed",
		})
		return err
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
		Prev:  query(currentPageInt - 1),
		Next:  query(currentPageInt + 1),
	}

	meta := Meta{
		TotalPage:   totalPage,
		CurrentPage: currentPageInt,
		PerPage:     pageSizeInt,
		Total:       total,
	}

	pageResult := PageResult{
		Data:  dest,
		Total: total,
		Links: links,
		Meta:  meta,
	}

	r.Ctx.JSON(http.StatusOK, pageResult)
	return nil
}

func NewPaginatorService(db *gorm.DB, ctx *gin.Context) paginateServiceImpl {
	return &PaginateService{Query: db, Ctx: ctx}
}
