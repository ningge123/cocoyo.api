package models

type CacheFields struct {
	ThreadsCount 		int `json:"threads_count"`
	CommentsCount 		int `json:"comments_count"`
	LikesCount 			int `json:"likes_count"`
	FollowingsCount 	int `json:"followings_count"`
	FollowersCount 		int `json:"followers_count"`
	SubscriptionsCount 	int `json:"subscriptions_count"`
}

type ExtendsFields struct {
	HomeUrl 	string `json:"home_url"`
	Github 		string `json:"github"`
	Weibo 		string `json:"weibo"`
}

type User struct {
	Id            	int     	`json:"id"`
	Email     		string  	`json:"email"`
	Username  		string  	`json:"username"`
	Gender        	string  	`json:"gender"`
	Phone         	string  	`json:"phone"`
	Password        string  	`json:"-"`
	Avatar      	string   	`json:"avatar"`
	BannedAt    	string 		`json:"banned_at"`
	ActivatedAt 	string 		`json:"activated_at"`
	Bio         	string    	`json:"bio"`
	Cache       	string    	`json:"cache"`
	Energy    		int       	`json:"energy"`
	Extends    		string    	`json:"extends"`
	IsAdmin       	int       	`json:"is_admin"`
	Level         	int       	`json:"level"`
	Settings	 	string     	`json:"settings"`
	LastActiveAt  	string  	`json:"last_active_at"`
	CreatedAt 		string  	`json:"created_at"`
	UpdatedAt 		string  	`json:"updated_at"`
}

//func (u *User) BeforeCreate() (err error)  {
//	var user User
//	db.Where("email = ?", u.Email).First(&user)
//
//	if user.Email != "" {
//		err = e.New(e.WARNING, "邮箱已存在!")
//
//		return
//	}
//
//	db.Where("username = ?", u.Username).First(&user)
//
//	if user.Username != "" {
//		err = e.New(e.WARNING, "用户名已存在!")
//
//		return
//	}
//
//	return nil
//}
//
//func (u *User) Create(username, email, password string) (token string, err error) {
//	now := time.Now()
//
//	hash := sha1.New()
//	hash.Write([]byte(password))
//	cipherStr := hash.Sum(nil)
//	password = hex.EncodeToString(cipherStr)
//
//	u.Username = username
//	u.Email = email
//	u.Password = password
//	u.CreatedAt = now
//	u.UpdatedAt = now
//
//	err = db.Create(u).Error
//
//	if err != nil {
//		return
//	}
//
//	// 发送邮箱 ~
//	// 创建access_token
//	token, err = jwt.GenerateToken(u.Username, u.Password)
//
//	return
//}
//
//func (u *User) ScopeWithoutBanned(db *gorm.DB) *gorm.DB {
//	return db.Where("banned_at is null")
//}