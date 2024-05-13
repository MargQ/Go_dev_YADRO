package wifi

import (
	"net"
	"testing"

	"github.com/mdlayher/wifi"
	"github.com/stretchr/testify/assert"
)

type MockWiFi struct {
	interfaces []*wifi.Interface
	err        error
}

func (m *MockWiFi) Interfaces() ([]*wifi.Interface, error) {
	return m.interfaces, m.err
}

func TestGetAddresses(t *testing.T) {
	// Подготовка тестовых данных
	expectedAddresses := []net.HardwareAddr{
		// Можно добавить дополнительные адреса, если необходимо
		{0xaa, 0xbb, 0xcc, 0xdd, 0xee, 0xff},
	}
	mockWiFi := &MockWiFi{
		interfaces: []*wifi.Interface{
			// Можно добавить дополнительные интерфейсы, если необходимо
			{HardwareAddr: expectedAddresses[0]},
		},
		err: nil,
	}
	wifiService := New(mockWiFi)

	// Запуск теста
	actualAddresses, actualErr := wifiService.GetAddresses()
	assert.Equal(t, actualAddresses, expectedAddresses)
	assert.Nil(t, actualErr)
}

func TestGetNames(t *testing.T) {
	// Подготовка тестовых данных
	expectedNames := []string{
		"eth0",
		"wlan0",
	}
	mockWiFi := &MockWiFi{
		interfaces: []*wifi.Interface{
			{Name: expectedNames[0]},
			{Name: expectedNames[1]},
			// Можно добавить дополнительные интерфейсы, если необходимо
		},
		err: nil,
	}
	wifiService := New(mockWiFi)

	// Запуск теста
	actualNames, actualErr := wifiService.GetNames()

	// Проверка результатов
	assert.Equal(t, actualNames, expectedNames)
	assert.Nil(t, actualErr)
}
