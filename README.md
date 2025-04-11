# CRUD Go Backend Project 🚀

This is a backend project built with [Golang](https://go.dev/). The project is designed for simple CRUD (Create, Read, Update, Delete) operations with scalability and maintainability in mind.

---

## 📂 Project Structure

```
/CRUD-Go  
│── /cmd                # Application entry point  
│    ├── main.go        # Main executable  
│── /config             # Configuration  
│    ├── database.go    # Database connection config  
│── /controllers        # Request handlers  
│    ├── todo-controler.go # Todo controller  
│── /models             # Data structure definitions  
│    ├── todo.go        # Todo model  
│── /repository         # Database interactions  
│    ├── todo_repository.go # Todo repository  
│── /routes             # Route definitions  
│    ├── router.go      # Application router  
│── /test               # Unit & integration tests  
│── .air.toml           # Air configuration (hot-reload)  
│── go.mod              # Go module definition  
│── go.sum              # Dependencies checksum  
│── Dockerfile          # Docker configuration  
│── .env                # Environment variables  
```

---

## 🛠️ Setup & Installation

### 1️⃣ Install Go
Ensure you have Go installed. Check the version:
```sh
go version
```
If not installed, download it from [golang.org/dl](https://golang.org/dl/).

### 2️⃣ Clone the repository
```sh
git clone https://github.com/your/repo.git
cd CRUD-Go
```

### 3️⃣ Initialize the Go module
```sh
go mod tidy
```

### 4️⃣ Run the application
```sh
go run cmd/main.go
```

Or use Air for hot-reloading during development:
```sh
air
```

The server will start at:
```
http://localhost:8080
```

---

## 📦 Dependencies

| Package                   | Description                     |
|---------------------------|---------------------------------|
| `github.com/joho/godotenv` | Load environment variables from `.env` |
| `github.com/gin-gonic/gin` | HTTP web framework for APIs |
| `gorm.io/gorm`             | ORM for database operations |
| `gorm.io/driver/mysql`     | MySQL driver for GORM |

To install all dependencies:
```sh
go mod tidy
```

---

## 🚀 Running with Docker

### Using Docker Compose (Recommended)

1. Create `.ENV` file from template:
   ```sh
   cp .ENV.example .ENV
   ```

2. Edit the `.ENV` file with your information

3. Run the application with Docker Compose:
   ```sh
   docker-compose up -d
   ```

4. Stop the application:
   ```sh
   docker-compose down
   ```

### Using Docker Individually

1. Build Docker image:
   ```sh
   docker build -t crud-go .
   ```
2. Run container:
   ```sh
   docker run -p 8080:8080 --env-file .ENV crud-go
   ```

---

## 📜 API Endpoints

| Method | Endpoint | Description |
|--------|---------|-------------|
| `GET`  | `/api/todos`     | Get all todos |
| `POST` | `/api/todos` | Create a new todo |
| `GET`  | `/api/todos/:id` | Get todo by ID |
| `PUT`  | `/api/todos/:id` | Update a todo |
| `DELETE` | `/api/todos/:id` | Delete a todo |

---

## 🛠️ Development

### 🧪 Run Tests
```sh
go test ./...
```

### 🧹 Format Code
```sh
go fmt ./...
```

---

## 📜 License

This project is licensed under the MIT License.

---

Let me know if you want any modifications! 🚀🔥

