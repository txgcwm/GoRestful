[TOC]

# 系统管理API

## 系统初始化

- URI: `/wapi/__system__/__init__`
- Method: `POST`
- Parameters:
```json
{
  "RootUsername": "root",
  "RootPassword": "root"
}
```

- Response:

```json
{
  "ErrCode": 0,
  "Result": null
}
```


## 清空数据

- URI: `/wapi/__system__/__cleanup__`
- Method: `POST`
- Parameters: `null`

- Response:

```json
{
  "ErrCode": 0,
  "Result": null
}
```


# 账号服务API

## 错误码

| 错误号      | 说明               |
| -------- | ---------------- |
| 20000001 | 用户未登录            |
| 20000002 | 用户不在此群组，也不在上级群组  |
| 20000003 | 无此资源权限记录         |
| 20000004 | 权限不足             |
| 20000005 | secret不合法        |
| 20000006 | 未找到nounce        |
| 20000007 | 未找到设备类型对应的secret |
| 20000008 | 用户名或密码错误         |
| 20000009 | 被删用户必须在指定群组中     |
| 20000010 | 不能操作高等级角色        |
| 20000011 | 角色不存在            |
| 20000012 | ID不合法            |

## 权限认证

### /wapi/auth/session

#### 用户登录

- URI: `/wapi/auth/session`

- Method: `POST`

- Parameters:

  ```json
  {
    "Username": "admin",
    "Password": "fae4be532ce"		// 经MD5摘要后的16进制字串
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,				// optional, 错误码
    "ErrMsg": "Success",		// optional, 错误说明 
    "Result": {
      "Nick": "Deadpool",		// 用户昵称
      "Groups": [{			// 用户所属群组信息列表
        "Group": 234,			// 用户在群组的ID
        "UserRole": 123		// 用户在群内的角色号
      }, {}],
      "Mqtt": [{				// 消息队列服务器信息
        "Host": "",			// optional, 主机名或IP
        "Type": "WebSocket",	// 协议类型, "WebSocket", "Tcp"
        "Port": 15675,		// 端口号
        "Path": "/ws"			// 访问路径
      }, {}]
    }
  }
  ```

#### 用户登出

- URI: `/wapi/auth/session`

- Method: `DELETE`

- Parameters:`null`

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": null
  }
  ```

### /wapi/auth/device/*

#### 设备端secret生成算法
```
{hash} = md5({device_id}.{salt}.{device_secret})
{secret} = {hash[0-5]}{salt[0-3]}{hash[6-17]}{slat[4-7]}{hash[18-23]}
```
- device_secret 统一分配

- salt 随机生成

#### 获取鉴权随机数

- URI: `/wapi/auth/device/nounce`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// 设备ID, 包含密钥类型
    "Secret": "a83b83ca9259c1sf8"		// 根据密钥类型对应的密钥与设备ID hash 生成
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Nounce": "8ac76d6efa5c5543d3"		// 用于获取 Token
    }
  }
  ```

