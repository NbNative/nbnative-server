package process

import apiV1 "k8s.io/api/core/v1"

type EventProcess interface {
	Process(*apiV1.Event) ([]bool, error)
}
