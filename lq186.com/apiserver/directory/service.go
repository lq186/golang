package directory

import (
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"strings"
	"github.com/pkg/errors"
	"fmt"
)

func Create(dir *db.Directory, tokenUser *db.User) error {

	uuid, err := common.UUID()
	if err != nil {
		log.Log.Errorf("Can not create uuid for directory: %v, more info: %v", dir, err.Error())
		return err
	}

	dir.ID = uuid
	dir.UserID = tokenUser.ID

	if "" != strings.Trim(dir.PID, " ") {

		var pDir db.Directory
		err = db.DB.First(&pDir, "id = ?", dir.PID).Error
		if err != nil {
			log.Log.Errorf("Query parent directory (ID: %s) error, more info: %v", dir.PID, err.Error())
			return err
		}

		if "" == strings.Trim(pDir.ID, " ") {
			log.Log.Warningf("Parent directory not found, not set parent directory id")
			dir.PID = ""
		}

	}

	return db.DB.Create(dir).Error
}

func Update(dir *db.Directory, tokenUser *db.User) error {

	if "" == dir.ID {
		err := "Directory ID should not be empty."
		log.Log.Error(err)
		return errors.New(err)
	}

	var oldDir db.Directory
	err := db.DB.First(&oldDir, "id = ? and user_id = ?", dir.ID, tokenUser.ID).Error
	if err != nil {
		log.Log.Errorf("Query directory (ID: %s, UserID: %s) error, more info: %v", dir.ID, tokenUser.ID, err)
		return err
	}

	if "" == oldDir.ID {
		err := fmt.Sprintf("Directory (%s) not found.", dir.ID)
		log.Log.Errorf(err)
		return errors.New(err)
	}

	if "" != strings.Trim(dir.PID, " ") {

		var pDir db.Directory
		err = db.DB.First(&pDir, "id = ?", dir.PID).Error
		if err != nil {
			log.Log.Errorf("Query parent directory (ID: %s) error, more info: %v", dir.PID, err.Error())
			return err
		}

		if "" == strings.Trim(pDir.ID, " ") {
			log.Log.Warningf("Parent directory not found, not set parent directory id")
			oldDir.PID = ""
		} else {
			oldDir.PID = dir.PID
		}

	}

	oldDir.DirName = dir.DirName

	return db.DB.Save(&oldDir).Error
}

func ListAll(tokenUser *db.User) ([]*db.Directory, error) {
	var dirs = []*db.Directory{}
	err := db.DB.Order("dir_name asc").Find(&dirs, "user_id = ?", tokenUser.ID).Error
	if err != nil {
		log.Log.Errorf("Query directory error, more info: %v", err)
		return nil, err
	}
	return dirs, nil
}

func Remove(dir *db.Directory, tokenUser *db.User) error {

	if "" == dir.ID {
		err := "Directory ID should not be empty."
		log.Log.Error(err)
		return errors.New(err)
	}

	tx := db.DB.Begin()

	// Remove Directory
	err := db.DB.Unscoped().Delete(dir, "user_id = ?", tokenUser.ID).Error
	if err != nil {
		log.Log.Errorf("Delete directory (ID: %s, UserID: %s) error, more info: %v", dir.ID, tokenUser.ID, err)
		tx.Rollback()
		return err
	}

	// Remove Article and ArticleDetail


	return tx.Commit().Error
}