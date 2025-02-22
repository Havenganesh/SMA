package db

import (
	"context"
	"fmt"
	"reflect"

	Err "sma/cErrors"

	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *analyticsDB

type analyticsDB struct {
	client   *mongo.Client
	dataBase *mongo.Database
}

// getCollection securely extracts the struct type name and returns the MongoDB collection.
func (db *analyticsDB) getCollection(doc interface{}) (*mongo.Collection, error) {
	if doc == nil {
		return nil, Err.DOCUMENT_CANNOT_BE_NIL
	}

	t := reflect.TypeOf(doc)
	for t.Kind() == reflect.Ptr || t.Kind() == reflect.Slice || t.Kind() == reflect.Array {
		t = t.Elem()
	}
	if t.Kind() != reflect.Struct {
		return nil, Err.DOCUMENTS_MUST_BE_STRUCT_TYPE
	}
	colName := t.Name()
	return db.dataBase.Collection(colName), nil
}

func (db *analyticsDB) InsertOne(doc interface{}) error {
	col, err := db.getCollection(doc)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = col.InsertOne(ctx, doc)
	if err != nil {
		return Err.DATABASE_INSERT_ONE_FAILED
	}
	return nil
}

func (db *analyticsDB) InsertMany(docs []interface{}) error {
	col, err := db.getCollection(docs)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = col.InsertMany(ctx, docs)
	if err != nil {
		log.Println("Mongo DB insert many error : ", err)
		return Err.DATABASE_INSERT_MANY_FAILED
	}
	return nil
}

func (db *analyticsDB) UpdateMany(filter interface{}, docs interface{}) error {
	col, err := db.getCollection(docs)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	_, err = col.UpdateMany(ctx, filter, docs)
	if err != nil {
		log.Println("Mongo DB insert many error : ", err)
		return Err.DATABASE_UPDATE_MANY_FAILED
	}
	return nil
}

func (db *analyticsDB) FindAll(filter interface{}, docs interface{}, opts ...*options.FindOptions) error {
	if !isPointer(docs) {
		return Err.DOCUMENTS_MUST_BE_POINTER_TYPE
	}
	col, err := db.getCollection(docs)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cur, err := col.Find(ctx, filter, opts...)
	if err != nil {
		log.Println("Mongo DB find all error : ", err)
		return Err.DATABASE_FIND_ALL_FAILED
	}
	defer cur.Close(ctx)
	err = cur.All(ctx, docs)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}
func (db *analyticsDB) FindOne(filter interface{}, doc interface{}, opts ...*options.FindOneOptions) error {
	if !isPointer(doc) {
		return Err.DOCUMENTS_MUST_BE_POINTER_TYPE
	}
	col, err := db.getCollection(doc)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cur := col.FindOne(ctx, filter, opts...)
	if cur == nil {
		log.Println("Mongo DB find one error : ", err)
		return Err.DATABASE_FIND_ONE_FAILED
	}
	cur.Decode(doc)
	return nil
}

func (db *analyticsDB) Aggregate(pipeline interface{}, coll interface{}, doc interface{}, opts ...*options.AggregateOptions) error {
	if !isPointer(doc) {
		return Err.DOCUMENTS_MUST_BE_POINTER_TYPE
	}
	col, err := db.getCollection(coll)
	if err != nil {
		return err
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	cursor, err := col.Aggregate(ctx, pipeline, opts...)
	if err != nil {
		return err
	}
	defer cursor.Close(ctx)
	if cursor.Next(ctx) {
		if err := cursor.Decode(doc); err != nil {
			return err
		}
	}
	return nil
}

func (db *analyticsDB) CreateUniqueIndex(doc interface{}, field bson.D) error {
	indexModel := mongo.IndexModel{
		Keys:    field,
		Options: options.Index().SetUnique(true),
	}
	col, err := db.getCollection(doc)
	if err != nil {
		fmt.Println(err)
		return err
	}
	idxName, err := col.Indexes().CreateOne(context.TODO(), indexModel)
	if err != nil {
		fmt.Println(err)

		return err
	}
	fmt.Println("index created for : ", idxName)
	return nil
}

func isPointer(doc interface{}) bool {
	t := reflect.TypeOf(doc)
	if t == nil || t.Kind() != reflect.Ptr {
		return false
	}
	return true
}
