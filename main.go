package main

import (
	"flag"
	"io/ioutil"
	"strings"

	"github.com/gin-gonic/gin"
)

var configPath string

func main() {
	flag.StringVar(&configPath, "c", "config/config.yaml", "config")
	flag.Parse()

	g := gin.Default()
	g.GET("/api/version", Version)
	g.GET("/api/v", Version)
	g.GET("/api/conf", Config)
	g.GET("/api/c", Config)

	g.Static("/static", "static")
	g.StaticFile("/index", "static/index.html")

	g.Run(":8888")
}

func Config(c *gin.Context) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	c.String(200, string(buf))
	return
}

func Version(c *gin.Context) {
	buf, err := ioutil.ReadFile(configPath)
	if err != nil {
		c.String(400, err.Error())
		return
	}
	var v = ""
	var strs = strings.Split(string(buf), "\n")
	for _, s := range strs {
		var ss = strings.Split(s, ":")
		if len(ss) > 1 {
			v = "version" + ss[1]
		}
		break
	}

	c.String(200, v)
}
