# MOREDOC - 魔豆文库 产品文档

> 版本：v1.0.0  
> 更新日期：2026-02-21  
> 文档状态：正式发布

---

## 目录

1. [产品概述](#1-产品概述)
2. [技术架构](#2-技术架构)
3. [核心功能模块](#3-核心功能模块)
4. [API 接口规范](#4-api-接口规范)
5. [使用指南](#5-使用指南)
6. [常见问题解答](#6-常见问题解答)

---

## 1. 产品概述

### 1.1 产品介绍

**MOREDOC（魔豆文库）** 是由深圳市摩枫网络科技（MNT.Ltd）使用 Golang 开发的开源文库系统，类似于百度文库、新浪爱问文库。它是 dochub 文库的重构版本，旨在提供功能完善、性能优越的文档在线预览与管理解决方案。

### 1.2 核心特性

- **多格式支持**：支持 TXT、PDF、EPUB、MOBI、Office（Word/Excel/PPT）等格式文档的在线预览
- **全文搜索**：基于结巴分词实现文档内容的智能搜索
- **用户系统**：完善的用户注册、登录、权限管理机制
- **内容管理**：文档分类、标签、推荐、审核等完整的内容管理流程
- **社交功能**：评论、收藏、关注、用户动态等互动功能
- **积分系统**：支持积分下载、签到奖励等激励机制
- **响应式设计**：基于 Nuxt4 + Element Plus + SCSS 的前端框架，支持多端适配

### 1.3 应用场景

- 企业内部知识库管理
- 在线教育平台的课件管理
- 个人文档分享站点
- 行业资料共享平台

---

## 2. 技术架构

### 2.1 技术栈

| 层级 | 技术选型 | 说明 |
|------|----------|------|
| **后端** | Golang 1.18+ | 主要开发语言 |
| **Web 框架** | Gin | HTTP Web 框架 |
| **RPC 框架** | gRPC | 高性能 RPC 通信 |
| **ORM** | GORM | 数据库对象关系映射 |
| **前端框架** | Nuxt4 | 服务端渲染前端框架 |
| **UI 组件** | Element Plus | 桌面端组件库 |
| **样式预处理** | SCSS | CSS 预处理器 |
| **数据库** | MySQL 5.7+ | 关系型数据库 |
| **缓存** | 内置缓存 | 文档预览缓存 |
| **文档转换** | LibreOffice / 其他 | Office 文档转 PDF |

### 2.2 系统架构图

```
┌─────────────────────────────────────────────────────────────┐
│                        前端层 (Frontend)                      │
│  ┌─────────────┐  ┌─────────────┐  ┌─────────────────────┐  │
│  │   Web 端    │  │   管理后台   │  │      用户中心        │  │
│  │  (Nuxt4)    │  │(Element Plus)│  │                     │  │
│  └─────────────┘  └─────────────┘  └─────────────────────┘  │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      API 网关层 (Gateway)                     │
│              RESTful API / gRPC / Protocol Buffers           │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      业务逻辑层 (Service)                     │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌────────────────┐ │
│  │  用户服务 │ │  文档服务 │ │  文章服务 │ │    系统服务     │ │
│  └──────────┘ └──────────┘ └──────────┘ └────────────────┘ │
└─────────────────────────────────────────────────────────────┘
                              │
                              ▼
┌─────────────────────────────────────────────────────────────┐
│                      数据存储层 (Storage)                     │
│  ┌──────────┐ ┌──────────┐ ┌──────────┐ ┌────────────────┐ │
│  │  MySQL   │ │  文件存储 │ │  缓存层   │ │   搜索引擎      │ │
│  └──────────┘ └──────────┘ └──────────┘ └────────────────┘ │
└─────────────────────────────────────────────────────────────┘
```

### 2.3 项目目录结构

```
moredoc/
├── api/                    # Protocol Buffers API 定义
│   └── v1/                 # API v1 版本
├── biz/                    # 业务逻辑层
├── cmd/                    # 命令行工具
├── cache/                  # 缓存相关
├── conf/                   # 配置定义
├── dict/                   # 结巴分词字典
├── docs/                   # 文档目录
├── documents/              # 用户上传的文档存储目录
├── middleware/             # 中间件
├── model/                  # 数据库模型（GORM）
├── service/                # 服务层
├── sitemap/                # 站点地图
├── third_party/            # 第三方依赖
├── uploads/                # 其他文件存储目录
├── util/                   # 工具函数
├── main.go                 # 项目入口
├── app.example.toml        # 配置文件示例
└── go.mod                  # Go 依赖管理
```

---

## 3. 核心功能模块

### 3.1 用户管理模块

#### 3.1.1 功能概述

用户管理模块负责用户的注册、登录、权限控制和用户信息管理。

#### 3.1.2 核心功能

| 功能 | 说明 |
|------|------|
| 用户注册 | 支持用户名、邮箱注册，可选邮箱验证 |
| 用户登录 | 支持密码登录、验证码登录 |
| 密码管理 | 密码修改、找回密码 |
| 用户组管理 | 基于角色的权限控制（RBAC）|
| 用户资料 | 头像、签名、个人信息的维护 |
| 用户动态 | 记录用户操作动态 |
| 每日签到 | 积分签到奖励机制 |

#### 3.1.3 用户状态

- **正常状态**：用户可以正常使用系统功能
- **禁用状态**：用户被禁止登录和使用
- **待验证**：邮箱待验证状态

### 3.2 文档管理模块

#### 3.2.1 功能概述

文档管理是系统的核心模块，负责文档的上传、转换、预览、下载等全生命周期管理。

#### 3.2.2 支持的文档格式

| 格式类型 | 扩展名 | 在线预览 | 下载 |
|----------|--------|----------|------|
| 文本文件 | .txt | ✅ | ✅ |
| PDF 文档 | .pdf | ✅ | ✅ |
| Word 文档 | .doc, .docx | ✅ | ✅ |
| Excel 表格 | .xls, .xlsx | ✅ | ✅ |
| PPT 演示 | .ppt, .pptx | ✅ | ✅ |
| EPUB 电子书 | .epub | ✅ | ✅ |
| MOBI 电子书 | .mobi | ✅ | ✅ |

#### 3.2.3 文档生命周期

```
上传 → 格式转换 → 内容提取 → 审核 → 发布 → 预览/下载 → 归档/删除
```

#### 3.2.4 文档状态

| 状态 | 说明 |
|------|------|
| 待审核 | 文档上传后等待管理员审核 |
| 已通过 | 审核通过，正常显示 |
| 已拒绝 | 审核未通过，需修改 |
| 已删除 | 进入回收站 |
| 已推荐 | 首页推荐展示 |

#### 3.2.5 文档属性

- **基础信息**：标题、描述、关键词、作者
- **文件信息**：大小、页数、格式、封面
- **统计信息**：浏览量、下载量、收藏量、评分
- **商业属性**：价格（积分）、是否收费

### 3.3 分类管理模块

#### 3.3.1 功能概述

分类管理支持多级分类体系，用于组织和归类文档与文章。

#### 3.3.2 分类特性

- **多级分类**：支持无限级分类层级
- **分类图标**：支持自定义分类图标和封面
- **分类排序**：自定义排序，值越大越靠前
- **分类统计**：显示分类下的文档数量
- **分类启用/禁用**：灵活控制分类显示

### 3.4 文章管理模块

#### 3.4.1 功能概述

文章管理用于发布和管理站点公告、帮助文档、新闻资讯等内容。

#### 3.4.2 文章特性

- **富文本编辑**：支持富文本内容编辑
- **文章分类**：支持多分类关联
- **文章推荐**：首页推荐展示
- **文章审核**：发布审核机制
- **唯一标识**：支持自定义文章标识（URL 友好）

### 3.5 评论管理模块

#### 3.5.1 功能概述

评论系统支持用户对文档和文章进行评论互动。

#### 3.5.2 评论特性

- **多级评论**：支持回复评论（二级评论）
- **评论审核**：支持评论审核机制
- **评论类型**：支持文档评论和文章评论
- **验证码保护**：防止恶意刷评论

### 3.6 收藏管理模块

#### 3.6.1 功能概述

收藏功能允许用户收藏感兴趣的文档和文章。

#### 3.6.2 收藏特性

- **类型区分**：支持文档收藏和文章收藏
- **快速访问**：从个人中心快速访问收藏内容
- **收藏统计**：显示收藏数量

### 3.7 系统配置模块

#### 3.7.1 配置分类

| 配置类型 | 说明 |
|----------|------|
| 系统配置 | 站点基本信息、SEO 配置 |
| 显示配置 | 页面显示选项、样式配置 |
| 安全配置 | 验证码、注册限制、访问控制 |
| 底链配置 | 页脚链接配置 |
| 语言配置 | 多语言支持配置 |

#### 3.7.2 系统监控

- **系统统计**：用户、文档、文章等数据统计
- **设备信息**：CPU、内存、磁盘使用情况
- **环境检测**：依赖环境检查
- **版本管理**：程序版本和更新检查

### 3.8 内容运营模块

#### 3.8.1 轮播图管理

- 支持 PC 端、移动端等不同位置的轮播图
- 支持排序和启用/禁用控制
- 支持跳转链接配置

#### 3.8.2 友情链接管理

- 站点友情链接管理
- 支持排序和启用/禁用控制

#### 3.8.3 导航管理

- 顶部导航、底部导航配置
- 支持多级导航
- 支持自定义颜色和打开方式

#### 3.8.4 广告管理

- 支持多种广告位配置
- 支持广告投放时间控制
- 支持广告内容自定义

### 3.9 权限管理模块

#### 3.9.1 权限体系

- **基于角色的访问控制（RBAC）**
- **用户组权限**：不同用户组拥有不同权限
- **功能权限**：细粒度的功能操作权限
- **数据权限**：数据级别的访问控制

#### 3.9.2 用户组特性

- **默认用户组**：新用户自动分配
- **权限继承**：用户组权限继承机制
- **功能开关**：上传、评论、发布文章等权限控制
- **审核设置**：评论审核、文档审核配置

### 3.10 搜索与记录模块

#### 3.10.1 文档搜索

- **全文搜索**：基于文档内容的智能搜索
- **多条件筛选**：分类、格式、时间、语言等筛选
- **搜索建议**：热门搜索词推荐
- **搜索结果排序**：相关度、时间、热度排序

#### 3.10.2 搜索记录

- 记录用户搜索行为
- 搜索统计和分析
- 搜索热词统计

### 3.11 举报与处罚模块

#### 3.11.1 举报功能

- 支持文档举报
- 多种举报原因
- 举报处理流程

#### 3.11.2 处罚功能

- 用户处罚管理
- 处罚类型：禁止上传、禁止评论等
- 处罚期限设置

---

## 4. API 接口规范

### 4.1 接口设计原则

- 统一使用 Protocol Buffers 定义接口
- RESTful API 风格，通过 gRPC-Gateway 暴露 HTTP 接口
- 统一接口前缀：`/api/v1/`
- 统一返回格式，包含状态码、消息和数据

### 4.2 通用请求参数

#### 4.2.1 分页参数

| 参数 | 类型 | 说明 |
|------|------|------|
| page | int64 | 页码，从 1 开始 |
| size | int64 | 每页数量，默认 20 |

#### 4.2.2 搜索参数

| 参数 | 类型 | 说明 |
|------|------|------|
| wd | string | 搜索关键词 |
| field | repeated string | 查询字段 |
| order | string | 排序字段 |

#### 4.2.3 通用响应结构

```protobuf
message ListXXXReply {
  int64 total = 1;                    // 总数
  repeated XXX items = 2;             // 数据列表
}
```

### 4.3 用户管理 API

#### 4.3.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| Register | POST | /api/v1/user/register | 用户注册 |
| Login | POST | /api/v1/user/login | 用户登录 |
| Logout | DELETE | /api/v1/user/logout | 退出登录 |
| GetUser | GET | /api/v1/user | 获取用户信息 |
| ListUser | GET | /api/v1/user/list | 用户列表 |
| AddUser | POST | /api/v1/user | 新增用户 |
| SetUser | PUT | /api/v1/user | 设置用户 |
| DeleteUser | DELETE | /api/v1/user | 删除用户 |
| UpdateUserPassword | PUT | /api/v1/user/password | 修改密码 |
| UpdateUserProfile | PUT | /api/v1/user/profile | 更新资料 |
| GetUserCaptcha | GET | /api/v1/user/captcha | 获取验证码 |
| GetUserPermissions | GET | /api/v1/user/permission | 获取权限 |
| CanIUploadDocument | GET | /api/v1/user/caniuploaddocument | 检查上传权限 |
| CanIPublishArticle | GET | /api/v1/user/canipublisharticle | 检查发布权限 |
| ListUserDynamic | GET | /api/v1/user/dynamic | 用户动态 |
| SignToday | PUT | /api/v1/user/sign | 每日签到 |
| GetSignedToday | GET | /api/v1/user/sign | 获取签到状态 |
| ListUserDownload | GET | /api/v1/user/download | 下载记录 |
| ListUserGroup | GET | /api/v1/user/group | 用户组列表 |
| FindPasswordStepOne | POST | /api/v1/user/findpassword/stepone | 找回密码第一步 |
| FindPasswordStepTwo | PUT | /api/v1/user/findpassword/steptwo | 找回密码第二步 |
| SendEmailCode | POST | /api/v1/user/email/code | 发送邮箱验证码 |

#### 4.3.2 核心数据结构

```protobuf
// 用户信息
message User {
  int64 id = 1;                       // 用户ID
  string username = 2;                // 用户名
  string mobile = 3;                  // 手机号
  string email = 4;                   // 邮箱
  string address = 5;                 // 地址
  string signature = 6;               // 个性签名
  string last_login_ip = 7;           // 最后登录IP
  string register_ip = 8;             // 注册IP
  int32 doc_count = 9;                // 文档数量
  int32 follow_count = 10;            // 关注数量
  int32 fans_count = 11;              // 粉丝数量
  int32 favorite_count = 12;          // 收藏数量
  int32 comment_count = 13;           // 评论数量
  int32 status = 14;                  // 用户状态
  string avatar = 15;                 // 头像
  string identity = 16;               // 身份证
  string realname = 17;               // 真实姓名
  repeated int64 group_id = 18;       // 用户组ID
  google.protobuf.Timestamp login_at = 19;  // 最后登录时间
  google.protobuf.Timestamp created_at = 20; // 注册时间
  google.protobuf.Timestamp updated_at = 21; // 更新时间
  int32 credit_count = 22;            // 积分
  int32 article_count = 23;           // 文章数量
  string remark = 24;                 // 备注
}

// 登录请求
message RegisterAndLoginRequest {
  string username = 1;                // 用户名
  string password = 2;                // 密码
  string captcha = 3;                 // 验证码
  string captcha_id = 4;              // 验证码ID
  string email = 5;                   // 邮箱
  string code = 6;                    // 邮箱验证码
}

// 登录响应
message LoginReply {
  string token = 1;                   // JWT Token
  User user = 2;                      // 用户信息
}
```

### 4.4 文档管理 API

#### 4.4.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| CreateDocument | POST | /api/v1/document | 创建文档 |
| UpdateDocument | PUT | /api/v1/document | 更新文档 |
| DeleteDocument | DELETE | /api/v1/document | 删除文档 |
| GetDocument | GET | /api/v1/document | 获取文档详情 |
| ListDocument | GET | /api/v1/document/list | 文档列表 |
| SearchDocument | GET | /api/v1/document/search | 搜索文档 |
| ListDocumentForHome | GET | /api/v1/document/home | 首页文档 |
| GetRelatedDocuments | GET | /api/v1/document/related | 相关文档 |
| DownloadDocument | GET | /api/v1/document/download | 下载文档 |
| CheckDocument | PUT | /api/v1/document/check | 审核文档 |
| SetDocumentRecommend | PUT | /api/v1/document/recommend | 设置推荐 |
| SetDocumentScore | POST | /api/v1/document/score | 文档评分 |
| GetDocumentScore | GET | /api/v1/document/score | 获取评分 |
| SetDocumentReconvert | PUT | /api/v1/document/reconvert | 重新转换 |
| SetDocumentsCategory | PUT | /api/v1/document/category | 批量设置分类 |
| SetDocumentsLanguage | PUT | /api/v1/document/language | 批量设置语言 |
| DownloadDocumentToBeReviewed | GET | /api/v1/document/download/bereviewed | 下载待审核文档 |

#### 4.4.2 回收站接口

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| ListRecycleDocument | GET | /api/v1/document/recycle | 回收站列表 |
| RecoverRecycleDocument | PUT | /api/v1/document/recycle | 恢复文档 |
| DeleteRecycleDocument | DELETE | /api/v1/document/recycle | 彻底删除 |
| ClearRecycleDocument | DELETE | /api/v1/document/recycle/all | 清空回收站 |

#### 4.4.3 核心数据结构

```protobuf
// 文档
message Document {
  int64 id = 1;                       // 文档ID
  string title = 2;                   // 文档标题
  string keywords = 3;                // 文档关键字
  string description = 4;             // 文档描述
  int64 user_id = 5;                  // 文档作者ID
  string cover = 6;                   // 文档封面
  int32 width = 7;                    // 文档宽度
  int32 height = 8;                   // 文档高度
  int32 preview = 9;                  // 可预览页数
  int32 pages = 10;                   // 文档页数
  string uuid = 11;                   // 文档UUID
  int32 download_count = 12;          // 下载次数
  int32 view_count = 13;              // 浏览次数
  int32 favorite_count = 14;          // 收藏次数
  int32 comment_count = 15;           // 评论次数
  int32 score = 16;                   // 文档评分
  int32 score_count = 17;             // 评分次数
  int32 price = 18;                   // 文档价格
  int64 size = 19;                    // 文档大小
  int32 status = 20;                  // 文档状态
  google.protobuf.Timestamp created_at = 21;  // 创建时间
  google.protobuf.Timestamp updated_at = 22;  // 更新时间
  google.protobuf.Timestamp deleted_at = 23;  // 删除时间
  google.protobuf.Timestamp recommend_at = 29; // 推荐时间
  int64 deleted_user_id = 24;         // 删除用户ID
  string username = 25;               // 作者用户名
  repeated int64 category_id = 26;    // 分类ID
  string deleted_username = 27;       // 删除用户名
  string ext = 28;                    // 扩展名
  Attachment attachment = 30;         // 附件信息
  User user = 31;                     // 作者信息
  bool enable_gzip = 32;              // 是否启用gzip
  string convert_error = 33;          // 转换错误信息
  string preview_ext = 34;            // 预览扩展名
  string language = 35;               // 文档语言
  string content = 36;                // 文档内容
  repeated Category category = 37;    // 分类信息
  string source = 38;                 // 来源
  string source_url = 39;             // 来源链接
}

// 创建文档请求
message CreateDocumentRequest {
  bool overwrite = 1;                 // 是否覆盖
  repeated int64 category_id = 2;     // 分类ID
  repeated CreateDocumentItem document = 3; // 文档列表
}

message CreateDocumentItem {
  string title = 1;                   // 文档标题
  int64 attachment_id = 2;            // 附件ID
  int32 price = 3;                    // 文档价格
  string language = 4;                // 文档语言
  string keywords = 5;                // 关键字
  string description = 6;             // 描述
  string source = 7;                  // 来源
  string source_url = 8;              // 来源链接
}

// 文档搜索请求
message SearchDocumentRequest {
  int32 page = 1;                     // 页码
  int32 size = 2;                     // 每页数量
  string wd = 3;                      // 搜索关键词
  repeated int64 category_id = 4;     // 分类筛选
  string sort = 5;                    // 排序方式
  string ext = 7;                     // 格式筛选
  int64 user_id = 8;                  // 用户筛选
  repeated string created_at = 9;     // 时间范围
  repeated string language = 10;      // 语言筛选
}
```

### 4.5 分类管理 API

#### 4.5.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| CreateCategory | POST | /api/v1/category | 创建分类 |
| UpdateCategory | PUT | /api/v1/category | 更新分类 |
| DeleteCategory | DELETE | /api/v1/category | 删除分类 |
| GetCategory | GET | /api/v1/category | 获取分类 |
| ListCategory | GET | /api/v1/category/list | 分类列表 |

#### 4.5.2 核心数据结构

```protobuf
message Category {
  int64 id = 1;                       // 分类ID
  int64 parent_id = 2;                // 父分类ID
  string title = 3;                   // 分类标题
  int32 doc_count = 4;                // 文档数量
  int32 sort = 5;                     // 排序
  bool enable = 6;                    // 是否启用
  google.protobuf.Timestamp created_at = 7;  // 创建时间
  google.protobuf.Timestamp updated_at = 8;  // 更新时间
  string cover = 9;                   // 分类封面
  string icon = 10;                   // 分类图标
  string description = 11;            // 分类描述
  bool show_description = 12;         // 是否显示描述
  int32 type = 13;                    // 分类类型
}
```

### 4.6 文章管理 API

#### 4.6.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| CreateArticle | POST | /api/v1/article | 创建文章 |
| UpdateArticle | PUT | /api/v1/article | 更新文章 |
| DeleteArticle | DELETE | /api/v1/article | 删除文章 |
| GetArticle | GET | /api/v1/article | 获取文章 |
| ListArticle | GET | /api/v1/article/list | 文章列表 |
| SetArticlesCategory | PUT | /api/v1/article/category | 批量设置分类 |
| RecommendArticles | PUT | /api/v1/article/recommend | 批量推荐 |
| CheckArticles | PUT | /api/v1/article/check | 批量审核 |
| ListRecycleArticle | GET | /api/v1/article/recycle/list | 回收站列表 |
| RestoreRecycleArticle | POST | /api/v1/article/recycle/restore | 恢复文章 |
| DeleteRecycleArticle | DELETE | /api/v1/article/recycle | 彻底删除 |
| EmptyRecycleArticle | DELETE | /api/v1/article/recycle/empty | 清空回收站 |

#### 4.6.2 核心数据结构

```protobuf
message Article {
  int64 id = 1;                       // 文章ID
  string identifier = 2;              // 唯一标识
  string author = 3;                  // 作者
  int64 view_count = 4;               // 浏览次数
  string title = 5;                   // 标题
  string keywords = 6;                // 关键字
  string description = 7;             // 描述
  string content = 8;                 // 内容
  google.protobuf.Timestamp created_at = 9;   // 创建时间
  google.protobuf.Timestamp updated_at = 10;  // 更新时间
  google.protobuf.Timestamp deleted_at = 11;  // 删除时间
  repeated int64 category_id = 12;    // 分类ID
  int64 favorite_count = 13;          // 收藏数
  int64 comment_count = 14;           // 评论数
  google.protobuf.Timestamp recommend_at = 15; // 推荐时间
  bool is_recommend = 16;             // 是否推荐
  int64 user_id = 17;                 // 作者ID
  User user = 18;                     // 作者信息
  int32 status = 19;                  // 状态
  string reject_reason = 20;          // 拒绝原因
  repeated Category category = 21;    // 分类信息
  string source = 22;                 // 来源
  string source_url = 23;             // 来源链接
}
```

### 4.7 评论管理 API

#### 4.7.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| CreateComment | POST | /api/v1/comment | 创建评论 |
| UpdateComment | PUT | /api/v1/comment | 更新评论 |
| DeleteComment | DELETE | /api/v1/comment | 删除评论 |
| GetComment | GET | /api/v1/comment | 获取评论 |
| ListComment | GET | /api/v1/comment/list | 评论列表 |
| CheckComment | POST | /api/v1/comment/check | 审核评论 |

#### 4.7.2 核心数据结构

```protobuf
message Comment {
  google.protobuf.Timestamp created_at = 1;   // 创建时间
  google.protobuf.Timestamp updated_at = 2;   // 更新时间
  int64 id = 3;                       // 评论ID
  int64 parent_id = 4;                // 父评论ID
  string content = 5;                 // 评论内容
  int64 document_id = 6;              // 文档ID
  int32 status = 7;                   // 状态
  int32 comment_count = 8;            // 回复数量
  int64 user_id = 9;                  // 用户ID
  User user = 10;                     // 用户信息
  string document_title = 11;         // 文档标题
  int32 type = 12;                    // 类型：0文档，1文章
  string document_uuid = 13;          // 文档UUID
}
```

### 4.8 收藏管理 API

#### 4.8.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| CreateFavorite | POST | /api/v1/favorite | 添加收藏 |
| DeleteFavorite | DELETE | /api/v1/favorite | 取消收藏 |
| GetFavorite | GET | /api/v1/favorite | 获取收藏状态 |
| ListFavorite | GET | /api/v1/favorite/list | 收藏列表 |

#### 4.8.2 核心数据结构

```protobuf
message Favorite {
  int64 id = 1;                       // 收藏ID
  int64 user_id = 2;                  // 用户ID
  int64 document_id = 3;              // 文档ID
  google.protobuf.Timestamp created_at = 4;   // 创建时间
  google.protobuf.Timestamp updated_at = 5;   // 更新时间
  string title = 6;                   // 标题
  string ext = 7;                     // 扩展名
  int32 score = 8;                    // 评分
  int64 size = 9;                     // 大小
  int32 pages = 10;                   // 页数
  string document_uuid = 11;          // 文档UUID
  int32 type = 12;                    // 类型：0文档，1文章
}
```

### 4.9 附件管理 API

#### 4.9.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| UpdateAttachment | PUT | /api/v1/attachment | 更新附件 |
| DeleteAttachment | DELETE | /api/v1/attachment | 删除附件 |
| GetAttachment | GET | /api/v1/attachment | 获取附件 |
| ListAttachment | GET | /api/v1/attachment/list | 附件列表 |

#### 4.9.2 核心数据结构

```protobuf
message Attachment {
  int64 id = 1;                       // 附件ID
  string hash = 2;                    // 文件哈希（MD5）
  int64 user_id = 3;                  // 上传用户ID
  int64 type_id = 4;                  // 类型ID
  int32 type = 5;                     // 附件类型
  bool enable = 6;                    // 是否启用
  string path = 7;                    // 附件路径
  string name = 8;                    // 附件名称
  int64 size = 9;                     // 附件大小
  int64 width = 10;                   // 宽度（图片）
  int64 height = 11;                  // 高度（图片）
  string ext = 12;                    // 扩展名
  string ip = 13;                     // 上传IP
  google.protobuf.Timestamp created_at = 14;  // 创建时间
  google.protobuf.Timestamp updated_at = 15;  // 更新时间
  string username = 16;               // 用户名
  string type_name = 17;              // 类型名称
  string description = 18;            // 描述
}
```

### 4.10 用户组管理 API

#### 4.10.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| CreateGroup | POST | /api/v1/group | 创建用户组 |
| UpdateGroup | PUT | /api/v1/group | 更新用户组 |
| DeleteGroup | DELETE | /api/v1/group | 删除用户组 |
| GetGroup | GET | /api/v1/group | 获取用户组 |
| ListGroup | GET | /api/v1/group/list | 用户组列表 |
| GetGroupPermission | GET | /api/v1/group/permission | 获取权限 |
| UpdateGroupPermission | PUT | /api/v1/group/permission | 更新权限 |

#### 4.10.2 核心数据结构

```protobuf
message Group {
  int64 id = 1;                       // 用户组ID
  string title = 2;                   // 用户组名称
  string color = 3;                   // 颜色
  bool is_default = 4;                // 是否默认
  bool is_display = 5;                // 是否显示
  string description = 6;             // 描述
  int32 user_count = 7;               // 用户数量
  int32 sort = 8;                     // 排序
  google.protobuf.Timestamp created_at = 9;   // 创建时间
  google.protobuf.Timestamp updated_at = 10;  // 更新时间
  bool enable_upload = 11;            // 允许上传
  bool enable_comment_approval = 12;  // 评论需审核
  bool enable_comment = 13;           // 允许评论
  bool enable_document_review = 14;   // 文档需审核
  bool enable_article = 15;           // 允许发布文章
  bool enable_article_approval = 16;  // 文章需审核
}
```

### 4.11 系统配置 API

#### 4.11.1 接口列表

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| GetSettings | GET | /api/v1/settings | 获取系统配置 |
| UpdateConfig | PUT | /api/v1/config | 更新配置 |
| ListConfig | GET | /api/v1/config/list | 配置列表 |
| GetStats | GET | /api/v1/stats | 系统统计 |
| GetEnvs | GET | /api/v1/envs | 环境检测 |
| UpdateSitemap | PUT | /api/v1/sitemap | 更新站点地图 |
| GetDeviceInfo | GET | /api/v1/device | 设备信息 |
| SetSQLMode | PUT | /api/v1/sqlmode | 设置SQL模式 |
| GetLatestRelease | GET | /api/v1/release | 最新版本 |
| RefreshLatestRelease | POST | /api/v1/release | 刷新版本 |
| IgnoreRelease | PUT | /api/v1/release/ignore | 忽略版本 |
| SetReleaseSource | PUT | /api/v1/release/source | 设置更新源 |

#### 4.11.2 核心数据结构

```protobuf
// 系统配置
message ConfigSystem {
  string domain = 1;                  // 站点域名
  string title = 2;                   // 站点标题
  string keywords = 3;                // 站点关键词
  string description = 4;             // 站点描述
  string logo = 5;                    // Logo
  string favicon = 6;                 // Favicon
  string icp = 7;                     // 备案号
  string analytics = 8;               // 统计代码
  string sitename = 9;                // 站点名称
  string copyright_start_year = 10;   // 版权起始年
  string register_background = 11;    // 注册页背景
  string login_background = 12;       // 登录页背景
  repeated string recommend_words = 13; // 推荐搜索词
  string version = 14;                // 版本号
  string credit_name = 15;            // 积分名称
  string sec_icp = 16;                // 公安备案
}

// 安全配置
message ConfigSecurity {
  bool is_close = 1;                  // 是否关闭站点
  string close_statement = 2;         // 关闭说明
  bool enable_register = 3;           // 允许注册
  bool enable_captcha_login = 4;      // 登录验证码
  bool enable_captcha_register = 5;   // 注册验证码
  bool enable_captcha_comment = 6;    // 评论验证码
  bool enable_captcha_find_password = 7; // 找回密码验证码
  bool enable_captcha_upload = 8;     // 上传验证码
  int32 max_document_size = 9;        // 最大文档大小
  repeated string document_allowed_ext = 10; // 允许格式
  bool login_required = 11;           // 需登录访问
  bool enable_verify_register_email = 12; // 邮箱验证
}

// 系统统计
message Stats {
  int64 user_count = 1;               // 用户数量
  int64 document_count = 2;           // 文档数量
  int64 category_count = 3;           // 分类数量
  int64 article_count = 4;            // 文章数量
  int64 comment_count = 5;            // 评论数量
  int64 banner_count = 6;             // Banner数量
  int64 friendlink_count = 7;         // 友情链接数量
  string os = 8;                      // 操作系统
  string version = 9;                 // 版本号
  string hash = 10;                   // Git Hash
  string build_at = 11;               // 构建时间
  int64 report_count = 12;            // 举报数量
}
```

### 4.12 其他管理 API

#### 4.12.1 轮播图管理

| 接口 | 方法 | 路径 |
|------|------|------|
| CreateBanner | POST | /api/v1/banner |
| UpdateBanner | PUT | /api/v1/banner |
| DeleteBanner | DELETE | /api/v1/banner |
| GetBanner | GET | /api/v1/banner |
| ListBanner | GET | /api/v1/banner/list |

#### 4.12.2 友情链接管理

| 接口 | 方法 | 路径 |
|------|------|------|
| CreateFriendlink | POST | /api/v1/friendlink |
| UpdateFriendlink | PUT | /api/v1/friendlink |
| DeleteFriendlink | DELETE | /api/v1/friendlink |
| GetFriendlink | GET | /api/v1/friendlink |
| ListFriendlink | GET | /api/v1/friendlink/list |

#### 4.12.3 导航管理

| 接口 | 方法 | 路径 |
|------|------|------|
| CreateNavigation | POST | /api/v1/navigation |
| UpdateNavigation | PUT | /api/v1/navigation |
| DeleteNavigation | DELETE | /api/v1/navigation |
| GetNavigation | GET | /api/v1/navigation |
| ListNavigation | GET | /api/v1/navigation/list |

#### 4.12.4 广告管理

| 接口 | 方法 | 路径 |
|------|------|------|
| CreateAdvertisement | POST | /api/v1/advertisement |
| UpdateAdvertisement | PUT | /api/v1/advertisement |
| DeleteAdvertisement | DELETE | /api/v1/advertisement |
| GetAdvertisement | GET | /api/v1/advertisement |
| GetAdvertisementByPosition | GET | /api/v1/advertisement/position |
| ListAdvertisement | GET | /api/v1/advertisement/list |

#### 4.12.5 语言管理

| 接口 | 方法 | 路径 |
|------|------|------|
| CreateLanguage | POST | /api/v1/language |
| UpdateLanguage | PUT | /api/v1/language |
| UpdateLanguageStatus | PUT | /api/v1/language/status |
| ListLanguage | GET | /api/v1/language/list |
| DeleteLanguage | DELETE | /api/v1/language |

#### 4.12.6 举报管理

| 接口 | 方法 | 路径 |
|------|------|------|
| CreateReport | POST | /api/v1/report |
| UpdateReport | PUT | /api/v1/report |
| DeleteReport | DELETE | /api/v1/report |
| ListReport | GET | /api/v1/report/list |

#### 4.12.7 处罚管理

| 接口 | 方法 | 路径 |
|------|------|------|
| CreatePunishment | POST | /api/v1/punishment |
| UpdatePunishment | PUT | /api/v1/punishment |
| GetPunishment | GET | /api/v1/punishment |
| ListPunishment | GET | /api/v1/punishment/list |
| CancelPunishment | PUT | /api/v1/punishment/cancel |

#### 4.12.8 搜索记录管理

| 接口 | 方法 | 路径 |
|------|------|------|
| DeleteSearchRecord | DELETE | /api/v1/searchrecord |
| ListSearchRecord | GET | /api/v1/searchrecord/list |

#### 4.12.9 权限管理

| 接口 | 方法 | 路径 |
|------|------|------|
| UpdatePermission | PUT | /api/v1/permission |
| GetPermission | GET | /api/v1/permission |
| ListPermission | GET | /api/v1/permission/list |

#### 4.12.10 健康检查

| 接口 | 方法 | 路径 | 说明 |
|------|------|------|------|
| Health | GET | /health | 健康检查 |
| Ping | GET | /api/v1/ping | 连通性测试 |

---

## 5. 使用指南

### 5.1 环境要求

#### 5.1.1 后端环境

- **Go 版本**：1.18 或更高
- **MySQL 版本**：5.7 或更高
- **操作系统**：Linux / macOS / Windows

#### 5.1.2 前端环境

- **Node.js**：14.16.0（推荐使用 nvm 管理）
- **npm**：6.x 或更高

#### 5.1.3 可选依赖

- **LibreOffice**：用于 Office 文档转换
- **Nginx**：用于反向代理和静态资源服务

### 5.2 安装部署

#### 5.2.1 后端部署

```bash
# 1. 克隆代码
git clone https://github.com/mnt-ltd/moredoc.git
cd moredoc

# 2. 安装依赖
go mod tidy

# 3. 初始化工程依赖
make init

# 4. 编译 proto api
make api

# 5. 复制配置文件
cp app.example.toml app.toml

# 6. 修改 app.toml 中的数据库配置
# 编辑 app.toml 文件，配置 MySQL 连接信息

# 7. 编译后端
go build -o moredoc main.go

# 8. 初始化数据库
./moredoc syncdb

# 9. 运行后端
go run main.go serve
```

#### 5.2.2 前端部署

```bash
# 1. 进入前端目录
cd web

# 2. 安装依赖
npm install

# 3. 运行开发环境
npm run dev

# 4. 构建生产环境
npm run generate
```

### 5.3 配置文件说明

#### 5.3.1 app.toml 示例

```toml
# 程序运行级别：debug、info、warn、error
level = "debug"

# 日志编码方式：json、console
logEncoding = "console"

# 后端监听端口
port = "8880"

# 数据库配置
[database]
driver = "mysql"
dsn = "root:root@tcp(localhost:3306)/moredoc?charset=utf8mb4&loc=Local&parseTime=true"
showSQL = true
maxOpen = 10
maxIdle = 10

# JWT 配置
[jwt]
secret = "moredoc"
expireDays = 365
```

### 5.4 管理员账号

- **默认账号**：admin
- **默认密码**：mnt.ltd

### 5.5 发布版本

```bash
# 1. 打标签
git tag -a v1.0.0 -m "release v1.0.0"

# 2. 推送标签
git push origin v1.0.0

# 3. 编译前端
cd web && npm run generate

# 4. 编译后端（Linux）
make buildlinux

# 5. 编译后端（Windows）
make buildwin
```

---

## 6. 常见问题解答

### 6.1 部署相关问题

#### Q1: 数据库连接失败怎么办？

**A**: 请检查以下几点：
- 确认 MySQL 服务已启动
- 检查 app.toml 中的数据库连接字符串是否正确
- 确认数据库用户有相应的权限
- 检查防火墙设置是否允许连接

#### Q2: 文档转换失败怎么办？

**A**: 
- 确认 LibreOffice 已正确安装
- 检查系统 PATH 环境变量是否包含 LibreOffice 路径
- 查看日志中的具体错误信息
- 确认文档格式是否受支持

#### Q3: 前端无法访问后端 API？

**A**:
- 检查后端服务是否正常运行
- 确认前端配置中的 API 地址是否正确
- 检查跨域配置（CORS）
- 查看浏览器控制台的网络请求错误

### 6.2 功能使用问题

#### Q4: 如何修改站点名称和 Logo？

**A**: 登录管理后台，进入"系统配置"-"系统设置"，可以修改站点标题、Logo、Favicon 等信息。

#### Q5: 如何设置文档审核？

**A**: 
1. 进入管理后台的"系统配置"-"安全配置"
2. 开启"文档需要审核"选项
3. 或者在"用户组管理"中为特定用户组设置"文档需要审核"

#### Q6: 如何配置用户权限？

**A**:
1. 进入管理后台的"用户组管理"
2. 创建或编辑用户组
3. 设置用户组的权限（上传、评论、发布文章等）
4. 将用户分配到相应的用户组

#### Q7: 如何添加轮播图？

**A**:
1. 进入管理后台的"轮播图管理"
2. 点击"新增"按钮
3. 上传图片并设置标题、链接等信息
4. 设置排序和启用状态

### 6.3 性能优化问题

#### Q8: 如何提高文档预览速度？

**A**:
- 启用 gzip 压缩（文档属性中设置）
- 优化服务器带宽
- 使用 CDN 加速静态资源
- 合理设置预览页数限制

#### Q9: 数据库查询慢怎么办？

**A**:
- 为常用查询字段添加索引
- 优化 SQL 查询语句
- 增加数据库连接池大小
- 考虑使用读写分离

### 6.4 安全问题

#### Q10: 如何防止恶意注册？

**A**:
- 开启注册验证码
- 开启邮箱验证
- 设置注册频率限制
- 使用 IP 黑名单机制

#### Q11: 如何备份数据？

**A**:
- 定期备份 MySQL 数据库
- 备份 documents 和 uploads 目录下的文件
- 备份 app.toml 配置文件
- 建议使用自动化脚本定期备份

### 6.5 开发相关问题

#### Q12: 如何添加新的 API 接口？

**A**:
1. 在 api/v1/ 目录下编辑或创建 proto 文件
2. 定义 message 和 service
3. 运行 `make api` 生成代码
4. 在 biz/ 目录下实现业务逻辑
5. 在 service/ 目录下注册服务

#### Q13: 如何修改前端页面？

**A**:
- 前端代码位于 web/ 目录
- 使用 Nuxt4 + Element Plus + SCSS 框架
- 修改后运行 `npm run dev` 查看效果
- 使用 `npm run generate` 构建生产版本

### 6.6 其他问题

#### Q14: 如何获取技术支持？

**A**:
- 查看官方文档：https://www.bookstack.cn/books/moredoc
- 提交 GitHub Issue：https://github.com/mnt-ltd/moredoc/issues
- 加入微信交流群（添加微信：进击的皇虫，备注"魔豆文库加群"）

#### Q15: 如何贡献代码？

**A**:
1. Fork 项目仓库
2. 创建特性分支
3. 提交代码变更
4. 创建 Pull Request
5. 等待代码审核

---

## 附录

### A. 开源协议

本项目基于 Apache License 2.0 协议开源。

### B. 相关链接

- **GitHub**: https://github.com/mnt-ltd/moredoc
- **Gitee**: https://gitee.com/mnt-ltd/moredoc
- **官方文档**: https://www.bookstack.cn/books/moredoc
- **演示站点**: https://moredoc.mnt.ltd

### C. 鸣谢

感谢以下开源项目为魔豆文库的开发奠定基础：
- Golang 生态：Gin、GORM、gRPC 等
- Vue.js 生态：Nuxt4、Element Plus、SCSS 等
- 以及其他众多开源项目

---

**文档维护**：MNT.Ltd  
**最后更新**：2026-02-21
