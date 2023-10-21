# cloud-isk

# 创建API服务
goctl api new core
# 启动服务
go run core.go -f etc/core-api.yaml
# 使用api文件生成代码
goctl api go -api core.api -dir . -style go_zero

腾讯云COS后台地址：https://console.cloud.tencent.com/cos/bucket

腾讯云COS帮助文档：https://cloud.tencent.com/document/product/436/31215

系统模块：

 用户模块：
 密码登录
 刷新Authorization
 邮箱注册
 用户详情
 用户容量
 
 存储池模块：
 中心存储池资源管理
 文件上传
 文件秒传
 文件分片上传
 对接 MinIO
 对接阿里对象存储
 个人存储池资源管理
 文件关联存储
 文件列表
 文件名称修改
 文件夹创建
 文件删除
 文件移动
 
 文件分享模块：
 创建分享记录
 获取资源详情
 资源保存
 
### 1. "用户登陆"

1. route definition

- Url: /user/login
- Method: POST
- Request: `LoginRequest`
- Response: `LoginReply`

2. request definition



```golang
type LoginRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
}
```


3. response definition



```golang
type LoginReply struct {
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
```

### 2. "用户详情"

1. route definition

- Url: /user/detail
- Method: GET
- Request: `UserDetailRequest`
- Response: `UserDetailReply`

2. request definition



```golang
type UserDetailRequest struct {
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type UserDetailReply struct {
	Name string `json:"name"`
	Email string `json:"email"`
}
```

### 3. "验证码发送"

1. route definition

- Url: /mail/code/send/register
- Method: POST
- Request: `MailCodeRequest`
- Response: `MailCodeReply`

2. request definition



```golang
type MailCodeRequest struct {
	Email string `json:"email"`
}
```


3. response definition



```golang
type MailCodeReply struct {
}
```

### 4. "用户注册"

1. route definition

- Url: /user/register
- Method: POST
- Request: `UserRegisterRequest`
- Response: `UserRegisterReply`

2. request definition



```golang
type UserRegisterRequest struct {
	Name string `json:"name"`
	Password string `json:"password"`
	Email string `json:"email"`
	Code string `json:"code"`
}
```


3. response definition



```golang
type UserRegisterReply struct {
}
```

### 5. "分享资源详情"

1. route definition

- Url: /share/basic/detail
- Method: GET
- Request: `ShareBasicDetailRequest`
- Response: `ShareBasicDetailReply`

2. request definition



```golang
type ShareBasicDetailRequest struct {
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type ShareBasicDetailReply struct {
	RepositoryIdentity string `json:"repository_identity"`
	Name string `json:"name"`
	Ext string `json:"ext"`
	Size int64 `json:"size"`
	Path string `json:"path"`
}
```

### 6. "文件上传"

1. route definition

- Url: /file/upload
- Method: POST
- Request: `FileUploadRequest`
- Response: `FileUploadReply`

2. request definition



```golang
type FileUploadRequest struct {
	Hash string `json:"hash,optional"`
	Name string `json:"name,optional"`
	Ext string `json:"ext,optional"`
	Size int64 `json:"size,optional"`
	Path string `json:"path,optional"`
}
```


3. response definition



```golang
type FileUploadReply struct {
	Identity string `json:"identity"`
	Name string `json:"name"`
	Ext string `json:"ext"`
}
```

### 7. "用户文件关联存储"

1. route definition

- Url: /user/repository/save
- Method: POST
- Request: `UserRepositoryRequest`
- Response: `UserRepositoryReply`

2. request definition



```golang
type UserRepositoryRequest struct {
	ParentId int `json:"parentId"`
	RepositoryIdentity string `json:"repositoryIdentity"`
	Ext string `json:"ext"`
	Name string `json:"name"`
}
```


3. response definition



```golang
type UserRepositoryReply struct {
}
```

### 8. "用户文件文件列表"

1. route definition

- Url: /user/file/list
- Method: GET
- Request: `UserFileListRequest`
- Response: `UserFileListReply`

2. request definition



```golang
type UserFileListRequest struct {
	Id int `json:"id,optional"`
	Page int `json:"page,optional"`
	Size int `json:"size,optional"`
}
```


3. response definition



```golang
type UserFileListReply struct {
	List []*UserFile `json:"list"`
	Count int `json:"count"`
}
```

### 9. "用户文件名称修改"

1. route definition

