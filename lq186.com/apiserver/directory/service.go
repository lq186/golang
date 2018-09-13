package directory

import (
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"strings"
	"github.com/pkg/errors"
	"fmt"
	"github.com/jinzhu/gorm"
)

func Create(dir *db.Directory, tokenUser *db.User) error {

	uuid, err := common.UUID()
	if err != nil {
		log.Log.Errorf("Can not create uuid for directory: %v, more info: %v", dir, err.Error())
		return err
	}

	dir.ID = uuid
	dir.UserID = tokenUser.ID

	if dir.SerNo <= 0 {
		maxSerNo, err := MaxSerNo()
		if err != nil {
			return err
		}
		dir.SerNo = maxSerNo
	}

	if "" != strings.Trim(dir.PID, " ") {

		var pDir db.Directory
		err = db.DB.First(&pDir, "id = ?", dir.PID).Error
		if err == gorm.ErrRecordNotFound {
			log.Log.Warningf("Parent directory not found, not set parent directory id")
			dir.PID = ""
		} else if err != nil {
			log.Log.Errorf("Query parent directory (ID: %s) error, more info: %v", dir.PID, err.Error())
			return err
		} else {
			if "" == strings.Trim(pDir.ID, " ") {
				log.Log.Warningf("Parent directory not found, not set parent directory id")
				dir.PID = ""
			} else {
				dir.PID = dir.PID
			}
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
	if err == gorm.ErrRecordNotFound {
		err := fmt.Sprintf("Directory (%s) not found.", dir.ID)
		log.Log.Errorf(err)
		return errors.New(err)
	}

	if err != nil {
		log.Log.Errorf("Query directory (ID: %s, UserID: %s) error, more info: %v", dir.ID, tokenUser.ID, err)
		return err
	}

	if "" != strings.Trim(dir.PID, " ") {

		var pDir db.Directory
		err = db.DB.First(&pDir, "id = ?", dir.PID).Error
		if err == gorm.ErrRecordNotFound {
			log.Log.Warningf("Parent directory not found, not set parent directory id")
			oldDir.PID = ""
		} else if err != nil {
			log.Log.Errorf("Query parent directory (ID: %s) error, more info: %v", dir.PID, err.Error())
			return err
		} else {
			if "" == strings.Trim(pDir.ID, " ") {
				log.Log.Warningf("Parent directory not found, not set parent directory id")
				oldDir.PID = ""
			} else {
				oldDir.PID = dir.PID
			}
		}

	}

	if dir.SerNo > 0 {
		oldDir.SerNo = dir.SerNo
	}

	oldDir.DirName = dir.DirName

	return db.DB.Save(&oldDir).Error
}

func ListAll(lang string) ([]*db.Directory, error) {
	var dirs = []*db.Directory{}
	err := db.DB.Order("ser_no desc, dir_name asc").Where("lang = ?", lang).Find(&dirs).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log.Errorf("Query directory error, more info: %v", err)
		return nil, err
	}
	return dirs, nil
}

func MaxSerNo() (uint, error) {
	var dir = db.Directory{}
	err := db.DB.Order("ser_no desc").First(&dir).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log.Errorf("Query directory error, more info: %v", err)
		return 0, err
	}
	if err == gorm.ErrRecordNotFound {
		return 1, nil
	}
	return dir.SerNo + 1, nil
}

func Remove(dir *db.Directory, tokenUser *db.User) error {

	if "" == dir.ID {
		err := "Directory ID should not be empty."
		log.Log.Error(err)
		return errors.New(err)
	}

	tx := db.DB.Begin()

	// Remove Directory
	err := tx.Unscoped().Delete(dir, "user_id = ?", tokenUser.ID).Error
	if err != nil {
		log.Log.Errorf("Delete directory (ID: %s, UserID: %s) error, more info: %v", dir.ID, tokenUser.ID, err)
		tx.Rollback()
		return err
	}

	// Remove sub Directories
	err = RemoveSubDir(tx, dir.ID)
	if err != nil {
		log.Log.Errorf("Delete sub directories error, more info: %v", err)
		tx.Rollback()
		return err
	}

	// Remove Article and ArticleDetail


	return tx.Commit().Error
}

func RemoveSubDir(tx *gorm.DB, pid string) error {
	dirs := []*db.Directory{}
	err := db.DB.Find(&dirs, "p_id = ?", pid).Error

	if err == gorm.ErrRecordNotFound {
		return nil
	}

	if err != nil {
		return err
	}

	for _, dir := range dirs {
		err := RemoveSubDir(tx, dir.ID)
		if err == nil {
			err = tx.Unscoped().Delete(dir).Error
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}
	return nil
}