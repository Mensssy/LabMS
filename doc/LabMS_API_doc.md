---
title: LabMS
language_tabs:
  - shell: Shell
  - http: HTTP
  - javascript: JavaScript
  - ruby: Ruby
  - python: Python
  - php: PHP
  - java: Java
  - go: Go
toc_footers: []
includes: []
search: true
code_clipboard: true
highlight_theme: darkula
headingLevel: 2
generator: "@tarslib/widdershins v4.0.23"

---

# LabMS

Base URLs:

# Authentication

# Default

## GET token鉴权

GET /any/url/except/login

用户可同时登录一台电脑和一台移动端设备
若同类型设备重复登录，最新的设备的token有效，之前的设备需要重新登录
目前除用户登录接口以外，其他接口在调用之前均会进行token鉴权

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 是 |none|

> 返回示例

```json
{
  "Data": null,
  "Msg": "token is malformed: token contains an invalid number of segments"
}
```

```json
{
  "Data": null,
  "Msg": "invalid token"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|400|[Bad Request](https://tools.ietf.org/html/rfc7231#section-6.5.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|

### 返回数据结构

状态码 **400**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|null|true|none||none|
|» Msg|string|true|none||none|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|null|true|none||none|
|» Msg|string|true|none||none|

# 用户操作

## POST 用户登录

POST /api/login

除用户id和密码外，还需传入设备信息（“Mobile” or “PC”）

> Body 请求参数

```yaml
userId: "17823997662"
password: "1234"
device: Mobile

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» userId|body|string| 是 |none|
|» password|body|string| 是 |none|
|» device|body|string| 是 |none|

> 返回示例

> 200 Response

```json
{
  "Data": {
    "token": "string",
    "tokenExTime": 0
  },
  "Msg": "string"
}
```

```json
{
  "Data": null,
  "Msg": "wrong password"
}
```

```json
{
  "Data": null,
  "Msg": "user not exists"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|object|true|none||none|
|»» token|string|true|none||none|
|»» tokenExTime|integer|true|none||token过期时间Unix时间戳|
|» Msg|string|true|none||none|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|null|true|none||none|
|» Msg|string|true|none||none|

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|null|true|none||none|
|» Msg|string|true|none||none|

## POST 用户注册

POST /api/signin

> Body 请求参数

```yaml
userId: mensssy
password: "123456"

```

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|body|body|object| 否 |none|
|» userId|body|string| 否 |none|
|» password|body|string| 否 |none|

> 返回示例

```json
{
  "Data": null,
  "Msg": "signin succeeded"
}
```

```json
{
  "Data": null,
  "Msg": "signin failed"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|500|[Internal Server Error](https://tools.ietf.org/html/rfc7231#section-6.6.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|null|true|none||none|
|» Msg|string|true|none||none|

状态码 **500**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|null|true|none||none|
|» Msg|string|true|none||none|

## GET 获取用户信息

GET /users

根据token获取用户信息
报错原因都与token有关，见token鉴权里的样例

### 请求参数

|名称|位置|类型|必选|说明|
|---|---|---|---|---|
|Authorization|header|string| 否 |none|

> 返回示例

```json
{
  "Data": {
    "UserId": "",
    "UserName": "BaconSandwich"
  },
  "Msg": "succeeded"
}
```

```json
{
  "Data": null,
  "Msg": "token has invalid claims: token is expired"
}
```

### 返回结果

|状态码|状态码含义|说明|数据模型|
|---|---|---|---|
|200|[OK](https://tools.ietf.org/html/rfc7231#section-6.3.1)|none|Inline|
|401|[Unauthorized](https://tools.ietf.org/html/rfc7235#section-3.1)|none|Inline|

### 返回数据结构

状态码 **200**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|object|true|none||none|
|»» UserId|string|true|none||none|
|»» UserName|string|true|none||none|
|» Msg|string|true|none||none|

状态码 **401**

|名称|类型|必选|约束|中文名|说明|
|---|---|---|---|---|---|
|» Data|null|true|none||none|
|» Msg|string|true|none||none|

# 数据模型

