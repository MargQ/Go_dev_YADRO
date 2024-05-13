# Задание №5: Тестирование и линтинг в Go

## Цель: Покрыть заранее предоставленный код тестами на уровне 80% покрытия, использовать mock-тесты для проверки различных сценариев, а также выполнить анализ кода с помощью golangci-lint и исправить выявленные замечания.

## Задание:
В данной работе предстоит выполнить следующие задачи:

1. Тестирование: Написать тесты для заранее предоставленного кода с целью достижения покрытия минимум 80%.
2. Линтинг с помощью golangci-lint: Запустить golangci-lint на предоставленном коде для выявления потенциальных проблем, стилевых нарушений и ошибок. Исправить все выявленные замечания.

## Выполнение задания
### Тесты
Разработаны тесты для пакетов [wifi](https://github.com/100thKing/Go_DEV_School/blob/main/5.%20%D0%A2%D0%B5%D1%81%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5%20%D0%B8%20%D0%BB%D0%B8%D0%BD%D1%82%D0%B8%D0%BD%D0%B3%20%D0%B2%20Go/internal/wifi/wi-fi_test.go) и [db](https://github.com/100thKing/Go_DEV_School/blob/main/5.%20%D0%A2%D0%B5%D1%81%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5%20%D0%B8%20%D0%BB%D0%B8%D0%BD%D1%82%D0%B8%D0%BD%D0%B3%20%D0%B2%20Go/internal/db/db_functions_test.go) с применением библиотек testify/assert и sqlmock и покрытием более 80%:

## Вывод:
<details><summary>db</summary>
  <img src = "https://github.com/100thKing/Go_DEV_School/blob/main/5.%20%D0%A2%D0%B5%D1%81%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5%20%D0%B8%20%D0%BB%D0%B8%D0%BD%D1%82%D0%B8%D0%BD%D0%B3%20%D0%B2%20Go/source/Pasted%20image%201.png">
</details>
<details><summary>wifi</summary>
  <img src = "https://github.com/100thKing/Go_DEV_School/blob/main/5.%20%D0%A2%D0%B5%D1%81%D1%82%D0%B8%D1%80%D0%BE%D0%B2%D0%B0%D0%BD%D0%B8%D0%B5%20%D0%B8%20%D0%BB%D0%B8%D0%BD%D1%82%D0%B8%D0%BD%D0%B3%20%D0%B2%20Go/source/Pasted%20image.png">
</details>
