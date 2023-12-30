# User-Profile API with Go-Lang
> [!NOTE]
> Postman Documentation: [here]()

# Prerequisite before testing API
1. Must have installed Go-Lang for checking use command :

    ```bash
    go version

2. Must have installed PostgreSQL

3. Create Database, you can use this steps, if you don't use PGAdmin 4 : 
    - Open Terminal
    - Entry PostgreSQL with this command : 
    ```bash
    psql -U your_username

    example:
    psql -U recodink
    ```

    - Create Database with this name: 
    ```bash
    CREATE DATABASE photo_profile_app;
    ```

    - for checking, you can typing this: 
    ```bash
    \l
    ```

    - exit psql, you can typing this:
    ```bash
    \q
    ```

    if you use PGAdmin 4 but doesn't know how to create, see this article: 
    Create Database with PGAdmin 4: [here](https://www.tutorialsteacher.com/postgresql/create-database) 

4. Create .env, you can see at .env.example

5. type this command in terminal or command prompt
    ```bash
    export GIN_MODE=release
    ```
6. type this command in terminal for run this app :
    ```bash
    go run main.go
    ```
 
7. that's it, you can use this feature api in this app: 
    - Auth
    ```bash
    - Register User
    - Login User
    ```
    - User
    ```bash
    - Update User
    - Delete User
    ```

    - Photos
    ```bash
    - Post Photos
    - Get Photos
    - Update Photos
    - Delete Photos
    ```

8. the last, this is all the endpoint of this app: 
    ## Auth Endpoint
    > Register User - 
    > POST http://localhost:8080/api/auth/register
    * Register User
    * Request Body: 
    {
    "Username": "recodink",
    "Email": "recodink16@gmail.com",
    "password": "recodink16"
    }

    * Response Body:
        * 201 Created
        ```
        {
            "message": "User Registered Successfully",
            "success": true
        }
        ```

    > Login User
    > POST http://localhost:8080/api/auth/login
    * Authentication Login User
    * Request Body :
    ```
    {
    "email": "recodink16@gmail.com",
    "password":"recodink16"
    }
    ```

    * Response Body :
        * 200 OK
        ```
        {
        "message": "User Login Successfully",
        "success": true,
        "token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE3MDM5Nzc3MjksImp0aSI6IjIifQ.Gbx5125_0HK7nkeRQRB-LZ-yMLl5DzpU56CCzIGP1Lw"
        }
        ```

    ## Photo Endpoint
    
    > Upload Photo -
    > POST http://localhost:8080/api/photos

    * Upload User Photo Profile
    * Requires authentication via JWT Token
    * Request Body :
    ```
    - Form-data
    {
        "title": "Screenshot",
        "caption": "ini screenshot kelas",
        "PhotoUrl": "/home/recodink/Pictures/Screenshots/Screenshot from 2023-12-29 18-42-21.png"
    }
    ```

    * Response Body :
        * 201 Created 
        ```
        {
            "message": "Photo Uploaded Successfully",
            "success": true
        }
        ```
    
    > Get Photo -
    > Get http://localhost:8080/api/photos

    * Get All User Photo Profile
    * Requires authentication JWT Token
    * Request Body :
    ```
    {
        
    }
    ```

    * Response Body :
        * 200 OK
        ```
        {
        "data": [
            {
                "ID": 2,
                "DeletedAt": null,
                "title": "Screenshot",
                "caption": "this logo created for final project with style black",
                "photo_url": "http://localhost:8080/api/photos/1703891341392573066-Screenshotfrom2023-12-2918-42-21.png",
                "user_id": "2",
                "user": {
                    "ID": 2,
                    "DeletedAt": null,
                    "username": "recodink",
                    "email": "recodink16@gmail.com",
                    "password": "$2a$10$BDw.eJ5VW3rQHITLA8sEfOw8T9gs0WFZLttv3tNZimTLFsQiqTl1O",
                    "photos": null,
                    "CreatedAt": "2023-12-30T06:08:11.727967+07:00",
                    "UpdatedAt": "2023-12-30T06:08:11.727967+07:00"
                },
                "CreatedAt": "2023-12-30T06:09:01.500456+07:00",
                "UpdatedAt": "2023-12-30T06:09:01.500456+07:00"
            }
        ],
        "message": "Successfully Retrieve Data",
        "success": true
        }
        ```
    
    > Update Photo -
    > PUT http://localhost:8080/api/photos/:id
    
    * Update User Photo Profile
    * Requires authetication via JWT Token
    * Request Body :
    ```
    {
        "title": "Screenshot",
        "caption": "this logo created for final project with style black",
        "photo_url": "http://localhost:8080/api/photos/1703891341392573066-Screenshotfrom2023-12-2918-42-21.png"
    }
        
    ```
    * Response Body :
        * 200 OK
        ```
        {
        "data": {
            "ID": 2,
            "DeletedAt": null,
            "title": "Screenshot",
            "caption": "this logo created for final project with style black",
            "photo_url": "http://localhost:8080/api/photos/1703891341392573066-Screenshotfrom2023-12-2918-42-21.png",
            "user_id": "2",
            "user": {
                "ID": 2,
                "DeletedAt": null,
                "username": "recodink",
                "email": "recodink16@gmail.com",
                "password": "$2a$10$G/wB1PEj2LxeBKdbXZDOUeq8Z5rKgVOfdMTUG2Vy3zbOUJZNo4vMO",
                "photos": null,
                "CreatedAt": "2023-12-30T06:08:11.727967+07:00",
                "UpdatedAt": "2023-12-30T06:08:11.727967+07:00"
            },
            "CreatedAt": "2023-12-30T06:09:01.500456+07:00",
            "UpdatedAt": "2023-12-30T06:11:04.83792719+07:00"
        },
        "message": "Photo Update Sucessfully",
        "success": true
        }

    > Delete Photo - 
    > DELETE http://localhost:8080/api/photos/:id

    * Delete User Photo Profile
    * Requires authentication via JWT Token
    * Request Body:
    ```
    {
        
    }
    ```
    * Response Body :
        * 200 OK
        ```
        {
            "message": "Photo Deleted Successfully",
            "success": true
        }

    
    ## User Endpoint

    > Update User -
    > PUT http://localhost:8080/api/users/:id
    
    * Update User
    * Requires authentication via JWT Token
    * Request Body :
    ```
    {
        "Username" : "recodink",
        "Email": "recodink16@gmail.com",
        "password": "recodink1607"
    }
    ```
    * Response Body :
        * 200 OK
        ```
        {
            "message": "User Update Successfully",
            "success": true
        }
        ```

    > Delete User - 
    > DELETE http://localhost:8080/api/users/:id
    
    * Delete User
    * Requires authentication via JWT Token
    * Request Body :
    ```
    {

    }
    ```

    * Response Body :
        * 200 OK
        {
            "message": "User Deleted Successfully",
            "success": true
        }