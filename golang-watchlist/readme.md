# 🎌 Anime Watch List API

> A clean, simple RESTful API to manage your anime watch list with Go and MySQL

[![Go Version](https://img.shields.io/badge/Go-1.24.2+-00ADD8?style=flat&logo=go)](https://golang.org/)
[![MySQL](https://img.shields.io/badge/MySQL-8.0+-4479A1?style=flat&logo=mysql&logoColor=white)](https://www.mysql.com/)
[![Docker](https://img.shields.io/badge/Docker-Required-2496ED?style=flat&logo=docker&logoColor=white)](https://www.docker.com/)

## 🛠️ Tech Stack

This project is built with:

- **[Gorilla Mux](https://github.com/gorilla/mux)** - HTTP router and URL matcher
- **Go's database/sql** - Standard database interface
- **[MySQL Driver](https://github.com/go-sql-driver/mysql)** - MySQL database driver
- **Docker** - Container platform for database


## ✨ Features

- **Complete CRUD Operations** - Create, read, update, and delete anime records
- **Episode Tracking** - Keep track of watched vs total episodes
- **Status Management** - Monitor your watching progress
- **Docker Ready** - Easy database setup with containerization
- **RESTful Design** - Clean, predictable API endpoints

## 🚀 Quick Start

### Prerequisites

Make sure you have these installed:
- [Go](https://golang.org/dl/) (v1.24.2 or higher)
- [Docker](https://www.docker.com/get-started)

### Installation

1. **Clone the repository**
   ```bash
   git clone https://github.com/utkarshkrsingh/keploy-api-fellowship.git
   cd keploy-api-fellowship/golang-watchlist/
   ```

2. **Start the MySQL database**
   ```bash
   docker run -d \
     --name mysql-anime-watch-list \
     --network mysql-network \
     -e MYSQL_ROOT_PASSWORD=qwerty \
     -e MYSQL_DATABASE=anime-watch-list \
     -p 5010:3306 \
     mysql
   ```

3. **Run the API server**
   ```bash
   go mod tidy
   go run ./cmd/server
   ```

🎉 **You're all set!** The API is now running at `http://localhost:8080`

## 📚 API Reference

### Endpoints Overview

| Method   | Endpoint              | Description                    |
|----------|-----------------------|--------------------------------|
| `GET`    | `/watchlist`          | Get all anime in your list    |
| `POST`   | `/watchlist`          | Add new anime to list          |
| `PUT`  | `/watchlist/{id}`     | Update existing anime          |
| `DELETE` | `/watchlist/{id}`     | Remove anime from list         |

### Data Structure

Each anime record contains:

```json
{
  "id": 1,
  "title": "Attack on Titan",
  "total_episodes": 24,
  "watched_episodes": 12,
  "type": "tv",
  "status": "watching"
}
```

## 💡 Usage Examples

### Get All Anime

**Request:**
```bash
curl -X GET http://localhost:8080/watchlist
```

**Response:**
```json
[
    {
        "id": 3,
        "title": "Block Lock Season 2",
        "total_episodes": 14,
        "watched_episodes": 14,
        "type": "tv",
        "status": "completed"
    },
    {
        "id": 4,
        "title": "Potemayo",
        "total_episodes": 12,
        "watched_episodes": 2,
        "type": "tv",
        "status": "watching"
    },
    {
        "id": 5,
        "title": "Blue Lock",
        "total_episodes": 24,
        "watched_episodes": 24,
        "type": "tv",
        "status": "completed"
    }
]
```

### Add New Anime

**Request:**
```bash
curl -X POST http://localhost:8080/watchlist \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Naruto: Shippuden",
    "totalepisodes": 500,
    "totalwatched": 500,
    "status": "completed",
    "type": "tv"
  }'
```

**Response:**
```json
{
    "id":6,
    "title":"Naruto: Shippuden",
    "total_episodes":500,
    "watched_episodes":500,
    "type":"tv",
    "status":"completed"
}
```

### Update Existing Anime

**Request:**
```bash
curl -X PATCH http://localhost:8080/watchlist/4 \
  -H "Content-Type: application/json" \
  -d '{
    "total_episodes": 12,
    "watched_episodes": 2,
    "status": "watching",
    "type": "tv"
  }'
```

**Response:**
```json
{
    "id": 4,
    "title": "Potemayo",
    "total_episodes": 12,
    "watched_episodes": 2,
    "type": "tv",
    "status": "watching"
}
```

### Delete Anime

**Request:**
```bash
curl -i -X DELETE http://localhost:8080/watchlist/5
```

**Response:**
```bash
HTTP/1.1 204 No Content
Date: Sun, 22 Jun 2025 15:17:23 GMT
```

## 📊 Status Types

Keep track of your watching progress:

| Status      | Description                    |
|-------------|--------------------------------|
| `watching`  | Currently watching             |
| `completed` | Finished all episodes          |
| `on-hold`   | Paused temporarily             |
| `dropped`   | Stopped watching               |
| `planning`  | Planning to watch later        |

## 🤝 Contributing

We welcome contributions! Here's how you can help:

1. Fork the repository
2. Create a feature branch (`git checkout -b feature/amazing-feature`)
3. Commit your changes (`git commit -m 'Add amazing feature'`)
4. Push to the branch (`git push origin feature/amazing-feature`)
5. Open a Pull Request

## 🐛 Issues & Support

Found a bug or need help? Please [open an issue](https://github.com/utkarshkrsingh/keploy-api-fellowship/issues) on GitHub.

---

<div align="center">
  <p>Made with ❤️ for anime fans</p>
  <p>⭐ Star this repo if you found it helpful!</p>
</div>
