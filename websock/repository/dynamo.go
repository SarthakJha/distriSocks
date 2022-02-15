package repository

import (
	"os"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
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
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"), ""),
			Region:      aws.String(region),
		},
	}))
	db := dynamo.New(sess)
	table := db.Table(tableName)
	rep.DB = db
	rep.Table = &table
}

func (rep *UserRepository) InitUserConnection(region, tableName string) {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Config: aws.Config{
			Credentials: credentials.NewStaticCredentials(os.Getenv("ACCESS_KEY_ID"), os.Getenv("ACCESS_KEY_SECRET"), ""),
			Region:      aws.String(region),
		},
	}))
	db := dynamo.New(sess)
	table := db.Table(tableName)
	rep.DB = db
	rep.Table = &table
}
