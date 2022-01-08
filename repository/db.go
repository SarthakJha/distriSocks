package repository

import (
	"github.com/SarthakJha/distr-websock/constants"
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

func (rep *MessageRepository) InitMessageConnection() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "dynamodb",
	}))
	db := dynamo.New(sess, &aws.Config{
		Region: aws.String(constants.AWS_REGION),
	})
	table := db.Table(constants.USER_TABLE_NAME)
	rep.DB = db
	rep.Table = &table
}

func (rep *UserRepository) InitUserConnection() {
	sess := session.Must(session.NewSessionWithOptions(session.Options{
		Profile: "dynamodb",
	}))
	db := dynamo.New(sess, &aws.Config{
		Region: aws.String(constants.AWS_REGION),
	})
	table := db.Table(constants.USER_TABLE_NAME)
	rep.DB = db
	rep.Table = &table
}