- Url: /user/file/name/update
- Method: POST
- Request: `UserFileNameUpdateRequest`
- Response: `UserFileNameUpdateReply`

2. request definition



```golang
type UserFileNameUpdateRequest struct {
	Name string `json:"name"`
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type UserFileNameUpdateReply struct {
}
```

### 10. "用户文件夹创建"

1. route definition

- Url: /user/folder/create
- Method: POST
- Request: `UserFolderCreateRequest`
- Response: `UserFolderCreateReply`

2. request definition



```golang
type UserFolderCreateRequest struct {
	ParentId int64 `json:"parent_id"`
	Name string `json:"name"`
}
```


3. response definition



```golang
type UserFolderCreateReply struct {
	Identity string `json:"identity"`
}
```

### 11. "用户文件删除"

1. route definition

- Url: /user/file/delete
- Method: DELETE
- Request: `UserFileDeleteRequest`
- Response: `UserFileDeleteReply`

2. request definition



```golang
type UserFileDeleteRequest struct {
	Identity string `json:"identity"`
}
```


3. response definition



```golang
type UserFileDeleteReply struct {
}
```

### 12. "用户文件修改"

1. route definition

- Url: /user/file/move
- Method: PUT
- Request: `UserFileMoveRequest`
- Response: `UserFileMoveReply`

2. request definition



```golang
type UserFileMoveRequest struct {
	Identity string `json:"identity"`
	ParentIdentity string `json:"parent_identity"`
}
```


3. response definition



```golang
type UserFileMoveReply struct {
}
```

### 13. "创建分享"

1. route definition

- Url: /share/basic/create
- Method: POST
- Request: `ShareBasicCreateRequest`
- Response: `ShareBasicCreateReply`

2. request definition



```golang
type ShareBasicCreateRequest struct {
	UserRepositoryIdentity string `json:"user_repository_identity"`
	ExpiredTime int `json:"expired_time"`
}
```


3. response definition



```golang
type ShareBasicCreateReply struct {
	Identity string `json:"identity"`
}
```

### 14. "分享资源的保存"

1. route definition

- Url: /share/basic/save
- Method: POST
- Request: `ShareBasicSaveRequest`
- Response: `ShareBasicSaveReply`

2. request definition



```golang
type ShareBasicSaveRequest struct {
	RepositoryIdentity string `json:"repository_identity"`
	ParentId int64 `json:"parent_id"`
}
```


3. response definition



```golang
type ShareBasicSaveReply struct {
	Identity string `json:"identity"`
}
```

### 15. "刷新Authorization"

1. route definition

- Url: /refresh/authorization
- Method: POST
- Request: `RefreshAuthorizationRequest`
- Response: `RefreshAuthorizationSaveReply`

2. request definition



```golang
type RefreshAuthorizationRequest struct {
}
```


3. response definition



```golang
type RefreshAuthorizationSaveReply struct {
	Token string `json:"token"`
	RefreshToken string `json:"refresh_token"`
}
```

### 16. "文件分片上传基本处理"

1. route definition

- Url: /file/upload/prepare
- Method: POST
- Request: `FileUploadPrepareRequest`
- Response: `FileUploadPrepareReply`

2. request definition



```golang
type FileUploadPrepareRequest struct {
	Md5 string `json:"md5"`
	Name string `json:"name"`
	Ext string `json:"ext"`
}
```


3. response definition



```golang
type FileUploadPrepareReply struct {
	Identity string `json:"identity"`
	UploadId string `json:"upload_id"`
	Key string `json:"key"`
}
```

### 17. "文件分片上传"

1. route definition

- Url: /file/upload/chunk
- Method: POST
- Request: `FileUploadChunkRequest`
- Response: `FileUploadChunkReply`

2. request definition



```golang
type FileUploadChunkRequest struct {
}
```


3. response definition



```golang
type FileUploadChunkReply struct {
	Etag string `json:"etag"`
}
```

### 18. "文件分片上传完成"

1. route definition

- Url: /file/upload/chunk/complete
- Method: POST
- Request: `FileUploadChunkCompleteRequest`
- Response: `FileUploadChunkCompleteReply`

2. request definition



```golang
type FileUploadChunkCompleteRequest struct {
	Key string `json:"key"`
	UploadId string `json:"upload_id"`
	CosObjects []CosObject `json:"cos_objects"`
}
```


3. response definition



```golang
type FileUploadChunkCompleteReply struct {
}
```

