## docker-compose

> docker-compose的简单实例与go测试实例

### 环境

- Linux
- Docker
- Docker-compose

### 配置文件

```yaml
version: "3"

services:
  mongo:
    image: mongo
    container_name: mongo
    ports:
      - "27017:27017"
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - ./mongodb:/data/db
    environment:
      MONGO_INITDB_ROOT_USERNAME: root      # 用户名
      MONGO_INITDB_ROOT_PASSWORD: 123456    # 密码

  mysql:
    image: mysql
    container_name: mysql
    ports:
      - "3306:3306"
    volumes:
      - "./mysql:/var/lib/mysql"
    environment:
      MYSQL_ROOT_PASSWORD: 123456           # 密码

  redis:
    image: redis
    container_name: redis
    command: redis-server --requirepass 123456 # 密码
    ports:
      - "6379:6379"
    volumes:
      - /etc/localtime:/etc/localtime:ro
      - ./redis:/data
    environment:
      - TZ=Asia/Shanghai
```
