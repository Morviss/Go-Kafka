# Go-Kafka


we have created the kafka impementaions with golang using kowl frame work
After created Docker-compose file we have confugired the kafka file and run the command

docker-compose up -d #to up the docker-compose

and added the producer and consumer files and run the commands for the producers and consumers

# producer

go run producer/producer.go producer/types.go producer/main.go 
or 
go run producer/*.go

# Consumer

go run consumer/connection.go consumer/types.go consumer/main.go 
or 
go run consumer/*.go
