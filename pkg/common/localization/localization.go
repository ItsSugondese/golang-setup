package localization

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/nicksnyder/go-i18n/v2/i18n"
	"github.com/pelletier/go-toml/v2"
	"golang.org/x/text/language"
)

func InitBundle() *i18n.Bundle {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("toml", toml.Unmarshal)

	// Load the translation files
	bundle.MustLoadMessageFile("locales/en.toml")
	bundle.MustLoadMessageFile("locales/ne-NP.toml")
	return bundle
}

func GetLocalizedMessage(c *gin.Context, messageID string, templateData map[string]interface{}) string {
	localizer, exists := c.Get("localizer")
	if !exists {
		panic("Localization error")
	}

	// Localize the message
	localizedMessage, err := localizer.(*i18n.Localizer).Localize(&i18n.LocalizeConfig{
		MessageID:    messageID,
		TemplateData: templateData,
		DefaultMessage: &i18n.Message{
			ID:    messageID,
			Other: "Message Not found",
		},
	})

	if err != nil {
		panic(fmt.Sprintf("Error localizing message: %v", err))
	}

	return localizedMessage
}
