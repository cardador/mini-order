package store

import (
	"context"
	"fmt"
	"interview/order/model"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb/types"
)

type DynamoStore struct {
	client    *dynamodb.Client
	tableName string
}

func NewDynamoStore(client *dynamodb.Client, table string) *DynamoStore {
	return &DynamoStore{client: client, tableName: table}
}

func (s *DynamoStore) SaveOrder(ctx context.Context, o model.Order) error {
	item, err := attributevalue.MarshalMap(o)
	if err != nil {
		return err
	}

	_, err = s.client.PutItem(ctx, &dynamodb.PutItemInput{
		TableName:           aws.String(s.tableName),
		Item:                item,
		ConditionExpression: aws.String("attribute_not_exists(id)"),
	})
	return err

}

func (s *DynamoStore) GetOrder(ctx context.Context, id string) (model.Order, error) {
	var o model.Order
	result, err := s.client.GetItem(ctx, &dynamodb.GetItemInput{
		TableName: aws.String(s.tableName),
		Key: map[string]types.AttributeValue{
			"id": &types.AttributeValueMemberS{Value: id},
		},
	})

	if err != nil {
		return o, err
	}

	if result.Item == nil {
		return o, fmt.Errorf("order not found")
	}

	err = attributevalue.UnmarshalMap(result.Item, &o)
	return o, err
}
