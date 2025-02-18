# 错误码

**注意：** 该组件错误码列表，由 `protoc --go-errors-code_out=paths=source_relative:./docs/guide/zh-CN/api/errors-code` 命令生成，不要对此文件做任何更改。

## 错误说明

如果请求结果返回格式如下：
```json
{
  "metadata": {},
  "message": "User already exists",
  "reason": "UserAlreadyExists",
  "code": 409
}
```

则表示调用 API 接口失败，可能需要客户端进行相应的错误处理。

## 错误码列表

当前组件支持的错误码列表如下：

| Reason | HTTP Status Code | Description |
| :----: | :--------------: | ----------- |
| Unknown | 500 |  未知错误，服务器内部错误 |
| InvalidParameter | 400 |  无效参数错误 |
| NotFound | 400 |  未找到错误 |
| Unauthorized | 401 |  未经授权错误 |
| Forbidden | 403 |  禁止访问错误 |
| IdempotentMissingToken | 400 |  缺少幂等性令牌错误 |
| IdempotentTokenExpired | 400 |  幂等性令牌已过期错误 |
| PageNotFound | 404 |  请求路径没有找到 |
| BindFailed | 500 |  Gin 请求参数绑定失败 |

## 参考

- [错误规范](https://github.com/superproj/zero/blob/master/docs/devel/zh-CN/conversions/errors.md)

