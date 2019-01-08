# Hello
my first repository
### 3）获取用户访问token
### GET /api/tokens/access_token 
#### req参数:Query
|字段|描述|类型|是否必填|
----|----|---|---|
|access_token|授权系统访问token|string|是|
|appid|APPID|string|是|
|appkey|appkey为用APP_SECRET对参数进行签名后的字符串|string|是|
|scope|应用授权作用域，拥有多个作用域用逗号（,）分隔|string|否，预留|
|auth_code|登录预授权码|string|是|
