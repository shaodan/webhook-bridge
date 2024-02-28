package handler

import (
	"github.com/go-lark/lark"
)

var (
	bots map[string]*lark.Bot
)

func init() {
	bots = make(map[string]*lark.Bot)
	bot := lark.NewNotificationBot("https://open.larksuite.com/open-apis/bot/v2/hook/7d9f268e-f162-4c3e-8b1f-0565e4bbe257")
	bots["aioquant"] = bot
	bot = lark.NewNotificationBot("https://open.larksuite.com/open-apis/bot/v2/hook/bd14ad86-cd94-4102-9cd9-801573598e3d")
	bots["blofin-mm"] = bot
}

func GetAvailableRepos() []string {
	keys := make([]string, 0, len(bots))
	for k := range bots {
		keys = append(keys, k)
	}
	return keys
}

func GetBotForRepo(repoName string) *lark.Bot {
	if bot, ok := bots[repoName]; ok {
		return bot
	}
	return nil
}
