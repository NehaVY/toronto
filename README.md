﻿# Toronto Time API

This Go application provides an HTTP API that retrieves the current time in the "America/Toronto" time zone, logs it, and stores it in a MySQL database. The current time is also returned in a JSON format when the `/current-time` endpoint is accessed.

## Features

- Fetches the current time in Toronto (with proper time zone handling).
- Stores the current time in a MySQL database (`toronto_time`).
- Logs all activities in a `server.log` file.
- Provides an API endpoint `/current-time` that returns the current time in JSON format.

## Prerequisites

- **Go (Golang)**: This project is built using Go 1.23.
- **MySQL**: A MySQL database should be set up and running.
- **Docker (Optional)**: For containerization of the application.

### MySQL Database Setup

The application requires a MySQL database to store the current time. You can create the necessary database and table by running the following commands:

```sql
CREATE DATABASE toronto_time;

USE toronto_time;

CREATE TABLE time_log (
    id INT AUTO_INCREMENT PRIMARY KEY,
    timestamp DATETIME NOT NULL
);
```

![alt text](image-5.png)

### Step 2: Configure Database Connection

In the `main.go` file, update the database connection parameters to match your MySQL setup. Open the `main.go` file and locate the following constants:

```go
const (
	DBUser     = "root"         // Replace with your MySQL username
	DBPassword = "Panda@sep18!" // Replace with your MySQL password
	DBName     = "toronto_time" // Database name you created
)
```
### Step 3: Install Dependencies

Before running the application, you need to install the required Go dependencies. 

To do this, run the following command in your terminal:

```bash
go mod tidy

```

### Step 4: Run the Application

After you have installed the required dependencies, you can run your Go application locally.

#### 1. **Run the Application**

In the terminal, navigate to the root directory of your Go project and execute the following command:

```bash
go run main.go
````

### Step 5: Test the API Endpoint

Now that the application is running, you can test the `/current-time` API endpoint to ensure that it is working as expected.
![alt text](image-3.png)

#### 1. **Send a Request to the API**

Open your web browser or a tool like [Postman](https://www.postman.com/) or [cURL](https://curl.se/) and make a request to the following URL:


### Step 6: Logs

Once the application is running, it's important to verify the application’s behavior by checking the logs. This will allow you to confirm that the application is handling requests and inserting data into the database as expected.

To check the logs, follow these steps:

![alt text](image.png)



### Step 7 : Dockerization


![alt text](image-1.png)

![alt text](image-2.png)
![alt text](image-4.png)




