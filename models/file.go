package models

type Files struct {
	Id       int    `json:"-"`
	Sha1     string `json:"sha1"`
	Name     string `json:"name"`
	Size     int64  `json:"size"`
	Location string `json:"location"`
	Status   int    `json:"status"` 
	CreateAt string `json:"create_at"`
	UpdateAt string `json:"-"`
}