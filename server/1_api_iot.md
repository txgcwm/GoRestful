[TOC]

# ϵͳ����API

## ϵͳ��ʼ��

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


## �������

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


# �˺ŷ���API

## ������

| �����      | ˵��               |
| -------- | ---------------- |
| 20000001 | �û�δ��¼            |
| 20000002 | �û����ڴ�Ⱥ�飬Ҳ�����ϼ�Ⱥ��  |
| 20000003 | �޴���ԴȨ�޼�¼         |
| 20000004 | Ȩ�޲���             |
| 20000005 | secret���Ϸ�        |
| 20000006 | δ�ҵ�nounce        |
| 20000007 | δ�ҵ��豸���Ͷ�Ӧ��secret |
| 20000008 | �û������������         |
| 20000009 | ��ɾ�û�������ָ��Ⱥ����     |
| 20000010 | ���ܲ����ߵȼ���ɫ        |
| 20000011 | ��ɫ������            |
| 20000012 | ID���Ϸ�            |

## Ȩ����֤

### /wapi/auth/session

#### �û���¼

- URI: `/wapi/auth/session`

- Method: `POST`

- Parameters:

  ```json
  {
    "Username": "admin",
    "Password": "fae4be532ce"		// ��MD5ժҪ���16�����ִ�
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,				// optional, ������
    "ErrMsg": "Success",		// optional, ����˵�� 
    "Result": {
      "Nick": "Deadpool",		// �û��ǳ�
      "Groups": [{			// �û�����Ⱥ����Ϣ�б�
        "Group": 234,			// �û���Ⱥ���ID
        "UserRole": 123		// �û���Ⱥ�ڵĽ�ɫ��
      }, {}],
      "Mqtt": [{				// ��Ϣ���з�������Ϣ
        "Host": "",			// optional, ��������IP
        "Type": "WebSocket",	// Э������, "WebSocket", "Tcp"
        "Port": 15675,		// �˿ں�
        "Path": "/ws"			// ����·��
      }, {}]
    }
  }
  ```

#### �û��ǳ�

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

#### �豸��secret�����㷨
```
{hash} = md5({device_id}.{salt}.{device_secret})
{secret} = {hash[0-5]}{salt[0-3]}{hash[6-17]}{slat[4-7]}{hash[18-23]}
```
- device_secret ͳһ����

- salt �������

#### ��ȡ��Ȩ�����

- URI: `/wapi/auth/device/nounce`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// �豸ID, ������Կ����
    "Secret": "a83b83ca9259c1sf8"		// ������Կ���Ͷ�Ӧ����Կ���豸ID hash ����
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Nounce": "8ac76d6efa5c5543d3"		// ���ڻ�ȡ Token
    }
  }
  ```

#### ��ȡ AccessToken
- URI: `/wapi/auth/device/token`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// �豸ID, �ɸ���ID�õ��豸���ͺ�
    "Secret": "a83b83ca9259c1sf8"		// ������Կ���Ͷ�Ӧ����Կ���豸ID �� Nounce hash ����
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Token": "8ac76d6efa5c5543d3",	// Access token
      "LifeTime": 1000,				// token ��Чʱ�䣬��λ��
      "Bound": true,					// �豸�Ƿ��Ѱ󶨣�δ�󶨵��豸����Ҫ�豸��������
      "BindGroup": 123,				// �豸δ��ʱ�����Ͱ���������Ⱥ�飬����0��Ч
      "Mqtt": [{					// ��Ϣ���з�������Ϣ
        "Host": "",				// optional, ��������IP
        "Type": "Tcp",			// Э������, "WebSocket", "Tcp"
        "Port": 1883,				// �˿ں�
        "Path": ""				// ����·��
      }, {}]
    }
  }
  ```

### /wapi/auth/rabbitmq/*

���� rabbitmq http auth backend

#### user_path

- URI: `/wapi/auth/rabbitmq/user`

- Method: `POST`

