run:
	docker-compose up --build -d
stop:
	docker-compose down -v
log:
	docker-compose logs --follow --tail 1 backend
test:
	cd go-server && go test ./...
mockery:
	cd go-server/domain && mockery --all --keeptree
swagger:
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/swagger.yaml \
    -l go-server \
    -o /local/go-server

mockery:
	cd go-server/domain && mockery --all --keeptree
