generate:
	openapi-generator generate -i https://date.nager.at/swagger/v3/swagger.json -g go -o ./dateapi/ --additional-properties=packageName=dateapi