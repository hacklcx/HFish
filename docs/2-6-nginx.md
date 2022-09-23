贡献用户：K龙



> 如果对编写 Nginx 配置文件不熟悉，推可以使用 [https://nginxconfig.io/](https://nginxconfig.io/) 辅助完成配置。（若使用[https://nginxconfig.io/](https://nginxconfig.io/) ，请于文件内使用include语句包含[https://nginxconfig.io/](https://nginxconfig.io/)生成的配置文件）

HFish 分别使用了 `TCP/4433` 和 `TCP/4434` 端口。其中 `TCP/4433` 用于对外提供 Web 服务和部分 API 接口；`TCP/4434` 端口用于节点通讯。因此，Nginx 需同时代理这两个端口，方可提供服务。

以下是一份简单的 Nginx 配置，仅供参考。

#### 常规方案

```
server {
    listen                  443 ssl http2;
    listen                  [::]:443 ssl http2;
    client_max_body_size 4G;										#设置传输大小限制，nginx默认文件上传大小为1M。
    server_name domain.com;											#需要将domain.com替换为您所使用的域名。

    ssl_certificate         /etc/nginx/cert/domain.com.pem;			#您的证书存储位置，通常位于/etc/nginc/cert文件夹下，需要将domain.com.pem替换为您的证书文件。
    ssl_certificate_key     /etc/nginx/cert/domain.com.key;			#您证书私钥的存储位置，通常位于/etc/nginc/cert文件夹下，需要将domain.com.key替换为您证书私钥文件。

    # 反向代理设置
    # 反向代理 /api/v1 用于与节点进行通讯
    location /api/v1 {
        proxy_pass https://127.0.0.1:4434$request_uri;
        proxy_set_header Host              $host;
        proxy_set_header X-Real-IP         $remote_addr;
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
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host  $host;
        proxy_set_header X-Forwarded-Port  $server_port;
        proxy_connect_timeout              60s;
        proxy_send_timeout                 60s;
        proxy_read_timeout                 60s;
    }
}

# HTTP 重定向至 HTTPS
server {
    listen      80;
    listen      [::]:80;
    server_name domain.com;
    location / {
        return 301 https://domain.com$request_uri? permanent;
    }
}

```

说明

* 部署时需要将文件内所有domain.com更改为所您所使用的域名。
* 本文件为已有证书使用，存放于/etc/nginx/conf.d/文件夹下（建议为命名为hfish.conf）。


#### Let's encrypto方案

> Let's Encrypt是一个数字证书认证机构，旨在以自动化流程消除手动创建和安装证书的复杂流程，并推广使万维网服务器的加密连接无所不在，为安全网站提供免费的SSL/TLS证书。

```
server {
    listen                  443 ssl http2;
    listen                  [::]:443 ssl http2;
    client_max_body_size 4G;												#设置传输大小限制，nginx默认文件上传大小为1M
    server_name domain.com;													#需要将domain.com替换为您所使用的域名

    ssl_certificate         /etc/letsencrypt/live/domain.com/fullchain.pem;	#您的证书存储位置，通常位于/etc/nginc/cert文件夹下，需要将domain.com.pem替换为您的证书文件
    ssl_certificate_key     /etc/letsencrypt/live/domain.com/privkey.pem;	#您证书私钥的存储位置，通常位于/etc/nginc/cert文件夹下，需要将domain.com.key替换为您证书私钥文件。

    # 反向代理设置
    # 反向代理 /api/v1 用于与节点进行通讯
    location /api/v1 {
        proxy_pass https://127.0.0.1:4434$request_uri;
        proxy_set_header Host              $host;
        proxy_set_header X-Real-IP         $remote_addr;
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
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host  $host;
        proxy_set_header X-Forwarded-Port  $server_port;
        proxy_connect_timeout              60s;
        proxy_send_timeout                 60s;
        proxy_read_timeout                 60s;
    }
}


server {
    listen      80;
    listen      [::]:80;
    server_name domain.com;

    #方向代理 / 用于let's encrypto验证域名归属
    location ^~ /.well-known/acme-challenge/{
        default_type "text/plain";
        root /;
    }

    # HTTP 重定向至 HTTPS
    location / {
        return 301 https://domain.com$request_uri? permanent;
    }
}

```

说明

* 本文件为生成证书后所使用的文件，存放于/etc/nginx/conf.d/文件夹下（建议为命名为hfish.conf）。
* 使用let's encrypto生成证书步骤请看下文。

前置步骤：

于/etc/nginx/conf.d/文件夹下，将以下内容写入文件，命名为*.conf（建议为命名为hfish.conf，星号内容可自定义）

```
server {
    listen      80;
    listen      [::]:80;
    server_name domain.com;

    #反向代理 / 用于let's encrypto验证域名归属
    location ^~ /.well-known/acme-challenge/{
        default_type "text/plain";
        root /;
    }
}
```

文件写入后运行命令

```
nginx -t && systemctl restart nginx
```

安装certbot

```
# ubuntu系统
apt install certbot
# centos系统
yum install certbot
```

完成域名验证并生成证书

```
certbot certonly --webroot domain.com											#将domain.com替换为需要申请证书的域名
```

运行该命令后根据提示输入验证文件地址。根据配置，验证地址为root /; ，即此处输入‘/’即可。

生成证书与密钥默认存放于/etc/letsencrypt/live/domain.com/文件夹下，此时将/etc/nginx/conf.d/文件夹下（建议为命名为hfish.conf）配置文件替换为完整的配置文件并运行以下命令即可。

```nginx
nginx -t && systemctl restart nginx
```

#### CERT证书自动续期方案

编辑定时任务

```
crontab -e																		#首次使用定时任务则选择“2”，使用vim编辑
```

写入以下命令，将domain.com替换为需要自动续期的域名

```
00 00 1 * * certbot renew --force-renewal && systemctl restart nginx			#每月一号强制续期所有使用certbot生成的证书
```

> certbot申请由Let's encrypto签发的证书通常有效期为三个月，剩余有效期大于三十日则判定为证书有效无法续期，此时使用“--force-renewal ”强制续期即可。
>
> 值得注意的是，Let's encrypto限制单个域名/IP每周生成证书最多五十个，每次生成证书时错误尝试次数小于等于五次。建议生成证书时先加入参数“--dry-run”测试。


#### Let's encrypto+Nginx鉴权方案

> 使用nginx鉴权可有效隐藏管理端界面，可有效防止hash匹配识别登录界面

```
server {
    listen                  443 ssl http2;
    listen                  [::]:443 ssl http2;
    client_max_body_size 4G;												#设置传输大小限制，nginx默认文件上传大小为1M
    server_name domain.com;													#需要将domain.com替换为您所使用的域名

    ssl_certificate         /etc/letsencrypt/live/domain.com/fullchain.pem;	#您的证书存储位置，通常位于/etc/nginc/cert文件夹下，需要将domain.com.pem替换为您的证书文件
    ssl_certificate_key     /etc/letsencrypt/live/domain.com/privkey.pem;	#您证书私钥的存储位置，通常位于/etc/nginc/cert文件夹下，需要将domain.com.key替换为您证书私钥文件。

    # 反向代理设置
    # 反向代理 /api/v1 用于与节点进行通讯
    location /api/v1 {
        proxy_pass https://127.0.0.1:4434$request_uri;
        proxy_set_header Host              $host;
        proxy_set_header X-Real-IP         $remote_addr;
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
        proxy_set_header X-Forwarded-For   $proxy_add_x_forwarded_for;
        proxy_set_header X-Forwarded-Proto $scheme;
        proxy_set_header X-Forwarded-Host  $host;
        proxy_set_header X-Forwarded-Port  $server_port;
        proxy_connect_timeout              60s;
        proxy_send_timeout                 60s;
        proxy_read_timeout                 60s;
    	auth_basic "Restricted";											#启用Nginx鉴权
    	auth_basic_user_file /etc/nginx/.htpasswd;							#密码文件存放地址，默认为/etc/nginx/.htpasswd
    }
}


server {
    listen      80;
    listen      [::]:80;
    server_name domain.com;

    #方向代理 / 用于let's encrypto验证域名归属
    location ^~ /.well-known/acme-challenge/{
        default_type "text/plain";
        root /;
    }

    # HTTP 重定向至 HTTPS
    location / {
        return 301 https://domain.com$request_uri? permanent;
    }
}

```

说明

* Nginx鉴权设置前置条件

  安装Nginx鉴权工具

  ```
  # ubuntu系统
  apt install apache2-utils
  # centos系统
  yum install httpd-tools
  ```

  生成密码

  ```
  htpasswd -c /etc/nginx/.htpasswd username							#username为登录用户名，运行此命令后按照步骤输入需要设置的密码即可
  ```

* 本文件为生成证书后所使用的文件，存放于/etc/nginx/conf.d/文件夹下（建议为命名为hfish.conf）。