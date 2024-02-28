package handler

import (
	"fmt"
	"log"
	"strings"
	"text/template"

	"github.com/go-lark/lark"
	"github.com/go-playground/webhooks/v6/gitlab"
)

type MyCommentEventPayload gitlab.CommentEventPayload

func HandleCommentEvent(payload gitlab.CommentEventPayload, bot *lark.Bot) {
	p := MyCommentEventPayload(payload)

	b := lark.NewCardBuilder()

	card := b.Card(
		b.Markdown(p.renderBody()),
		b.Note().AddText(b.Text(p.renderFooter()).LarkMd()),
	).Blue().Title(p.renderTitle())

	card.Blue()

	msg := lark.NewMsgBuffer(lark.MsgInteractive)

	om := msg.Card(card.String()).Build()
	_, err := bot.PostNotificationV2(om)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Send message to feishu success")
}

func (payload *MyCommentEventPayload) renderTitle() string {
	userName := payload.User.Name

	return fmt.Sprintf("%s comment", userName)
}

func (payload *MyCommentEventPayload) renderBody() string {
	t := template.Must(template.New("commentEvent").Parse(`
**{{.MergeRequest.Title}}**
{{if ne (len .ObjectAttributes.Description) 0}}

{{.ObjectAttributes.Description}}
Author: {{.User.Name}}
{{end}}
`))
	var buf strings.Builder

	if err := t.Execute(&buf, payload); err != nil {
		return ""
	}

	result := buf.String()
	// fmt.Printf("%+v\n", payload)

	return result
}

func (payload *MyCommentEventPayload) renderFooter() string {
	return fmt.Sprintf("%s", payload.Repository.Name)
}
