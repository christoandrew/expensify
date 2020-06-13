package db

type Registry interface {
	Register(k interface{}) *Registry
	RegisterAll(k interface{}) *Registry
}

type ModelRegistry interface {
	GetModels() []interface{}
}

type DefaultModelRegistry struct {
	models []interface{}
}

func NewModelRegistry() *DefaultModelRegistry {
	return &DefaultModelRegistry{}
}

func (r *DefaultModelRegistry) GetModels() []interface{} {
	return r.models
}

func (r *DefaultModelRegistry) Register(model interface{}) {
	r.models = append(r.models, model)
}

func (r *DefaultModelRegistry) RegisterAll(models ...interface{}) {
	for _, model := range models {
		r.Register(model)
	}
}
