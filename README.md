# GitLab Webhook to Feishu Bot

This project allows you to set up a webhook in GitLab that will send notifications to a Feishu bot when specific actions occur in your repository.

## 待开发功能
[] 更多event，优先级 comment
[] 消息格式优化, 带链接
[] 绑定gitlab & feishu user, 自动@
[] 应用机器人：支持发送飞书消息，实现操作命令
[] 支持多项目多群


## 补充信息

1. 飞书机器人配置 [自定义机器人](https://open.feishu.cn/document/client-docs/bot-v3/add-custom-bot) 每个群单独，不需要审核和 secret，直接发送到 webhook 即可
2. 应用机器人 https://open.feishu.cn/app/cli_a5eee6c7a87e500b/baseinfo
3. 飞书捷径
4. 飞书 Gitlab 连接器

# 参考

1. https://github.com/NinoFocus/gitlab-feishu-webhook
2. https://github.com/EalenXie/webhook
3. https://github.com/XUJINKAI/webhook-bridge/blob/09ad228a2a3e8a657a3461e03d35480d007eb0a0/scripts/api/modules/msg_parse_gitlab.py#L5

## Installation

1. Clone this repository in your local machine.
2. Create a new Feishu bot and copy the webhook url.
3. Copy `.env-example` to `.env`.
4. Open `.env` file and modify the `FEISHU_BOT_WEBHOOK_URL` value.
5. Run `docker-compose up -d --build` to build the container and start the service.
6. In your GitLab project, go to Settings > Integrations and enter the URL of your server followed by the endpoint https://YOUR_DOMAIN/gitlab/webhook (e.g. https://example.com/gitlab/webhook)
7. Select the events you want to trigger the webhook (e.g. Push events) and save the changes.

## Usage

When the events you selected are triggered, the webhook will be sent to your server and the server will be executed, sending a notification to your Feishu bot.

## TODO List

- [x] Push events
- [ ] Tag Push events
- [ ] Comments
- [ ] Confidential comments
- [ ] Issues events
- [ ] Confidential issues events
- [x] Merge request events
- [ ] Job events
- [ ] Pipeline events
- [ ] Wiki page events
- [ ] Deployment events
- [ ] Feature flag events
- [ ] Release events
