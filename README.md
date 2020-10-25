### 基于Gin框架

### Use

```
go env -w GO111MODULE=on
go env -w GOPROXY=https://goproxy.io,direct

# 设置不走 proxy 的私有仓库，多个用逗号相隔（可选）
go env -w GOPRIVATE=*.corp.example.com

# 设置不走 proxy 的私有组织（可选）
go env -w GOPRIVATE=example.com/org_name
```

```
export GOPROXY=https://mirrors.cloud.tencent.com/go/
export GOPROXY=https://mirrors.aliyun.com/goproxy/
```

```
go mod init projectName
```

```
go mod tidy
```

```
go mod vendor
```

```
go build -mod=vendor
```