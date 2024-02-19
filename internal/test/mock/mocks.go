package test

import (
	"totoconfigapi/internal/db"

	"github.com/stretchr/testify/mock"
)

// Mock for the ConfigDB interface
type MockConfigDB struct {
	mock.Mock
}

func (m *MockConfigDB) GetConfigs(pkg string, country string, randomPercentile int) ([]*db.Config, error) {
	args := m.Called(pkg, country, randomPercentile)
	return args.Get(0).([]*db.Config), args.Error(1)
}
