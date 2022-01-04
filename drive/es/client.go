package es

import (
	"context"
)

type ElasticsearchClient interface {
	Search(ctx context.Context, query SearchBody, size int, indices ...string) (*SearchResponse, error)
}
