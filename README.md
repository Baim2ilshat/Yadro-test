# Тестовое задание Yadro

## Требования
- Go (1.20 или выше)
- unit тесты (опционально)

## Инструкция по запуску
### Запуск программы (в папке Yadro-test)
```
go run main.go
```

### Запуск тестов
```
go test ./...
```

### Пример запуска программы
```
[09:31:49.285] The competitor(3) registered
[09:32:17.531] The competitor(2) registered
[09:37:47.892] The competitor(5) registered
[09:38:28.673] The competitor(1) registered
[09:39:25.079] The competitor(4) registered
[09:55:00.000] The start time for the competitor(1) was set by a draw to 10:00:00.000
[09:56:30.000] The start time for the competitor(2) was set by a draw to 10:01:30.000
[09:58:00.000] The start time for the competitor(3) was set by a draw to 10:03:00.000
...
```

### Пример запуска тестов
```
ok      Yadro-test/config       (cached)
ok      Yadro-test/models       2.161s
ok      Yadro-test/processor    2.301s
```
