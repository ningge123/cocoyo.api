package requests

import (
	"cocoyo/models"
	"github.com/astaxie/beego/validation"
)

type Register struct {
	Username string `valid:"Required;MinSize(5);MaxSize(100)"`
	Email    string `valid:"Required;Email"`
	Password string `valid:"Required;MinSize(5);MaxSize(100)"`
}

func (r *Register) Valid(v *validation.Validation)  {
	// 这里可以验证用户名是否重复
	var user models.User
	 if ! models.ReturnDB().Where("username = ?", r.Username).Or("email = ?", r.Email).First(&user).RecordNotFound() {
		 v.SetError("Username", "邮箱或用户名已存在")
	 }
}

type Login struct {
	Email string `valid:"Required;Email"`
	Password string `valid:"Required;MinSize(5);MaxSize(100)"`
}