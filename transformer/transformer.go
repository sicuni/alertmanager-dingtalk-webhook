package transformer

import (
	"bytes"
	"fmt"

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

	for _, alert := range notification.Alerts {
		annotations := alert.Annotations
		buffer.WriteString(fmt.Sprintf("\n> 告警时间：%s\n", alert.StartsAt.Format("15:04:05")))
		buffer.WriteString(fmt.Sprintf("\n> 告警内容：%s\n", annotations["description"]))
		buffer.WriteString(fmt.Sprintf("\n> 告警范围：%s\n", alert.Labels["clustername"]))

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
