package mongo

import (
	"context"

	"github.com/opentracing/opentracing-go"
	"github.com/pkg/errors"
	"github.com/rezaAmiri123/test-microservice/message_service/domain/email"
	"github.com/rezaAmiri123/test-microservice/pkg/logger"
	"github.com/rezaAmiri123/test-microservice/pkg/mongodb"
	"github.com/rezaAmiri123/test-microservice/pkg/pagnation"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	log logger.Logger
	cfg *mongodb.Config
	db  *mongo.Client
}

func NewMongoRepository(log logger.Logger, cfg *mongodb.Config, db *mongo.Client) *mongoRepository {
	return &mongoRepository{log: log, cfg: cfg, db: db}
}
func (p *mongoRepository) Create(ctx context.Context, e *email.Email) error {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.Create")
	defer span.Finish()

	//collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	_, err := p.getCollection().InsertOne(ctx, e, &options.InsertOneOptions{})
	if err != nil {
		p.traceErr(span, err)
		return errors.Wrap(err, "InsertOne")
	}

	return nil
}
func (p *mongoRepository) getCollection() *mongo.Collection {
	return p.db.Database(p.cfg.DB).Collection("emails")
}

func (p *mongoRepository) GetByUUID(ctx context.Context, uuid string) (*email.Email, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.GetProductById")
	defer span.Finish()

	//collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)

	var e email.Email
	if err := p.getCollection().FindOne(ctx, bson.M{"uuid": uuid}).Decode(&e); err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Decode")
	}

	return &e, nil
}

func (p *mongoRepository) List(ctx context.Context, query *pagnation.Pagination) (*email.EmailList, error) {
	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.List")
	defer span.Finish()

	//collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)
	//
	//var e email.Email
	//if err := p.getCollection().FindOne(ctx, bson.M{"uuid": uuid}).Decode(&e); err != nil {
	//	p.traceErr(span, err)
	//	return nil, errors.Wrap(err, "Decode")
	//}
	count, err := p.getCollection().CountDocuments(ctx, bson.D{})
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "CountDocuments")
	}
	if count == 0 {
		return &email.EmailList{}, nil
	}

	limit := int64(query.GetLimit())
	skip := int64(query.GetOffset())
	cursor, err := p.getCollection().Find(ctx, bson.D{}, &options.FindOptions{
		Limit: &limit,
		Skip:  &skip,
	})
	if err != nil {
		p.traceErr(span, err)
		return nil, errors.Wrap(err, "Find")
	}
	defer cursor.Close(ctx) // nolint: errcheck

	emails := make([]*email.Email, 0, query.GetSize())

	for cursor.Next(ctx) {
		var e email.Email
		if err := cursor.Decode(&e); err != nil {
			p.traceErr(span, err)
			return nil, errors.Wrap(err, "Find")
		}
		emails = append(emails, &e)
	}

	if err := cursor.Err(); err != nil {
		span.SetTag("error", true)
		span.LogKV("error_code", err.Error())
		return nil, errors.Wrap(err, "cursor.Err")
	}
	//return models.NewProductListWithPagination(products, count, pagination), nil
	//res := &article.ArticleList{}
	//	res.TotalCount = int64(totalCount)
	//	res.TotalPages = int64(query.GetTotalPages(int(totalCount)))
	//	res.Page = int64(query.GetPage())
	//	res.Size = int64(query.GetSize())
	//	res.HasMore = query.GetHasMore(int(totalCount))
	//	res.Articles = articles
	res := &email.EmailList{}
	res.TotalCount = count
	res.TotalPages = int64(query.GetTotalPages(int(count)))
	res.Page = int64(query.GetPage())
	res.Size = int64(query.GetSize())
	res.HasMore = query.GetHasMore(int(count))
	res.Emails = emails

	return res, nil
}

//func (p *mongoRepository) UpdateProduct(ctx context.Context, product *models.Product) (*models.Product, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.UpdateProduct")
//	defer span.Finish()
//
//	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)
//
//	ops := options.FindOneAndUpdate()
//	ops.SetReturnDocument(options.After)
//	ops.SetUpsert(true)
//
//	var updated models.Product
//	if err := collection.FindOneAndUpdate(ctx, bson.M{"_id": product.ProductID}, bson.M{"$set": product}, ops).Decode(&updated); err != nil {
//		p.traceErr(span, err)
//		return nil, errors.Wrap(err, "Decode")
//	}
//
//	return &updated, nil
//}
//
//func (p *mongoRepository) GetProductById(ctx context.Context, uuid uuid.UUID) (*models.Product, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.GetProductById")
//	defer span.Finish()
//
//	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)
//
//	var product models.Product
//	if err := collection.FindOne(ctx, bson.M{"_id": uuid.String()}).Decode(&product); err != nil {
//		p.traceErr(span, err)
//		return nil, errors.Wrap(err, "Decode")
//	}
//
//	return &product, nil
//}
//
//func (p *mongoRepository) DeleteProduct(ctx context.Context, uuid uuid.UUID) error {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.DeleteProduct")
//	defer span.Finish()
//
//	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)
//
//	return collection.FindOneAndDelete(ctx, bson.M{"_id": uuid.String()}).Err()
//}
//
//func (p *mongoRepository) Search(ctx context.Context, search string, pagination *utils.Pagination) (*models.ProductsList, error) {
//	span, ctx := opentracing.StartSpanFromContext(ctx, "mongoRepository.Search")
//	defer span.Finish()
//
//	collection := p.db.Database(p.cfg.Mongo.Db).Collection(p.cfg.MongoCollections.Products)
//
//	filter := bson.D{
//		{Key: "$or", Value: bson.A{
//			bson.D{{Key: "name", Value: primitive.Regex{Pattern: search, Options: "gi"}}},
//			bson.D{{Key: "description", Value: primitive.Regex{Pattern: search, Options: "gi"}}},
//		}},
//	}
//
//	count, err := collection.CountDocuments(ctx, filter)
//	if err != nil {
//		p.traceErr(span, err)
//		return nil, errors.Wrap(err, "CountDocuments")
//	}
//	if count == 0 {
//		return &models.ProductsList{Products: make([]*models.Product, 0)}, nil
//	}
//
//	limit := int64(pagination.GetLimit())
//	skip := int64(pagination.GetOffset())
//	cursor, err := collection.Find(ctx, filter, &options.FindOptions{
//		Limit: &limit,
//		Skip:  &skip,
//	})
//	if err != nil {
//		p.traceErr(span, err)
//		return nil, errors.Wrap(err, "Find")
//	}
//	defer cursor.Close(ctx) // nolint: errcheck
//
//	products := make([]*models.Product, 0, pagination.GetSize())
//
//	for cursor.Next(ctx) {
//		var prod models.Product
//		if err := cursor.Decode(&prod); err != nil {
//			p.traceErr(span, err)
//			return nil, errors.Wrap(err, "Find")
//		}
//		products = append(products, &prod)
//	}
//
//	if err := cursor.Err(); err != nil {
//		span.SetTag("error", true)
//		span.LogKV("error_code", err.Error())
//		return nil, errors.Wrap(err, "cursor.Err")
//	}
//
//	return models.NewProductListWithPagination(products, count, pagination), nil
//}

func (p *mongoRepository) traceErr(span opentracing.Span, err error) {
	span.SetTag("error", true)
	span.LogKV("error_code", err.Error())
}
