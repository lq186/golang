package article

import (
	"github.com/lq186/golang/lq186.com/apiserver/db"
	"github.com/lq186/golang/lq186.com/apiserver/common"
	"github.com/lq186/golang/lq186.com/apiserver/log"
	"time"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
)

func Create(requestBody *CreateRequestBody, tokenUser *db.User) (*db.Article, error) {

	uuid, err := common.UUID()
	if err != nil {
		log.Log.Errorf("can not build uuid, more info: %v", err)
		return nil, err
	}

	article := db.Article{}

	article.ID = uuid
	article.CreateAt = time.Now()
	article.Read = 0
	article.UserID = tokenUser.ID

	article.Title = requestBody.Title
	article.DirID = requestBody.DirID
	article.IsTop = requestBody.IsTop
	article.Lang = requestBody.Lang

	tx := db.DB.Begin()

	err = tx.Create(&article).Error
	if err != nil {
		tx.Rollback()
		log.Log.Errorf("create Article failed, more info: %v", err)
		return nil, err
	}

	detail := db.ArticleDetail{}
	detail.Title = requestBody.Title
	detail.Content = requestBody.Content
	detail.ID = article.ID

	err = tx.Create(&detail).Error
	if err != nil {
		tx.Rollback()
		log.Log.Errorf("create ArticleDetail (%s) failed, more info: %v", article.ID, err)
		return nil, err
	}

	tx.Commit()
	return &article, nil
}

func Update(requestBody *UpdateRequestBody, tokenUser *db.User) (*db.Article, error) {

	var article = db.Article{}
	err := db.DB.First(&article, "id = ? and user_id = ?", requestBody.ID, tokenUser.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log.Errorf("query Article (%s) failed, more info: %v", requestBody.ID, err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Can not found article.")
	}

	var dir = db.Directory{}
	err = db.DB.First(&dir, "id = ?", requestBody.DirID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log.Errorf("query Directory (%s) failed, more info: %v", requestBody.DirID, err)
		return nil, err
	}
	if err == gorm.ErrRecordNotFound {
		return nil, errors.New("Can not found directory.")
	}

	tx := db.DB.Begin()

	err = tx.Model(&article).Update(db.Article{Title: requestBody.Title, DirID: requestBody.DirID, IsTop: requestBody.IsTop}).Error
	if err != nil {
		tx.Rollback()
		log.Log.Errorf("update Article (%s) failed, more info: %v", requestBody.ID, err)
		return nil, err
	}

	var detail = db.ArticleDetail{}
	err = tx.First(&detail, "id = ?", requestBody.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		log.Log.Errorf("query ArticleDetail (%s) failed, more info: %v", requestBody.ID, err)
		return nil, err
	}

	if err == gorm.ErrRecordNotFound {
		err = tx.Create(&detail).Error
		if err != nil {
			tx.Rollback()
			log.Log.Errorf("create ArticleDetail (%s) failed, more info: %v", requestBody.ID, err)
			return nil, err
		}
	} else {
		err = tx.Model(&detail).Update(db.ArticleDetail{Title: requestBody.Title, Content: requestBody.Content}).Error
		if err != nil {
			tx.Rollback()
			log.Log.Errorf("update ArticleDetail (%s) failed, more info: %v", requestBody.ID, err)
			return nil, err
		}
	}

	tx.Commit()
	return &article, nil

}

func Remove(requestBody *RemoveRequestBody, tokenUser *db.User) error {

	var article = db.Article{}
	err := db.DB.First(&article, "id in (?) and user_id = ?", requestBody.IDS, tokenUser.ID).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log.Errorf("query Article (%v) failed, more info: %v", requestBody.IDS, err)
		return err
	}

	tx := db.DB.Begin()

	err = tx.Delete(&article).Error
	if err != nil {
		tx.Rollback()
		log.Log.Errorf("delete Article (%v) failed, more info: %v", requestBody.IDS, err)
		return err
	}

	err = tx.Delete(&db.ArticleDetail{}, "id in (?)", requestBody.IDS).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		tx.Rollback()
		log.Log.Errorf("delete ArticleDetail (%v) failed, more info: %v", requestBody.IDS, err)
		return err
	}

	tx.Commit()
	return nil
}

func ListPage(requestBody *ListPageRequestBody) (*common.Page, error) {

	var articles = []*db.Article{}
	var count = uint32(0)

	err := db.DB.Order("is_top desc, create_at desc").Where("lang = ?", requestBody.Lang).Find(&db.Article{}).Count(&count).Offset((requestBody.Page - 1) * requestBody.Size).Limit(requestBody.Size).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log.Errorf("query Articles failed, more info: %v", err)
		return nil, err
	}

	page := common.Page{}
	page.Page = requestBody.Page
	page.Size = requestBody.Size
	page.Count = count
	page.Content = articles

	return &page, nil
}

func ListAll(lang string, dirId string) (*[]*db.Article, error) {

	var articles = []*db.Article{}

	err := db.DB.Order("is_top desc, create_at desc").Where("lang = ? and dir_id = ?", lang, dirId).Find(&articles).Error
	if err != nil && err != gorm.ErrRecordNotFound {
		log.Log.Errorf("query Articles failed, more info: %v", err)
		return nil, err
	}

	return &articles, nil
}

func QueryDetail(id string) (*db.ArticleDetail, error) {

	var detail = db.ArticleDetail{}
	err := db.DB.First(&detail, "id = ?", id).Error
	if err != nil {
		log.Log.Errorf("query ArticleDetail (%s) failed, more info: %v", id, err)
		return nil, err
	}

	return &detail, nil
}
