#### Nginx 部署

`（该文档特殊致谢Github用户：JamCh01）`

>如果对编写 Nginx 配置文件不熟悉，推荐使用 https://nginxconfig.io/ 辅助完成配置

HFish 分别使用了 `TCP/4433` 和 `TCP/4434` 端口。其中 `TCP/4433` 用于对外提供 Web 服务和部分 API 接口；`TCP/4434` 端口用于节点通讯。因此，Nginx 需同时代理这两个端口，方可提供服务。

以下是一份简单的 Nginx 配置，仅供参考。

```
server {
    listen                  443 ssl http2;
    listen                  [::]:443 ssl http2;
    # 域名设置，需根据实际情况进行修改
    server_name             localhost.localdomain;

    # SSL设置，需根据实际情况进行修改
    ssl_certificate         /etc/letsencrypt/live/localhost.localdomain/fullchain.pem;
    ssl_certificate_key     /etc/letsencrypt/live/localhost.localdomain/privkey.pem;
    ssl_trusted_certificate /etc/letsencrypt/live/localhost.localdomain/chain.pem;


    # 日志设置，需根据实际情况进行修改
    access_log              /var/log/nginx/localhost.localdomain.access.log;
    error_log               /var/log/nginx/localhost.localdomain.error.log warn;

    # 反向代理设置
    # 反向代理 /api/v1 用于与节点进行通讯
    location /api/v1 {
        proxy_pass https://127.0.0.1:4434$request_uri;
        proxy_set_header Host              $host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header Forwarded         $proxy_add_forwarded;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host  $host;
        proxy_set_header X-Forwarded-Port  $server_port;
        proxy_connect_timeout              60s;
        proxy_send_timeout                 60s;
        proxy_read_timeout                 60s;
    }

    # 反向代理 /tmp 用于与节点进行下载
    location /tmp {
        proxy_pass https://127.0.0.1:4434$request_uri;
        proxy_set_header Host              $host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header Forwarded         $proxy_add_forwarded;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host  $host;
        proxy_set_header X-Forwarded-Port  $server_port;
        proxy_connect_timeout              60s;
        proxy_send_timeout                 60s;
        proxy_read_timeout                 60s;
    }

    # 反向代理 / 用于与 Web 服务
    location / {
        proxy_pass https://127.0.0.1:4433$request_uri;
        proxy_set_header Host              $host;
        proxy_set_header X-Real-IP         $remote_addr;
        proxy_set_header Forwarded         $proxy_add_forwarded;
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host  $host;
        proxy_set_header X-Forwarded-Port  $server_port;
        proxy_connect_timeout              60s;
        proxy_send_timeout                 60s;
        proxy_read_timeout                 60s;
    }
    

    # additional config
    include nginxconfig.io/general.conf;
}

# HTTP 重定向至 HTTPS
server {
    listen      80;
    listen      [::]:80;
    server_name localhost.localdomain;
    location ^~ /.well-known/acme-challenge/ {
        root /var/www/_letsencrypt;
    }

    location / {
        return 301 https://localhost.localdomain$request_uri;
    }
}
```