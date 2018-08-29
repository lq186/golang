package user

import (
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/satori/go.uuid"
	"strings"
	"time"
)

func Create(user *db.User) error {
	user.CreateAt = time.Now()
	user.Salt = common.RandomString(6)
	user.Pwd = password(user.Pwd, user.Salt)
	uuid, err := uuid.NewV4()
	if err != nil {
		return err
	}
	user.ID = strings.Replace(uuid.String(), "-", "", 4)
	return db.DB.Create(user).Error
}

func Login(body *LoginBody) (*db.User, error) {

	return nil, nil

}

func password(pwd string, salt string) string {
	return common.MD5(salt + pwd + salt)
}
