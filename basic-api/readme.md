# ⛩️ Anime Watch List API

A RESTful API for managing your anime watch list, built with Go and MySQL.

## 🛠️ Tech Stack

- **[GORM](https://gorm.io/gorm)** - Object-Relational Mapping
- **[MySQL Driver](https://gorm.io/driver/mysql)** - Database driver for MySQL
- **[Gin](https://github.com/gin-gonic/gin)** - HTTP web framework
- **Docker** - Containerized MySQL database

## 📊 Data Model

The API manages anime records with the following structure:

```go
type WatchList struct {
    gorm.Model
    Name          string `gorm:"not null" json:"name"`
    TotalEpisodes int    `gorm:"not null" json:"totalepisodes"`
    TotalWatched  int    `gorm:"not null" json:"totalwatched"`
    Status        string `json:"status"`
    Type          string `json:"type"`
}
```

## 🚀 API Endpoints

| Method | Endpoint | Description |
|--------|----------|-------------|
| `GET` | `/watchlist` | Retrieve all anime records or search by name |
| `POST` | `/watchlist` | Add a new anime to the watch list |
| `PATCH` | `/watchlist/:name` | Update an existing anime record |
| `DELETE` | `/watchlist/:name` | Delete an anime record |

### Query Parameters

- **GET `/watchlist?name={anime_name}`** - Search for a specific anime by name

## 🗄️ Database Setup

This project uses Docker to run a MySQL database container, eliminating the need for manual database setup. GORM provides the ORM layer for efficient database operations.

## ⚙️ Installation & Setup

### Prerequisites

- **Go** v1.24.2 or higher
- **Docker**

### Steps

1. **Clone the repository:**
   ```bash
   git clone http://github.com/utkarshkrsingh/keploy-api-fellowship.git
   cd keploy-api-fellowship/basic-api/
   ```

2. **Start MySQL container:**
   ```bash
   docker run -d \
     --name mysql-anime-watch-list \
     --network mysql-network \
     -e MYSQL_ROOT_PASSWORD=qwerty \
     -e MYSQL_DATABASE=anime-watch-list \
     -p 5010:3306 \
     mysql
   ```

3. **Install dependencies and run the API:**
   ```bash
   go mod download  # Download required packages
   go mod tidy      # Clean up dependencies
   go run .         # Start the API server
   ```

The API will be available at `http://localhost:8080`

## 📝 API Examples

### 1. Get All Records

**Request:**
```bash
curl -X GET http://localhost:8080/watchlist
```

**Response:**
```json
{
  "data": [
    {
      "ID": 1,
      "CreatedAt": "2025-06-20T17:21:20.819+05:30",
      "UpdatedAt": "2025-06-20T17:21:20.819+05:30",
      "DeletedAt": null,
      "name": "Kaiju No. 8",
      "totalepisodes": 12,
      "totalwatched": 12,
      "status": "completed",
      "type": "tv"
    },
    {
      "ID": 3,
      "CreatedAt": "2025-06-20T17:24:40.957+05:30",
      "UpdatedAt": "2025-06-20T17:24:40.957+05:30",
      "DeletedAt": null,
      "name": "Demon Slayer: Kimetsu No Yaiba",
      "totalepisodes": 26,
      "totalwatched": 26,
      "status": "completed",
      "type": "tv"
    }
  ]
}
```

### 2. Search by Name

**Request:**
```bash
curl -X GET "http://localhost:8080/watchlist?name=Naruto"
```

**Response:**
```json
{
  "data": [
    {
      "ID": 6,
      "CreatedAt": "2025-06-20T22:30:26.002+05:30",
      "UpdatedAt": "2025-06-20T22:30:26.002+05:30",
      "DeletedAt": null,
      "name": "Naruto",
      "totalepisodes": 220,
      "totalwatched": 220,
      "status": "completed",
      "type": "tv"
    }
  ]
}
```

### 3. Add New Anime

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
  "message": "Anime inserted successfully"
}
```

### 4. Update Existing Anime

**Request:**
```bash
curl -X PATCH http://localhost:8080/watchlist/Potemayo \
  -H "Content-Type: application/json" \
  -d '{
    "name": "Potemayo",
    "totalepisodes": 12,
    "totalwatched": 2,
    "status": "watching",
    "type": "tv"
  }'
```

**Response:**
```json
{
  "message": "Database updated successfully"
}
```

### 5. Delete Anime

**Request:**
```bash
curl -X DELETE http://localhost:8080/watchlist/Potemayo
```

**Response:**
```json
{
  "message": "Anime removed successfully"
}
```

## 📋 Status Values

Common status values for anime records:
- `watching` - Currently watching
- `completed` - Finished watching
- `on-hold` - Temporarily paused
- `dropped` - Discontinued watching
- `planning` - Planning to watch
