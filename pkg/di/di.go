package di

import "go.uber.org/dig"

var container = dig.New()

func Provide(constructor interface{}, opts ...dig.ProvideOption) any {
	if err := container.Provide(constructor, opts...); err != nil {
		panic(err)
	}
	return nil
}

func Invoke(function interface{}, opts ...dig.InvokeOption) error {
	return container.Invoke(function, opts...)
}
