run:
	docker-compose -f docker-compose.yml up --build -d
prod:
	docker-compose -f docker-compose.prod.yml up --build -d
stop:
	docker-compose down -v --remove-orphans
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
