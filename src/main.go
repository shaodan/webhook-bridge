package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/webhooks/v6/gitlab"
	"github.com/shaodan/webhook-bridge/src/handler"
)

func main() {
	server := gin.Default()

	loadAndCheck()

	server.GET("/ping", handlePing)
	server.GET("/gitlab/repos", handleListGitlabRepos)
	server.POST("/gitlab/webhook", handleGitlabWebhook)

	log.Fatal(server.Run(":8083"))
}

func loadAndCheck() {
	repos := handler.GetAvailableRepos()
	if len(repos) == 0 {
		log.Fatal("FEISHU_BOT_WEBHOOK_URL not found")
	}
}

func handlePing(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"message": "pong",
	})
}

func handleListGitlabRepos(ctx *gin.Context) {
	repos := handler.GetAvailableRepos()
	ctx.JSON(200, repos)
}

func handleGitlabWebhook(ctx *gin.Context) {
	hook, _ := gitlab.New()

	repo := ctx.Query("repo")
	if repo == "" {
		log.Println("no query param: repo")
		return
	}
	bot := handler.GetBotForRepo(repo)
	if bot == nil {
		log.Println("wrong repo name")
		return
	}

	payload, err := hook.Parse(ctx.Request, gitlab.PushEvents, gitlab.MergeRequestEvents, gitlab.CommentEvents)
	if err != nil {
		log.Println(err)
		return
	}
	switch payload.(type) {
	case gitlab.PushEventPayload:
		go handler.HandlePushEvent(payload.(gitlab.PushEventPayload), bot)
	case gitlab.MergeRequestEventPayload:
		go handler.HandleMergeRequestEvent(payload.(gitlab.MergeRequestEventPayload), bot)
	case gitlab.CommentEventPayload:
		go handler.HandleCommentEvent(payload.(gitlab.CommentEventPayload), bot)
	}
}
