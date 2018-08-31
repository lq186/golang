package directory

import (
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"strings"
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
