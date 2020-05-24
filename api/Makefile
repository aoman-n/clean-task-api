.PHONY: integ-test
integ-test:
	GO_ENV=test \
	DB_USER=root \
	DB_NAME=test_db \
	DB_PASSWORD=password \
	DB_HOST=mysql \
	DB_PORT=3306 \
	SIGNINGKEY=signingkey \
	FRONT_ORIGIN=http://localhost:3000 \
	BFF_ORIGIN=http://front:3000 \
	go test -v -tags=integration ./...

.PHONY: unit-test
unit-test:
	go test -v -tags=unit ./...
