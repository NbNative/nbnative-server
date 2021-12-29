package process

import apiV1 "k8s.io/api/core/v1"

type EsProcessor struct {
}

func NewEsProcessor() *EsProcessor {
	return &EsProcessor{}
}
func (e *EsProcessor) Process(event *apiV1.Event) ([]bool, error) {
	return nil, nil
}
