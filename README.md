# MoneyExchange    [![Test](https://github.com/lirprocs/moneyExchange/actions/workflows/test.yaml/badge.svg)](https://github.com/lirprocs/moneyExchange/actions/workflows/test.yaml)

## Описание
Проект MoneyExchange предоставляет REST сервис по расчету всех вариантов размена для указанной суммы денег.

На вход принимается HTTP запрос в формате:
```json
{
  "amount": 400,
  "banknotes": [5000, 2000, 1000, 500, 200, 100, 50]
}
```

## Установка
1. Клонируйте репозиторий
```bash
git clone https://github.com/lirprocs/moneyExchange.git
```
2. Перейдите в директорию проекта (Не нужно, еслу уже находитесь в ней):
```bash
cd moneyExchange
```
3. Установите зависимости:
```bash
go mod tidy
```

## Запуск
1. Перейдите в директорию проекта (Не нужно, еслу уже находитесь в ней):
```bash
cd moneyExchange
```
2. Запустите сервер (можно использовать флаг "--config" или переменную окружения "CONFIG_PATH"):
```bash
go run cmd/main.go --config ./config/local.yaml 
```
3. Отправьте POST запрос на сервер по адресу:
```
localhost:8088/exchange 
```
(Можно использовать Postman)

4. Для остановки сервера можно использовать комбинацию клавиш `Ctrl + C` в окне терминала, из которого был запущен сервер. (В проекте реализован graceful shutdown)

## Примичание
В проекте thumbnail настроен автоматический запуск тестов при пуше изменений в репозиторий. При отправке изменений на GitHub, GitHub Actions автоматически запустит тесты для проверки работоспособности кода.
