package db

import "time"

type User struct {
	ID 				string `gorm:"type:varchar(32);primary key:id;"`
	Email			string `gorm:"type:varchar(30);unique_index"`
	Nickname		string `gorm:"type:varchar(30);"`
	HeadImg			string `gorm:"type:varchar(128);"`
	Pwd				string `gorm:"type:varchar(32);not null"`
	Salt			string `gorm:"type:varchar(8);not null"`
	Err				uint
	CreateAt		time.Time
	LoginAt			time.Time
	LoginIp			string `gorm:"type:varchar(32);"`

	Token			string `gorm:"type:varchar(32);primary key:token;"`
	TokenExpirseAt	time.Time
}

type Directory struct {
	ID		string `gorm:"type:varchar(32);primary key:id;"`
	UserID	string `gorm:"type:varchar(32);not null;"`
	DirName	string `gorm:"type:varchar(64);not null;"`
	PID		string `gorm:"type:varchar(32);not null;"`
	SerNo	uint
}

type Article struct {
	ID			string `gorm:"type:varchar(32);primary key:id;"`
	UserID		string `gorm:"type:varchar(32);not null;"`
	DirID		string `gorm:"type:varchar(32);not null;"`
	Title		string `gorm:"type:varchar(128);not null;"`
	CreateAt	time.Time
	Read 		uint
	IsTop		bool
}

type ArticleDetail struct {
	ID 		string `gorm:"type:varchar(32);primary key:id;"`
	Title	string `gorm:"type:varchar(128);not null;"`
	Content	string `gorm:"type:text;"`
}
