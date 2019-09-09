etcd 中使用  grpc 的设计

首先在最外层设计 API 
```go
type Authenticator interface {
	AuthEnable(ctx context.Context, r *pb.AuthEnableRequest) (*pb.AuthEnableResponse, error)
	AuthDisable(ctx context.Context, r *pb.AuthDisableRequest) (*pb.AuthDisableResponse, error)
	Authenticate(ctx context.Context, r *pb.AuthenticateRequest) (*pb.AuthenticateResponse, error)
	UserAdd(ctx context.Context, r *pb.AuthUserAddRequest) (*pb.AuthUserAddResponse, error)
	UserDelete(ctx context.Context, r *pb.AuthUserDeleteRequest) (*pb.AuthUserDeleteResponse, error)
	UserChangePassword(ctx context.Context, r *pb.AuthUserChangePasswordRequest) (*pb.AuthUserChangePasswordResponse, error)
	UserGrantRole(ctx context.Context, r *pb.AuthUserGrantRoleRequest) (*pb.AuthUserGrantRoleResponse, error)
	UserGet(ctx context.Context, r *pb.AuthUserGetRequest) (*pb.AuthUserGetResponse, error)
	UserRevokeRole(ctx context.Context, r *pb.AuthUserRevokeRoleRequest) (*pb.AuthUserRevokeRoleResponse, error)
	RoleAdd(ctx context.Context, r *pb.AuthRoleAddRequest) (*pb.AuthRoleAddResponse, error)
	RoleGrantPermission(ctx context.Context, r *pb.AuthRoleGrantPermissionRequest) (*pb.AuthRoleGrantPermissionResponse, error)
	RoleGet(ctx context.Context, r *pb.AuthRoleGetRequest) (*pb.AuthRoleGetResponse, error)
	RoleRevokePermission(ctx context.Context, r *pb.AuthRoleRevokePermissionRequest) (*pb.AuthRoleRevokePermissionResponse, error)
	RoleDelete(ctx context.Context, r *pb.AuthRoleDeleteRequest) (*pb.AuthRoleDeleteResponse, error)
	UserList(ctx context.Context, r *pb.AuthUserListRequest) (*pb.AuthUserListResponse, error)
	RoleList(ctx context.Context, r *pb.AuthRoleListRequest) (*pb.AuthRoleListResponse, error)
}
```
API 设计为 interface ，可以使用 http , grpc 等进行扩展 API 调用

下面是 扩展 grpc 调用。

包装一层 AuthServer 来满足 grpc 的 server 方法定义

使用满足 Authenticator 接口 的 结构 作为调用 函数的入口

非常像 mvc 中 ， controller 对 model 的操作。

```go
type AuthServer struct {
	authenticator etcdserver.Authenticator
}

func NewAuthServer(s *etcdserver.EtcdServer) *AuthServer {
	return &AuthServer{authenticator: s}
}

func (as *AuthServer) AuthEnable(ctx context.Context, r *pb.AuthEnableRequest) (*pb.AuthEnableResponse, error) {
	resp, err := as.authenticator.AuthEnable(ctx, r)
	if err != nil {
		return nil, togRPCError(err)
	}
	return resp, nil
}

···
```
