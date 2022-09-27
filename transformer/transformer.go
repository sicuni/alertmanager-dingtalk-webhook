package transformer

import (
	"bytes"
	"fmt"
	"time"

	"github.com/sicuni/alertmanager-dingtalk-webhook/model"
)

// TransformToMarkdown transform alertmanager notification to dingtalk markdow message
func TransformToMarkdown(notification model.Notification) (markdown *model.DingTalkMarkdown, robotURL string, err error) {
	status := notification.Status
	if status == "resolved" {
		return
	}
	annotations := notification.CommonAnnotations
	robotURL = annotations["dingtalkRobot"]

	var buffer bytes.Buffer
	buffer.WriteString("### <font color=\"red\"> ！</font>星智云警报\n")

	var cstSh, _ = time.LoadLocation("Asia/Shanghai")
	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		alert.StartsAt = alert.StartsAt.In(cstSh)
		alert.EndsAt = alert.EndsAt.In(cstSh)
		buffer.WriteString(fmt.Sprintf("\n> 告警时间：%s\n", alert.StartsAt.Format("2006-01-02 15:04:05")))
		buffer.WriteString(fmt.Sprintf("\n> 告警内容：%s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 资源范围：%s\n", alert.Labels["clustername"]))

	}

	markdown = &model.DingTalkMarkdown{
		MsgType: "markdown",
		Markdown: &model.Markdown{
			Title: "星智云警报",
			Text:  buffer.String(),
		},
		At: &model.At{
			IsAtAll: false,
		},
	}

	return
}
