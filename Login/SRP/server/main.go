package main

import (
	"encoding/hex"
	"fmt"
	"github.com/gin-gonic/gin"
	srp "github.com/langwan/chi-go-srp"
)

type RegRequest struct {
	Username string `json:"username"`
	Salt     string `json:"salt"`
	Verifier string `json:"verifier"`
}

type Database struct {
	Username string
	Salt     string
	Verifier string
}

type LoginARequest struct {
	Username   string `json:"username"`
	EphemeralA string `json:"ephemeralA"`
}

type LoginAResponse struct {
	Salt       string `json:"salt"`
	EphemeralB string `json:"ephemeralB"`
}

type LoginMRequest struct {
	Username string `json:"username"`
	M1       string `json:"m1"`
}

var database Database
var server *srp.Server
var params *srp.Params

func init() {
	var err error
	params, err = srp.GetParams(2048)
	if err != nil {
		panic(err)
	}
}

func main() {
	g := gin.New()
	g.POST("reg", func(c *gin.Context) {
		req := RegRequest{}
		c.BindJSON(&req)
		database.Verifier = req.Verifier
		database.Salt = req.Salt
		database.Username = req.Username
		fmt.Printf("database %+v\n", database)
		c.JSON(200, "ok")
	})
	g.POST("login/a", func(c *gin.Context) {

		req := LoginARequest{}
		c.BindJSON(&req)
		secretServer, err := srp.GenKey()
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		ephemeralA, err := hex.DecodeString(req.EphemeralA)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		verifier, err := hex.DecodeString(database.Verifier)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}
		server = srp.NewServer(params, verifier, secretServer)

		server.SetA(ephemeralA)
		ephemeralB := server.ComputeB()

		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		hexEphemeralB := hex.EncodeToString(ephemeralB)
		fmt.Println("a, b", req.EphemeralA, hexEphemeralB)
		resp := LoginAResponse{Salt: database.Salt, EphemeralB: hexEphemeralB}
		c.JSON(200, resp)
	})
	g.POST("login/m", func(c *gin.Context) {
		req := LoginMRequest{}
		c.BindJSON(&req)
		m1, err := hex.DecodeString(req.M1)
		if err != nil {
			c.AbortWithError(500, err)
			return
		}

		_, ok := server.CheckM1(m1)
		if !ok {
			c.AbortWithStatus(403)
			return
		} else {
			c.JSON(200, "ok")
		}
	})
	g.Run(":3003")
}
