# URL Shortener

This is a simple URL shortener web application written in Go using the Fiber web framework and GORM ORM.

## Features

- Create short URLs
- Redirect to full URLs using short URLs

## API Endpoints

### Create a Short URL

```
POST /
```

** Request Body: **

```json
{
  "full": "https://example.com"
}
```

** Response Body: **

```json
{
  "ID": 1,
  "CreatedAt": "2023-04-12T12:00:00Z",
  "UpdatedAt": "2023-04-12T12:00:00Z",
  "DeletedAt": null,
  "slug": "AbCdEf",
  "full": "https://example.com"
}
```

### Resolve a Short URL

```
GET /:slug
```

** Response: **

The application will redirect to the full URL associated with the given slug.
