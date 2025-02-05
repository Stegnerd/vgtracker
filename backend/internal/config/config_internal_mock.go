package config

type InternalMock struct {
	NewConfigFunc  func() (*Config, error)
	LoadConfigFunc func() (*Config, error)
	SaveConfigFunc func(saveArr []byte) error
}

var _ ConfigInternalMethods = &InternalMock{}

func (i InternalMock) NewConfig() (*Config, error) {
	return i.NewConfigFunc()
}

func (i InternalMock) LoadConfig() (*Config, error) {
	return i.LoadConfigFunc()
}

func (i InternalMock) SaveConfig(saveArr []byte) error {
	return i.SaveConfigFunc(saveArr)
}
