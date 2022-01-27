package repository

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/guregu/dynamo"
)

type MessageRepository struct {
	DB    *dynamo.DB
	Table *dynamo.Table
}

type UserRepository struct {
	DB    *dynamo.DB
	Table *dynamo.Table
}

func (rep *MessageRepository) InitMessageConnection(region, tableName string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "dynamodb",
	}))
	db := dynamo.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	table := db.Table(tableName)
	rep.DB = db
	rep.Table = &table
}

func (rep *UserRepository) InitUserConnection(region, tableName string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "dynamodb",
	}))
	db := dynamo.New(sess, &aws.Config{
		Region: aws.String(region),
	})
	table := db.Table(tableName)
	rep.DB = db
	rep.Table = &table
}
