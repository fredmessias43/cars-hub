package models

import "github.com/fredmessias43/car-hub/src/contracts"

type List[T contracts.Model] struct {
	Value []T
}

func (l *List[T]) ToMap() []map[string]interface{} {
	var result []map[string]interface{}
	for _, item := range l.Value {
		result = append(result, item.ToMap())
	}
	return result
}
