package models

import "time"

type PathName struct {
	Name string `json:"name" bson:"name"`
}

type MenuModel struct {
	Id          string    `json:"id" bson:"id"`
	Role        int       `json:"role" bson:"role"`
	Name        string    `json:"name" bson:"name"`
	Icon        string    `json:"icon" bson:"icon"`
	Path        string    `json:"path" bson:"path"`
	PathName    PathName  `json:"pathname" bson:"pathname"`
	Description string    `json:"description" bson:"description"`
	Index       int       `json:"index" bson:"index"`
	Status      bool      `json:"status" bson:"status"`
	CreatedDate time.Time `json:"create_date" bson:"create_date" db:"create_date"`
	LastUpdate  time.Time `json:"update_date" bson:"update_date" db:"update_date"`
}

type MenuCreateReqModel struct {
	Role        int      `json:"role" bson:"role"`
	Name        string   `json:"name" bson:"name"`
	Icon        string   `json:"icon" bson:"icon"`
	Path        string   `json:"path" bson:"path"`
	PathName    PathName `json:"pathname" bson:"pathname"`
	Description string   `json:"description" bson:"description"`
}

type MenuCreateResModel struct {
	Name        string   `json:"name" bson:"name"`
	Icon        string   `json:"icon" bson:"icon"`
	Path        string   `json:"path" bson:"path"`
	PathName    PathName `json:"pathname" bson:"pathname"`
	Description string   `json:"description" bson:"description"`
}

type MenuResModel struct {
	Id          string   `json:"id" bson:"id"`
	Name        string   `json:"name" bson:"name"`
	Icon        string   `json:"icon" bson:"icon"`
	Path        string   `json:"path" bson:"path"`
	PathName    PathName `json:"pathname" bson:"pathname"`
	Description string   `json:"description" bson:"description"`
}
