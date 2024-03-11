package middleware

import (
	"github.com/gin-gonic/gin"
	"strconv"
)

func PagedWrapper(h func(c *gin.Context) *gin.Error) func(c *gin.Context) *gin.Error {
	return func(c *gin.Context) *gin.Error {
		c.Set("page_number", getPageFromQuery(c))
		c.Set("limit", getLimitFromQuery(c))
		c.Set("search_query", c.Query("search_query"))

		return h(c)
	}
}

func getLimitFromQuery(c *gin.Context) int {
	if r, err := strconv.Atoi(c.DefaultQuery("limit", "20")); err != nil {
		return 20
	} else {
		if r < 1 {
			r = 20
		}
		return r
	}
}

func getPageFromQuery(c *gin.Context) int {
	if r, err := strconv.Atoi(c.DefaultQuery("page_number", "1")); err != nil {
		return 1
	} else {
		if r < 1 {
			r = 1
		}
		return r
	}
}
