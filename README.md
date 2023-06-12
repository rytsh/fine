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

Push file
```sh
curl -X POST -F "file=@/path/to/file" http://localhost:8080/api/v1/file?path={file_path}
```

Push or update file
```sh
curl -X PUT -F "file=@/path/to/file" http://localhost:8080/api/v1/file?path={file_path}
```

Get file
```sh
curl -o output -X GET http://localhost:8080/api/v1/file?path={file_path}
```

Delete file
```sh
curl -X DELETE http://localhost:8080/api/v1/file?path={file_path}
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
Generate swagger docs

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
