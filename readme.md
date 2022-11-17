# Privy Cake API
Build using `go 1.19` and `mysql` as db engine.

### Entry Point
* Unfortunately the application unable to run on docker container, due to when I found an issue and unable to solve it, and the issue is error on database connection as follow.

  ``dial tcp 127.0.0.1:3306: connect: connection refused``


* To run the application through application entry point, make sure to adjust and confirm your database configuration is correct in `config/db.go`.


* Then, run the application in your terminal with `go run main.go`.


### Unit Testing
* Run `go test -v`


### List of Endpoint
* To migrate the database table.
  * [GET] `/migrate`


* To list all cake.
    * [GET] `/cakes/`


* To get specific cake.
    * [GET] `/cakes/:id`


* To insert new cake.
  * [POST] `/cakes/`
  * ```
    {
        "title": "string",
        "description": "string",
        "rating": int,
        "image": "string"
    }
    ```

* To update a cake.
  * [PATCH] `/cakes/:id`
  * ```
    {
        "title": "string",
        "description": "string",
        "rating": int,
        "image": "string"
    }
    ```

* To delete a cake.
  * [DELETE] `/cakes/:id`