# ==== Main Default ====

.PHONY: all
all: reset run-admin wait run-proxy

# Default target when just typing "make"
.DEFAULT_GOAL := all

# ==== Reset everything ====

.PHONY: reset
reset:
	docker compose down
	docker compose up --build -d

# ==== Run fake admin server ====

.PHONY: run-admin
run-admin:
	go run admin_mock.go

# ==== Wait for admin server to be ready ====

.PHONY: wait
wait:
	@echo "Waiting for admin server to be ready..."
	@sleep 2

# ==== Run Prequal proxy server ====

.PHONY: run-proxy
run-proxy:
	go run cmd/main.go

# ==== Cleanup ====

.PHONY: down
down:
	docker compose down
	@if [ -f admin.pid ]; then kill `cat admin.pid` && rm admin.pid; fi

# ==== RabbitMQ (optional) ====

.PHONY: run-rabbitmq
run-rabbitmq:
	docker run -d --rm --name rabbitmq \
		-p 5672:5672 -p 15672:15672 \
		rabbitmq:4.0-management
