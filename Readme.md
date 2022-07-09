# Base-go

## DB
### Mongo:
  - Run with docker:
    ```bash
    docker run --name mongodb -p 27017:27017 -v /home/loh/mongo/data:/data/db -e MONGO_INITDB_ROOT_USERNAME=root -e MONGO_INITDB_ROOT_PASSWORD=12345678  -d mongo
    ```
