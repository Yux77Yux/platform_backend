package client

import (
	"log"
	"sync"
)

var (
	service_address = ":8080"

	aggregator_client  *AggregatorClient
	user_client        *UserClient
	auth_client        *AuthClient
	creation_client    *CreationClient
	interaction_client *InteractionClient
	comment_client     *CommentClient
	review_client      *ReviewClient
)

func Close() {
	user_client.Close()
	auth_client.Close()
	creation_client.Close()
	interaction_client.Close()
	comment_client.Close()
	aggregator_client.Close()
	review_client.Close()
}

func GetAggregatorClient() (*AggregatorClient, error) {
	if aggregator_client != nil {
		return aggregator_client, nil
	}
	aggregator_client, err := NewAggregatorClient()
	if err != nil {
		log.Printf("error: creation client %v", err)
		return nil, err
	}
	return aggregator_client, nil
}

func GetCreationClient() (*CreationClient, error) {
	if creation_client != nil {
		return creation_client, nil
	}
	creation_client, err := NewCreationClient()
	if err != nil {
		log.Printf("error: creation client %v", err)
		return nil, err
	}
	return creation_client, nil
}

func GetUserClient() (*UserClient, error) {
	if user_client != nil {
		return user_client, nil
	}
	user_client, err := NewUserClient()
	if err != nil {
		log.Printf("error: user client %v", err)
		return nil, err
	}
	return user_client, nil
}

func GetReviewClient() (*ReviewClient, error) {
	if review_client != nil {
		return review_client, nil
	}
	review_client, err := NewReviewClient()
	if err != nil {
		log.Printf("error: review client %v", err)
	}
	return review_client, nil
}

func GetInteractionClient() (*InteractionClient, error) {
	if interaction_client != nil {
		return interaction_client, nil
	}
	interaction_client, err := NewInteractionClient()
	if err != nil {
		log.Printf("error: interaction client %v", err)
	}
	return interaction_client, nil
}

func GetCommentClient() (*CommentClient, error) {
	if comment_client != nil {
		return comment_client, nil
	}
	comment_client, err := NewCommentClient()
	if err != nil {
		log.Printf("error: comment client %v", err)
	}
	return comment_client, nil
}

func GetAuthClient() (*AuthClient, error) {
	if auth_client != nil {
		return auth_client, nil
	}
	auth_client, err := NewAuthClient()
	if err != nil {
		log.Printf("error: auth client %v", err)
	}
	return auth_client, nil
}

// 使用了envoy，所以使用envoy地址即可
func Run() func() {
	var wg sync.WaitGroup

	wg.Add(7)

	go func() {
		defer wg.Done()
		var err error
		aggregator_client, err = NewAggregatorClient()
		if err != nil {
			log.Printf("error: aggregator client %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		user_client, err = NewUserClient()
		if err != nil {
			log.Printf("error: user client %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		auth_client, err = NewAuthClient()
		if err != nil {
			log.Printf("error: auth client %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		creation_client, err = NewCreationClient()
		if err != nil {
			log.Printf("error: creation client %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		interaction_client, err = NewInteractionClient()
		if err != nil {
			log.Printf("error: interaction client %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		comment_client, err = NewCommentClient()
		if err != nil {
			log.Printf("error: comment client %v", err)
		}
	}()

	go func() {
		defer wg.Done()
		var err error
		review_client, err = NewReviewClient()
		if err != nil {
			log.Printf("error: review client %v", err)
		}
	}()

	wg.Wait() // 等待所有 gRPC 客户端初始化完成

	return func() {
		Close()
	}
}
