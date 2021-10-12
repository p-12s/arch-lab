# БД пусть будет отдельно (в моем случае - на локале)
### При запуске сервиса в docker-compose/swarm это делать не нужно

Запускаем миграции:
```
migrate -path ./schema -database 'postgres://postgres:qwerty@localhost:5432/sync-monolit?sslmode=disable' up
# вместо up -> down, если что-то пошло не так
```
Можно убедиться, что БД с таблицами созданы:
```
psql -U postgres
\d
select * from users;
```

Запуск приложения:
```
go run ./cmd/main.go
```
