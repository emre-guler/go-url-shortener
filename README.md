# Go URL Shortener

A simple and efficient URL shortener service written in Go that allows you to create shortened URLs. The service uses Google Sheets as a database for storing URL mappings and Redis for caching.

## Features

- Create shortened URLs with custom paths
- URL validation to ensure valid input
- Google Sheets integration for persistent storage
- Redis caching for improved performance
- Dockerized application for easy deployment

## Prerequisites

- Go 1.18 or higher
- Redis server
- Google Sheets API credentials
- Docker (optional)

## Dependencies

- github.com/asaskevich/govalidator - For URL validation
- github.com/go-redis/redis/v9 - Redis client for caching
- google.golang.org/api - Google Sheets API client

## Setup

1. Clone the repository:
```bash
git clone https://github.com/emre-guler/url-shortener.git
cd url-shortener
```

2. Set up Google Sheets:
   - Create a new Google Sheet
   - Set up Google Cloud Project and enable Google Sheets API
   - Download credentials.json and place it in the `db` directory
   - Set the environment variable for your spreadsheet ID:
   ```bash
   export URL_SHORTENER_PROJECT_SPREADSHEET_ID=your_spreadsheet_id
   ```

3. Install dependencies:
```bash
go mod download
```

4. Start Redis server:
```bash
redis-server
```

5. Run the application:
```bash
go run main.go
```

## Docker Setup

The project includes a Dockerfile for containerized deployment:

```bash
docker build -t url-shortener .
docker run -e URL_SHORTENER_PROJECT_SPREADSHEET_ID=your_spreadsheet_id url-shortener
```

## Usage

1. Run the application
2. Enter the URL you want to shorten when prompted
3. Enter your desired custom path
4. The application will return your shortened URL in the format: `https://www.emreguler.dev/your-custom-path`

## Features Implementation Status

- ✅ Google Sheets as DB
- ✅ Environment Variables for SpreadsheetId
- ✅ Dockerization
- ✅ Redis Caching

## Contributing

Feel free to submit issues, fork the repository, and create pull requests for any improvements.