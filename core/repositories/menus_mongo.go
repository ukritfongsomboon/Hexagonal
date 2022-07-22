package repositories

import (
	"context"
	"errors"
	"hexagonal/core/models"

	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var collection_menu = "hexagonal_menus"

type menuRepositoryDB struct {
	db *mongo.Database
}

func NewMenuRepositoryDB(db *mongo.Database) MenuRepository {
	return menuRepositoryDB{db: db}
}

func (r menuRepositoryDB) GetAllMenu() ([]models.MenuModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := bson.A{
		bson.D{{"$match", bson.D{{"status", true}}}},
		bson.D{{"$sort", bson.D{{"index", 1}}}},
	}

	result := []models.MenuModel{}
	cursor, err := r.db.Collection(collection_menu).Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r menuRepositoryDB) GetMenuByPer() error {
	return nil
}

func (r menuRepositoryDB) CreateMenu(menu models.MenuCreateReqModel) (*models.MenuModel, error) {
	// # data tranfer object service --> repository
	newMenu := models.MenuModel{
		Id:          uuid.New().String(),
		Role:        menu.Role,
		Name:        menu.Name,
		Icon:        menu.Icon,
		Path:        menu.Path,
		PathName:    menu.PathName,
		Description: menu.Description,
		Index:       0,
		Status:      false,
		CreatedDate: time.Now(),
		LastUpdate:  time.Now(),
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := r.db.Collection(collection_menu).InsertOne(ctx, &newMenu)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("menu already exist")
		}
		return nil, err
	}

	var nMenu models.MenuModel
	query := bson.M{"_id": res.InsertedID}

	err = r.db.Collection(collection_menu).FindOne(ctx, query).Decode(&nMenu)
	if err != nil {
		return nil, err
	}

	return &nMenu, nil
}

func (r menuRepositoryDB) UpdateMenu() error {
	return nil
}

func (r menuRepositoryDB) DeleteMenu() error {
	return nil
}
