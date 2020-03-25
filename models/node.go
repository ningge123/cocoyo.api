package models

import (
	"time"
)

type Node struct {
	ID          int `json:"id"`
	Banner 		string `json:"banner"`
	Cache 		string `json:"-"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
	NodeID      int64 `json:"node_id"`
	Settings    string `json:"-"`
	Title     	string `json:"title"`
	Children    []Node `json:"children"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt 	time.Time `json:"updated_at"`
	
	CacheMap 	interface{} `json:"cache"`
	SettingsMap interface{} `json:"settings"`
}