#### 获取 AccessToken
- URI: `/wapi/auth/device/token`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// 设备ID, 可根据ID得到设备类型号
    "Secret": "a83b83ca9259c1sf8"		// 根据密钥类型对应的密钥与设备ID 和 Nounce hash 生成
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Token": "8ac76d6efa5c5543d3",	// Access token
      "LifeTime": 1000,				// token 有效时间，单位秒
      "Bound": true,					// 设备是否已绑定，未绑定的设备，需要设备发绑定请求
      "BindGroup": 123,				// 设备未绑定时，发送绑定请求至此群组，大于0有效
      "Mqtt": [{					// 消息队列服务器信息
        "Host": "",				// optional, 主机名或IP
        "Type": "Tcp",			// 协议类型, "WebSocket", "Tcp"
        "Port": 1883,				// 端口号
        "Path": ""				// 访问路径
      }, {}]
    }
  }
  ```

### /wapi/auth/rabbitmq/*

用于 rabbitmq http auth backend

#### user_path

- URI: `/wapi/auth/rabbitmq/user`

- Method: `POST`

- Parameter:

  ```json
  {
    "username": "1#admin",
    "password": "843afe55627"		// hash 之后的密码
  }
  ```

- Response: `allow`

#### vhost_path

- URI: `/wapi/auth/rabbitmq/vhost`

- Method: `POST`

- Parameter:

  ```json
  {
    "username": "1#admin",
    "vhost": "/",
    "ip": "192.168.1.10"
  }
  ```

- Response: `allow`

#### resource_path

- URI: `/wapi/auth/rabbitmq/resource`

- Method: `POST`

- Parameter:

  ```json
  {
    "username": "1#admin",
    "vhost": "/",
    "resource": "queue",
    "name": "mqtt-subscription-1#adminqos1",
    "permission": "configure"
  }
  ```

- Response: `deny`

#### topic_path

- URI: `/wapi/auth/rabbitmq/topic`

- Method: `POST`

- Parameter:

  ```json
  {
    "username": "1#admin",
    "vhost": "/",
    "resource": "topic",
    "name": "/device/bind/123456/4",
    "permission": "read",
    "routing_key": "xxx"
  }
  ```

- Response: `allow`




## 账号管理


### /wapi/account/password

#### 修改自身密码

- URI: `/wapi/account/password`
- Method: `PUT`
- Parameter: 

```json
{
  "Old": "admin",     // 旧密码
  "New": "admin"      // 新密码
}
```

- Response:

```json
{
  "ErrCode": 0
}
```

### /wapi/account/active

#### 更新用户状态（接口待定，暂无使用）

- URI: `/wapi/account/active`
- Method: `PUT`

##### 启用用户

- Parameter:

  ```json
  {
    "Action": "Enable",		// 启用用户
    "Username": "admin"		// 用户ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "ActiveCode": "123456"	// 激活码
    }
  }
  ```

##### 激活用户
- Parameter:

  ```json
  {
    "Action": "Activate",		// 启用用户
    "Username": "admin"		// 用户ID
    "ActiveCode": "123456",    // 激活码
    "Password": "888888"		// 新密码明文
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

##### 禁用用户
- Parameter:
  ```json
  {
    "Action": "Disable",		// 启用用户
    "Username": "admin"		// 用户ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

## 群组管理

### /wapi/group/nodes

#### 查询群组子节点树

- URI: `/wapi/group/nodes`

- Method: `GET`

- Parameter:

  ```json
  {
    "Root": 12345，	// 群组根节点ID
    "Deep": 1			// 递归深度，0表示不递归，只查询本身信息，-1表示全递归
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Node": 111,
      "Name": "node111",
      "Memo": {},			// 附加信息
      "Children": [{
        "Node": 222,
        "Name": "node222",
        "Memo": {},
        "Children": null
      }, {}]
    }
  }
  ```

#### 创建群组节点

- URI: `/wapi/group/nodes`

- Method: `POST`

- Parameter:

  ```json
  {
    "Parent": 111,
    "Name": "node111",
    "Memo": {}			// 附加信息，不同应用有不同格式，见下表
  }
  ```
  | 应用名称     | 备注信息格式                                   | 说明      |
  | -------- | ---------------------------------------- | ------- |
  | 佐水污水处理系统 | {"Address":"杭州市西湖区","Owner":[{"Name":"张三","Phone":"13966990066"}]} | 企业责任人信息 |
  |          |                                          |         |

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Node": 999			// 新建的节点ID
    }
  }
  ```

#### 删除空群组节点

只能删除空群组节点，空节点是无任何下属的节点

- URI: `/wapi/group/nodes`

- Method: `DELETE`

