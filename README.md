# envchecker
环境配置检测工具

- 注意：
    - 如果需要部署https，则需要在JavaScript初始化WebSocket时使用 wss 替代 ws 模式
        ```javascript
            var url = "ws://" + window.location.host + "/ws";
        ``` 
        修改为 
        ```javascript
            var url = "wss://" + window.location.host + "/ws";
        ``` 
    - 如果需要部署多处，则需要将HTML静态文件打包进可执行文件中。可使用 [packr](https://github.com/gobuffalo/packr) 工具
-  默认开放端口: 8080
-  配置文件: config.ini
-  启动: go run *.go
-  支持组件：
    - APM
    - Apollo
    - Consul
    - Elasticsearch
    - Kafka
    - MySQL
    - Redis
    - Syslog
    - Clickhouse
    - PostgreSQL
-  操作界面：
![envchecker.png](https://i.loli.net/2021/05/26/SdXfFGArQExsiJt.png)
