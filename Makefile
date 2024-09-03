REPOSITORY=europe-west1-docker.pkg.dev/optidate-411610/optidate

generate:
	openapi-generator generate -i https://date.nager.at/swagger/v3/swagger.json -g go -o ./dateapi/ --additional-properties=packageName=dateapi

build-docker:
	docker build \
	--build-arg PORT=8000 \
	-t ${REPOSITORY}/optidate:latest \
	.

build-frontend:
	bun run --cwd ./frontend build

deploy-docker:
	docker push ${REPOSITORY}/optidate:latest
