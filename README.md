# myMusic
## 项目目录

```
music-platform/
├── cmd/                        # 主程序入口
│   ├── server/                 # 服务启动相关
│   │   ├── main.go             # 主程序入口
│   │   └── config.yaml         # 服务配置文件
│   └── migrate/                # 数据迁移工具
│       └── migrate.go          # 数据库迁移脚本
├── configs/                    # 配置文件目录
│   ├── app.yaml                # 应用配置
│   ├── database.yaml           # 数据库配置
│   └── redis.yaml              # Redis 配置
├── initialize/                       # 初始化
│   ├── init_all.go             # 调用包内所有模块的初始化函数
│   ├── init_oss.go             # 初始化oss配置
│   ├── init_db.go              # 初始化数据库
│   ├── init_redis.go           # 初始化 Redis
│   └── init_config.go          # 初始化配置
├── internal/                   # 内部实现逻辑
│   ├── auth/                   # 用户认证模块
│   │   ├── handler.go          # 登录/注册接口处理
│   │   ├── jwt.go              # JWT 生成与校验
│   │   └── sso.go              # 单点登录实现
│   ├── music/                  # 音乐相关功能
│   │   ├── handler.go          # 音乐接口处理
│   │   ├── service.go          # 音乐逻辑实现
│   │   ├── storage.go          # 阿里云OSS存储逻辑
│   │   └── search.go           # 歌词搜索逻辑（OpenSearch）
│   ├── comments/               # 评论模块
│   │   ├── handler.go          # 评论接口处理
│   │   ├── service.go          # 评论逻辑实现
│   │   └── model.go            # 评论数据模型
│   ├── db/                     # 数据库操作
│   │   ├── mysql.go            # MySQL 连接与操作
│   │   └── redis.go            # Redis 连接与缓存操作
│   ├── middleware/             # 中间件
│   │   ├── logger.go           # 日志中间件
│   │   ├── cors.go             # 跨域处理
│   │   ├── auth.go             # 用户认证中间件
│   │   └── rate_limit.go       # 限流中间件
│   ├── util/                   # 工具函数
│   │   ├── hash.go             # 哈希算法
│   │   ├── json.go             # JSON 工具
│   │   ├── response.go         # 统一响应格式
│   │   └── time.go             # 时间处理工具
│   ├── router/                 # 路由模块
│   │   ├── api.go              # API 路由
│   │   └── router.go           # 路由注册
│   └── config/                 # 配置加载模块
│       └── loader.go           # 加载配置
├── pkg/                        # 可复用的库代码
│   ├── opensearch/             # OpenSearch 客户端封装
│   ├── oss/                    # 阿里云OSS客户端封装
│   ├── jwt/                    # JWT 工具封装
│   └── logger/                 # 日志模块封装
├── test/                       # 测试代码
│   ├── auth_test.go            # 用户认证测试
│   ├── music_test.go           # 音乐模块测试
│   └── comments_test.go        # 评论模块测试
├── web/                        # 前端代码（如有前后端分离）
│   ├── index.html              # 主页面
│   ├── js/                     # JavaScript 文件
│   └── css/                    # CSS 文件
├── go.mod                      # Go 模块文件
├── go.sum                      # Go 依赖锁定文件
└── README.md                   # 项目说明文档

```

## 第三方库

### 阿里云 oss 存储音乐

* 注意点
  * 性能消耗，涉及到服务器工作，总是会占用一些资源
  * 马内消耗，好的服务很多，但是贵
  * 安全问题，资源被盗、被篡改、被恶意消耗问题

* 分发音乐模式
  * oss 签名 url，安全，可设置限时访问，需要服务器生成，消耗性能
  * oss + cnd，贵，cdn 分发进一步支持高并发
