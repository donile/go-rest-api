package services

import "github.com/sarulabs/di"

type Builder struct {
	builder *di.Builder
}

func NewBuilder() *Builder {
	builder, _ := di.NewBuilder()
	return &Builder{
		builder: builder,
	}
}

func (b *Builder) Build() *Container {
	container := b.builder.Build()
	return &Container{
		container: container,
	}
}

func (b *Builder) AddSingleton(implementationType string, factoryFunc func() (interface{}, error)) {
	definition := di.Def{
		Name:  implementationType,
		Build: func(ctn di.Container) (interface{}, error) { return factoryFunc() },
	}
	b.builder.Add(definition)
}

type Container struct {
	container di.Container
}

func (c *Container) GetRequiredService(implementationType string) interface{} {
	return c.container.Get(implementationType)
}
