package factories

import "github.com/brianvoe/gofakeit/v6"

type UserFactory struct {
}

// Definition Defines the model's default state.
func (f *UserFactory) Definition() any {
	faker := gofakeit.New(0)
	return map[string]interface{}{
		"name":       faker.Name(),
		"avatar":     faker.Email(),
		"created_at": faker.Date(),
		"updated_at": faker.Date(),
	}
}
