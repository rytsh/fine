# fine

Rest-API file management service.

## Usage

Add configuration file and give `CONFIG_FILE` to your environment variables. Default is `config.yaml` in the current directory.

Configuration file could be json, yaml or toml.

```sh
export CONFIG_FILE=/path/to/config.yaml
./fine
```

### API

Go to swagger-ui page: http://localhost:8080/api/v1/swagger/index.html

__fine__ server automatically create folders if not exists in PUT and POST requests.

#### POST

```sh
curl -X POST -F "file=@/path/to/file" http://localhost:8080/api/v1/file?path={file_path}
```

#### PUT

```sh
curl -X PUT -F "file=@/path/to/file" http://localhost:8080/api/v1/file?path={file_path}
```

#### GET

```sh
curl -o output -X GET http://localhost:8080/api/v1/file?path={file_path}
```

#### DELETE

```sh
curl -X DELETE http://localhost:8080/api/v1/file?path={file_path}
```

When deleting a directory add `force=true` query parameter to delete recursively.

```sh
curl -X DELETE http://localhost:8080/api/v1/file?path={directory_path}&force=true
```

### Configuration

```yaml
log:
  level: info

server:
  port: 8080
  host: 0.0.0.0
  base_path: "" # example as "/fine"

storage:
  local:
    path: /path/to/storage
```

## Development

<details><summary>Click here to see details...</summary>

Generate swagger docs if you change api definitions

```sh
make docs
```

Run service

```sh
make run
```

Test with curl

```sh
# download file
curl -o example.txt -X GET http://localhost:8080/api/v1/file?path=example.txt
# upload file
echo "Merhaba dünya" > 👋.txt
curl -X PUT -F "file=@👋.txt" http://localhost:8080/api/v1/file?path=👋.txt
# get file
curl -o hi.txt -X GET http://localhost:8080/api/v1/file?path=👋.txt
# delete file
curl -X DELETE http://localhost:8080/api/v1/file?path=👋.txt
```

</details>
