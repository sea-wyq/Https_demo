# Https_demo

## 证书生成（域名是localhost）
```bash
openssl genrsa -out server.key 2048
openssl req -nodes -new -key server.key -subj "/CN=localhost" -out server.csr
openssl x509 -req -sha256 -days 365 -in server.csr -signkey server.key -out server.crt
```


## Curl 测试
```bash
curl --cacert ../cert/server.crt https://localhost
```
