# Animal REST

## How to Run the Program

1. **Clone this repository:**
    ```bash
    git clone https://github.com/abimo2020/animal-rest
    ```
2. **Navigate to the project directory:**
    ```bash
    cd animal-rest
    ```
3. **Install the dependencies:**
    ```bash
    go mod download
    ```
4. **Update the environment variables in the `.env` file as needed.**
5. **Run the application:**
    ```bash
    go run main.go
    ```

## Alternative: Using Docker Compose

1. **Ensure that Docker is installed and running on your computer.**
2. **Open the `docker-compose.yml` file and modify the environment variables if desired (optional).**
3. **Run Docker Compose:**
    ```bash
    docker-compose up -d
    ```

## API Documentation

To easily access the REST API, please refer to the documentation:
- [API Documentation](https://documenter.getpostman.com/view/27529365/2sAXxS9CRE)

You can also import the Postman collection directly into your Postman application, as this repository includes a prepared Postman JSON file.

## Tech Stack

1. **Golang**
2. **MySQL**
3. **Redis**
4. **Echo Framework**
5. **GORM**
6. **Docker Compose**