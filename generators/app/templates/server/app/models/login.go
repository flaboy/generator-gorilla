package models

import "gorm.io/gorm"

type Login struct {
	gorm.Model
	Channel     string
	ChannelId   string
	AccessToken string
	Data        string
}

func (c *Login) ToString() string {
	return c.Channel + ":" + c.ChannelId
}

func (c *Login) RetrieveByAccessToken(token string) error {
	return DB.Where("access_token = ?", token).First(c).Error
}