- Parameter:

  ```json
  {
    "username": "1#admin",
    "password": "843afe55627"		// hash ֮�������
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




## �˺Ź���


### /wapi/account/password

#### �޸���������

- URI: `/wapi/account/password`
- Method: `PUT`
- Parameter: 

```json
{
  "Old": "admin",     // ������
  "New": "admin"      // ������
}
```

- Response:

```json
{
  "ErrCode": 0
}
```

### /wapi/account/active

#### �����û�״̬���ӿڴ���������ʹ�ã�

- URI: `/wapi/account/active`
- Method: `PUT`

##### �����û�

- Parameter:

  ```json
  {
    "Action": "Enable",		// �����û�
    "Username": "admin"		// �û�ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "ActiveCode": "123456"	// ������
    }
  }
  ```

##### �����û�
- Parameter:

  ```json
  {
    "Action": "Activate",		// �����û�
    "Username": "admin"		// �û�ID
    "ActiveCode": "123456",    // ������
    "Password": "888888"		// ����������
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

##### �����û�
- Parameter:
  ```json
  {
    "Action": "Disable",		// �����û�
    "Username": "admin"		// �û�ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

## Ⱥ�����

### /wapi/group/nodes

#### ��ѯȺ���ӽڵ���

- URI: `/wapi/group/nodes`

- Method: `GET`

- Parameter:

  ```json
  {
    "Root": 12345��	// Ⱥ����ڵ�ID
    "Deep": 1			// �ݹ���ȣ�0��ʾ���ݹ飬ֻ��ѯ������Ϣ��-1��ʾȫ�ݹ�
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Node": 111,
      "Name": "node111",
      "Memo": {},			// ������Ϣ
      "Children": [{
        "Node": 222,
        "Name": "node222",
        "Memo": {},
        "Children": null
      }, {}]
    }
  }
  ```

#### ����Ⱥ��ڵ�

- URI: `/wapi/group/nodes`

- Method: `POST`

- Parameter:

  ```json
  {
    "Parent": 111,
    "Name": "node111",
    "Memo": {}			// ������Ϣ����ͬӦ���в�ͬ��ʽ�����±�
  }
  ```
  | Ӧ������     | ��ע��Ϣ��ʽ                                   | ˵��      |
  | -------- | ---------------------------------------- | ------- |
  | ��ˮ��ˮ����ϵͳ | {"Address":"������������","Owner":[{"Name":"����","Phone":"13966990066"}]} | ��ҵ��������Ϣ |
  |          |                                          |         |

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Node": 999			// �½��Ľڵ�ID
    }
  }
  ```

#### ɾ����Ⱥ��ڵ�

ֻ��ɾ����Ⱥ��ڵ㣬�սڵ������κ������Ľڵ�

- URI: `/wapi/group/nodes`

- Method: `DELETE`

