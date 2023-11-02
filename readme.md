# Rectangles Application

This Go application is a simple REST API service that allows you to manage and perform operations on rectangles. It uses the [Gin](https://github.com/gin-gonic/gin) web framework for routing and [GORM](https://gorm.io/) for database access.

## Features

- Create and manage rectangles with attributes such as position, size, and timestamp.
- Calculate intersecting rectangles based on a provided main rectangle and a list of input rectangles.
- Retrieve a list of all stored rectangles from the database.

## Installation

1. Clone the repository to your local machine:

   ```bash
   git clone https://github.com/yazdanifar/rectangles.git
   ```

2. Navigate to the project directory:

   ```bash
   cd rectangles
   ```

3. Install dependencies:

   ```bash
   go get -u
   ```

4. Run the application:

   ```bash
   go run main.go
   ```

The application will be available at [http://localhost:8080](http://localhost:8080).

## Usage

### Persist Intersecting Rectangles

To persist intersecting rectangles, make a POST request to the root endpoint with the following JSON payload:

```json
{
   "main": {"x": 0, "y": 0, "width": 10, "height": 20},
   "input": [
      {"x": 2, "y": 18, "width": 5, "height": 4},
      {"x": 12, "y": 18, "width": 5, "height": 4},
      {"x": -1, "y": -1, "width": 5, "height": 4}
   ]
}
```
Each "input" specifies a rectangle, and any input rectangles that intersect with the "main" rectangle are saved, while 
others are discarded.


### Get All Rectangles

To retrieve a list of all rectangles, make a GET request to the root endpoint.

## Database

This application uses an SQLite database (`rectangles.db`) to store rectangles. The database is automatically migrated when the application starts.

