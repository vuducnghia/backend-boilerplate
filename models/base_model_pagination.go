package models

import (
	"github.com/gin-gonic/gin"
	"github.com/uptrace/bun"
	"math"
)

type PaginationMeta struct {
	CurrentPage int `json:"current_page"`
	PageCount   int `json:"page_count"`
	TotalCount  int `json:"total_count"`
}

type PaginationWrapper struct {
	PaginationMeta `json:"metadata"`
	Data           interface{} `json:"data"`
}

func NewPaginationWrapper(data interface{}, count int, c *gin.Context) *PaginationWrapper {
	return &PaginationWrapper{
		PaginationMeta{
			CurrentPage: c.GetInt("page_number"),
			PageCount:   int(math.Ceil(float64(count) / float64(c.GetInt("limit")))),
			TotalCount:  count,
		},
		data,
	}
}

func ApplyPagination(q *bun.SelectQuery, c *gin.Context) *bun.SelectQuery {
	p := c.GetInt("page_number")
	l := c.GetInt("limit")

	return q.Limit(l).Offset((p - 1) * l)
}
