package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"math"
	"net/http"
)

/**
UI -> Java Server
UI -> Go Server
*/

func main() {
	g := gin.New()

	type ApiResponse struct {
		MoneyA int64  `json:"moneyA"`
		MoneyB uint64 `json:"moneyB"`
		MoneyC string `json:"moneyC"`
	}

	g.Any("api", func(c *gin.Context) {
		resp := ApiResponse{
			MoneyA: math.MaxInt64 - 1,
			MoneyB: math.MaxUint64 - 1,
			MoneyC: fmt.Sprintf("%d", math.MaxInt64-1),
		}
		fmt.Println(resp)
		c.JSON(http.StatusOK, &resp)
	})

	g.Run(":9100")
}
