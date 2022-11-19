package util

import (
	"context"
	"encoding/json"
	"log"

	"github.com/aws/aws-lambda-go/lambda"

	"github.com/caospinac/aws-sls-go/domain"
)

type Handler func(context.Context, domain.EventRequest) (Response, ApiError)
type lambdaHandler func(context.Context, domain.EventRequest) (*domain.EventResponse, error)

func getLambdaHandler(handler Handler) lambdaHandler {
	return func(ctx context.Context, event domain.EventRequest) (*domain.EventResponse, error) {
		eventJSON, _ := json.Marshal(event)
		log.Printf("event:%s", string(eventJSON))

		var response *domain.EventResponse

		resBuilder, err := handler(ctx, event)
		if err != nil {
			log.Printf("request_id:%s, error:%s", event.RequestContext.RequestID, err.getMessage())
			response = err.build()
		} else {
			response = resBuilder.build()
		}

		log.Printf("request_id:%s, status:%d", event.RequestContext.RequestID, response.StatusCode)

		return response, nil
	}
}

func Start(handler Handler) {
	lambda.Start(getLambdaHandler(handler))
}
