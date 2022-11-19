package domain

import (
	"github.com/aws/aws-lambda-go/events"
)

type EventResponse events.APIGatewayV2HTTPResponse
type EventRequest events.APIGatewayV2HTTPRequest

type Response struct {
	Status  int         `json:"status"`
	Message string      `json:"message"`
	Body    interface{} `json:"body,omitempty"`
	Offset  *int64      `json:"offset,omitempty"`
	Limit   *int64      `json:"limit,omitempty"`
}

type ApiError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}
