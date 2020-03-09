package models

import "time"

type AddView struct {
	Data struct {
		ArticleID string `json:"article_id"`
	} `json:"data"`
}

type Count struct {
	Reference string `json:"reference"`
	Count     int    `json"count"`
}

type GetView struct {
	Data struct {
		ArticleID  string `json:"article_id"`
		Type       string `json:"type"`
		Attributes struct {
			Count []Count `json:"count"`
		} `json:"attributes"`
	} `json:"data"`
}

type SingleView struct {
	Id        int       `json:id`
	ArticleID string    `article_id`
	TimeStamp time.Time `json:t_timestamp`
}
