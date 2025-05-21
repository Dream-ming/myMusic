# 数据库设计

## 数据库表设计

* 歌曲信息表（宽表）
  * 冗余 artist_name 等字段，避免连表查询，提升读取速度
* 歌曲播放计数表（窄表）
  * 仅有歌曲 id 和播放次数两个字段，避免数据冗余，提升查询速度
  * 会反复被查询和修改，配合 Redis 做排行榜缓存

### 表结构
![alt text](e36c2ef16ee0065c1baec6179f452b5.png)
### 用户信息表
![alt text](73ab1c738428848a8a3fa02c757e18b.png)
### 歌曲评论表
![alt text](a1510883ff1f4ba45dbac9e1153a5b4.png)
### 热歌排行表
![alt text](eb3cda28f90251c365e424159debc9f.png)
### 歌曲信息表（宽表）
![alt text](72b3fb95de19baa00ffb08f305b9f8a.png)
### 歌单表
![alt text](b99047f0619ae59538942ab8c71d9eb.png)
### 歌单歌曲关联表
![alt text](871b24504d703e95bc06a9c9d48e6ed.png)
### 歌曲播放计数表（窄表）
![alt text](6d87fad309454ab2bf98dffa3b34583.png)
