# Variáveis de configuração
COVERAGE=coverage.out
PKGS_NO_MOCKS=$(shell go list ./... | grep -v '/mocks')
MIN_COVERAGE=80

default: start

start: 
	go run cmd/main.go


.PHONY: coverage
coverage: test-covout
	@echo "Coverage total (sem mocks):"
	@go tool cover -func=$(COVERAGE) | grep total

.PHONY: coverage-check
coverage-check: coverage
	@echo "Verificando se a cobertura atende o mínimo de $(MIN_COVERAGE)%..."
	@COVERAGE_PCT=$$(go tool cover -func=$(COVERAGE) | grep total | awk '{print $$3}' | sed 's/%//'); \
	if [ $$(echo "$$COVERAGE_PCT < $(MIN_COVERAGE)" | bc -l) -eq 1 ]; then \
		echo "❌ Coverage atual: $$COVERAGE_PCT% (mínimo exigido: $(MIN_COVERAGE)%)"; \
		echo "❌ Não é possível iniciar o ambiente de desenvolvimento!"; \
		exit 1; \
	else \
		echo "✅ Coverage atual: $$COVERAGE_PCT% (mínimo exigido: $(MIN_COVERAGE)%)"; \
	fi

test-cov:
	go test -cover $(PKGS_NO_MOCKS)

test-covout:
	go test -coverprofile=coverage.out $(PKGS_NO_MOCKS)

test-html:
	go tool cover -html=coverage.out

test-all: 
	make test-cov && make test-covout && make test-html