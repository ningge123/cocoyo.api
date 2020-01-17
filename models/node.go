package models

import (
	"cocoyo/pkg/e"
	"encoding/json"
	"github.com/jinzhu/gorm"
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

func ScopeRoot(db *gorm.DB) *gorm.DB {
	return db.Where("node_id = ?", 0)
}

func ScopeLeaf(db *gorm.DB) *gorm.DB  {
	return db.Where("node_id <> ?", 0)
}

func GetNodes(all bool, page int, limit int) []*Node {
	var nodes []*Node

	offset := (page - 1) * limit

	chain := db.Order("created_at").Set("gorm:auto_preload", true).Offset(offset).Limit(limit)

	if all {
		chain = chain.Scopes(ScopeRoot)
	} else {
		chain = chain.Scopes(ScopeLeaf)
	}

	err := chain.Find(&nodes).Error

	e.Throw(err)

	// 对格式进行转化
	for key, node := range nodes {
		json.Unmarshal([]byte(node.Settings), &nodes[key].SettingsMap)
		json.Unmarshal([]byte(node.Cache), &nodes[key].CacheMap)

		for k, children := range node.Children {
			json.Unmarshal([]byte(children.Settings), &nodes[key].Children[k].SettingsMap)
			json.Unmarshal([]byte(children.Cache), &nodes[key].Children[k].CacheMap)
		}
	}

	return nodes
}