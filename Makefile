run:
	docker-compose up --build -d
stop:
	docker-compose down -v
swagger:
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/swagger.yaml \
    -l go-server \
    -o /local/go-server
