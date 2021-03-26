package users

import (
	"fmt"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

func RetrieveByHashKey(hashKey string) (User, error)  {
	sess := session.Must(session.NewSession())
	svc := dynamodb.New(sess)
	user := User{}

}