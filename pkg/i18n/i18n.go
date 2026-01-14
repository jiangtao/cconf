package i18n

import (
	"embed"
	"fmt"
	"os"
	"strings"

	"github.com/nicksnyder/go-i18n/v2/i18n"
	"golang.org/x/text/language"
	"gopkg.in/yaml.v3"
)

//go:embed *.yaml
var fs embed.FS

var localizer *i18n.Localizer

// Init initializes i18n with the specified language
func Init(lang string) error {
	bundle := i18n.NewBundle(language.English)
	bundle.RegisterUnmarshalFunc("yaml", yaml.Unmarshal)

	// Load all translation files
	files, err := fs.ReadDir(".")
	if err != nil {
		return fmt.Errorf("failed to read i18n files: %w", err)
	}

	for _, file := range files {
		if file.IsDir() || !strings.HasSuffix(file.Name(), ".yaml") {
			continue
		}
		data, err := fs.ReadFile(file.Name())
		if err != nil {
			return fmt.Errorf("failed to read %s: %w", file.Name(), err)
		}
		if _, err := bundle.ParseMessageFileBytes(data, file.Name()); err != nil {
			return fmt.Errorf("failed to parse %s: %w", file.Name(), err)
		}
	}

	// Map short codes to full language tags
	langTag := language.English
	switch strings.ToLower(lang) {
	case "zh", "zh-cn", "zh_cn":
		langTag = language.SimplifiedChinese
	case "en", "en-us", "en_us":
		langTag = language.English
	}

	localizer = i18n.NewLocalizer(bundle, langTag.String())
	return nil
}

// InitFromEnv initializes i18n from environment variables
func InitFromEnv() error {
	// Check command line flag first (handled by caller)
	// Then check config file (handled by caller)
	// Then check environment variables
	lang := os.Getenv("LANG")
	if lang == "" {
		lang = os.Getenv("LC_ALL")
	}

	// Default to English
	if lang == "" {
		return Init("en")
	}

	// Extract language code from LANG (e.g., "zh_CN.UTF-8" -> "zh")
	parts := strings.Split(lang, ".")
	if len(parts) > 0 {
		langCode := strings.ReplaceAll(parts[0], "_", "-")
		if strings.HasPrefix(langCode, "zh") {
			return Init("zh")
		}
	}

	return Init("en")
}

// T returns the translated message for the given ID
func T(msgID string, templateData map[string]interface{}) string {
	if localizer == nil {
		return msgID
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    msgID,
		TemplateData: templateData,
	})
	if err != nil {
		return msgID
	}
	return msg
}

// MustT returns the translated message or panics
func MustT(msgID string, templateData map[string]interface{}) string {
	if localizer == nil {
		return msgID
	}

	msg, err := localizer.Localize(&i18n.LocalizeConfig{
		MessageID:    msgID,
		TemplateData: templateData,
	})
	if err != nil {
		panic(fmt.Sprintf("i18n error for %s: %v", msgID, err))
	}
	return msg
}
