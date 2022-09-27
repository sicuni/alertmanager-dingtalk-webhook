## Alertmanager Dingtalk Webhook

Webhook service support send Prometheus 2.0 alert message to Dingtalk.

## How To Use

将alertmanager的webhook指向该服务，并且在prometheus中配置对应的规则需要上报钉钉的URL。具体的匹配可以按照alertmanager的alertmanagerConfig配置


```
groups:
alert:test-load-1
expr:node_load1 > 1
for: 2m
labels:
    dingtalk: https://oapi.dingtalk.com/robot/send?access_token=xxxxx
annotations:
description: {{$labels.instance}}: job {{$labels.job}} 测试测试 负载大于1
summary: {{$labels.instance}}: load1 >1
```
