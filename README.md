# Object-Mocker

`OM` = *ObjectMocker*

- [English](./doc/en/README.md)
- [Other](./doc/other/README.md)

提供请求模拟返回工具，并提供了 JSON 的存储能力。

## 使用

### HTTP 请求模拟

`OM` 使用 `gin` 框架

#### RestFul 支持

`OM` 提供了四种 RestFul 的请求方法支持，可以根据实际场景选择。

- GET 获取单个或列表
- POST 创建数据
- DELETE 删除数据
- PUT 更新数据

#### JSON 对象处理

`OM` 提供了 JSON 对象的处理操作。JSON 对象被放在不同的 `PATH` 下，根据请求的 `PATH` 自动归类。`PATH` 本意模拟不同的资源类型。

比如：

- 模拟数据库中的用户信息，`PATH` 可以为 `/database/user`

可以认为 `PATH` 指向的地址为某种资源。

##### POST 创建 JSON 对象

- 请求路由：`json/{json path}`
- 请求方式：`POST`
- 请求体：`application/json`
- 请求体格式：`JSON`

###### 举例：

模拟创建用户信息

```shell
curl --location --request POST 'http://127.0.0.1:3000/json/database/user' \
--header 'Content-Type: application/json' \
--data-raw '{
    "username":"test",
    "age": 20,
    "sex": "unknown",
    "email": "test@test.test"
}'
```

返回信息：

```json
{
  "id": "558237d4-63e8-11ec-8000-acde48001122",
  "create_at": 1640261005980396000,
  "update_at": 1640261005980396000,
  "delete_at": -1,
  "data": {
    "age": 20,
    "email": "test@test.test",
    "sex": "unknown",
    "username": "test"
  }
}
```

##### GET 指定 JSON 对象

目前 GET 方法只提供了 ID 的查找方式

- 请求路由：`json/{json path}`
- 请求方式：`GET`

###### 举例

获取 `/database/user` 下的 `id` 为 `768f26b2-63e8-11ec-8000-acde48001122` 的对象

```shell
curl --location --request GET 'http://127.0.0.1:3000/json/database/user?id=768f26b2-63e8-11ec-8000-acde48001122'
```

返回信息

```json
{
  "id": "768f26b2-63e8-11ec-8000-acde48001122",
  "create_at": 1640261061429830000,
  "update_at": 1640261061429830000,
  "delete_at": -1,
  "data": {
    "age": 20,
    "email": "test@test.test",
    "sex": "unknown",
    "username": "test"
  }
}
```

##### GET 获取指定 PATH 下所有对象

- 请求路由：`json/{json path}`
- 请求方式：`POST`

###### 举例

获取 `/database/user` 下所有 JSON 对象

```shell
curl --location --request GET 'http://127.0.0.1:3000/json/database/user'
```

返回信息

```json
[
  {
    "id": "ca7d095e-63ea-11ec-8000-acde48001122",
    "create_at": 1640262061233188000,
    "update_at": 1640262061233188000,
    "delete_at": -1,
    "data": {
      "age": 20,
      "email": "test@test.test",
      "sex": "unknown",
      "username": "test"
    }
  },
  {
    "id": "caabdaea-63ea-11ec-8000-acde48001122",
    "create_at": 1640262061540018000,
    "update_at": 1640262061540018000,
    "delete_at": -1,
    "data": {
      "age": 20,
      "email": "test@test.test",
      "sex": "unknown",
      "username": "test"
    }
  },
  {
    "id": "cb0ec114-63ea-11ec-8000-acde48001122",
    "create_at": 1640262062188163000,
    "update_at": 1640262062188163000,
    "delete_at": -1,
    "data": {
      "age": 20,
      "email": "test@test.test",
      "sex": "unknown",
      "username": "test"
    }
  }
]
```

##### DELETE 删除 JSON 对象

- 请求路由：`json/{json path}`
- 请求方式：`DELETE`

###### 举例

删除 `/database/user` 下的 `id` 为 `768f26b2-63e8-11ec-8000-acde48001122` 的对象

