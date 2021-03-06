package es

import "encoding/json"

type BulkResponse struct {
	Errors bool               `json:"errors"`
	Items  []BulkResponseItem `json:"items"`
}

type BulkResponseItem struct {
	Index BulkIndexResponse `json:"index"`
}

type BulkIndexResponse struct {
	ID     string `json:"_id"`
	Result string `json:"result"`
	Status int    `json:"status"`
	Error  struct {
		Type   string `json:"type"`
		Reason string `json:"reason"`
		Cause  struct {
			Type   string `json:"type"`
			Reason string `json:"reason"`
		} `json:"caused_by"`
	} `json:"error"`
}

type es7searchResponse struct {
	Hits es7its                         `json:"hits"`
	Aggs map[string]AggregationResponse `json:"aggregations,omitempty"`
}

type es7its struct {
	Total struct {
		Value int `json:"value"`
	} `json:"total"`
	Hits []Hit `json:"hits"`
}

type SearchBody struct {
	Indices        []string           `json:"-"`
	Aggregations   json.RawMessage    `json:"aggs,omitempty"`
	Query          *Query             `json:"query,omitempty"`
	Sort           []map[string]Order `json:"sort,omitempty"`
	Size           int                `json:"size"`
	TerminateAfter int                `json:"terminate,omitempty"`
	SearchAfter    []interface{}      `json:"search_after,omitempty"`
}

type Order string

const (
	AscOrder Order = "asc"
)

type Query struct {
	Term         *Terms                        `json:"term,omitempty"`
	RangeQueries map[string]RangeQuery         `json:"range,omitempty"`
	BoolQuery    map[BoolQueryType][]BoolQuery `json:"bool,omitempty"`
}

type BoolQuery struct {
	Term         map[string]string             `json:"term,omitempty"`
	Regexp       map[string]TermQuery          `json:"regexp,omitempty"`
	Nested       *NestedQuery                  `json:"nested,omitempty"`
	BoolQuery    map[BoolQueryType][]BoolQuery `json:"bool,omitempty"`
	RangeQueries map[string]RangeQuery         `json:"range,omitempty"`
	MatchQueries map[string]MatchQuery         `json:"match,omitempty"`
}

type MatchQuery struct {
	Query string `json:"query"`
}

type NestedQuery struct {
	Path  string `json:"path"`
	Query Query  `json:"query"`
}

type Terms map[string]TermQuery

type TermQuery struct {
	Value string `json:"value"`
}

type BoolQueryType string

// Must defines must bool query type.
const Must BoolQueryType = "must"

// Should defines should bool query type.
const Should BoolQueryType = "should"

type RangeQuery struct {
	GTE interface{} `json:gte`
	LTE interface{} `json:lte`
}

type SearchResponse struct {
	Hits  Hits                           `json:"hits"`
	Aggs  map[string]AggregationResponse `json:"aggregations,omitempty"`
	Error *SearchResponseError           `json:"error,omitempty"`
}
type Hits struct {
	Total int   `json:"total"`
	Hits  []Hit `json:"hits"`
}
type Hit struct {
	Source *json.RawMessage `json:"_source"`
}

type AggregationResponse struct {
	Buckets []struct {
		Key string `json:"key"`
	} `json:"buckets"`
}
type SearchResponseError struct {
	json.RawMessage
}