- Parameter:

  ```json
  {
    "Node": 111			// ɾ���Ľڵ�ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

#### �޸�Ⱥ��ڵ�

- URI: `/wapi/group/nodes`

- Method: `PUT`

- Parameter:

  ```json
  {
    "Node": 222,
    "Parent": 111,
    "Name": "node111",
    "Memo": {}			// ������Ϣ
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

### /wapi/group/users

#### ��ѯ�����û��б�

- URI: `/wapi/group/users`

- Method: `GET`

- Parameter: 

  ```json
  {
    "Group": 2			// ����Ⱥ��
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "Username": "user1",
      "Nick": "spock",
      "Role": 5				// ��Ⱥ�Ľ�ɫ
    }, {}]
  }
  ```

#### ���������û�

- URI: `/wapi/group/users`

- Method: `POST`

- Parameter:

  ```json
  {
    "Username": "guest",	// �û�ID
    "Password": "mima",
    "Nick": "kerk",
    "Group": 2,			// ����Ⱥ��ID
    "Role": 2				// ��Ⱥ�Ľ�ɫ
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

#### ɾ�������û�

�ȴ�Ⱥ�����Ƴ��û�������Ƴ����û��������κ�Ⱥ�飬�ٴ�ϵͳ��ɾ���û��˺š�ֻ��ɾ��Ⱥ����Ա�������˺š�

- URI: `/wapi/group/users`

- Method: `DELETE`

- Parameter:

  ```json
  {
    "Group": 5,			// Ⱥ��ID
    "Username": "user1"	// Ҫ�Ƴ����û�ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```


### /wapi/group/role

#### ���������û���ɫ

- URI: `/wapi/group/role`
- Method: `PUT`
- Parameter: 

```json
{
  "Username": "robin",	// �û�ID
  "Group": 5,			// Ⱥ��ID
  "Role": 2				// ��ɫ���
}
```
- Response:

```json
{
  "ErrCode": 0
}
```

### /wapi/group/password

#### ���������û�����

�����޸��������룬�����޸ķ�Ⱥ����Ա�������û�������

- URI: `/wapi/group/password`
- Method: `PUT`
- Parameter: 

```json
{
  "Group": 4,				// �û�����Ⱥ��
  "Username": "robin",		// �û�ID
  "Password": "password"	// ����
}
```

- Response:

```json
{
  "ErrCode": 0
}
```


## �豸����
### �豸ID�������

����`${�豸���ͱ��}_${���к����ʹ���}${���к�}`
ʾ����`W001_P13866886688`

| ���ͱ�� | ˵��            |
| ---- | ------------- |
| W001 | ��ˮ��������ڲ�����001 |
|      |               |

| ���к����ʹ��� | ˵��    |
| ------- | ----- |
| P       | �ֻ���   |
| M       | MAC��ַ |


### /wapi/dev/devices

#### ��ѯ�豸�б�

- URI: `/wapi/dev/devices`

- Method: `GET`

- Parameter:

  ```json
  {
    "Groups": [2],	// �豸����Ⱥ�ڵ�ID�б�
    "Position": false	// �Ƿ񷵻����õ���λ��
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "Group": 2345,				// ����Ⱥ��
    	"Device": "W001_P13866886688",	// �豸ID
      "Name": "���������",		// �豸����
      "Model": "��ˮ�����",		// �豸�ͺţ��ɳ����Զ���
      "BindTime": 12345,			// �豸��ʱ��(UTC)����λ��
      "Memo": {},					// ������Ϣ����ͬӦ���в�ͬ��ʽ�����±�
      "Position": [120.0, 30.5],	// (���ȣ�γ��)
      "PositionTime": 1432344983	// λ�ø���ʱ��(UTC)����λ��
    }, {}]
  }
  ```

#### �����豸
- URI: `/wapi/dev/devices`

- Method: `POST`

- Parameter:

  ```json
  {
    "Group": 2345,					// ����Ⱥ��
    "Device": "W001_P13866886688",	// �豸ID
    "Name": "MachineI",				// �豸����
    "Model": "��ˮ�����",				// �豸�ͺţ��ɳ����Զ���
    "Memo": {}						// ������Ϣ����ͬӦ���в�ͬ��ʽ�����±�
  }
  ```

  | Ӧ������     | ��ע��Ϣ��ʽ                                   | ˵��      |
  | -------- | ---------------------------------------- | ------- |
  | ��ˮ��ˮ����ϵͳ | {"Owner":[{"Name":"����","Phone":"13966990066"}]} | ��¼��������Ϣ |
  |          |                                          |         |

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": null
  }
  ```

#### ɾ���豸
- URI: `/wapi/dev/devices`

- Method: `DELETE`

- Parameter:

  ```json
  {
    "Device": "0_P_15858275538"	// �豸ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": null
  }
  ```

#### �޸��豸��Ϣ

- URI: `/wapi/dev/devices`

- Method: `PUT`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// �豸ID
    "Name": "MachineI",				// �豸����
    "Group": 2345,						// ����Ⱥ��
    "Model": "��ˮ�����",				// �豸�ͺţ��ɳ����Զ���
    "Memo": {}						// ������Ϣ
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
#### �¼���ѯ

- URI: `/wapi/dev/events`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// ��ѯָ���豸
    "Groups": [],				// ��ѯȺ���б��������豸��Device�ֶβ�Ϊ��ʱ��Ч
    "Name": "Motion",			// ��ѯ���¼����ƣ�Ϊ�ձ�ʾ��ѯ�����¼�
    "Action": 1,				// ֻ��ѯָ���Ķ���
    "Limit": 10,				// ��ҳ���ƣ�����ѯ����
    "Offset": 30				// ��ҳ���ƣ��ӵڼ�����ʼ��ѯ
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "ID": 123456,			// �¼���¼ID��
      "Name": "Motion",		// �¼����ƣ�UI�ϵ��¼�ԭ��ݴ���ʾ
      "Action": 0,			// ����յ��Ķ�������
      "Time": [1000, 2000],	// �¼���ֹʱ��
      "Device": "W001_P13866886688",	// �¼��������豸ID
      "Sensor": "1",			// �¼������Ĵ�������ʶ�ţ�����ͬ����ͨ����
      "Process": 0			// �¼������־, 0-δ����1-����
    }, {}]
  }
  ```

#### �¼��ϱ�

- URI: `/wapi/dev/events`

- Method: `POST`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// �豸ID
    "Name":"Motion",			// �¼�����
    "Sensor": "1",			// ��⵽�¼��Ĵ�������ʶ�ţ�����ͬ����ͨ����
    "Action": 1,				// ��������: 0-stop, 1-start, 2-pulse
    "Time": [1000, 2000]		// ��ʼ�ͽ���ʱ�䣬ֻ��stop����Ҫ������ʱ�䣬��λΪ����
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```
- Ȩ��˵������̨ϵͳ�û��ſ��ϱ�

#### �¼������־����

- URI: `/wapi/dev/events`

- Method: `PUT`

- Parameter:

  ```json
  {
    "ID": 12345,			// �¼���¼ID
    "Process": 1			// �¼������־, 0-δ����1-����
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```

### /wapi/dev/logs
#### ��־��ѯ

- URI: `/wapi/dev/logs`

- Method: `GET`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// �豸ID
    "Name": "Log1",			// ��ѯ����־����
    "Time": [1000, 2000]		// ��ֹʱ��(UTC)����λΪ���룬��Ϊ�գ���ʾֻ��ѯ�����¼�
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": [{
      "ID": 123456,		// ��־��¼ID��
    	"Name": "Log1",		// ��־����
      "Time": 1000,		// ��־����ʱ�䣬UTCMS
      "Brief": "xxx",		// ��־��Ҫ��Ϣ
      "Detail": "{}"		// ��־��ϸ��Ϣ
    }, {}]
  }
  ```

#### ��־�ϱ�

- URI: `/wapi/dev/logs`

- Method: `POST`

- Parameter:

  ```json
  {
    "Device": "W001_P13866886688",	// �豸ID
    "Name": "Log1",		// ��־����
    "Time": 1000,			// ��־����ʱ�䣬UTCMS
    "Brief": "xxx",		// ��־��Ҫ��Ϣ
    "Detail": "{}"		// ��־��ϸ��Ϣ
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```
# MQTT ����API
## MQTT ����Ȩ�޿���
### Client ��ʶ

- ��ʽ��`{type}#{id}`

- ˵����

  | type     | id   | ʾ��                  |
  | -------- | ---- | ------------------- |
  | 1-user   | �û�ID | 1#admin             |
  | 2-device | �豸ID | 2#W001_P13866885566 |
  | 3-group  | Ⱥ��ID | 3#5                 |

### ��¼

- �û�����ͬClient��ʶ
- ���룺�豸ʹ��token��Ϊ���룬�û��˺�ʹ������MD5ժҪ֮����Ϊ����

## ������Ϣ

### /bind/request/{group}/{device}

#### ˵��

�豸�˷����İ󶨵�Ⱥ������

####Ȩ��
- ����Ȩ��ӵ���ߣ�����Ⱥ����"device"��Դ��Ȩ�޵��û�
- ����Ȩ��ӵ���ߣ�{device}��ָ�豸
#### ����

��

### /bind/update/{group}/{device}

#### ˵��

�ͻ��˷������豸Ⱥ��󶨹�ϵ���֪ͨ

####Ȩ��
- ����Ȩ��ӵ���ߣ�{device}��ָ�豸���Լ�����Ⱥ����"device"��Դ��Ȩ�޵��û�

- ����Ȩ��ӵ���ߣ�����Ⱥ����"device"��ԴдȨ�޵��û�

#### ����

```json
{
  "Type": "Bind",	// �������: Bind, Unbind, Move
  "From": 12		// ԭ����Ⱥ�飬�������ΪMoveʱ��Ч
}
```


## �豸����Ϣ

### /device/info/{device}

#### ˵��

���豸����������״̬���֪ͨ

####Ȩ��
- ����Ȩ��ӵ���ߣ�����{device}����Ⱥ����"device"��Դ��Ȩ�޵��û�
- ����Ȩ��ӵ���ߣ�{device}��ָ�豸
#### ����

```json
[{
    "Key": "Temperature",
  	"Value": "35"
}, {}]
```

#### ״̬�б�

| Key  | Value Format           | ˵��      |
| ---- | ---------------------- | ------- |
| Gps  | {longitude},{latitude} | ����λ������� |
|      |                        |         |
|      |                        |         |

### /device/query/{device}

#### ˵��

�ɿͻ��˷������豸״̬��ѯ��Ϣ���豸�յ���Ϣ��ظ� /device/info ��Ϣ

####Ȩ��
- ����Ȩ��ӵ���ߣ�{device}��ָ�豸
- ����Ȩ��ӵ���ߣ�����{device}����Ⱥ����"device"��Դ��Ȩ�޵��û�
#### ����

```json
[
  "Temperature", "Humidity"		// ����Ϊ�����ʾ��ѯ��������
]
```

### /device/control/{device}

#### ˵��

�ɿͻ��˷����豸��������

####Ȩ��
- ����Ȩ��ӵ���ߣ�{device}��ָ�豸
- ����Ȩ��ӵ���ߣ�����{device}����Ⱥ����"device"��Դִ��Ȩ�޵��û�
#### ����

```json
[{
  "Command": "Switch1",	// ������
  "Params": ["On"]
}, {}]
```

#### �����б�

| Command | Parameters | ˵��   |
| ------- | ---------- | ---- |
|         |            |      |
|         |            |      |

### /device/event/{device}

#### ˵��

���豸�����ı����¼�

####Ȩ��
- ����Ȩ��ӵ���ߣ�����{device}����Ⱥ����"device"��Դ��Ȩ�޵��û�
- ����Ȩ��ӵ���ߣ�{device}��ָ�豸
#### ����

```json
{
    "Name": "Motion",
  	"Action": 1,			// 0-stop, 1-start, 2-pulse
    "Time": [1000, 2000]	// ��ʼ�ͽ���ʱ�䣬"Start" "Pulse" ʱ���������ʱ�䣬��λΪ����
}
```


## �Ի�����Ϣ

### /talk/{from_id}/{to_id}[/{group}]
#### ˵��

���豸���û����͵ĶԻ���Ϣ��Ⱥ�鲻�ܷ�����Ϣ

| Field    | Example | Comment                       |
| -------- | ------- | ----------------------------- |
| frome_id | 1#admin | ���ͷ� Client ID                 |
| to_id    | 3#5     | ���շ� Client ID                 |
| group    | Ⱥ��      | ���շ�����Ⱥ��ID�����շ�����Ⱥ��ʱ������ϣ�����Ȩ����֤ |

####Ȩ��
- ����Ȩ��ӵ���ߣ����޽����ߣ�Ⱥ���û����ɶ�����Ϣ
- ����Ȩ��ӵ���ߣ��߽���������Ⱥ��鿴������Ȩ��
#### ����

```json
{
    "Type": 1,		// ��Ϣ���ͣ�1-SMS, 2-WebRTC ...
    "Data": {}		// ��Ϣ����
}
```

##### SMS ��Ϣ��ʽ

- Type: `1`
- Data: �����ַ���

##### WebRTC ��Ϣ��ʽ

- Type: `2`
- Data: ͬ WebRTC ʾ�������������ʽ



# ����λ�÷���API

## ������

| �����  | ˵��   |
| ---- | ---- |
|      |      |

## λ�ù���
### /wapi/gps/info

#### ��ѯλ��

- URI: `/wapi/gps/info`

- Method: `GET`

- Parameter:

  ```json
  {
    "Type": 2,			// ʵ������: 1-user, 2-device
    "ID": "W001_P13688660055"	// ʵ��ID
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0,
    "Result": {
      "Position": [120.0, 30.5],	// (���ȣ�γ��)
      "Time": 1432344983			// λ�ø���ʱ��(UTC)����λ��
    }
  }
  ```

#### ����λ��

- URI: `/wapi/gps/info`

- Method: `PUT`

- Parameter:

  ```json
  {
    "Type": 2,				// ʵ������: 1-user, 2-device
    "ID": "W001_P13688660055",	// ʵ��ID
    "Position": [120.0, 30.5]		// (���ȣ�γ��)
  }
  ```

- Response:

  ```json
  {
    "ErrCode": 0
  }
  ```




