package main

import (
    "flag"
    "fmt"
    "github.com/gin-gonic/gin"
    "io/ioutil"
    "net"
    "os"
)

var configfile string
var app_version="0.0.1"

func main() {

    var port int
    flag.IntVar(&port, "p", 80, "端口号,默认为空")
    flag.StringVar(&configfile, "c", "/etc/config", "配置文件，默认/etc/config")
    flag.Parse()

    r := gin.Default()

    r.GET("/", homepage)
    r.GET("/version", version)
    r.GET("/config", config)
    r.GET("/whoami", whoami)

    r.Run(fmt.Sprintf(":%d", port))
}

func homepage(c *gin.Context) {
    c.String(200, "Hello, World!")
}

func version(c *gin.Context) {
    c.String(200, app_version)
}

func config(c *gin.Context) {
    file, err := os.Stat(configfile)
    if err != nil || file.IsDir() {
        c.String(500, "config file does not exist")
        return
    }
    config, err := ioutil.ReadFile(configfile)
    if err != nil {
        c.String(500, err.Error())
        return
    }
    c.String(200, string(config))
}

func whoami(c *gin.Context) {
    var ips string

    addrs, err := net.InterfaceAddrs()

    if err != nil {
        c.String(500, err.Error())
        return
    }

    for _, address := range addrs {

        // 检查ip地址判断是否回环地址
        if ipnet, ok := address.(*net.IPNet); ok && !ipnet.IP.IsLoopback() {
            if ipnet.IP.To4() != nil {
                ips += ipnet.IP.String() + "    "
            }

        }
    }

    c.String(200, ips)
}
