package controllers

import (
	"encoding/json"
	"io/ioutil"

	"github.com/astaxie/beego"
	"github.com/thoratou/organize-jds/models"
)

//DeserializeCompanyFromJSONFile deserialize company settings
func DeserializeCompanyFromJSONFile(filePath string) (*models.Settings, error) {
	raw, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	settings := &models.Settings{}
	err = json.Unmarshal(raw, settings)

	beego.Info("company css:", settings.CSS)
	beego.Info("company name:", settings.CompanyName)
	beego.Info("company mail extension:", settings.MailExtension)
	beego.Info("company sender mail:", settings.SenderMail)

	return settings, err
}

var globalSettings *models.Settings

//SetSettings register global settings
func SetSettings(settings *models.Settings) {
	globalSettings = settings
}

//GetSettings get golabl settings
func GetSettings() *models.Settings {
	return globalSettings
}
