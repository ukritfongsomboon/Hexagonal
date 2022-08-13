package repositories

import (
	"context"
	"errors"
	"hexagonal/core/models"

	"time"

	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var collection_user = "hexagonal_users_test"

type userRepositoryDB struct {
	db *mongo.Database
}

func NewUserRepositoryDB(db *mongo.Database) UserRepository {
	return userRepositoryDB{db: db}
}

func (r userRepositoryDB) GetAll(p models.UserPaginationModel) ([]models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	if p.Row == 0 {
		p.Row = 10
	}

	// refs https://www.codementor.io/@arpitbhayani/fast-and-efficient-pagination-in-mongodb-9095flbqr
	query := bson.A{
		bson.D{{"$skip", p.Row * (p.Page - 1)}},
		bson.D{{"$limit", p.Row}},
		bson.D{{"$sort", bson.D{{"create_date", 1}}}},
	}
	result := []models.UserModel{}
	cursor, err := r.db.Collection(collection_user).Aggregate(ctx, query)
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &result); err != nil {
		return nil, err
	}

	return result, nil
}

func (r userRepositoryDB) GetById(UserID string) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := bson.D{{"user_id", UserID}}
	result := models.UserModel{}
	err := r.db.Collection(collection_user).FindOne(ctx, query).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r userRepositoryDB) Create(payload models.UserCreateModel) (*models.UserModel, error) {
	// # New version 2
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	id := uuid.New().String()
	user := models.UserModel{
		CreatedDate: time.Now(),
		LastUpdate:  time.Now(),
		Email:       payload.Email,
		Name:        payload.Name,
		// Password:    "",
		Status: true,
		Role:   0,
		UserID: id,
		Oauth: []models.UserOauthModel{
			{
				Provider: payload.Provider,
				Id:       id,
				Email:    payload.Email,
				Password: payload.Password,
			},
		},
	}

	res, err := r.db.Collection(collection_user).InsertOne(ctx, &user)
	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("email already exist")
		}
		return nil, err
	}

	// refs https://pkg.go.dev/go.mongodb.org/mongo-driver/mongo#IndexView.CreateMany
	// TODO Declare an array of bsonx models for the indexes
	indexField := []mongo.IndexModel{
		{
			Keys: bson.D{{Key: "email", Value: 1}}, Options: options.Index().SetUnique(true),
		},
		{
			Keys: bson.D{{Key: "user_id", Value: 1}}, Options: options.Index().SetUnique(true),
		},
	}

	if _, err := r.db.Collection(collection_user).Indexes().CreateMany(ctx, indexField); err != nil {
		// TODO แก้ไข เมื่อ setindex ไม่ได้ ไม่ต้อง errorออกมา หรือแก้เป็น log warning
		return nil, errors.New("could not create index")
	}
	var newUser models.UserModel
	query := bson.M{"_id": res.InsertedID}

	err = r.db.Collection(collection_user).FindOne(ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r userRepositoryDB) GetCountAll() (int64, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	filter := bson.D{{"status", true}}
	count, err := r.db.Collection(collection_user).CountDocuments(ctx, filter)
	if err != nil {
		return 0, err
	}
	return count, err
}

func (r userRepositoryDB) GetByEmail(email string) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	query := bson.D{{"email", email}}
	result := models.UserModel{}
	err := r.db.Collection(collection_user).FindOne(ctx, query).Decode(&result)
	if err != nil {
		return nil, err
	}
	return &result, nil
}

func (r userRepositoryDB) Update(u models.UserUpdateReqModel) (*models.UserModel, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	filter := bson.D{{"user_id", u.UserID}}
	update := bson.D{{"$set", bson.D{{"email", u.Email}, {"name", u.Name}, {"update_date", time.Now()}}}}

	_, err := r.db.Collection(collection_user).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUser models.UserModel
	query := bson.M{"user_id": u.UserID}

	err = r.db.Collection(collection_user).FindOne(ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}

func (r userRepositoryDB) UpdateImage(u models.UserUpdateImgReqModel) (*models.UserModel, error) {
	// # Update file name to Database
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	filter := bson.D{{"user_id", u.UserID}}
	update := bson.D{{"$set", bson.D{{"image", u.Filename}, {"update_date", time.Now()}}}}

	_, err := r.db.Collection(collection_user).UpdateOne(ctx, filter, update)
	if err != nil {
		return nil, err
	}

	var newUser models.UserModel
	query := bson.M{"user_id": u.UserID}

	err = r.db.Collection(collection_user).FindOne(ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}

	return &newUser, nil
}
