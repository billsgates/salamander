run:
	docker-compose up --build -d
stop:
	docker-compose down -v
swagger_v1:
	docker run --rm -v ${PWD}:/local swaggerapi/swagger-codegen-cli-v3 generate \
    -i /local/swagger_v1.yaml \
    -l go-server \
    -o /local/go-server
