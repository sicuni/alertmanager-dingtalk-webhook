package main

import (
	"flag"
	"net/http"
	"github.com/sicuni/alertmanager-dingtalk-webhook/model"
	"github.com/sicuni/alertmanager-dingtalk-webhook/notifier"
	"github.com/gin-gonic/gin"
)

var (
	h            bool
	defaultRobot string
	listen		 string
)

func init() {
	flag.BoolVar(&h, "h", false, "help")
	flag.StringVar(&defaultRobot, "defaultRobot", "", "global dingtalk robot webhook, you can overwrite by alert rule with annotations dingtalkRobot")
	flag.StringVar(&listen, "listen", ":9998", "server listen port")
}

func main() {

	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	router := gin.Default()
	router.POST("/webhook", func(c *gin.Context) {
		var notification model.Notification

		err := c.BindJSON(&notification)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		notifyDingTalk := notification.CommonLabels["dingtalk"]
		err = notifier.Send(notification, notifyDingTalk)

		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})

		}

		c.JSON(http.StatusOK, gin.H{"message": "send to dingtalk successful!"})

	})
	router.Run(listen)
}
