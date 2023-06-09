<h1 align="center">nft-crud-service<br></h1>
<h4 align="center">NFT list, buy and create app.</h4>

## Layers
This project was builded in hexagonal architecture and hace the next layers:
* Handler
* Dependencies
* Config
* Service
* Domain
* Repository

### External dependencies
* _MySQL_


## How To Use

### Running the project

Before running the project please ensure that all the external dependencies are installed in your system. Then follow the next:

1. First step, in root project path run docker compose to start database

    ```
    docker compose -f docker-compose.yml up -d
    ```

2. Second step, go to mysql docker container terminal to create database tables and also insert users
    
   ```
   > mysql -u root -p 
   Enter password: <insert-root-here>
   ```
   ```
   # Once inside
   > use marketplace; 
   ```
   ```
   # Run the following querys
   CREATE TABLE user (
        id VARCHAR(150) NOT NULL,
        balance FLOAT NOT NULL DEFAULT 100,
        PRIMARY KEY (id)
   );
   
   CREATE TABLE nft (
        id VARCHAR(150) NOT NULL, 
        image VARCHAR(500) NOT NULL,
        description VARCHAR(150) NOT NULL,
        owner VARCHAR(150) NOT NULL,
        co_creators VARCHAR(500),
        created_at VARCHAR(150) NOT NULL,
        created_by VARCHAR(150) NOT NULL,
        PRIMARY KEY (id),
        FOREIGN KEY (owner) REFERENCES user(id) ON UPDATE CASCADE 
   );

   # Mock users
   INSERT INTO user(id) VALUES('user_1'),('user_2'),('user_3'),('user_4'),('user_5');
    ```

3. Run the project itself

    ```
    go run .
    ```

<br>

## Endpoints
### Create NFT 
#### Request
- `user`and `co_creators` must have users valid and inserted in db
- `co_creators` could be nil
   ```shell
   POST /v1/nft
   {
      "image": "some image",
      "description": "some description",
      "user": "user_1",
      "co_creators": ["user_2"]
   }
   ```
#### Response
- NFT entity
   ```shell
   {
      "id": "c0b4e40b-795a-4c3f-9881-79de0e85f600",
      "image": "some image",
      "description": "some description",
      "owner": "user_1",
      "co_creators": [
          "user_2"
      ],
      "created_at": "2023-02-14T20:57:27-03:00",
      "created_by": "user_1"
   }
   ```

### List NFT 
#### Request
- No body required
- List all NFT's in case you dont use query params
- Have `next` and `took` query params to use pagination
- `next` is used to indicate the number of next elements you are going to take
- `took` is used to indicate from where should start to return in order from 0 to n
   ```shell
   GET /v1/nft
   or
   GET /v1/nft?next=1&took=0
   ```
#### Response
- NFT list entity and args to next call
- `next` indicates how much registries left in db to list
- `took` indicates how much elements did you took in last request
   ```shell
   {
      "data": [
          {
              "id": "c0b4e40b-795a-4c3f-9881-79de0e85f600",
              "image": "some image",
              "description": "some description",
              "owner": "user_1",
              "co_creators": [
                  "user_2"
              ],
              "created_at": "2023-02-14T20:57:27-03:00",
              "created_by": "user_1"
          }
      ],
      "next": 0,
      "took": 1
   }
   ```

### Buy NFT 
#### Request
- All fields are mandatory
- `buyer_id` must be a valid one
- `amount` is a float
   ```shell
   POST /v1/nft/buy
   {
      "nft_id": "c4913807-c41a-495b-bc69-d19b9b112f64",
      "buyer_id": "user_3",
      "amount": 20
   }
   ```
#### Response
- NFT entity updated and users list involved in transaction with new balances
   ```shell
   {
      "nft": {
          "id": "c0b4e40b-795a-4c3f-9881-79de0e85f600",
          "image": "some image",
          "description": "some description",
          "owner": "user_3",
          "co_creators": [
              "user_2",
              "user_1"
          ],
          "created_at": "2023-02-14T20:57:27-03:00",
          "created_by": "user_1"
      },
      "users": [
          {
              "id": "user_1",
              "balance": 116
          },
          {
              "id": "user_2",
              "balance": 104
          },
          {
              "id": "user_3",
              "balance": 80
          }
      ]
   }
   ```

<br>

### Running the tests

In order to run the project tests you need to execute the following command:

   ```shell
   go test ./...
   ```