# SNS BACKEND

## 起動方法
```shell
docker compose up
```

## API
### sign up
    POST http://localhost:8080/auth/signup
    Content-Type: application/json

    {
        "userid": null,
        "nickname": "test_nickname",
        "username": "test_username",
        "password": "test_password"
    }

## DB
