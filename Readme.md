# RSS Agregator

This RSS Aggregator is a web service designed to provide an efficient and flexible way to manage and interact with RSS feeds. It operates exclusively through API endpoints, allowing developers to integrate its functionalities into their applications seamlessly.

## Key Features
* User Management
    - Create User: Register new users to the service.
    - Get User: Retrieve user details with ease.
    - Delete User: Remove user accounts along with associated data.

* Feed Management
    - Post Feed: Add new RSS feeds to the system.
    - Delete Feed: Remove existing feeds managed by users.

* Follow/Unfollow Feeds
    - Users can follow specific feeds to receive updates.
    - Option to unfollow feeds when no longer interested.

* Auto-Scraping
    
    The service automatically scrapes followed feeds, ensuring users always have access to the latest content.

* Post Management
    
    Users can post custom content to their feeds or delete existing posts

## Setup and Starting the web service
### *Installation*
* Start with adding dependencies
    > go mod tidy

    > go mod vendor

### *Database*
In my project I've used PostgresDB so i would recommend you to use the same. Now to create the database and use other sql queries.
* Create a postgres databse using 
    > createdb rssdb

* Create **.env** and add **PORTNO=8080** and **DB_URL=<your_db_url>**

* Go automatically downloads the required dependencies from the vendor directory, which inlcudes **sqlc, goose**. 

* Use goose to do all the migrations, here my project database has six  migrations to do
![alt text](/images/image.png)
   
    > goose <db_url> up

* After migrations use sqlc to make sure the go sql query functions our handled properly
    > sqlc generate

### *Start the web server*

* Run the go build command and start your web server
    > go build

    now execute the .exe file
    > ./rssagg

    this spins off the rss web service ready to be used
    ![alt text](/images/image-1.png)

### *Test the web service using postman*

* Using postman send http requests to the server and see it working
![alt text](/images/image-2.png)

* The api endpoints include:
    * Users endpoints:
### **Users Endpoints:**

- **`v1/user`**: **GET** request to retrieve user details by their API key
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Fetches user details based on the API key provided in the authorization header.
  - **Route**:
    ```go
    v1Router.Get("/user", apiCfg.middlewareAuth(apiCfg.handlerGetUserByApiKey))
    ```

- **`v1/user`**: **POST** request to create a new user
  - **Requires**: Body `{"name": "username"}`
  - **Description**: Creates a new user with the specified name.
  - **Route**:
    ```go
    v1Router.Post("/user", apiCfg.handlerCreateUser)
    ```

- **`v1/user`**: **DELETE** request to delete an existing user
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Deletes a user based on the provided API key.
  - **Route**:
    ```go
    v1Router.Delete("/user", apiCfg.middlewareAuth(apiCfg.handlerDeleteUser))
    ```

- **`v1/users`**: **GET** request to retrieve all users
  - **Description**: Fetches a list of all users.
  - **Route**:
    ```go
    v1Router.Get("/users", apiCfg.handlerGetUsers)
    ```

- **`v1/posts`**: **GET** request to retrieve posts by the authenticated user
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Fetches all posts created by the authenticated user.
  - **Route**:
    ```go
    v1Router.Get("/posts", apiCfg.middlewareAuth(apiCfg.handlerGetUserPosts))
    ```

---

### **Feeds Endpoints:**

- **`v1/feeds`**: **GET** request to retrieve all available feeds
  - **Description**: Fetches a list of all feeds available in the system.
  - **Route**:
    ```go
    v1Router.Get("/feeds", apiCfg.handlerGetFeeds)
    ```

- **`v1/feed`**: **POST** request to create a new feed
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Creates a new feed for the authenticated user.
  - **Route**:
    ```go
    v1Router.Post("/feed", apiCfg.middlewareAuth(apiCfg.handlerCreateFeed))
    ```

- **`v1/feed`**: **DELETE** request to delete an existing feed
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Deletes a feed by its ID.
  - **Route**:
    ```go
    v1Router.Delete("/feed", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeed))
    ```

---
### **Feed Follows Endpoints:**

- **`v1/feed-follows`**: **GET** request to retrieve the feed follows of the authenticated user
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Fetches a list of all feeds followed by the authenticated user.
  - **Route**:
    ```go
    v1Router.Get("/feed-follows", apiCfg.middlewareAuth(apiCfg.handlerGetFeedFollows))
    ```

- **`v1/feed-follow`**: **POST** request to follow a new feed
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Allows the authenticated user to follow a new feed.
  - **Route**:
    ```go
    v1Router.Post("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerCreateFeedFollows))
    ```

- **`v1/feed-follow`**: **DELETE** request to unfollow a feed
  - **Requires**: Header `Authorization: ApiKey <apikey>`
  - **Description**: Allows the authenticated user to unfollow a feed.
  - **Route**:
    ```go
    v1Router.Delete("/feed-follow", apiCfg.middlewareAuth(apiCfg.handlerDeleteFeedFollows))
    ```

---
## Authentication

- All requests to endpoints that require authentication should include an `Authorization` header in the following format:
> curl -H "Authorization: ApiKey <apikey>" http://yourdomain/v1/user

- Instead of using curl you can use postman to test api by sending appropriate auth header and json body where required.

