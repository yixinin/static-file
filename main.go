package main

import (
	"flag"
	"io/ioutil"
	"os/exec"
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
	var gitLog = ""
	var cmd = exec.Command("bash", "-c", "git log")
	out, err := cmd.CombinedOutput()
	if err != nil {
		if x := string(out); x != "" {
			c.String(400, x+"\n"+err.Error())
		} else {
			c.String(400, err.Error())
		}
		return
	}
	var strs = strings.Split(string(out), "commit ")
	for _, v := range strs {
		if len(v) > 54 && v[41] == 65 {
			gitLog = "commit " + v
			break
		}
	}

	c.String(200, gitLog)
}
