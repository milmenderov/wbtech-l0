## App - L0
Сервис для записи в базу заказов и быстрого их просмотра
### Необходимые переменные для запуска:
| Name            | Default value              |
|-----------------|----------------------------|
| DB_HOST         | db                         |
| DB_PORT         | 5432                       |
| DB_USER         | postgres                   |
| DB_NAME         | postgres                   |
| DB_PASSWORD     | qwerty                     |
| APP_PORT        | 8080                       |
| NATS_URL        | nats://nats-streaming:4222 |
| NATS_SUB_NAME   | Orders                     |
| NATS_CLUSTER_ID | test-cluster               |
| NATS_CLIENT_ID  | client-sub                 |
| HTTP_PORT       | 8080                       |
| HTTP_HOST       | 0.0.0.0                    |


### Как запускать
Для старта прописать команду(она сбилдит и запустит необходимые контейнеры)   
```make re```
