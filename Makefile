# Makefile

# Задача по умолчанию
.DEFAULT_GOAL := lint

# Линтинг
lint:
    golangci-lint run