- Parameter:

  ```json
  {
    "Node": 111			// 删除的节点ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

#### 修改群组节点

- URI: `/wapi/group/nodes`

- Method: `PUT`

- Parameter:

  ```json
  {
    "Node": 222,
    "Parent": 111,
    "Name": "node111",
    "Memo": {}			// 附加信息
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

### /wapi/group/users

#### 查询组内用户列表

- URI: `/wapi/group/users`

- Method: `GET`

- Parameter: 

  ```json
  {
    "Group": 2			// 所属群组
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "Username": "user1",
      "Nick": "spock",
      "Role": 5				// 在群的角色
    }, {}]
  }
  ```

#### 创建组内用户

- URI: `/wapi/group/users`

- Method: `POST`

- Parameter:

  ```json
  {
    "Username": "guest",	// 用户ID
    "Password": "mima",
    "Nick": "kerk",
    "Group": 2,			// 所属群组ID
    "Role": 2				// 在群的角色
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

#### 删除组内用户

先从群组中移除用户，如果移除后用户不属于任何群组，再从系统中删除用户账号。只能删除群管理员创建的账号。

- URI: `/wapi/group/users`

- Method: `DELETE`

- Parameter:

  ```json
  {
    "Group": 5,			// 群组ID
    "Username": "user1"	// 要移除的用户ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```


### /wapi/group/role

#### 更新组内用户角色

- URI: `/wapi/group/role`
- Method: `PUT`
- Parameter: 

```json
{
  "Username": "robin",	// 用户ID
  "Group": 5,			// 群组ID
  "Role": 2				// 角色编号
}
```
- Response:

```json
{
  "ErrCode": 0
}
```

### /wapi/group/password

#### 更新组内用户密码

不能修改自身密码，不能修改非群管理员创建的用户的密码

- URI: `/wapi/group/password`
- Method: `PUT`
- Parameter: 

```json
{
  "Group": 4,				// 用户所在群组
  "Username": "robin",		// 用户ID
  "Password": "password"	// 密码
}
```

- Response:

```json
{
  "ErrCode": 0
}
```


## 设备管理
### 设备ID编码规则

规则：`${设备类型编号}_${序列号类型代码}${序列号}`
示例：`W001_P13866886688`

| 类型编号 | 说明            |
| ---- | ------------- |
| W001 | 污水处理机，内部代号001 |
|      |               |

| 序列号类型代码 | 说明    |
| ------- | ----- |
| P       | 手机号   |
| M       | MAC地址 |


### /wapi/dev/devices

#### 查询设备列表

- URI: `/wapi/dev/devices`

- Method: `GET`

- Parameter:

  ```json
  {
    "Groups": [2],	// 设备所在群节点ID列表
    "Position": false	// 是否返回设置地理位置
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "Group": 2345,				// 所属群组
    	"Device": "W001_P13866886688",	// 设备ID
      "Name": "豪华处理机",		// 设备名称
      "Model": "污水处理机",		// 设备型号，由厂家自定义
      "BindTime": 12345,			// 设备绑定时间(UTC)，单位秒
      "Memo": {},					// 附加信息，不同应用有不同格式，见下表
      "Position": [120.0, 30.5],	// (经度，纬度)
      "PositionTime": 1432344983	// 位置更新时间(UTC)，单位秒
    }, {}]
  }
  ```

#### 创建设备
- URI: `/wapi/dev/devices`

- Method: `POST`

- Parameter:

  ```json
  {
    "Group": 2345,					// 所属群组
    "Device": "W001_P13866886688",	// 设备ID
    "Name": "MachineI",				// 设备名称
    "Model": "污水处理机",				// 设备型号，由厂家自定义
    "Memo": {}						// 附加信息，不同应用有不同格式，见下表
  }
  ```

  | 应用名称     | 备注信息格式                                   | 说明      |
  | -------- | ---------------------------------------- | ------- |
  | 佐水污水处理系统 | {"Owner":[{"Name":"张三","Phone":"13966990066"}]} | 记录责任人信息 |
  |          |                                          |         |

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": null
  }
  ```

#### 删除设备
- URI: `/wapi/dev/devices`

- Method: `DELETE`

- Parameter:

  ```json
  {
    "Device": "0_P_15858275538"	// 设备ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": null
  }
  ```

#### 修改设备信息

- URI: `/wapi/dev/devices`

- Method: `PUT`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// 设备ID
    "Name": "MachineI",				// 设备名称
    "Group": 2345,						// 所属群组
    "Model": "污水处理机",				// 设备型号，由厂家自定义
    "Memo": {}						// 附加信息
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": null
  }
  ```

### /wapi/dev/events
#### 事件查询

- URI: `/wapi/dev/events`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// 查询指定设备
    "Groups": [],				// 查询群组列表内所有设备，Device字段不为空时有效
    "Name": "Motion",			// 查询的事件名称，为空表示查询所有事件
    "Action": 1,				// 只查询指定的动作
    "Limit": 10,				// 分页控制，最大查询条数
    "Offset": 30				// 分页控制，从第几条开始查询
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "ID": 123456,			// 事件记录ID号
      "Name": "Motion",		// 事件名称，UI上的事件原因据此显示
      "Action": 0,			// 最后收到的动作类型
      "Time": [1000, 2000],	// 事件起止时间
      "Device": "W001_P13866886688",	// 事件产生的设备ID
      "Sensor": "1",			// 事件产生的传感器标识号，意义同报警通道号
      "Process": 0			// 事件处理标志, 0-未处理，1-处理
    }, {}]
  }
  ```

#### 事件上报

- URI: `/wapi/dev/events`

- Method: `POST`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// 设备ID
    "Name":"Motion",			// 事件名称
    "Sensor": "1",			// 检测到事件的传感器标识号，意义同报警通道号
    "Action": 1,				// 动作类型: 0-stop, 1-start, 2-pulse
    "Time": [1000, 2000]		// 开始和结束时间，只有stop才需要带结束时间，单位为毫秒
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```
- 权限说明：后台系统用户才可上报

