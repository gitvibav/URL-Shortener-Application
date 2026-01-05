## URL Shortener Application

### Local Development Setup

#### Prerequisites
- Go 1.25 or later
- Redis (local installation)

#### Installation

1. **Install Redis locally:**
   ```bash
   # macOS
   brew install redis
   
   # Start Redis
   brew services start redis
   # OR
   redis-server
   ```

2. **Install Go dependencies:**
   ```bash
   go mod download
   ```

3. **Set up environment variables:**
   Edit `.env` if needed (defaults should work for local Redis):
   - `APP_PORT`: Port for the API server (default: 8080)
   - `DB_ADDR`: Redis address (default: localhost:6379)
   - `DB_PASS`: Redis password (leave empty for local Redis without password)

4. **Run the application:**
   ```bash
   go run .
   ```

   The API will be available at `http://localhost:8080`

#### API Endpoints

- `POST /api/v1` - Shorten a URL
- `GET /api/v1/:shortID` - Get original URL by short ID
- `PUT /api/v1/:shortID` - Edit a shortened URL
- `DELETE /api/v1/:shortID` - Delete a shortened URL
- `POST /api/v1/addTag` - Add a tag to a URL