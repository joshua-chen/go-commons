/*
 * @Descripttion:
 * @version:
 * @Author: joshua
 * @Date: 2020-05-27 16:03:56
 * @LastEditors: joshua
 * @LastEditTime: 2020-05-27 16:39:46
 */

package msg

const (

	// msg define
	Success                  = "恭喜, 成功"
	OperateSuccess     string = "恭喜, 操作成功"
	OperateFailed      string = "抱歉, 操作失败"
	ParseParamsFailed string = "解析参数失败"

	RegisterSuccess     string = "恭喜, 注册用户成功"
	RegisterFailed      string = "注册失败"
	LoginSuccess       string = "恭喜, 登录成功"
	LoginFailed        string = "登录失败"
	DeleteUsersSuccess string = "删除用户成功"
	DeleteUsersFailed  string = "删除用户错误"

	DeleteRolesSuccess string = "删除角色成功"
	DeleteRolesFailed  string = "删除角色错误"

	UsernameFailed             string = "用户名错误"
	PasswordFailed             string = "密码错误"
	TokenCreateFailed          string = "生成token错误"
	TokenExactFailed           string = "token不存在或header设置不正确"
	TokenExpired               string = "回话已过期"
	TokenParseFailed           string = "token解析错误"
	TokenParseFailedAndEmpty   string = "解析错误,token为空"
	TokenParseFailedAndInvalid string = "解析错误,token无效"
	NotFound                   string = "您请求的url不存在"
	PermissionsLess            string = "权限不足"

	RoleCreateFailed  string = "创建角色失败"
	RoleCreateSuccess string = "创建角色成功"

	// value define

)