```shell
curl --location --request DELETE 'http://127.0.0.1:3000/json/database/user/?id=768f26b2-63e8-11ec-8000-acde48001122'
```

返回信息
```json
{
    "id": "768f26b2-63e8-11ec-8000-acde48001122",
    "create_at": 1640261061429830000,
    "update_at": 1640261061429830000,
    "delete_at": 1640262173219044000,
    "data": {
        "age": 20,
        "email": "test@test.test",
        "sex": "unknown",
        "username": "test"
    }
}
```

##### PUT 创建 JSON 对象

- 请求路由：`json/{json path}`
- 请求方式：`PUT`
- 请求体：`application/json`
- 请求体格式：`JSON`

###### 举例
```shell
curl --location --request PUT '127.0.0.1:3000/json/database/user?id=ca7d095e-63ea-11ec-8000-acde48001122' \
--header 'Content-Type: application/json' \
--data-raw '{
    "age": 20,
    "email": "test@test.test",
    "sex": "no sex",
    "username": "test"
}'
```

返回信息

```json
{
  "id": "ca7d095e-63ea-11ec-8000-acde48001122",
  "create_at": 1640262061233188000,
  "update_at": 1640262289549003000,
  "delete_at": -1,
  "data": {
    "age": 20,
    "email": "test@test.test",
    "sex": "no sex",
    "username": "test"
  }
}
```

#### 操作 PATH 对象
所有 JSON 对象都被存放在了对应的 PATH 下，因此有时可能需要查阅`PATH`结构来确定数据状态。

    PATH 结构的操作是比较慢的，因此不建议经常获取 PATH 信息。

PATH 的内存结构本质上是 多路树结构。

##### 获取指定 PATH 结构
```shell
curl --location --request GET 'http://127.0.0.1:3000/node/database/user'
```

返回信息
```json
{
    "scope": "user",
    "children": {
        "username": {
            "scope": "username",
            "children": {},
            "data": {
                "558237d4-63e8-11ec-8000-acde48001122": {
                    "id": "558237d4-63e8-11ec-8000-acde48001122",
                    "create_at": 1640261005980396000,
                    "update_at": 1640261005980396000,
                    "delete_at": -1,
                    "data": {
                        "age": 20,
                        "email": "test@test.test",
                        "sex": "unknown",
                        "username": "test"
                    }
                }
            }
        }
    },
    "data": {
        "c7c3579a-63ea-11ec-8000-acde48001122": {
            "id": "c7c3579a-63ea-11ec-8000-acde48001122",
            "create_at": 1640262056660777000,
            "update_at": 1640262056660777000,
            "delete_at": -1,
            "data": {
                "age": 20,
                "email": "test@test.test",
                "sex": "unknown",
                "username": "test"
            }
        },
        "c817ce4c-63ea-11ec-8000-acde48001122": {
            "id": "c817ce4c-63ea-11ec-8000-acde48001122",
            "create_at": 1640262057214319000,
            "update_at": 1640262057214319000,
            "delete_at": -1,
            "data": {
                "age": 20,
                "email": "test@test.test",
                "sex": "unknown",
                "username": "test"
            }
        },
        "c86f58ec-63ea-11ec-8000-acde48001122": {
            "id": "c86f58ec-63ea-11ec-8000-acde48001122",
            "create_at": 1640262057788031000,
            "update_at": 1640262057788031000,
            "delete_at": -1,
            "data": {
                "age": 20,
                "email": "test@test.test",
                "sex": "unknown",
                "username": "test"
            }
        },
        "c8ba6d3c-63ea-11ec-8000-acde48001122": {
            "id": "c8ba6d3c-63ea-11ec-8000-acde48001122",
            "create_at": 1640262058280071000,
            "update_at": 1640262058280071000,
            "delete_at": -1,
            "data": {
                "age": 20,
                "email": "test@test.test",
                "sex": "unknown",
                "username": "test"
            }
        }
    }
}
```

##### 获取当前全局 PATH 信息

```shell
curl --location --request GET 'http://127.0.0.1:3000/nodes/'
```

返回信息
由于返回信息过多，这里不列出，结构与 **获取指定 PATH 数据** 返回结构相同。