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

## 技术亮点

### 宽窄表设计
* 宽表
  * 适用于频繁查询且字段较多的场景
  * 比如歌曲信息表，做成宽表，减少 join 操作
* 窄表
  * 适用于字段较少，且具有高扩展性的数据
  * 比如歌曲计数表，网站每一次访问，都需要修改，做成窄表，降低IO消耗
* 权衡
  * 考虑查询频率、数据规模、字段数量

### Redis缓存设计
* 延时双删步骤
  * 更新数据库中的数据
  * 删除缓存中的数据
  * 设置延迟（500ms）
  * 再次删除缓存中的数据
* 原理
  * 避免因并发请求在第一步和第二步之间读取到旧缓存数据
  * 通过延迟再删，确保缓存在并发场景下不会污染数据库的最新数据
  * 延时的时间依据系统的写操作时间而定

### OpenSearch 搜索
* 选择 OpenSearch 的原因
  * 阿里云的产品，能与其他阿里系服务如 OSS 无缝集成
  * 提供友好的中文分词支持，能满足歌曲名和歌词的精准搜索需求
* 技术
  * 使用分词技术，对歌名和歌词进行索引
  * 定期回收和更新索引，确保新增和修改的歌曲能及时被检索到

### oss 存储音乐

* 分发音乐模式
  * oss 签名 url，安全，可设置限时访问，需要服务器生成，消耗性能
  * oss + cnd，贵，cdn 分发进一步支持高并发
* 分流高并发压力
  * 将服务器需要面对的高并发压力转移给oss
  * oss存放音乐、歌词、封面，通过分布式存储架构，应对高并发
  * cdn可以进行缓存，进一步减少访问oss的频次

### jwt + sso 单点登录
* jwt（json with token）
  * 用于用户登录后的身份验证
  * 登录成功后生成一个 jwt，包含用户的身份信息和签名
  * 客户端每次请求都会携带该jwt
* sso（单点登录）
  * 通过统一的认证中心管理用户登陆状态
  * 用户只需要登录一次就可以访问整个系统（如 music 界面和 player 界面）
  * 单设备登录，对每个用户仅维护一个 token，新设备登录时，更新 token，即可挤掉旧设备
  * 多设备登录，维护一个 user_tokens 表，存储用户的每个 jwt 即可

* token
  * 存放在客户端，每次请求时，带上token
  * 跨平台支持，可在 web、移动端、第三方应用之间通用
  * 服务器不需要存储用户登录状态，token本身携带用户状态
  * token可以自定义内容
* cookie
  * 服务器通过 Set-Cookie 响应头将 cookie 发送给客户端
  * 浏览器自动发送 cookie，无需开发者处理
  * 容易被窃取
* session
  * 用户登录后，服务器生成一个 session id，并存储用户的状态信息
  * 通过 cookie 将 id 发送给浏览器
  * 浏览器发送消息时，通过 cookie 将 session id 带上，服务器根据 id 查找对应的用户信息

### kafka 消息队列
* 优点
  * Kafka 具备高吞吐量和高可用性，适合处理大规模并发的消息
  * Kafka 提供了分区和副本机制，保障了消息的可靠性和分布式处理能力
* 应用场景
  * 用户评论的异步处理：当用户提交评论时，评论数据会先写入 Kafka 队列，后台消费者再异步写入数据库
  * 数据统计任务：将用户的行为日志（如播放、搜索）写入 Kafka，后续进行大数据分析
  * 异步通知：如用户收到点赞或回复时，系统通过 Kafka 进行消息处理和延迟发送


## 项目扩展