#### 事件处理标志更新

- URI: `/wapi/dev/events`

- Method: `PUT`

- Parameter:

  ```json
  {
    "ID": 12345,			// 事件记录ID
    "Process": 1			// 事件处理标志, 0-未处理，1-处理
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

### /wapi/dev/logs
#### 日志查询

- URI: `/wapi/dev/logs`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// 设备ID
    "Name": "Log1",			// 查询的日志名称
    "Time": [1000, 2000]		// 起止时间(UTC)，单位为毫秒，若为空，表示只查询最新事件
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "ID": 123456,		// 日志记录ID号
    	"Name": "Log1",		// 日志名称
      "Time": 1000,		// 日志产生时间，UTCMS
      "Brief": "xxx",		// 日志简要信息
      "Detail": "{}"		// 日志详细信息
    }, {}]
  }
  ```

#### 日志上报

- URI: `/wapi/dev/logs`

- Method: `POST`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// 设备ID
    "Name": "Log1",		// 日志名称
    "Time": 1000,			// 日志产生时间，UTCMS
    "Brief": "xxx",		// 日志简要信息
    "Detail": "{}"		// 日志详细信息
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```
# MQTT 服务API
## MQTT 访问权限控制
### Client 标识

- 格式：`{type}#{id}`

- 说明：

  | type     | id   | 示例                  |
  | -------- | ---- | ------------------- |
  | 1-user   | 用户ID | 1#admin             |
  | 2-device | 设备ID | 2#W001_P13866885566 |
  | 3-group  | 群组ID | 3#5                 |

### 登录

- 用户名：同Client标识
- 密码：设备使用token作为密码，用户账号使用密码MD5摘要之后作为密码

## 绑定类消息

### /bind/request/{group}/{device}

#### 说明

设备端发出的绑定到群组请求

####权限
- 订阅权限拥有者：具有群组内"device"资源读权限的用户
- 发布权限拥有者：{device}所指设备
#### 内容

空

### /bind/update/{group}/{device}

#### 说明

客户端发出的设备群组绑定关系变更通知

####权限
- 订阅权限拥有者：{device}所指设备，以及具有群组内"device"资源读权限的用户

- 发布权限拥有者：具有群组内"device"资源写权限的用户

