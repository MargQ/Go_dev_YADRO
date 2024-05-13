package db

import (
	"strings"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	// Создаём новую db для теста в виде мока
	db, _, _ := sqlmock.New()
	service := New(db)
	assert.Equal(t, service.DB, db)
}

func TestGetNames(t *testing.T) {
	// Подготовка тестовых данных
	expectedNames := []string{"John", "Jane", "Jack"}
	expectedType := []string{}

	db, mock, _ := sqlmock.New()

	// Создаём ожидание запроса к БД
	mock.ExpectQuery("SELECT name FROM users").WillReturnRows(
		sqlmock.NewRows([]string{"name"}).FromCSVString(strings.Join(expectedNames, "\n")),
	)
	service := Service{DB: db}

	gotNames, _ := service.GetNames()

	// Запуск теста
	assert.Equal(t, gotNames, expectedNames)
	assert.IsType(t, expectedType, gotNames)
}

func TestSelectUniqueValues(t *testing.T) {
	// Подготовка тестовых данных
	columnName := "unique_column"
	tableName := "unique_table"
	expectedValues := []string{"value1", "value2", "value3"}

	db, mock, _ := sqlmock.New()

	// Создаём ожидание запроса к БД
	mock.ExpectQuery("SELECT DISTINCT " + columnName + " FROM " + tableName).WillReturnRows(
		sqlmock.NewRows([]string{"unique_column"}).FromCSVString(strings.Join(expectedValues, "\n")),
	)
	service := Service{DB: db}

	gotValues, _ := service.SelectUniqueValues(columnName, tableName)

	// Запуск теста
	assert.Equal(t, gotValues, expectedValues)
}
