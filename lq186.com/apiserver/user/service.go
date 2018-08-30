package user

import (
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"time"
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"github.com/pkg/errors"
)

func Create(user *db.User) error {
	user.CreateAt = time.Now()
	user.Salt = common.RandomString(6)
	user.Pwd = password(user.Pwd, user.Salt)
	uuid, err := common.UUID()
	if err != nil {
		return err
	}
	user.ID = uuid
	return db.DB.Create(user).Error
}

func Update(user *db.User) error {
	var oldUser db.User
	err := db.DB.First(&oldUser, "token = ?", user.Token).Error
	if err != nil {
		log.Log.Debugf("Query by token (%s) error, more info: %v", user.Token, err.Error())
		return err
	}

	err = db.DB.Model(&oldUser).Update(db.User{Nickname: user.Nickname, HeadImg: user.HeadImg}).Error
	if err != nil {
		log.Log.Errorf("Update user (id: %s) error, more info: %v", user.ID, err.Error())
		return err
	}
	return nil
}

func Login(body *LoginBody) (*db.User, error) {

	var user db.User
	err := db.DB.First(&user, "email = ? and ( err <= ? or login_at < ? )", body.Username, 5, time.Now().Add(-10 * time.Minute)).Error
	if err != nil {
		log.Log.Errorf("Query user error, more info: %v", err.Error())
		return nil, err
	}

	if "" == user.ID {
		return nil, errors.New("User not found.")
	}

	if password(body.Password, user.Salt) != user.Pwd {
		return nil, errors.New("Password miss match.")
	}

	token := common.MD5(user.ID + time.Now().String())
	user.Token = token
	user.TokenExpirseAt = time.Now().Add(2 * time.Hour)

	err = db.DB.Model(&user).Update(db.User{Token:user.Token, TokenExpirseAt:user.TokenExpirseAt, LoginIp: body.Ip, LoginAt:time.Now()}).Error
	if err != nil {
		log.Log.Errorf("Update user error, more info: %v", err.Error())
		return nil, err
	}

	user.Pwd = "";
	user.Salt = "";

	return &user, nil

}

func ExistsValidToken(token string) bool {
	var user db.User
	err := db.DB.First(&user, "token = ? and token_expirse_at > ?", token, time.Now()).Error
	if err != nil {
		log.Log.Debugf("Query token error, more info: %v", err.Error())
		return false
	}

	return "" != user.ID
}

func password(pwd string, salt string) string {
	return common.MD5(salt + pwd + salt)
}