#### 内容

```json
{
  "Type": "Bind",	// 变更类型: Bind, Unbind, Move
  "From": 12		// 原所属群组，变更类型为Move时有效
}
```


## 设备类消息

### /device/info/{device}

#### 说明

由设备发布的自身状态变更通知

####权限
- 订阅权限拥有者：具有{device}所在群组内"device"资源读权限的用户
- 发布权限拥有者：{device}所指设备
#### 内容

```json
[{
    "Key": "Temperature",
  	"Value": "35"
}, {}]
```

#### 状态列表

| Key  | Value Format           | 说明      |
| ---- | ---------------------- | ------- |
| Gps  | {longitude},{latitude} | 地理位置坐标点 |
|      |                        |         |
|      |                        |         |

### /device/query/{device}

#### 说明

由客户端发布的设备状态查询消息，设备收到消息后回复 /device/info 消息

####权限
- 订阅权限拥有者：{device}所指设备
- 发布权限拥有者：具有{device}所在群组内"device"资源读权限的用户
#### 内容

```json
[
  "Temperature", "Humidity"		// 数组为空则表示查询所有属性
]
```

### /device/control/{device}

#### 说明

由客户端发布设备控制命令

####权限
- 订阅权限拥有者：{device}所指设备
- 发布权限拥有者：具有{device}所在群组内"device"资源执行权限的用户
#### 内容

```json
[{
  "Command": "Switch1",	// 命令名
  "Params": ["On"]
}, {}]
```

#### 命令列表

| Command | Parameters | 说明   |
| ------- | ---------- | ---- |
|         |            |      |
|         |            |      |

### /device/event/{device}

#### 说明

由设备发布的报警事件

####权限
- 订阅权限拥有者：具有{device}所在群组内"device"资源读权限的用户
- 发布权限拥有者：{device}所指设备
#### 内容

```json
{
    "Name": "Motion",
  	"Action": 1,			// 0-stop, 1-start, 2-pulse
    "Time": [1000, 2000]	// 开始和结束时间，"Start" "Pulse" 时无须带结束时间，单位为毫秒
}
```


## 对话类消息

### /talk/{from_id}/{to_id}[/{group}]
#### 说明

由设备或用户发送的对话消息，群组不能发送消息

| Field    | Example | Comment                       |
| -------- | ------- | ----------------------------- |
| frome_id | 1#admin | 发送方 Client ID                 |
| to_id    | 3#5     | 接收方 Client ID                 |
| group    | 群组      | 接收方所在群组ID，接收方不是群组时必须加上，用于权限验证 |

####权限
- 订阅权限拥有者：仅限接收者，群内用户均可订阅消息
- 发布权限拥有者：具接收者所在群组查看（读）权限
#### 内容

```json
{
    "Type": 1,		// 消息类型，1-SMS, 2-WebRTC ...
    "Data": {}		// 消息内容
}
```

##### SMS 消息格式

- Type: `1`
- Data: 短信字符串

##### WebRTC 消息格式

- Type: `2`
- Data: 同 WebRTC 示例代码中信令格式



# 地理位置服务API

## 错误码

| 错误号  | 说明   |
| ---- | ---- |
|      |      |

## 位置管理
### /wapi/gps/info

#### 查询位置

- URI: `/wapi/gps/info`

- Method: `GET`

- Parameter:

  ```json
  {
    "Type": 2,			// 实体类型: 1-user, 2-device
    "ID": "W001_P13688660055"	// 实体ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Position": [120.0, 30.5],	// (经度，纬度)
      "Time": 1432344983			// 位置更新时间(UTC)，单位秒
    }
  }
  ```

#### 更新位置

- URI: `/wapi/gps/info`

- Method: `PUT`

- Parameter:

  ```json
  {
    "Type": 2,				// 实体类型: 1-user, 2-device
    "ID": "W001_P13688660055",	// 实体ID
    "Position": [120.0, 30.5]		// (经度，纬度)
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```




