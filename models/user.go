package models

import (
	"cocoyo/pkg/e"
	"github.com/jinzhu/gorm"
	"time"
)

type CacheFields struct {
	ThreadsCount 		int `json:"threads_count"`
	CommentsCount 		int `json:"comments_count"`
	LikesCount 			int `json:"likes_count"`
	FollowingsCount 	int `json:"followings_count"`
	FollowersCount 		int `json:"followers_count"`
	SubscriptionsCount 	int `json:"subscriptions_count"`
}

type ExtendsFields struct {
	Company 	string `json:"company"`
	Location 	string `json:"location"`
	HomeUrl 	string `json:"home_url"`
	Github 		string `json:"github"`
	Twitter 	string `json:"twitter"`
	Facebook 	string `json:"facebook"`
	Instagram 	string `json:"instagram"`
	Telegram 	string `json:"telegram"`
	Coding 		string `json:"coding"`
	Steam 		string `json:"steam"`
	Weibo 		string `json:"weibo"`
}

type User struct {
	ActivatedAt 	time.Time `gorm:"default:NULL"`
	Avatar      	string `gorm:"default:NULL"`
	BannedAt    	time.Time `gorm:"default:NULL"`
	Bio         	string `gorm:"default:NULL"`
	Cache       	string `gorm:"default:NULL"`
	CreatedAt 		time.Time `gorm:"default:NULL"`
	DeletedAt 		time.Time `gorm:"default:NULL"`
	Email     		string
	Password        string `gorm:"default:NULL"`
	Energy    		int `gorm:"default:0"`
	Extends    		string `gorm:"default:NULL"`
	Gender        	string `gorm:"default:male"`
	ID            	int
	IsAdmin       	int `gorm:"default:0"`
	LastActiveAt  	time.Time `gorm:"default:NULL"`
	Level         	int `gorm:"default:0"`
	Phone         	string `gorm:"default:NULL"`
	Realname      	string `gorm:"default:NULL"`
	Settings	 	string `gorm:"default:NULL"`
	UpdatedAt 		time.Time `gorm:"default:NULL"`
	Username  		string
}

func (u *User) BeforeCreate() (err error)  {
	var user User
	db.Where("username = ?", u.Username).First(&user)

	if user.Username != "" {
		err = e.New("用户名已存在")

		return
	}

	return nil
}


func (u *User) Create(username string, email string, password string) (err error) {
	now := time.Now()

	u.Username = username
	u.Email = email
	u.Password = password
	u.CreatedAt = now
	u.UpdatedAt = now

	return db.Create(u).Error
}

func (u *User) ScopeWithoutBanned(db *gorm.DB) *gorm.DB {
	return db.Where("banned_at is null")
}