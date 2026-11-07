# backoffice  (:8080)

REST API (BFF) + trang admin backoffice (`web/`, embed vào binary): `POST /api/login`, `GET /api/me`, `GET /api/wallets/{id}`, `POST /api/wallets/{id}/adjust`.

Gọi auth-service (REST: login, JWKS, client_credentials), wallet-service (gRPC, forward token admin) và user-service (REST, token của backoffice). Địa chỉ trong `config.yml`: `oauth.auth_url`, `wallet_service.addr`, `user_service.url`.

## Chạy

Folder này là 1 project Go độc lập (repo git, module riêng), mở folder lên rồi:

```bash
go run ./cmd/rest    # hoặc: make run
make test             # go test -race ./...
```

Cấu hình ở `config/config.yml`, ghi đè bằng biến môi trường theo quy tắc `.` → `_`
(ví dụ `REST_SERVER_ADDR=:9999`, `OAUTH_JWKS_URL=...`).

## Thư viện (make mod)

Code dùng chung (`authkit`, `apperr`) nằm ở repo **go-lib** (`github.com/phillong97254/go-lib`), giống `go-library`
ở các dự án trước. Thư viện được vendor vào `vendor/` nên build/Docker không cần mạng hay repo khác.

```bash
make mod                          # go get -u thư viện + go mod tidy + go mod vendor
make mod GO_LIB_VERSION=v0.2.0    # chốt go-lib ở 1 version (khi dùng remote)
make mod-local                    # dùng ../go-lib (replace) – mặc định hiện tại vì go-lib chưa push
make mod-remote                   # bỏ replace, lấy go-lib theo version từ remote (cần make netrc nếu repo private)
```

## Docker

```bash
make build            # docker build -t ghcr.io/phillong97254/backoffice:dev-latest .
make up               # chạy riêng service này (gọi service khác qua host.docker.internal)
make ci-docker-build ci-docker-push TAG=v1.0.0   # push lên ghcr.io/phillong97254 (GitHub Container Registry)
```

## Cấu trúc

```
cmd/rest/main.go, routes/   # bootstrap + đăng ký route
config/                      # config.go (viper) + config.yml
vendor/                      # thư viện (kể cả go-lib), tạo bằng make mod
internal/
    domain/
    internal/gateway/{auth,user,wallet}  # client gọi sang service khác
    internal/service/{auth,account,wallet}
    internal/handler/rest/v1/{public/auth,authz/account,authz/wallet}
```

## gRPC contract

`api/proto/wallet/v1/wallet.proto` → code sinh ra ở `pkg/gen/wallet/v1`. Đây là bản copy hợp đồng của wallet-service: sửa proto ở wallet-service trước rồi copy file `.proto` sang đây.

```bash
make gen-proto        # buf generate trong container go-tools
# hoặc cài buf + protoc-gen-go + protoc-gen-go-grpc rồi: cd api && buf generate
```
