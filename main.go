package botwaygo

import (
	"bytes"
	"errors"
	"io/ioutil"

	"github.com/abdfnx/botway/constants"
	"github.com/spf13/viper"
	"github.com/tidwall/gjson"
)

func GetBotInfo(value string) string {
	viper.SetConfigType("yaml")

	BotConfig, err := ioutil.ReadFile(".botway.yaml")

	if err != nil {
		panic(err)
	}

	viper.ReadConfig(bytes.NewBuffer(BotConfig))

	return viper.GetString("bot." + value)
}

func GetToken() string {
	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Your Bot language is not Go"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_token").String()

		return data
	}
}

func GetAppId() string {
	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Your Bot language is not Go"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots." + GetBotInfo("name") + ".bot_app_id").String()

		return data
	}
}

func GetGuildId(serverName string) string {
	if GetBotInfo("lang") != "go" {
		panic(errors.New("ERROR: Your Bot language is not Go"))
	} else if GetBotInfo("type") != "discord" {
		panic(errors.New("ERROR: This function/feature is only working with discord bots"))
	} else {
		data := gjson.Get(string(constants.BotwayConfig), "botway.bots." + GetBotInfo("name") + ".guilds." + serverName + ".server_id").String()

		return data
	}
}
