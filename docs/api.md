# API Documentation

## List Nginx Configurations

- **Endpoint:** `/api/config/list`
- **Method:** `GET`

### Sample Requests

#### curl

```bash
curl -X GET http://localhost:8090/api/config/list
```

#### JavaScript (fetch)

```javascript
fetch('http://localhost:8090/api/config/list')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

## Get Nginx Configuration

- **Endpoint:** `/api/config/get`
- **Method:** `POST`
- **Payload:**
  - `filePath`: string

### Sample Requests

#### curl

```bash
curl -X POST http://localhost:8090/api/config/get \
  -H "Content-Type: application/x-www-form-urlencoded" \
  -d "filePath=/etc/nginx/sites-available/default"
```

#### JavaScript (fetch)

```javascript
fetch('http://localhost:8090/api/config/get', {
  method: 'POST',
  headers: {
    'Content-Type': 'application/json',
  },
  body: JSON.stringify({ filePath: '/etc/nginx/sites-available/default' })
})
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

## Check Nginx Installation

- **Endpoint:** `/api/check/nginx`
- **Method:** `GET`

### Sample Requests

#### curl
```bash
curl -X GET http://localhost:8090/api/check/nginx
```

#### JavaScript (fetch)
```javascript
fetch('http://localhost:8090/api/check/nginx')
  .then(response => response.json())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

## Install Nginx

- **Endpoint:** `/api/install/nginx`
- **Method:** `GET`

### Sample Requests

#### curl
```bash
curl -X GET http://localhost:8090/api/install/nginx
```

#### JavaScript (fetch)
```javascript
fetch('http://localhost:8090/api/install/nginx')
  .then(response => response.text())
  .then(data => console.log(data))
  .catch(error => console.error('Error:', error));
```

Note: The install endpoint returns a stream of data. The JavaScript example above will log the entire response once it's complete. For real-time updates, you might want to use Server-Sent Events (SSE) instead.
