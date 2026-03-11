# Protocol Documentation
<a name="top"></a>

## Table of Contents

- [api/v1/group.proto](#api_v1_group-proto)
    - [DeleteGroupRequest](#api-v1-DeleteGroupRequest)
    - [GetGroupPermissionRequest](#api-v1-GetGroupPermissionRequest)
    - [GetGroupRequest](#api-v1-GetGroupRequest)
    - [Group](#api-v1-Group)
    - [GroupPermissions](#api-v1-GroupPermissions)
    - [ListGroupReply](#api-v1-ListGroupReply)
    - [ListGroupRequest](#api-v1-ListGroupRequest)
    - [UpdateGroupPermissionRequest](#api-v1-UpdateGroupPermissionRequest)
  
    - [GroupAPI](#api-v1-GroupAPI)
  
- [api/v1/category.proto](#api_v1_category-proto)
    - [Category](#api-v1-Category)
    - [DeleteCategoryRequest](#api-v1-DeleteCategoryRequest)
    - [GetCategoryRequest](#api-v1-GetCategoryRequest)
    - [ListCategoryReply](#api-v1-ListCategoryReply)
    - [ListCategoryRequest](#api-v1-ListCategoryRequest)
  
    - [CategoryAPI](#api-v1-CategoryAPI)
  
- [api/v1/document.proto](#api_v1_document-proto)
    - [CheckDocumentRequest](#api-v1-CheckDocumentRequest)
    - [CreateDocumentItem](#api-v1-CreateDocumentItem)
    - [CreateDocumentRequest](#api-v1-CreateDocumentRequest)
    - [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest)
    - [Document](#api-v1-Document)
    - [DocumentScore](#api-v1-DocumentScore)
    - [DownloadDocumentReply](#api-v1-DownloadDocumentReply)
    - [GetDocumentRequest](#api-v1-GetDocumentRequest)
    - [ListDocumentForHomeItem](#api-v1-ListDocumentForHomeItem)
    - [ListDocumentForHomeRequest](#api-v1-ListDocumentForHomeRequest)
    - [ListDocumentForHomeResponse](#api-v1-ListDocumentForHomeResponse)
    - [ListDocumentReply](#api-v1-ListDocumentReply)
    - [ListDocumentRequest](#api-v1-ListDocumentRequest)
    - [RecoverRecycleDocumentRequest](#api-v1-RecoverRecycleDocumentRequest)
    - [SearchDocumentReply](#api-v1-SearchDocumentReply)
    - [SearchDocumentRequest](#api-v1-SearchDocumentRequest)
    - [SetDocumentRecommendRequest](#api-v1-SetDocumentRecommendRequest)
    - [SetDocumentsCategoryRequest](#api-v1-SetDocumentsCategoryRequest)
    - [SetDocumentsLanguageRequest](#api-v1-SetDocumentsLanguageRequest)
  
    - [DocumentAPI](#api-v1-DocumentAPI)
    - [RecycleAPI](#api-v1-RecycleAPI)
  
- [api/v1/article.proto](#api_v1_article-proto)
    - [Article](#api-v1-Article)
    - [CheckArticlesRequest](#api-v1-CheckArticlesRequest)
    - [DeleteArticleRequest](#api-v1-DeleteArticleRequest)
    - [GetArticleRequest](#api-v1-GetArticleRequest)
    - [ListArticleReply](#api-v1-ListArticleReply)
    - [ListArticleRequest](#api-v1-ListArticleRequest)
    - [RecommendArticlesRequest](#api-v1-RecommendArticlesRequest)
    - [RestoreArticleRequest](#api-v1-RestoreArticleRequest)
    - [SearchArticleReply](#api-v1-SearchArticleReply)
    - [SetArticlesCategoryRequest](#api-v1-SetArticlesCategoryRequest)
  
    - [ArticleAPI](#api-v1-ArticleAPI)
  
- [api/v1/searchrecord.proto](#api_v1_searchrecord-proto)
    - [DeleteSearchRecordRequest](#api-v1-DeleteSearchRecordRequest)
    - [ListSearchRecordReply](#api-v1-ListSearchRecordReply)
    - [ListSearchRecordRequest](#api-v1-ListSearchRecordRequest)
    - [SearchRecord](#api-v1-SearchRecord)
  
    - [SearchRecordAPI](#api-v1-SearchRecordAPI)
  
- [api/v1/banner.proto](#api_v1_banner-proto)
    - [Banner](#api-v1-Banner)
    - [DeleteBannerRequest](#api-v1-DeleteBannerRequest)
    - [GetBannerRequest](#api-v1-GetBannerRequest)
    - [ListBannerReply](#api-v1-ListBannerReply)
    - [ListBannerRequest](#api-v1-ListBannerRequest)
  
    - [BannerAPI](#api-v1-BannerAPI)
  
- [api/v1/favorite.proto](#api_v1_favorite-proto)
    - [DeleteFavoriteRequest](#api-v1-DeleteFavoriteRequest)
    - [Favorite](#api-v1-Favorite)
    - [GetFavoriteRequest](#api-v1-GetFavoriteRequest)
    - [ListFavoriteReply](#api-v1-ListFavoriteReply)
    - [ListFavoriteRequest](#api-v1-ListFavoriteRequest)
  
    - [FavoriteAPI](#api-v1-FavoriteAPI)
  
- [api/v1/punishment.proto](#api_v1_punishment-proto)
    - [CancelPunishmentRequest](#api-v1-CancelPunishmentRequest)
    - [CreatePunishmentRequest](#api-v1-CreatePunishmentRequest)
    - [GetPunishmentRequest](#api-v1-GetPunishmentRequest)
    - [ListPunishmentReply](#api-v1-ListPunishmentReply)
    - [ListPunishmentRequest](#api-v1-ListPunishmentRequest)
    - [Punishment](#api-v1-Punishment)
  
    - [PunishmentAPI](#api-v1-PunishmentAPI)
  
- [api/v1/language.proto](#api_v1_language-proto)
    - [DeleteLanguageRequest](#api-v1-DeleteLanguageRequest)
    - [Language](#api-v1-Language)
    - [ListLanguageReply](#api-v1-ListLanguageReply)
    - [ListLanguageRequest](#api-v1-ListLanguageRequest)
    - [UpdateLanguageStatusRequest](#api-v1-UpdateLanguageStatusRequest)
  
    - [LanguageAPI](#api-v1-LanguageAPI)
  
- [api/v1/advertisement.proto](#api_v1_advertisement-proto)
    - [Advertisement](#api-v1-Advertisement)
    - [DeleteAdvertisementRequest](#api-v1-DeleteAdvertisementRequest)
    - [GetAdvertisementByPositionRequest](#api-v1-GetAdvertisementByPositionRequest)
    - [GetAdvertisementRequest](#api-v1-GetAdvertisementRequest)
    - [ListAdvertisementReply](#api-v1-ListAdvertisementReply)
    - [ListAdvertisementRequest](#api-v1-ListAdvertisementRequest)
  
    - [AdvertisementAPI](#api-v1-AdvertisementAPI)
  
- [api/v1/health.proto](#api_v1_health-proto)
    - [PingRequest](#-PingRequest)
    - [PongReply](#-PongReply)
  
    - [HealthAPI](#-HealthAPI)
  
- [api/v1/report.proto](#api_v1_report-proto)
    - [DeleteReportRequest](#api-v1-DeleteReportRequest)
    - [ListReportReply](#api-v1-ListReportReply)
    - [ListReportRequest](#api-v1-ListReportRequest)
    - [Report](#api-v1-Report)
  
    - [ReportAPI](#api-v1-ReportAPI)
  
- [api/v1/friendlink.proto](#api_v1_friendlink-proto)
    - [DeleteFriendlinkRequest](#api-v1-DeleteFriendlinkRequest)
    - [Friendlink](#api-v1-Friendlink)
    - [GetFriendlinkRequest](#api-v1-GetFriendlinkRequest)
    - [ListFriendlinkReply](#api-v1-ListFriendlinkReply)
    - [ListFriendlinkRequest](#api-v1-ListFriendlinkRequest)
  
    - [FriendlinkAPI](#api-v1-FriendlinkAPI)
  
- [api/v1/permission.proto](#api_v1_permission-proto)
    - [GetPermissionReply](#api-v1-GetPermissionReply)
    - [GetPermissionRequest](#api-v1-GetPermissionRequest)
    - [ListPermissionReply](#api-v1-ListPermissionReply)
    - [ListPermissionRequest](#api-v1-ListPermissionRequest)
    - [Permission](#api-v1-Permission)
  
    - [PermissionAPI](#api-v1-PermissionAPI)
  
- [api/v1/navigation.proto](#api_v1_navigation-proto)
    - [DeleteNavigationRequest](#api-v1-DeleteNavigationRequest)
    - [GetNavigationRequest](#api-v1-GetNavigationRequest)
    - [ListNavigationReply](#api-v1-ListNavigationReply)
    - [ListNavigationRequest](#api-v1-ListNavigationRequest)
    - [Navigation](#api-v1-Navigation)
  
    - [NavigationAPI](#api-v1-NavigationAPI)
  
- [api/v1/comment.proto](#api_v1_comment-proto)
    - [CheckCommentRequest](#api-v1-CheckCommentRequest)
    - [Comment](#api-v1-Comment)
    - [CreateCommentRequest](#api-v1-CreateCommentRequest)
    - [DeleteCommentRequest](#api-v1-DeleteCommentRequest)
    - [GetCommentRequest](#api-v1-GetCommentRequest)
    - [ListCommentReply](#api-v1-ListCommentReply)
    - [ListCommentRequest](#api-v1-ListCommentRequest)
  
    - [CommentAPI](#api-v1-CommentAPI)
  
- [api/v1/attachment.proto](#api_v1_attachment-proto)
    - [Attachment](#api-v1-Attachment)
    - [DeleteAttachmentRequest](#api-v1-DeleteAttachmentRequest)
    - [GetAttachmentRequest](#api-v1-GetAttachmentRequest)
    - [ListAttachmentReply](#api-v1-ListAttachmentReply)
    - [ListAttachmentRequest](#api-v1-ListAttachmentRequest)
  
    - [AttachmentAPI](#api-v1-AttachmentAPI)
  
- [api/v1/download.proto](#api_v1_download-proto)
    - [Download](#api-v1-Download)
    - [ListDownloadReply](#api-v1-ListDownloadReply)
    - [ListDownloadRequest](#api-v1-ListDownloadRequest)
  
    - [DownloadAPI](#api-v1-DownloadAPI)
  
- [api/v1/config.proto](#api_v1_config-proto)
    - [CPUInfo](#api-v1-CPUInfo)
    - [Config](#api-v1-Config)
    - [ConfigCaptcha](#api-v1-ConfigCaptcha)
    - [ConfigDisplay](#api-v1-ConfigDisplay)
    - [ConfigFooter](#api-v1-ConfigFooter)
    - [ConfigSecurity](#api-v1-ConfigSecurity)
    - [ConfigSystem](#api-v1-ConfigSystem)
    - [Configs](#api-v1-Configs)
    - [DeviceInfo](#api-v1-DeviceInfo)
    - [DiskInfo](#api-v1-DiskInfo)
    - [EnvDependent](#api-v1-EnvDependent)
    - [Envs](#api-v1-Envs)
    - [ListConfigRequest](#api-v1-ListConfigRequest)
    - [MemoryInfo](#api-v1-MemoryInfo)
    - [Release](#api-v1-Release)
    - [Settings](#api-v1-Settings)
    - [Stats](#api-v1-Stats)
  
    - [ConfigAPI](#api-v1-ConfigAPI)
  
- [api/v1/user.proto](#api_v1_user-proto)
    - [DeleteUserRequest](#api-v1-DeleteUserRequest)
    - [Dynamic](#api-v1-Dynamic)
    - [FindPasswordRequest](#api-v1-FindPasswordRequest)
    - [GetUserCaptchaReply](#api-v1-GetUserCaptchaReply)
    - [GetUserCaptchaRequest](#api-v1-GetUserCaptchaRequest)
    - [GetUserPermissionsReply](#api-v1-GetUserPermissionsReply)
    - [GetUserRequest](#api-v1-GetUserRequest)
    - [ListUserDownloadReply](#api-v1-ListUserDownloadReply)
    - [ListUserDownloadRequest](#api-v1-ListUserDownloadRequest)
    - [ListUserDynamicReply](#api-v1-ListUserDynamicReply)
    - [ListUserDynamicRequest](#api-v1-ListUserDynamicRequest)
    - [ListUserGroupReply](#api-v1-ListUserGroupReply)
    - [ListUserReply](#api-v1-ListUserReply)
    - [ListUserRequest](#api-v1-ListUserRequest)
    - [LoginReply](#api-v1-LoginReply)
    - [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest)
    - [SendEmailCodeRequest](#api-v1-SendEmailCodeRequest)
    - [SetUserRequest](#api-v1-SetUserRequest)
    - [Sign](#api-v1-Sign)
    - [UpdateUserPasswordRequest](#api-v1-UpdateUserPasswordRequest)
    - [User](#api-v1-User)
  
    - [EmailCodeType](#api-v1-EmailCodeType)
  
    - [UserAPI](#api-v1-UserAPI)
  
- [Scalar Value Types](#scalar-value-types)



<a name="api_v1_group-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/group.proto



<a name="api-v1-DeleteGroupRequest"></a>

### DeleteGroupRequest
删除用户组，可以批量删除


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetGroupPermissionRequest"></a>

### GetGroupPermissionRequest
获取用户组权限


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-GetGroupRequest"></a>

### GetGroupRequest
根据组名或者ID获取用户组


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |






<a name="api-v1-Group"></a>

### Group
用户组，角色


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户组ID |
| title | [string](#string) |  | 用户组名称 |
| color | [string](#string) |  | 用户组颜色 |
| is_default | [bool](#bool) |  | 是否是默认用户组 |
| is_display | [bool](#bool) |  | 是否显示 |
| description | [string](#string) |  | 用户组描述 |
| user_count | [int32](#int32) |  | 用户组下的用户数量 |
| sort | [int32](#int32) |  | 排序 |
| enable_upload | [bool](#bool) |  | 是否允许上传文档 |
| enable_comment_approval | [bool](#bool) |  | 是否需要审核评论 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |
| enable_comment | [bool](#bool) |  | 是否允许评论 |
| enable_document_review | [bool](#bool) |  | 文档是否需要审核 |
| enable_article | [bool](#bool) |  | 是否允许发布文章 |
| enable_article_approval | [bool](#bool) |  | 是否需要审核文章 |






<a name="api-v1-GroupPermissions"></a>

### GroupPermissions
用户组权限


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission_id | [int64](#int64) | repeated |  |






<a name="api-v1-ListGroupReply"></a>

### ListGroupReply
用户组列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [Group](#api-v1-Group) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="api-v1-ListGroupRequest"></a>

### ListGroupRequest
查询用户组列表。不需要分页，直接返回全部用户组，只是可以指定查询哪些字段


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wd | [string](#string) |  |  |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| sort | [string](#string) |  |  |
| field | [string](#string) | repeated |  |






<a name="api-v1-UpdateGroupPermissionRequest"></a>

### UpdateGroupPermissionRequest
更新用户组权限


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group_id | [int64](#int64) |  |  |
| permission_id | [int64](#int64) | repeated |  |





 

 

 


<a name="api-v1-GroupAPI"></a>

### GroupAPI
用户组 API 服务，提供用户组和组权限的维护能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateGroup | [Group](#api-v1-Group) | [Group](#api-v1-Group) | 创建用户组：新增一个用户组并返回创建后的用户组信息 |
| UpdateGroup | [Group](#api-v1-Group) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户组：根据用户组 ID 修改名称、权限开关和展示属性 |
| DeleteGroup | [DeleteGroupRequest](#api-v1-DeleteGroupRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除用户组：支持按用户组 ID 批量删除用户组 |
| GetGroup | [GetGroupRequest](#api-v1-GetGroupRequest) | [Group](#api-v1-Group) | 获取用户组：根据用户组 ID 或名称查询单个用户组详情 |
| ListGroup | [ListGroupRequest](#api-v1-ListGroupRequest) | [ListGroupReply](#api-v1-ListGroupReply) | 用户组列表：分页查询用户组并可指定返回字段 |
| GetGroupPermission | [GetGroupPermissionRequest](#api-v1-GetGroupPermissionRequest) | [GroupPermissions](#api-v1-GroupPermissions) | 获取用户组权限：返回指定用户组绑定的权限 ID 列表 |
| UpdateGroupPermission | [UpdateGroupPermissionRequest](#api-v1-UpdateGroupPermissionRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户组权限：为指定用户组批量设置权限集合 |

 



<a name="api_v1_category-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/category.proto



<a name="api-v1-Category"></a>

### Category
文档分类


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 分类ID |
| parent_id | [int64](#int64) |  | 父分类ID |
| title | [string](#string) |  | 分类标题 |
| doc_count | [int32](#int32) |  | 文档数量 |
| sort | [int32](#int32) |  | 排序，倒序排序，值越大越靠前 |
| enable | [bool](#bool) |  | 是否启用 |
| cover | [string](#string) |  | 分类封面 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |
| icon | [string](#string) |  | 分类图标 |
| description | [string](#string) |  | 分类描述 |
| show_description | [bool](#bool) |  | 是否在分类筛选文档列表处显示分类描述 |
| type | [int32](#int32) |  | 分类类型 |






<a name="api-v1-DeleteCategoryRequest"></a>

### DeleteCategoryRequest
删除分类请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetCategoryRequest"></a>

### GetCategoryRequest
获取分类请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListCategoryReply"></a>

### ListCategoryReply
分类列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| category | [Category](#api-v1-Category) | repeated | 分类列表 |






<a name="api-v1-ListCategoryRequest"></a>

### ListCategoryRequest
分类列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| parent_id | [int64](#int64) | repeated | 父分类ID |
| wd | [string](#string) |  | 搜索关键字 |
| enable | [bool](#bool) | repeated | 是否启用 |
| field | [string](#string) | repeated | 查询字段 |
| type | [int32](#int32) | repeated | 分类类型 |





 

 

 


<a name="api-v1-CategoryAPI"></a>

### CategoryAPI
分类 API 服务，提供文档分类的维护与列表查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建分类：新增一个文档分类节点并保存分类属性 |
| UpdateCategory | [Category](#api-v1-Category) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新分类：根据分类 ID 修改名称、封面、排序和展示配置 |
| DeleteCategory | [DeleteCategoryRequest](#api-v1-DeleteCategoryRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除分类：支持按分类 ID 批量删除分类节点 |
| GetCategory | [GetCategoryRequest](#api-v1-GetCategoryRequest) | [Category](#api-v1-Category) | 获取分类：根据分类 ID 查询单个分类详情 |
| ListCategory | [ListCategoryRequest](#api-v1-ListCategoryRequest) | [ListCategoryReply](#api-v1-ListCategoryReply) | 分类列表：按父分类、启用状态和类型分页查询分类 |

 



<a name="api_v1_document-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/document.proto



<a name="api-v1-CheckDocumentRequest"></a>

### CheckDocumentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated | 文档ID |
| status | [int32](#int32) |  | 文档状态 |






<a name="api-v1-CreateDocumentItem"></a>

### CreateDocumentItem
创建文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| title | [string](#string) |  | 文档标题 |
| attachment_id | [int64](#int64) |  | 文档附件ID |
| price | [int32](#int32) |  | 文档价格 |
| language | [string](#string) |  | 文档语言 |
| keywords | [string](#string) |  | 文档关键字 |
| description | [string](#string) |  | 文档描述 |
| source | [string](#string) |  | 来源 |
| source_url | [string](#string) |  | 来源链接 |






<a name="api-v1-CreateDocumentRequest"></a>

### CreateDocumentRequest
创建文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| overwrite | [bool](#bool) |  | 是否覆盖。暂时用不到 |
| category_id | [int64](#int64) | repeated | 文档分类ID |
| document | [CreateDocumentItem](#api-v1-CreateDocumentItem) | repeated | 文档列表 |






<a name="api-v1-DeleteDocumentRequest"></a>

### DeleteDocumentRequest
删除文档，放入回收站


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Document"></a>

### Document
文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文档ID |
| title | [string](#string) |  | 文档标题 |
| keywords | [string](#string) |  | 文档关键字 |
| description | [string](#string) |  | 文档描述 |
| user_id | [int64](#int64) |  | 文档作者 |
| cover | [string](#string) |  | 文档封面 |
| width | [int32](#int32) |  | 文档宽度 |
| height | [int32](#int32) |  | 文档高度 |
| preview | [int32](#int32) |  | 文档可预览页数，0表示不限制 |
| pages | [int32](#int32) |  | 文档页数 |
| uuid | [string](#string) |  | 文档UUID |
| download_count | [int32](#int32) |  | 文档下载次数 |
| view_count | [int32](#int32) |  | 文档浏览次数 |
| favorite_count | [int32](#int32) |  | 文档收藏次数 |
| comment_count | [int32](#int32) |  | 文档评论次数 |
| score | [int32](#int32) |  | 文档评分 |
| score_count | [int32](#int32) |  | 文档评分次数 |
| price | [int32](#int32) |  | 文档价格 |
| size | [int64](#int64) |  | 文档大小 |
| status | [int32](#int32) |  | 文档状态，见 web/utils/enum.js |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档更新时间 |
| deleted_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档删除时间 |
| recommend_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文档推荐时间 |
| deleted_user_id | [int64](#int64) |  | 删除文档的用户 |
| username | [string](#string) |  | 文档作者用户名 |
| category_id | [int64](#int64) | repeated | 文档分类ID |
| deleted_username | [string](#string) |  | 删除文档的用户名 |
| ext | [string](#string) |  | 文档扩展名 |
| attachment | [Attachment](#api-v1-Attachment) |  | 文档附件 |
| user | [User](#api-v1-User) |  | 文档作者 |
| enable_gzip | [bool](#bool) |  | 是否启用gzip压缩 |
| convert_error | [string](#string) |  | 转换错误信息 |
| preview_ext | [string](#string) |  | 预览文件扩展名 |
| language | [string](#string) |  | 文档语言 |
| content | [string](#string) |  | 文档内容 |
| category | [Category](#api-v1-Category) | repeated | 文档分类 |
| source | [string](#string) |  | 来源 |
| source_url | [string](#string) |  | 来源链接 |






<a name="api-v1-DocumentScore"></a>

### DocumentScore
文档评分


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 评分ID |
| document_id | [int64](#int64) |  | 文档ID |
| user_id | [int64](#int64) |  | 用户ID |
| score | [int32](#int32) |  | 评分，100~500，100为1分，500为5分 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 评分时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DownloadDocumentReply"></a>

### DownloadDocumentReply
文档下载


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| url | [string](#string) |  |  |






<a name="api-v1-GetDocumentRequest"></a>

### GetDocumentRequest
查询文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文档ID |
| with_author | [bool](#bool) |  | 是否查询作者信息 |
| uuid | [string](#string) |  | 文档UUID |
| with_all_content | [bool](#bool) |  | 是否查询所有文本内容 |






<a name="api-v1-ListDocumentForHomeItem"></a>

### ListDocumentForHomeItem
首页文档查询返回项


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category_id | [int64](#int64) |  | 分类ID |
| category_cover | [string](#string) |  | 分类封面 |
| category_name | [string](#string) |  | 分类名称 |
| document | [Document](#api-v1-Document) | repeated | 文档列表 |






<a name="api-v1-ListDocumentForHomeRequest"></a>

### ListDocumentForHomeRequest
查询文档（针对首页的查询）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| limit | [int64](#int64) |  |  |
| field | [string](#string) | repeated |  |






<a name="api-v1-ListDocumentForHomeResponse"></a>

### ListDocumentForHomeResponse
查询文档（针对首页的查询）


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document | [ListDocumentForHomeItem](#api-v1-ListDocumentForHomeItem) | repeated | 文档列表 |






<a name="api-v1-ListDocumentReply"></a>

### ListDocumentReply
文档列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文档总数 |
| document | [Document](#api-v1-Document) | repeated | 文档列表 |






<a name="api-v1-ListDocumentRequest"></a>

### ListDocumentRequest
文档列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段 |
| order | [string](#string) |  | 排序 |
| category_id | [int64](#int64) | repeated | 分类ID |
| user_id | [int64](#int64) | repeated | 用户ID |
| status | [int32](#int32) | repeated | 文档状态 |
| is_recommend | [bool](#bool) | repeated | 是否推荐 |
| limit | [int64](#int64) |  | 查询数量显示。当该值大于0时，page和size无效 |
| ext | [string](#string) |  | 文档扩展名 |
| created_at | [string](#string) | repeated | 创建时间 |
| fee_type | [string](#string) |  | 收费类型 |
| language | [string](#string) | repeated | 文档语言 |






<a name="api-v1-RecoverRecycleDocumentRequest"></a>

### RecoverRecycleDocumentRequest
恢复文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-SearchDocumentReply"></a>

### SearchDocumentReply
文档搜索响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文档总数 |
| spend | [string](#string) |  | 搜索耗时 |
| document | [Document](#api-v1-Document) | repeated | 文档列表 |






<a name="api-v1-SearchDocumentRequest"></a>

### SearchDocumentRequest
文档搜索


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  | 页码 |
| size | [int32](#int32) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| category_id | [int64](#int64) | repeated | 分类 |
| sort | [string](#string) |  | 排序 |
| ext | [string](#string) |  | 类型 |
| user_id | [int64](#int64) |  | 用户ID |
| created_at | [string](#string) | repeated | 时间区间

时间范围 |
| language | [string](#string) | repeated | 文档语言 |






<a name="api-v1-SetDocumentRecommendRequest"></a>

### SetDocumentRecommendRequest
设置文档推荐


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated | 文档ID |
| type | [int32](#int32) |  | 0, 取消推荐，1:推荐 2:重新推荐 |






<a name="api-v1-SetDocumentsCategoryRequest"></a>

### SetDocumentsCategoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) | repeated | 文档ID |
| category_id | [int64](#int64) | repeated | 分类ID |






<a name="api-v1-SetDocumentsLanguageRequest"></a>

### SetDocumentsLanguageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) | repeated | 文档ID |
| language | [string](#string) |  | 文档语言代码 |





 

 

 


<a name="api-v1-DocumentAPI"></a>

### DocumentAPI
文档 API 服务，提供文档创建、查询、检索、评分和推荐管理能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListDocumentForHome | [ListDocumentForHomeRequest](#api-v1-ListDocumentForHomeRequest) | [ListDocumentForHomeResponse](#api-v1-ListDocumentForHomeResponse) | 首页文档推荐：返回首页场景下按分类聚合的文档列表 |
| SetDocumentRecommend | [SetDocumentRecommendRequest](#api-v1-SetDocumentRecommendRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置文档推荐：批量调整文档推荐状态或刷新推荐时间 |
| CreateDocument | [CreateDocumentRequest](#api-v1-CreateDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建文档：批量提交附件生成文档记录并写入基础信息 |
| UpdateDocument | [Document](#api-v1-Document) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新文档：根据文档 ID 修改标题、描述、价格和分类等信息 |
| DeleteDocument | [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除文档：支持按文档 ID 批量删除文档到回收站 |
| GetDocument | [GetDocumentRequest](#api-v1-GetDocumentRequest) | [Document](#api-v1-Document) | 获取文档：根据文档 ID 或 UUID 查询文档详情 |
| GetRelatedDocuments | [Document](#api-v1-Document) | [ListDocumentReply](#api-v1-ListDocumentReply) | 相关文档：根据当前文档推荐相似或相关的文档列表 |
| DownloadDocument | [Document](#api-v1-Document) | [DownloadDocumentReply](#api-v1-DownloadDocumentReply) | 获取下载链接：根据文档信息生成当前用户可用的下载地址 |
| CheckDocument | [CheckDocumentRequest](#api-v1-CheckDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 审核文档：批量更新文档审核状态 |
| ListDocument | [ListDocumentRequest](#api-v1-ListDocumentRequest) | [ListDocumentReply](#api-v1-ListDocumentReply) | 文档列表：按关键词、分类、作者、状态和语言分页查询文档 |
| SearchDocument | [SearchDocumentRequest](#api-v1-SearchDocumentRequest) | [SearchDocumentReply](#api-v1-SearchDocumentReply) | 搜索文档：根据全文检索条件返回匹配文档和耗时信息 |
| SetDocumentScore | [DocumentScore](#api-v1-DocumentScore) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置文档评分：提交当前用户对文档的评分结果 |
| GetDocumentScore | [DocumentScore](#api-v1-DocumentScore) | [DocumentScore](#api-v1-DocumentScore) | 获取我的文档评分：查询当前登录用户对指定文档的评分记录 |
| SetDocumentReconvert | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 触发文档重转：将待处理文档标记为重新转换预览文件 |
| SetDocumentsCategory | [SetDocumentsCategoryRequest](#api-v1-SetDocumentsCategoryRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 批量设置文档分类：为多篇文档统一更新所属分类 |
| SetDocumentsLanguage | [SetDocumentsLanguageRequest](#api-v1-SetDocumentsLanguageRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 批量设置文档语言：为多篇文档统一更新语言标签 |
| DownloadDocumentToBeReviewed | [Document](#api-v1-Document) | [DownloadDocumentReply](#api-v1-DownloadDocumentReply) | 下载待审核文档：为待审核或审核拒绝的文档生成下载地址 |


<a name="api-v1-RecycleAPI"></a>

### RecycleAPI
回收站 API 服务，提供已删除文档的查询、恢复和清理能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| ListRecycleDocument | [ListDocumentRequest](#api-v1-ListDocumentRequest) | [ListDocumentReply](#api-v1-ListDocumentReply) | 回收站列表：分页查询已进入回收站的文档 |
| RecoverRecycleDocument | [RecoverRecycleDocumentRequest](#api-v1-RecoverRecycleDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 恢复回收站文档：支持单个或批量恢复已删除文档 |
| DeleteRecycleDocument | [DeleteDocumentRequest](#api-v1-DeleteDocumentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 彻底删除回收站文档：按文档 ID 从回收站永久移除 |
| ClearRecycleDocument | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 清空回收站：永久删除回收站中的全部文档记录 |

 



<a name="api_v1_article-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/article.proto



<a name="api-v1-Article"></a>

### Article
文章


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文章ID |
| identifier | [string](#string) |  | 文章唯一标识 |
| author | [string](#string) |  | 文章作者。如果为空，则使用网站名称作为作者 |
| view_count | [int64](#int64) |  | 文章浏览次数 |
| title | [string](#string) |  | 文章标题 |
| keywords | [string](#string) |  | 文章关键字 |
| description | [string](#string) |  | 文章描述 |
| content | [string](#string) |  | 文章内容 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章更新时间 |
| deleted_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章删除时间 |
| category_id | [int64](#int64) | repeated | 文章分类ID |
| favorite_count | [int64](#int64) |  | 文章收藏次数 |
| comment_count | [int64](#int64) |  | 文章评论次数 |
| recommend_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 文章推荐时间 |
| is_recommend | [bool](#bool) |  | 是否推荐 |
| user_id | [int64](#int64) |  | 文章作者ID |
| user | [User](#api-v1-User) |  | 文章作者 |
| status | [int32](#int32) |  | 文章状态 |
| reject_reason | [string](#string) |  | 文章拒绝原因 |
| category | [Category](#api-v1-Category) | repeated | 文章分类 |
| source | [string](#string) |  | 文章来源 |
| source_url | [string](#string) |  | 文章来源链接 |






<a name="api-v1-CheckArticlesRequest"></a>

### CheckArticlesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| article_id | [int64](#int64) | repeated | 文章ID |
| status | [int32](#int32) |  | 文章状态 |
| rejeact_reason | [string](#string) |  | 拒绝原因 |






<a name="api-v1-DeleteArticleRequest"></a>

### DeleteArticleRequest
删除文章请求，传入单个或者多个文章ID


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetArticleRequest"></a>

### GetArticleRequest
根据ID或者文章标识获取文章，二选一


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 文章ID |
| identifier | [string](#string) |  | 文章唯一标识 |






<a name="api-v1-ListArticleReply"></a>

### ListArticleReply
文章列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文章总数 |
| article | [Article](#api-v1-Article) | repeated | 文章列表 |






<a name="api-v1-ListArticleRequest"></a>

### ListArticleRequest
文章列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段 |
| order | [string](#string) |  | 排序字段，根据指定的字段倒序排序 |
| category_id | [int64](#int64) | repeated | 分类ID |
| is_recommend | [bool](#bool) | repeated | 是否推荐 |
| status | [int32](#int32) | repeated | 文章状态 |
| user_id | [int64](#int64) | repeated | 用户ID |
| created_at | [string](#string) | repeated | 时间范围 |
| sort | [string](#string) |  | 排序：latest-最新 |






<a name="api-v1-RecommendArticlesRequest"></a>

### RecommendArticlesRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| article_id | [int64](#int64) | repeated | 文章ID |
| is_recommend | [bool](#bool) |  | 是否推荐 |






<a name="api-v1-RestoreArticleRequest"></a>

### RestoreArticleRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-SearchArticleReply"></a>

### SearchArticleReply
文档搜索响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 文档总数 |
| spend | [string](#string) |  | 搜索耗时 |
| article | [Article](#api-v1-Article) | repeated | 文档列表 |






<a name="api-v1-SetArticlesCategoryRequest"></a>

### SetArticlesCategoryRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| article_id | [int64](#int64) | repeated | 文章ID |
| category_id | [int64](#int64) | repeated | 分类ID |





 

 

 


<a name="api-v1-ArticleAPI"></a>

### ArticleAPI
文章 API 服务，提供文章的创建、审核、检索与回收站管理能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateArticle | [Article](#api-v1-Article) | [Article](#api-v1-Article) | 创建文章：提交文章标题、内容和分类信息，返回新建后的文章详情 |
| UpdateArticle | [Article](#api-v1-Article) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新文章：根据文章 ID 修改文章内容、状态和展示信息 |
| DeleteArticle | [DeleteArticleRequest](#api-v1-DeleteArticleRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除文章：支持按文章 ID 批量删除文章到回收站 |
| GetArticle | [GetArticleRequest](#api-v1-GetArticleRequest) | [Article](#api-v1-Article) | 获取文章：根据文章 ID 或唯一标识查询单篇文章详情 |
| ListArticle | [ListArticleRequest](#api-v1-ListArticleRequest) | [ListArticleReply](#api-v1-ListArticleReply) | 文章列表：按关键词、分类、作者和状态分页查询文章 |
| SetArticlesCategory | [SetArticlesCategoryRequest](#api-v1-SetArticlesCategoryRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 批量设置文章分类：为多篇文章统一更新所属分类 |
| RecommendArticles | [RecommendArticlesRequest](#api-v1-RecommendArticlesRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 批量推荐文章：统一设置文章推荐状态或重新推荐 |
| CheckArticles | [CheckArticlesRequest](#api-v1-CheckArticlesRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 批量审核文章：统一更新文章审核状态并记录拒绝原因 |
| ListRecycleArticle | [ListArticleRequest](#api-v1-ListArticleRequest) | [ListArticleReply](#api-v1-ListArticleReply) | 回收站文章列表：分页查询已删除的文章记录 |
| RestoreRecycleArticle | [RestoreArticleRequest](#api-v1-RestoreArticleRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 恢复回收站文章：支持按文章 ID 批量恢复已删除文章 |
| DeleteRecycleArticle | [DeleteArticleRequest](#api-v1-DeleteArticleRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 彻底删除回收站文章：按文章 ID 从回收站永久删除文章 |
| EmptyRecycleArticle | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 清空文章回收站：永久删除回收站中的全部文章 |
| SearchArticle | [ListArticleRequest](#api-v1-ListArticleRequest) | [SearchArticleReply](#api-v1-SearchArticleReply) | 搜索文章：根据关键词和筛选条件执行全文检索 |
| GetRelatedArticles | [GetArticleRequest](#api-v1-GetArticleRequest) | [ListArticleReply](#api-v1-ListArticleReply) | 相关文章：根据当前文章返回推荐的相关文章列表 |

 



<a name="api_v1_searchrecord-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/searchrecord.proto



<a name="api-v1-DeleteSearchRecordRequest"></a>

### DeleteSearchRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-ListSearchRecordReply"></a>

### ListSearchRecordReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| search_record | [SearchRecord](#api-v1-SearchRecord) | repeated |  |






<a name="api-v1-ListSearchRecordRequest"></a>

### ListSearchRecordRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| keywords | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |
| user_id | [int64](#int64) | repeated |  |
| ip | [string](#string) |  |  |






<a name="api-v1-SearchRecord"></a>

### SearchRecord
这里是proto文件中的结构体，可以根据需要删除或者调整


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| ip | [string](#string) |  |  |
| total | [int32](#int32) |  |  |
| page | [int32](#int32) |  |  |
| user_agent | [string](#string) |  |  |
| keywords | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| spend_time | [float](#float) |  |  |
| username | [string](#string) |  |  |





 

 

 


<a name="api-v1-SearchRecordAPI"></a>

### SearchRecordAPI
搜索记录 API 服务，提供搜索日志的查询和清理能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| DeleteSearchRecord | [DeleteSearchRecordRequest](#api-v1-DeleteSearchRecordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除搜索记录：支持按记录 ID 批量删除搜索日志 |
| ListSearchRecord | [ListSearchRecordRequest](#api-v1-ListSearchRecordRequest) | [ListSearchRecordReply](#api-v1-ListSearchRecordReply) | 搜索记录列表：按关键词、用户和 IP 分页查询搜索日志 |

 



<a name="api_v1_banner-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/banner.proto



<a name="api-v1-Banner"></a>

### Banner
banner，轮播图


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 主键 |
| title | [string](#string) |  | 标题，名称 |
| path | [string](#string) |  | 图片地址 |
| sort | [int32](#int32) |  | 排序，值越大越靠前 |
| enable | [bool](#bool) |  | 是否启用 |
| type | [int32](#int32) |  | 类型，如PC轮播图、小程序轮播图等，见 web/utils/enum.js 中的枚举 |
| url | [string](#string) |  | 跳转地址 |
| description | [string](#string) |  | 描述 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DeleteBannerRequest"></a>

### DeleteBannerRequest
删除轮播图


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetBannerRequest"></a>

### GetBannerRequest
获取轮播图


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListBannerReply"></a>

### ListBannerReply
轮播图列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| banner | [Banner](#api-v1-Banner) | repeated | 轮播图数组 |






<a name="api-v1-ListBannerRequest"></a>

### ListBannerRequest
轮播图列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| type | [int32](#int32) | repeated | 类型 |
| enable | [bool](#bool) | repeated | 是否启用 |
| wd | [string](#string) |  | 搜索关键字 |
| field | [string](#string) | repeated | 查询字段，不指定，则查询全部 |





 

 

 


<a name="api-v1-BannerAPI"></a>

### BannerAPI
轮播图 API 服务，提供轮播图内容的新增、编辑、删除和查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateBanner | [Banner](#api-v1-Banner) | [Banner](#api-v1-Banner) | 创建轮播图：新增一条轮播图并返回创建后的数据 |
| UpdateBanner | [Banner](#api-v1-Banner) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新轮播图：根据轮播图 ID 修改标题、图片和展示状态 |
| DeleteBanner | [DeleteBannerRequest](#api-v1-DeleteBannerRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除轮播图：支持按轮播图 ID 批量删除配置 |
| GetBanner | [GetBannerRequest](#api-v1-GetBannerRequest) | [Banner](#api-v1-Banner) | 查询轮播图：根据轮播图 ID 获取单条详情 |
| ListBanner | [ListBannerRequest](#api-v1-ListBannerRequest) | [ListBannerReply](#api-v1-ListBannerReply) | 轮播图列表：按类型、启用状态和关键词分页查询轮播图 |

 



<a name="api_v1_favorite-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/favorite.proto



<a name="api-v1-DeleteFavoriteRequest"></a>

### DeleteFavoriteRequest
取消收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Favorite"></a>

### Favorite
文档收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| document_id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| ext | [string](#string) |  |  |
| score | [int32](#int32) |  |  |
| size | [int64](#int64) |  |  |
| pages | [int32](#int32) |  |  |
| document_uuid | [string](#string) |  |  |
| type | [int32](#int32) |  | 0: 文档 1: 文章 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-GetFavoriteRequest"></a>

### GetFavoriteRequest
根据文章id，查询用户是否有收藏某篇文档


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) |  |  |
| type | [int32](#int32) |  |  |






<a name="api-v1-ListFavoriteReply"></a>

### ListFavoriteReply
查询用户的收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| favorite | [Favorite](#api-v1-Favorite) | repeated |  |






<a name="api-v1-ListFavoriteRequest"></a>

### ListFavoriteRequest
查询用户的收藏


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| type | [int32](#int32) |  | 0: 文档 1: 文章 |





 

 

 


<a name="api-v1-FavoriteAPI"></a>

### FavoriteAPI
收藏 API 服务，提供收藏的创建、取消和列表查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateFavorite | [Favorite](#api-v1-Favorite) | [Favorite](#api-v1-Favorite) | 添加收藏：将指定文档或文章加入当前用户收藏夹 |
| DeleteFavorite | [DeleteFavoriteRequest](#api-v1-DeleteFavoriteRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 取消收藏：支持按收藏记录 ID 批量移除收藏 |
| GetFavorite | [GetFavoriteRequest](#api-v1-GetFavoriteRequest) | [Favorite](#api-v1-Favorite) | 查询收藏状态：根据文档或文章 ID 判断当前用户是否已收藏 |
| ListFavorite | [ListFavoriteRequest](#api-v1-ListFavoriteRequest) | [ListFavoriteReply](#api-v1-ListFavoriteReply) | 收藏列表：分页查询当前用户的文档或文章收藏记录 |

 



<a name="api_v1_punishment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/punishment.proto



<a name="api-v1-CancelPunishmentRequest"></a>

### CancelPunishmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-CreatePunishmentRequest"></a>

### CreatePunishmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) | repeated |  |
| type | [int32](#int32) | repeated |  |
| enable | [bool](#bool) |  |  |
| remark | [string](#string) |  |  |
| reason | [string](#string) |  |  |






<a name="api-v1-GetPunishmentRequest"></a>

### GetPunishmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListPunishmentReply"></a>

### ListPunishmentReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| punishment | [Punishment](#api-v1-Punishment) | repeated |  |






<a name="api-v1-ListPunishmentRequest"></a>

### ListPunishmentRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |
| type | [int32](#int32) | repeated |  |
| enable | [int32](#int32) | repeated |  |
| user_id | [int64](#int64) | repeated |  |






<a name="api-v1-Punishment"></a>

### Punishment
这里是proto文件中的结构体，可以根据需要删除或者调整


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| type | [int32](#int32) |  |  |
| enable | [bool](#bool) |  |  |
| operators | [string](#string) |  |  |
| remark | [string](#string) |  |  |
| reason | [string](#string) |  |  |
| username | [string](#string) |  |  |





 

 

 


<a name="api-v1-PunishmentAPI"></a>

### PunishmentAPI
惩罚 API 服务，提供处罚记录的创建、查询、更新和解除能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreatePunishment | [CreatePunishmentRequest](#api-v1-CreatePunishmentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建惩罚：批量为用户设置禁言、限制等处罚措施 |
| UpdatePunishment | [Punishment](#api-v1-Punishment) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新惩罚：修改处罚记录的状态、原因或截止时间 |
| GetPunishment | [GetPunishmentRequest](#api-v1-GetPunishmentRequest) | [Punishment](#api-v1-Punishment) | 获取惩罚：根据处罚记录 ID 查询单条处罚详情 |
| ListPunishment | [ListPunishmentRequest](#api-v1-ListPunishmentRequest) | [ListPunishmentReply](#api-v1-ListPunishmentReply) | 惩罚列表：按用户、处罚类型和状态分页查询处罚记录 |
| CancelPunishment | [CancelPunishmentRequest](#api-v1-CancelPunishmentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 取消惩罚：按处罚记录 ID 批量解除已生效的处罚 |

 



<a name="api_v1_language-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/language.proto



<a name="api-v1-DeleteLanguageRequest"></a>

### DeleteLanguageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Language"></a>

### Language



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| language | [string](#string) |  |  |
| enable | [bool](#bool) |  |  |
| code | [string](#string) |  |  |
| total | [int32](#int32) |  |  |
| sort | [int32](#int32) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |






<a name="api-v1-ListLanguageReply"></a>

### ListLanguageReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| language | [Language](#api-v1-Language) | repeated |  |






<a name="api-v1-ListLanguageRequest"></a>

### ListLanguageRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| wd | [string](#string) |  |  |
| enable | [bool](#bool) | repeated |  |
| field | [string](#string) | repeated |  |
| page | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |






<a name="api-v1-UpdateLanguageStatusRequest"></a>

### UpdateLanguageStatusRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |
| enable | [bool](#bool) |  |  |





 

 

 


<a name="api-v1-LanguageAPI"></a>

### LanguageAPI
语言 API 服务，提供站点语言配置的维护和列表查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateLanguage | [Language](#api-v1-Language) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建语言：新增一种站点支持的语言配置 |
| UpdateLanguage | [Language](#api-v1-Language) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新语言：根据语言 ID 修改名称、编码和排序信息 |
| UpdateLanguageStatus | [UpdateLanguageStatusRequest](#api-v1-UpdateLanguageStatusRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 批量更新语言状态：统一启用或停用指定语言项 |
| ListLanguage | [ListLanguageRequest](#api-v1-ListLanguageRequest) | [ListLanguageReply](#api-v1-ListLanguageReply) | 语言列表：按关键词和启用状态分页查询语言配置 |
| DeleteLanguage | [DeleteLanguageRequest](#api-v1-DeleteLanguageRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除语言：支持按语言 ID 批量删除语言配置 |

 



<a name="api_v1_advertisement-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/advertisement.proto



<a name="api-v1-Advertisement"></a>

### Advertisement



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| start_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| end_time | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| position | [string](#string) |  |  |
| content | [string](#string) |  |  |
| enable | [bool](#bool) |  |  |
| remark | [string](#string) |  |  |
| title | [string](#string) |  |  |






<a name="api-v1-DeleteAdvertisementRequest"></a>

### DeleteAdvertisementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetAdvertisementByPositionRequest"></a>

### GetAdvertisementByPositionRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| position | [string](#string) | repeated |  |






<a name="api-v1-GetAdvertisementRequest"></a>

### GetAdvertisementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListAdvertisementReply"></a>

### ListAdvertisementReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| advertisement | [Advertisement](#api-v1-Advertisement) | repeated |  |






<a name="api-v1-ListAdvertisementRequest"></a>

### ListAdvertisementRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |
| position | [string](#string) | repeated |  |
| enable | [bool](#bool) | repeated |  |





 

 

 


<a name="api-v1-AdvertisementAPI"></a>

### AdvertisementAPI
广告位 API 服务，提供广告配置的新增、更新、删除和查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateAdvertisement | [Advertisement](#api-v1-Advertisement) | [Advertisement](#api-v1-Advertisement) | 创建广告：新增一条广告位配置并返回创建后的广告详情 |
| UpdateAdvertisement | [Advertisement](#api-v1-Advertisement) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新广告：根据广告 ID 修改广告位内容、投放时间和启用状态 |
| DeleteAdvertisement | [DeleteAdvertisementRequest](#api-v1-DeleteAdvertisementRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除广告：支持按广告 ID 批量删除广告位配置 |
| GetAdvertisement | [GetAdvertisementRequest](#api-v1-GetAdvertisementRequest) | [Advertisement](#api-v1-Advertisement) | 查询广告：根据广告 ID 获取单条广告位详情 |
| GetAdvertisementByPosition | [GetAdvertisementByPositionRequest](#api-v1-GetAdvertisementByPositionRequest) | [ListAdvertisementReply](#api-v1-ListAdvertisementReply) | 按广告位查询广告：根据广告位标识返回对应的广告列表 |
| ListAdvertisement | [ListAdvertisementRequest](#api-v1-ListAdvertisementRequest) | [ListAdvertisementReply](#api-v1-ListAdvertisementReply) | 广告列表：按关键词、位置和启用状态筛选广告位配置 |

 



<a name="api_v1_health-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/health.proto



<a name="-PingRequest"></a>

### PingRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |






<a name="-PongReply"></a>

### PongReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |





 

 

 


<a name="-HealthAPI"></a>

### HealthAPI
健康检查 API 服务，提供服务存活探测和连通性测试能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Health | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 服务健康检查：用于探测服务是否正常存活并可响应请求 |
| Ping | [.PingRequest](#PingRequest) | [.PongReply](#PongReply) | Ping 测试：回显请求信息并返回服务当前响应时间 |

 



<a name="api_v1_report-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/report.proto



<a name="api-v1-DeleteReportRequest"></a>

### DeleteReportRequest
删除举报请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-ListReportReply"></a>

### ListReportReply
举报列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| report | [Report](#api-v1-Report) | repeated |  |






<a name="api-v1-ListReportRequest"></a>

### ListReportRequest
举报列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |
| status | [bool](#bool) | repeated |  |






<a name="api-v1-Report"></a>

### Report
举报


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 举报ID |
| document_id | [int64](#int64) |  | 文档ID |
| user_id | [int64](#int64) |  | 举报人ID |
| reason | [int32](#int32) |  | 举报原因 |
| status | [bool](#bool) |  | 举报处理状态 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 举报时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 处理时间 |
| document_title | [string](#string) |  | 文档标题 |
| remark | [string](#string) |  | 处理备注 |
| username | [string](#string) |  | 举报人 |





 

 

 


<a name="api-v1-ReportAPI"></a>

### ReportAPI
举报 API 服务，提供举报提交、处理和列表查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateReport | [Report](#api-v1-Report) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建举报：提交对文档内容的举报信息 |
| UpdateReport | [Report](#api-v1-Report) | [.google.protobuf.Empty](#google-protobuf-Empty) | 处理举报：更新举报状态并记录处理备注 |
| DeleteReport | [DeleteReportRequest](#api-v1-DeleteReportRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除举报：支持按举报 ID 批量删除举报记录 |
| ListReport | [ListReportRequest](#api-v1-ListReportRequest) | [ListReportReply](#api-v1-ListReportReply) | 举报列表：按处理状态和关键词分页查询举报记录 |

 



<a name="api_v1_friendlink-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/friendlink.proto



<a name="api-v1-DeleteFriendlinkRequest"></a>

### DeleteFriendlinkRequest
删除友情链接


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-Friendlink"></a>

### Friendlink
友情链接


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int32](#int32) |  | 主键 |
| title | [string](#string) |  | 标题 |
| link | [string](#string) |  | 链接 |
| description | [string](#string) |  | 描述 |
| sort | [int32](#int32) |  | 排序 |
| enable | [bool](#bool) |  | 是否启用 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-GetFriendlinkRequest"></a>

### GetFriendlinkRequest
获取友情链接


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListFriendlinkReply"></a>

### ListFriendlinkReply
友情链接列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| friendlink | [Friendlink](#api-v1-Friendlink) | repeated |  |
| total | [int64](#int64) |  |  |






<a name="api-v1-ListFriendlinkRequest"></a>

### ListFriendlinkRequest
友情链接列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int32](#int32) |  |  |
| size | [int32](#int32) |  |  |
| wd | [string](#string) |  |  |
| enable | [bool](#bool) | repeated |  |
| field | [string](#string) | repeated |  |





 

 

 


<a name="api-v1-FriendlinkAPI"></a>

### FriendlinkAPI
友情链接 API 服务，提供友情链接的新增、编辑、删除和查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateFriendlink | [Friendlink](#api-v1-Friendlink) | [Friendlink](#api-v1-Friendlink) | 创建友情链接：新增一条友情链接配置并返回创建结果 |
| UpdateFriendlink | [Friendlink](#api-v1-Friendlink) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新友情链接：根据链接 ID 修改名称、地址和展示状态 |
| DeleteFriendlink | [DeleteFriendlinkRequest](#api-v1-DeleteFriendlinkRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除友情链接：支持按链接 ID 批量删除友情链接 |
| GetFriendlink | [GetFriendlinkRequest](#api-v1-GetFriendlinkRequest) | [Friendlink](#api-v1-Friendlink) | 获取友情链接：根据链接 ID 查询单条友情链接详情 |
| ListFriendlink | [ListFriendlinkRequest](#api-v1-ListFriendlinkRequest) | [ListFriendlinkReply](#api-v1-ListFriendlinkReply) | 友情链接列表：按关键词和启用状态分页查询友情链接 |

 



<a name="api_v1_permission-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/permission.proto



<a name="api-v1-GetPermissionReply"></a>

### GetPermissionReply
权限响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission | [Permission](#api-v1-Permission) |  |  |






<a name="api-v1-GetPermissionRequest"></a>

### GetPermissionRequest
权限请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListPermissionReply"></a>

### ListPermissionReply
权限列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| permission | [Permission](#api-v1-Permission) | repeated |  |






<a name="api-v1-ListPermissionRequest"></a>

### ListPermissionRequest
权限列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| method | [string](#string) | repeated |  |
| path | [string](#string) |  |  |






<a name="api-v1-Permission"></a>

### Permission
权限


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 权限ID |
| method | [string](#string) |  | 请求方法 |
| path | [string](#string) |  | 请求路径 |
| title | [string](#string) |  | 权限名称 |
| description | [string](#string) |  | 权限描述 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |





 

 

 


<a name="api-v1-PermissionAPI"></a>

### PermissionAPI
权限 API 服务，提供权限元数据的查询和维护能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdatePermission | [Permission](#api-v1-Permission) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新权限：维护权限名称和描述等展示信息 |
| GetPermission | [GetPermissionRequest](#api-v1-GetPermissionRequest) | [Permission](#api-v1-Permission) | 获取权限：根据权限 ID 查询单个权限详情 |
| ListPermission | [ListPermissionRequest](#api-v1-ListPermissionRequest) | [ListPermissionReply](#api-v1-ListPermissionReply) | 权限列表：按请求方法、路径和关键词分页查询权限 |

 



<a name="api_v1_navigation-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/navigation.proto



<a name="api-v1-DeleteNavigationRequest"></a>

### DeleteNavigationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetNavigationRequest"></a>

### GetNavigationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListNavigationReply"></a>

### ListNavigationReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| navigation | [Navigation](#api-v1-Navigation) | repeated |  |






<a name="api-v1-ListNavigationRequest"></a>

### ListNavigationRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |






<a name="api-v1-Navigation"></a>

### Navigation



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| title | [string](#string) |  |  |
| href | [string](#string) |  |  |
| target | [string](#string) |  |  |
| color | [string](#string) |  |  |
| sort | [int32](#int32) |  |  |
| enable | [bool](#bool) |  |  |
| parent_id | [int64](#int64) |  |  |
| description | [string](#string) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| fixed | [bool](#bool) |  |  |





 

 

 


<a name="api-v1-NavigationAPI"></a>

### NavigationAPI
导航 API 服务，提供导航链接的新增、编辑、删除和查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateNavigation | [Navigation](#api-v1-Navigation) | [Navigation](#api-v1-Navigation) | 创建导航：新增一条站点导航并返回创建后的数据 |
| UpdateNavigation | [Navigation](#api-v1-Navigation) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新导航：根据导航 ID 修改名称、链接、排序和固定状态 |
| DeleteNavigation | [DeleteNavigationRequest](#api-v1-DeleteNavigationRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除导航：支持按导航 ID 批量删除导航项 |
| GetNavigation | [GetNavigationRequest](#api-v1-GetNavigationRequest) | [Navigation](#api-v1-Navigation) | 获取导航：根据导航 ID 查询单个导航项详情 |
| ListNavigation | [ListNavigationRequest](#api-v1-ListNavigationRequest) | [ListNavigationReply](#api-v1-ListNavigationReply) | 导航列表：按关键词和排序规则分页查询导航项 |

 



<a name="api_v1_comment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/comment.proto



<a name="api-v1-CheckCommentRequest"></a>

### CheckCommentRequest
审核评论，修改评论状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated | 评论ID |
| status | [int32](#int32) |  | 状态，见 web/utils/enum.js 枚举 |






<a name="api-v1-Comment"></a>

### Comment
评论


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |
| id | [int64](#int64) |  | 评论ID |
| parent_id | [int64](#int64) |  | 父评论ID |
| content | [string](#string) |  | 评论内容 |
| document_id | [int64](#int64) |  | 文档ID |
| status | [int32](#int32) |  | 状态，见 web/utils/enum.js 枚举 |
| comment_count | [int32](#int32) |  | 回复数量 |
| user_id | [int64](#int64) |  | 用户ID |
| user | [User](#api-v1-User) |  | 用户信息 |
| document_title | [string](#string) |  | 文档标题 |
| type | [int32](#int32) |  | 评论类型：0, 文档评论；1 文章评论 |
| document_uuid | [string](#string) |  | 唯一标识 |






<a name="api-v1-CreateCommentRequest"></a>

### CreateCommentRequest
创建评论请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| document_id | [int64](#int64) |  | 文档ID |
| parent_id | [int64](#int64) |  | 父评论ID |
| content | [string](#string) |  | 评论内容 |
| captcha_id | [string](#string) |  | 验证码ID |
| captcha | [string](#string) |  | 验证码 |
| type | [int32](#int32) |  |  |






<a name="api-v1-DeleteCommentRequest"></a>

### DeleteCommentRequest
删除评论请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetCommentRequest"></a>

### GetCommentRequest
获取评论请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListCommentReply"></a>

### ListCommentReply
获取评论列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| comment | [Comment](#api-v1-Comment) | repeated | 评论列表 |






<a name="api-v1-ListCommentRequest"></a>

### ListCommentRequest
获取评论列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键词 |
| field | [string](#string) | repeated | 查询的数据字段 |
| order | [string](#string) |  | 排序字段 |
| status | [int32](#int32) | repeated | 状态，见 web/utils/enum.js 枚举 |
| document_id | [int64](#int64) |  | 文档ID |
| user_id | [int64](#int64) |  | 用户ID |
| parent_id | [int64](#int64) | repeated | 父评论ID |
| with_document_title | [bool](#bool) |  | 是否返回文档标题 |
| type | [int32](#int32) | repeated | 0 文档，1文章 |





 

 

 


<a name="api-v1-CommentAPI"></a>

### CommentAPI
评论 API 服务，提供评论发布、审核、删除和查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| CreateComment | [CreateCommentRequest](#api-v1-CreateCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建评论：提交评论内容并关联到文档或文章 |
| UpdateComment | [Comment](#api-v1-Comment) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新评论：管理员修改评论内容或状态时使用 |
| DeleteComment | [DeleteCommentRequest](#api-v1-DeleteCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除评论：管理员或评论作者可按评论 ID 删除评论 |
| GetComment | [GetCommentRequest](#api-v1-GetCommentRequest) | [Comment](#api-v1-Comment) | 获取评论详情：根据评论 ID 查询单条评论信息 |
| ListComment | [ListCommentRequest](#api-v1-ListCommentRequest) | [ListCommentReply](#api-v1-ListCommentReply) | 评论列表：按状态、文档、用户和评论类型分页查询评论 |
| CheckComment | [CheckCommentRequest](#api-v1-CheckCommentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 审核评论：批量更新评论审核状态 |

 



<a name="api_v1_attachment-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/attachment.proto



<a name="api-v1-Attachment"></a>

### Attachment
附件


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 附件ID |
| hash | [string](#string) |  | 附件哈希值，MD5 |
| user_id | [int64](#int64) |  | 上传用户ID |
| type_id | [int64](#int64) |  | 附件类型ID，如果是文档类型，则为文档ID |
| type | [int32](#int32) |  | 附件类型，见 web/utils/enum.js |
| enable | [bool](#bool) |  | 是否启用 |
| path | [string](#string) |  | 附件路径 |
| name | [string](#string) |  | 附件名称 |
| size | [int64](#int64) |  | 附件大小，单位：字节 |
| width | [int64](#int64) |  | 附件宽度，单位：像素。针对图片附件 |
| height | [int64](#int64) |  | 附件高度，单位：像素。针对图片附件 |
| ext | [string](#string) |  | 扩展名，如：.docx |
| ip | [string](#string) |  | 上传IP地址 |
| username | [string](#string) |  | 用户名称 |
| type_name | [string](#string) |  | 附件类型名称 |
| description | [string](#string) |  | 附件描述、备注 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-DeleteAttachmentRequest"></a>

### DeleteAttachmentRequest
删除附件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |






<a name="api-v1-GetAttachmentRequest"></a>

### GetAttachmentRequest
获取附件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListAttachmentReply"></a>

### ListAttachmentReply
列出附件响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| attachment | [Attachment](#api-v1-Attachment) | repeated |  |






<a name="api-v1-ListAttachmentRequest"></a>

### ListAttachmentRequest
列出附件请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键字 |
| enable | [bool](#bool) | repeated | 是否启用 |
| user_id | [int64](#int64) | repeated | 用户ID |
| type | [int64](#int64) | repeated | 类型 |
| ext | [string](#string) |  | 扩展名 |





 

 

 


<a name="api-v1-AttachmentAPI"></a>

### AttachmentAPI
附件 API 服务，提供附件信息的维护、软删除和后台查询能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| UpdateAttachment | [Attachment](#api-v1-Attachment) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新附件：修改附件名称、描述、状态等元数据信息 |
| DeleteAttachment | [DeleteAttachmentRequest](#api-v1-DeleteAttachmentRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除附件：执行软删除，附件会进入延迟清理流程而不是立即物理删除 |
| GetAttachment | [GetAttachmentRequest](#api-v1-GetAttachmentRequest) | [Attachment](#api-v1-Attachment) | 查询附件：根据附件 ID 获取单个附件详情 |
| ListAttachment | [ListAttachmentRequest](#api-v1-ListAttachmentRequest) | [ListAttachmentReply](#api-v1-ListAttachmentReply) | 附件列表：按关键词、上传用户、类型和扩展名分页筛选附件 |

 



<a name="api_v1_download-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/download.proto



<a name="api-v1-Download"></a>

### Download



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |
| user_id | [int64](#int64) |  |  |
| document_id | [int64](#int64) |  |  |
| ip | [string](#string) |  |  |
| is_pay | [bool](#bool) |  |  |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  |  |
| title | [string](#string) |  |  |
| ext | [string](#string) |  |  |
| score | [int32](#int32) |  |  |
| size | [int64](#int64) |  |  |
| pages | [int32](#int32) |  |  |
| document_uuid | [string](#string) |  |  |






<a name="api-v1-ListDownloadReply"></a>

### ListDownloadReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  |  |
| download | [Download](#api-v1-Download) | repeated |  |






<a name="api-v1-ListDownloadRequest"></a>

### ListDownloadRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  |  |
| size | [int64](#int64) |  |  |
| wd | [string](#string) |  |  |
| field | [string](#string) | repeated |  |
| order | [string](#string) |  |  |





 

 

 


<a name="api-v1-DownloadAPI"></a>

### DownloadAPI
rpc CreateDownload(Download) returns (Download) {
    option (google.api.http) = {
      post : &#39;/api/v1/download&#39;,
      body : &#39;*&#39;,
    };
  }

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|

 



<a name="api_v1_config-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/config.proto



<a name="api-v1-CPUInfo"></a>

### CPUInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cores | [int32](#int32) |  | CPU核心数 |
| model_name | [string](#string) |  | CPU型号 |
| mhz | [float](#float) |  | CPU主频 |
| percent | [float](#float) |  | CPU使用率 |






<a name="api-v1-Config"></a>

### Config
配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 配置ID |
| label | [string](#string) |  | 配置标签 |
| name | [string](#string) |  | 配置名称 |
| value | [string](#string) |  | 配置值 |
| placeholder | [string](#string) |  | 配置占位符 |
| input_type | [string](#string) |  | 输入类型，如：textarea、number、switch等 |
| category | [string](#string) |  | 配置分类，如：system、footer、security等，见 web/utils/enum.js |
| sort | [int32](#int32) |  | 排序，越小越靠前 |
| options | [string](#string) |  | 配置项枚举，一个一行，如select的option选项，用 key=value 的形式 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |
| col_num | [int32](#int32) |  | 配置项占据的列数 |






<a name="api-v1-ConfigCaptcha"></a>

### ConfigCaptcha
验证码配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| length | [int32](#int32) |  | 验证码长度 |
| width | [int32](#int32) |  | 验证码宽度 |
| height | [int32](#int32) |  | 验证码高度 |
| type | [string](#string) |  | 验证码类型，见 web/utils/enum.js |






<a name="api-v1-ConfigDisplay"></a>

### ConfigDisplay



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| show_register_user_count | [bool](#bool) |  | 是否显示注册用户数量 |
| show_index_categories | [bool](#bool) |  | 是否显示首页分类 |
| pages_per_read | [int32](#int32) |  | 每次阅读的页数 |
| copyright_statement | [string](#string) |  | 在页面最底部的版权声明 |
| show_document_descriptions | [bool](#bool) |  | 是否显示文档描述 |
| hide_keywords_on_lists | [bool](#bool) |  | 是否在文档列表页隐藏右侧关键词 |
| show_document_count | [bool](#bool) |  | 是否显示文档数量 |
| show_document_view_count | [bool](#bool) |  | 是否显示文档阅读数量 |
| show_document_download_count | [bool](#bool) |  | 是否显示文档评论数量 |
| show_document_favorite_count | [bool](#bool) |  | 是否显示文档收藏数量 |
| hide_category_without_document | [bool](#bool) |  | 是否隐藏没有文档的分类 |
| wechat_tip | [string](#string) |  | 微信公众号提示文案 |
| wechat_qrcode | [string](#string) |  | 微信公众号链接二维码 |
| contact_tip | [string](#string) |  | 联系我们提示文案 |
| contact_link | [string](#string) |  | 联系我们跳转地址 |
| index_document_style | [string](#string) |  | 首页文档样式 |
| home_version | [string](#string) |  |  |






<a name="api-v1-ConfigFooter"></a>

### ConfigFooter
底链配置项，为跳转的链接地址


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| about | [string](#string) |  | 关于我们 |
| contact | [string](#string) |  | 联系我们 |
| agreement | [string](#string) |  | 用户协议、文库协议 |
| copyright | [string](#string) |  | 版权声明 |
| feedback | [string](#string) |  | 意见和建议反馈 |






<a name="api-v1-ConfigSecurity"></a>

### ConfigSecurity
安全配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| is_close | [bool](#bool) |  | 是否关闭站点 |
| close_statement | [string](#string) |  | 关闭站点的说明，支持HTML |
| enable_register | [bool](#bool) |  | 是否开放注册 |
| enable_captcha_login | [bool](#bool) |  | 是否开启登录验证码 |
| enable_captcha_register | [bool](#bool) |  | 是否开启注册验证码 |
| enable_captcha_comment | [bool](#bool) |  | 是否开启评论验证码 |
| enable_captcha_find_password | [bool](#bool) |  | 是否开启找回密码验证码 |
| enable_captcha_upload | [bool](#bool) |  | 是否开启上传验证码 |
| max_document_size | [int32](#int32) |  | 文档最大大小 |
| document_allowed_ext | [string](#string) | repeated | 文档允许的扩展名 |
| login_required | [bool](#bool) |  | 是否登录才能访问 |
| enable_verify_register_email | [bool](#bool) |  | 是否开启注册邮箱验证 |






<a name="api-v1-ConfigSystem"></a>

### ConfigSystem
系统配置项


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| domain | [string](#string) |  | 站点域名，如： https://moredoc.mnt.ltd |
| title | [string](#string) |  | 站点标题，首页显示 |
| keywords | [string](#string) |  | 站点关键词，SEO用 |
| description | [string](#string) |  | 站点描述，SEO用 |
| logo | [string](#string) |  | 站点logo |
| favicon | [string](#string) |  | 站点favicon |
| icp | [string](#string) |  | 站点备案号 |
| analytics | [string](#string) |  | 站点统计代码，目前只支持百度统计 |
| sitename | [string](#string) |  | 站点名称 |
| copyright_start_year | [string](#string) |  | 站点版权起始年份，如：2018，则底部显示 2018 - 2023 |
| register_background | [string](#string) |  | 注册页背景图 |
| login_background | [string](#string) |  | 登录页背景图 |
| recommend_words | [string](#string) | repeated | 推荐搜索词，首页展示 |
| version | [string](#string) |  | 程序版本号 |
| credit_name | [string](#string) |  | 积分名称 |
| sec_icp | [string](#string) |  | 京公网安备 |






<a name="api-v1-Configs"></a>

### Configs
配置列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| config | [Config](#api-v1-Config) | repeated |  |






<a name="api-v1-DeviceInfo"></a>

### DeviceInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| cpu | [CPUInfo](#api-v1-CPUInfo) |  | CPU信息 |
| memory | [MemoryInfo](#api-v1-MemoryInfo) |  | 内存信息 |
| disk | [DiskInfo](#api-v1-DiskInfo) | repeated | 磁盘信息 |






<a name="api-v1-DiskInfo"></a>

### DiskInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| disk_name | [string](#string) |  | 磁盘名称 |
| total | [int64](#int64) |  | 总空间 |
| used | [int64](#int64) |  | 已用空间 |
| free | [int64](#int64) |  | 空闲空间 |
| percent | [float](#float) |  | 磁盘使用率 |






<a name="api-v1-EnvDependent"></a>

### EnvDependent
依赖项


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| name | [string](#string) |  | 依赖名称 |
| description | [string](#string) |  | 依赖描述 |
| is_installed | [bool](#bool) |  | 是否已安装 |
| error | [string](#string) |  | 错误信息 |
| checked_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 检测时间 |
| cmd | [string](#string) |  | 检测命令 |
| is_required | [bool](#bool) |  | 是否必须 |
| version | [string](#string) |  | 版本号 |






<a name="api-v1-Envs"></a>

### Envs
依赖项列表


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| envs | [EnvDependent](#api-v1-EnvDependent) | repeated |  |






<a name="api-v1-ListConfigRequest"></a>

### ListConfigRequest
查询配置项请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| category | [string](#string) | repeated |  |






<a name="api-v1-MemoryInfo"></a>

### MemoryInfo



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总内存 |
| available | [int64](#int64) |  | 可用内存 |
| used | [int64](#int64) |  | 已用内存 |
| free | [int64](#int64) |  | 空闲内存 |
| percent | [float](#float) |  | 内存使用率 |






<a name="api-v1-Release"></a>

### Release



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| tag_name | [string](#string) |  | 版本号 |
| name | [string](#string) |  | 版本名称 |
| body | [string](#string) |  | 版本说明 |
| release_at | [string](#string) |  | 发布时间 |
| ignore | [string](#string) |  | 忽略版本 |
| source | [string](#string) |  | 更新源 |
| current | [string](#string) |  | 当前版本 |






<a name="api-v1-Settings"></a>

### Settings
系统配置


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| system | [ConfigSystem](#api-v1-ConfigSystem) |  | 系统配置 |
| footer | [ConfigFooter](#api-v1-ConfigFooter) |  | 底链配置 |
| security | [ConfigSecurity](#api-v1-ConfigSecurity) |  | 安全配置 |
| display | [ConfigDisplay](#api-v1-ConfigDisplay) |  | 显示配置 |
| language | [Language](#api-v1-Language) | repeated | 语言配置 |






<a name="api-v1-Stats"></a>

### Stats
系统状态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| user_count | [int64](#int64) |  | 用户数量 |
| document_count | [int64](#int64) |  | 文档数量 |
| category_count | [int64](#int64) |  | 分类数量 |
| article_count | [int64](#int64) |  | 文章数量 |
| comment_count | [int64](#int64) |  | 评论数量 |
| banner_count | [int64](#int64) |  | banner数量 |
| friendlink_count | [int64](#int64) |  | 友情链接数量 |
| os | [string](#string) |  | 操作系统 |
| version | [string](#string) |  | 程序版本号 |
| hash | [string](#string) |  | 程序构建时的 git hash |
| build_at | [string](#string) |  | 程序构建时间 |
| report_count | [int64](#int64) |  | 举报数量 |





 

 

 


<a name="api-v1-ConfigAPI"></a>

### ConfigAPI
配置 API 服务，提供站点配置、运行状态和版本信息管理能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| GetSettings | [.google.protobuf.Empty](#google-protobuf-Empty) | [Settings](#api-v1-Settings) | 获取站点设置：返回前台可读取的系统、展示、安全和语言配置 |
| UpdateConfig | [Configs](#api-v1-Configs) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新配置：批量保存后台配置项的键值内容 |
| ListConfig | [ListConfigRequest](#api-v1-ListConfigRequest) | [Configs](#api-v1-Configs) | 配置项列表：按配置分类查询后台配置项明细 |
| GetStats | [.google.protobuf.Empty](#google-protobuf-Empty) | [Stats](#api-v1-Stats) | 获取系统统计：返回用户、文档、评论等核心统计指标 |
| GetEnvs | [.google.protobuf.Empty](#google-protobuf-Empty) | [Envs](#api-v1-Envs) | 获取环境依赖：检查程序运行依赖及安装状态 |
| UpdateSitemap | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新站点地图：重新生成或刷新站点 sitemap 数据 |
| GetDeviceInfo | [.google.protobuf.Empty](#google-protobuf-Empty) | [DeviceInfo](#api-v1-DeviceInfo) | 获取设备信息：返回当前服务所在机器的 CPU、内存和磁盘状态 |
| SetSQLMode | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置 SQL Mode：修正数据库 only_full_group_by 等兼容性配置 |
| GetLatestRelease | [.google.protobuf.Empty](#google-protobuf-Empty) | [Release](#api-v1-Release) | 获取最新版本：读取远程发布源中的最新版本信息 |
| RefreshLatestRelease | [.google.protobuf.Empty](#google-protobuf-Empty) | [Release](#api-v1-Release) | 刷新最新版本：主动重新拉取远程版本发布信息 |
| IgnoreRelease | [Release](#api-v1-Release) | [.google.protobuf.Empty](#google-protobuf-Empty) | 忽略版本：记录指定版本为忽略状态，避免重复提示 |
| SetReleaseSource | [Release](#api-v1-Release) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置发布源：切换版本检查使用的来源地址或渠道 |

 



<a name="api_v1_user-proto"></a>
<p align="right"><a href="#top">Top</a></p>

## api/v1/user.proto



<a name="api-v1-DeleteUserRequest"></a>

### DeleteUserRequest
删除用户


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) | repeated |  |
| password | [string](#string) |  | 管理员删除用户时需要验证密码 |






<a name="api-v1-Dynamic"></a>

### Dynamic
用户动态


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 动态ID |
| user_id | [int64](#int64) |  | 用户ID |
| content | [string](#string) |  | 内容 |
| type | [int32](#int32) |  | 类型 |
| username | [string](#string) |  | 用户名 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |






<a name="api-v1-FindPasswordRequest"></a>

### FindPasswordRequest
找回密码


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | 邮箱 |
| token | [string](#string) |  | 签名token |
| password | [string](#string) |  | 新密码 |
| captcha | [string](#string) |  | 验证码 |
| captcha_id | [string](#string) |  | 验证码ID |






<a name="api-v1-GetUserCaptchaReply"></a>

### GetUserCaptchaReply
验证码响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| enable | [bool](#bool) |  | 是否启用验证码 |
| id | [string](#string) |  | 验证码ID |
| captcha | [string](#string) |  | 验证码 |
| type | [string](#string) |  | 验证码类型 |






<a name="api-v1-GetUserCaptchaRequest"></a>

### GetUserCaptchaRequest
查询验证码请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| type | [string](#string) |  | 验证码类型：register、login、comment、find_password、upload |






<a name="api-v1-GetUserPermissionsReply"></a>

### GetUserPermissionsReply
用户权限信息查询


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| permission | [Permission](#api-v1-Permission) | repeated |  |






<a name="api-v1-GetUserRequest"></a>

### GetUserRequest
获取用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  |  |






<a name="api-v1-ListUserDownloadReply"></a>

### ListUserDownloadReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| download | [Download](#api-v1-Download) | repeated | 下载列表 |






<a name="api-v1-ListUserDownloadRequest"></a>

### ListUserDownloadRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |






<a name="api-v1-ListUserDynamicReply"></a>

### ListUserDynamicReply
用户动态列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| dynamic | [Dynamic](#api-v1-Dynamic) | repeated | 动态列表 |






<a name="api-v1-ListUserDynamicRequest"></a>

### ListUserDynamicRequest
用户动态列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |






<a name="api-v1-ListUserGroupReply"></a>

### ListUserGroupReply



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| group | [Group](#api-v1-Group) | repeated |  |






<a name="api-v1-ListUserReply"></a>

### ListUserReply
用户列表响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| total | [int64](#int64) |  | 总数 |
| user | [User](#api-v1-User) | repeated | 用户列表 |






<a name="api-v1-ListUserRequest"></a>

### ListUserRequest
用户列表请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| page | [int64](#int64) |  | 页码 |
| size | [int64](#int64) |  | 每页数量 |
| wd | [string](#string) |  | 搜索关键词 |
| sort | [string](#string) |  | 排序字段 |
| id | [int64](#int64) | repeated | 用户ID |
| group_id | [int64](#int64) | repeated | 用户组ID |
| status | [int32](#int32) | repeated | 用户状态 |
| limit | [int64](#int64) |  | 请求数量限制，大于0时，page和size无效 |
| field | [string](#string) | repeated | 返回字段 |






<a name="api-v1-LoginReply"></a>

### LoginReply
用户登录响应


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| token | [string](#string) |  |  |
| user | [User](#api-v1-User) |  |  |






<a name="api-v1-RegisterAndLoginRequest"></a>

### RegisterAndLoginRequest
用户注册登录请求


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |
| captcha | [string](#string) |  | 验证码 |
| captcha_id | [string](#string) |  | 验证码ID |
| email | [string](#string) |  | 邮箱 |
| code | [string](#string) |  | 邮箱验证码(如果用户注册需要验证邮箱，则必须传递该参数) |






<a name="api-v1-SendEmailCodeRequest"></a>

### SendEmailCodeRequest



| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| email | [string](#string) |  | 邮箱 |
| captcha | [string](#string) |  |  |
| captcha_id | [string](#string) |  |  |
| type | [EmailCodeType](#api-v1-EmailCodeType) |  |  |






<a name="api-v1-SetUserRequest"></a>

### SetUserRequest
管理后台设置用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户ID |
| username | [string](#string) |  | 用户名 |
| password | [string](#string) |  | 密码 |
| group_id | [int64](#int64) | repeated | 用户组ID |
| email | [string](#string) |  | 邮箱 |






<a name="api-v1-Sign"></a>

### Sign
用户签到


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 签到ID |
| user_id | [int64](#int64) |  | 用户ID |
| sign_at | [int32](#int32) |  | 签到日期 |
| ip | [string](#string) |  | 签到IP |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 创建时间 |
| award | [int32](#int32) |  | 签到积分奖励 |






<a name="api-v1-UpdateUserPasswordRequest"></a>

### UpdateUserPasswordRequest
修改用户密码


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| id | [int64](#int64) |  | 用户ID |
| old_password | [string](#string) |  | 旧密码 |
| new_password | [string](#string) |  | 新密码 |






<a name="api-v1-User"></a>

### User
用户信息


| Field | Type | Label | Description |
| ----- | ---- | ----- | ----------- |
| login_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 最后登录时间 |
| created_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 注册时间 |
| updated_at | [google.protobuf.Timestamp](#google-protobuf-Timestamp) |  | 更新时间 |
| id | [int64](#int64) |  | 用户ID |
| username | [string](#string) |  | 用户名 |
| mobile | [string](#string) |  | 手机号 |
| email | [string](#string) |  | 邮箱，唯一 |
| address | [string](#string) |  | 地址 |
| signature | [string](#string) |  | 个性签名 |
| last_login_ip | [string](#string) |  | 最后登录IP |
| register_ip | [string](#string) |  | 注册IP |
| doc_count | [int32](#int32) |  | 文档数量 |
| follow_count | [int32](#int32) |  | 关注数量 |
| fans_count | [int32](#int32) |  | 粉丝数量 |
| favorite_count | [int32](#int32) |  | 收藏数量 |
| comment_count | [int32](#int32) |  | 评论数量 |
| status | [int32](#int32) |  | 用户状态，见 web/utils/enum.js，当前没有使用 |
| avatar | [string](#string) |  | 头像 |
| identity | [string](#string) |  | 身份证 |
| realname | [string](#string) |  | 真实姓名 |
| group_id | [int64](#int64) | repeated | 用户组ID |
| credit_count | [int32](#int32) |  | 积分 |
| article_count | [int32](#int32) |  | 文章数量 |
| remark | [string](#string) |  | 备注 |





 


<a name="api-v1-EmailCodeType"></a>

### EmailCodeType


| Name | Number | Description |
| ---- | ------ | ----------- |
| EmailCodeTypeRegister | 0 | 注册 |
| EmailCodeTypeLogin | 1 | 登录 |


 

 


<a name="api-v1-UserAPI"></a>

### UserAPI
用户 API 服务，提供注册登录、资料维护、权限查询和用户行为能力

| Method Name | Request Type | Response Type | Description |
| ----------- | ------------ | ------------- | ------------|
| Register | [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest) | [LoginReply](#api-v1-LoginReply) | 用户注册：通过用户名、密码、邮箱等信息创建账号并返回登录态 |
| Login | [RegisterAndLoginRequest](#api-v1-RegisterAndLoginRequest) | [LoginReply](#api-v1-LoginReply) | 用户登录：校验账号凭证并返回登录令牌和用户信息 |
| Logout | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 退出登录：使当前登录态失效 |
| GetUser | [GetUserRequest](#api-v1-GetUserRequest) | [User](#api-v1-User) | 获取用户信息：可查询指定用户公开资料，或查询当前用户的完整资料 |
| UpdateUserPassword | [UpdateUserPasswordRequest](#api-v1-UpdateUserPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户密码：可修改当前用户密码，管理员也可代为重置指定用户密码 |
| UpdateUserProfile | [User](#api-v1-User) | [.google.protobuf.Empty](#google-protobuf-Empty) | 更新用户资料：修改头像、签名、地址等基础资料信息 |
| DeleteUser | [DeleteUserRequest](#api-v1-DeleteUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 删除用户：支持按用户 ID 批量删除用户，需校验操作权限 |
| AddUser | [SetUserRequest](#api-v1-SetUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 创建用户：管理员直接新增后台用户账号 |
| SetUser | [SetUserRequest](#api-v1-SetUserRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 设置用户：管理员更新用户基础信息和所属用户组 |
| ListUser | [ListUserRequest](#api-v1-ListUserRequest) | [ListUserReply](#api-v1-ListUserReply) | 用户列表：分页查询用户，可按用户组、状态和关键词筛选 |
| GetUserCaptcha | [GetUserCaptchaRequest](#api-v1-GetUserCaptchaRequest) | [GetUserCaptchaReply](#api-v1-GetUserCaptchaReply) | 获取验证码：根据业务场景返回是否启用验证码及验证码内容 |
| GetUserPermissions | [.google.protobuf.Empty](#google-protobuf-Empty) | [GetUserPermissionsReply](#api-v1-GetUserPermissionsReply) | 获取用户权限：返回当前登录用户拥有的权限集合 |
| CanIUploadDocument | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 检查上传权限：判断当前用户是否具备上传文档资格 |
| CanIPublishArticle | [.google.protobuf.Empty](#google-protobuf-Empty) | [.google.protobuf.Empty](#google-protobuf-Empty) | 检查发文权限：判断当前用户是否具备发布文章资格 |
| ListUserDynamic | [ListUserDynamicRequest](#api-v1-ListUserDynamicRequest) | [ListUserDynamicReply](#api-v1-ListUserDynamicReply) | 用户动态列表：查询当前用户及其关注对象的动态内容 |
| SignToday | [.google.protobuf.Empty](#google-protobuf-Empty) | [Sign](#api-v1-Sign) | 每日签到：完成当天签到并返回签到记录和奖励信息 |
| GetSignedToday | [.google.protobuf.Empty](#google-protobuf-Empty) | [Sign](#api-v1-Sign) | 查询今日签到：获取当前用户当天的签到记录 |
| ListUserDownload | [ListUserDownloadRequest](#api-v1-ListUserDownloadRequest) | [ListUserDownloadReply](#api-v1-ListUserDownloadReply) | 下载记录列表：分页查询当前用户的文档下载记录 |
| ListUserGroup | [.google.protobuf.Empty](#google-protobuf-Empty) | [ListUserGroupReply](#api-v1-ListUserGroupReply) | 用户组列表：返回当前站点可选的用户组集合 |
| FindPasswordStepOne | [FindPasswordRequest](#api-v1-FindPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 找回密码第一步：校验账号并发送找回密码验证码 |
| FindPasswordStepTwo | [FindPasswordRequest](#api-v1-FindPasswordRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 找回密码第二步：校验验证码后重置账号密码 |
| SendEmailCode | [SendEmailCodeRequest](#api-v1-SendEmailCodeRequest) | [.google.protobuf.Empty](#google-protobuf-Empty) | 发送邮箱验证码：用于注册或登录等场景的邮箱校验 |

 



## Scalar Value Types

| .proto Type | Notes | C++ | Java | Python | Go | C# | PHP | Ruby |
| ----------- | ----- | --- | ---- | ------ | -- | -- | --- | ---- |
| <a name="double" /> double |  | double | double | float | float64 | double | float | Float |
| <a name="float" /> float |  | float | float | float | float32 | float | float | Float |
| <a name="int32" /> int32 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint32 instead. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="int64" /> int64 | Uses variable-length encoding. Inefficient for encoding negative numbers – if your field is likely to have negative values, use sint64 instead. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="uint32" /> uint32 | Uses variable-length encoding. | uint32 | int | int/long | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="uint64" /> uint64 | Uses variable-length encoding. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum or Fixnum (as required) |
| <a name="sint32" /> sint32 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int32s. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sint64" /> sint64 | Uses variable-length encoding. Signed int value. These more efficiently encode negative numbers than regular int64s. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="fixed32" /> fixed32 | Always four bytes. More efficient than uint32 if values are often greater than 2^28. | uint32 | int | int | uint32 | uint | integer | Bignum or Fixnum (as required) |
| <a name="fixed64" /> fixed64 | Always eight bytes. More efficient than uint64 if values are often greater than 2^56. | uint64 | long | int/long | uint64 | ulong | integer/string | Bignum |
| <a name="sfixed32" /> sfixed32 | Always four bytes. | int32 | int | int | int32 | int | integer | Bignum or Fixnum (as required) |
| <a name="sfixed64" /> sfixed64 | Always eight bytes. | int64 | long | int/long | int64 | long | integer/string | Bignum |
| <a name="bool" /> bool |  | bool | boolean | boolean | bool | bool | boolean | TrueClass/FalseClass |
| <a name="string" /> string | A string must always contain UTF-8 encoded or 7-bit ASCII text. | string | String | str/unicode | string | string | string | String (UTF-8) |
| <a name="bytes" /> bytes | May contain any arbitrary sequence of bytes. | string | ByteString | str | []byte | ByteString | string | String (ASCII-8BIT) |

