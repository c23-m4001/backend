http:
	go run -tags http . http

# e.g. make jwt-key-gen flag="--force"
jwt-key-gen:
	go run -tags tools . jwt-key-gen $(flag)

# e.g. make migrate flag="--rollback --steps=12"
migrate:
	go run -tags tools . migrate $(flag)

migrate-fresh:
	go run -tags tools . migrate-fresh

# e.g. make migrate-gen filename=create_table_name
migrate-gen:
	go run -tags tools . migrate-gen -f $(filename)

# e.g. make seed name=table_name
seed:
	go run -tags tools . seed $(name)

# e.g. make seed-prodution name=table_name
seed-production:
	go run -tags tools . seed --production $(name)

generate-docs:
	go run capstone/tool/swag fmt -d main.go,./delivery/api && go run capstone/tool/swag init -d ./,./delivery/api --outputTypes json,yaml
