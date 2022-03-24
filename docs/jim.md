| DATABASE | CHARACTER SET |    COLLATION    |
|----------|---------------|-----------------|
| jim      | utf8          | utf8_general_ci |

### `device`
设备

|     COLUMN     |    DATA TYPE    | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |  COLLATION  |                 COMMENT                 |
|----------------|-----------------|----------|-----|-------------------|---------------|-------------|-----------------------------------------|
| app_id         | bigint unsigned | NO       | MUL |                   |               |             | app_id                                  |
| brand          | varchar(20)     | NO       |     |                   | utf8mb4       | utf8mb4_bin | 手机厂商                                |
| conn_addr      | varchar(25)     | NO       |     |                   | utf8mb4       | utf8mb4_bin | 连接层服务器地址                        |
| conn_fd        | bigint          | NO       |     |                   |               |             | TCP连接对应的文件描述符                 |
| create_time    | datetime        | NO       |     | CURRENT_TIMESTAMP |               |             | 创建时间                                |
| device_id      | bigint          | NO       | UNI |                   |               |             | 设备id                                  |
| id             | bigint unsigned | NO       | PRI |                   |               |             | 自增主键                                |
| model          | varchar(20)     | NO       |     |                   | utf8mb4       | utf8mb4_bin | 机型                                    |
| sdk_version    | varchar(10)     | NO       |     |                   | utf8mb4       | utf8mb4_bin | app版本                                 |
| status         | tinyint         | NO       |     |                 0 |               |             | 在线状态，0：离线；1：在线              |
| system_version | varchar(10)     | NO       |     |                   | utf8mb4       | utf8mb4_bin | 系统版本                                |
| type           | tinyint         | NO       |     |                   |               |             | 设备类型,1:Android；2：IOS；3：Windows; |
|                |                 |          |     |                   |               |             | 4：MacOS；5：Web                        |
| update_time    | datetime        | NO       |     | CURRENT_TIMESTAMP |               |             | 更新时间                                |
| user_id        | bigint unsigned | NO       |     |                 0 |               |             | 账户id                                  |

### `device_ack`
设备消息同步序列号

|   COLUMN    |    DATA TYPE    | NULLABLE | KEY |      DEFAULT      | CHARACTER SET | COLLATION |    COMMENT     |
|-------------|-----------------|----------|-----|-------------------|---------------|-----------|----------------|
| ack         | bigint unsigned | NO       |     |                 0 |               |           | 收到消息确认号 |
| create_time | datetime        | NO       |     | CURRENT_TIMESTAMP |               |           | 创建时间       |
| device_id   | bigint unsigned | NO       | UNI |                   |               |           | 设备id         |
| id          | bigint unsigned | NO       | PRI |                   |               |           | 自增主键       |
| update_time | datetime        | NO       |     | CURRENT_TIMESTAMP |               |           | 更新时间       |

### `emp`


| COLUMN |  DATA TYPE   | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      | COMMENT |
|--------|--------------|----------|-----|---------|---------------|--------------------|---------|
| age    | int          | NO       | MUL |         |               |                    |         |
| id     | int          | NO       | PRI |         |               |                    |         |
| name   | varchar(100) | NO       |     |         | utf8mb4       | utf8mb4_0900_ai_ci |         |
| salary | int          | YES      |     |         |               |                    |         |

### `feed`
动态

|    COLUMN     |    DATA TYPE     | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      |            COMMENT             |
|---------------|------------------|----------|-----|---------|---------------|--------------------|--------------------------------|
| comment_count | int unsigned     | NO       |     |       0 |               |                    | 动态评论数                     |
| content       | text             | NO       |     |         | utf8mb4       | utf8mb4_unicode_ci | 动态内容                       |
| created_at    | datetime         | YES      |     |         |               |                    |                                |
| deleted_at    | datetime         | YES      |     |         |               |                    |                                |
| hot           | bigint unsigned  | NO       |     |       0 |               |                    | 热门排序值                     |
| id            | bigint unsigned  | NO       | PRI |       0 |               |                    | 动态 ID                        |
| is_enable     | tinyint unsigned | NO       |     |       1 |               |                    | 是否启用                       |
| like_count    | int unsigned     | NO       |     |       0 |               |                    | 动态点赞数                     |
| operator      | bigint unsigned  | NO       |     |       0 |               |                    | 审核人                         |
| remark        | varchar(500)     | NO       |     |         | utf8mb4       | utf8mb4_unicode_ci | 备注                           |
| review_status | tinyint unsigned | NO       |     |       0 |               |                    | 审核状态 0-未审核 1-已审核     |
|               |                  |          |     |         |               |                    | 2-未通过                       |
| type          | tinyint          | NO       |     |       0 |               |                    | 动态类型                       |
| updated_at    | datetime         | YES      |     |         |               |                    |                                |
| user_id       | bigint unsigned  | NO       | MUL |       0 |               |                    | 用户ID                         |
| view_count    | int unsigned     | NO       |     |       0 |               |                    | 动态阅读数                     |

### `feed_image`
动态图片

|   COLUMN   |    DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      | COMMENT |
|------------|-----------------|----------|-----|---------|---------------|--------------------|---------|
| created_at | datetime        | YES      |     |         |               |                    |         |
| deleted_at | datetime        | YES      |     |         |               |                    |         |
| feed_id    | bigint unsigned | NO       | MUL |       0 |               |                    | 动态ID  |
| id         | bigint unsigned | NO       | PRI |       0 |               |                    | ID      |
| img_url    | varchar(2000)   | NO       |     |         | utf8mb4       | utf8mb4_unicode_ci | 图片URL |
| updated_at | datetime        | YES      |     |         |               |                    |         |

### `feed_like`
动态点赞

|   COLUMN   |    DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET | COLLATION | COMMENT |
|------------|-----------------|----------|-----|---------|---------------|-----------|---------|
| created_at | datetime        | YES      |     |         |               |           |         |
| deleted_at | datetime        | YES      |     |         |               |           |         |
| feed_id    | bigint unsigned | NO       | MUL |       0 |               |           | 动态ID  |
| id         | bigint unsigned | NO       | PRI |       0 |               |           | ID      |
| updated_at | datetime        | YES      |     |         |               |           |         |
| user_id    | bigint unsigned | NO       |     |       0 |               |           | 用户ID  |

### `feed_video`
动态视频

|   COLUMN   |   DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      |   COMMENT   |
|------------|----------------|----------|-----|---------|---------------|--------------------|-------------|
| cover_url  | varchar(2048)  | NO       |     |         | utf8mb4       | utf8mb4_unicode_ci | 视频封面URL |
| created_at | datetime       | YES      |     |         |               |                    |             |
| deleted_at | datetime       | YES      |     |         |               |                    |             |
| duration   | float unsigned | NO       |     |       0 |               |                    | 视频时长    |
| feed_id    | int unsigned   | NO       | MUL |       0 |               |                    | 所属动态id  |
| height     | int unsigned   | NO       |     |       0 |               |                    | 视频高度    |
| id         | int unsigned   | NO       | PRI |         |               |                    |             |
| updated_at | datetime       | YES      |     |         |               |                    |             |
| video_url  | varchar(2048)  | NO       |     |         | utf8mb4       | utf8mb4_unicode_ci | 视频URL     |
| width      | int unsigned   | NO       |     |       0 |               |                    | 视频宽度    |

### `gid`
分布式自增主键

|   COLUMN    |    DATA TYPE    | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |  COLLATION  | COMMENT  |
|-------------|-----------------|----------|-----|-------------------|---------------|-------------|----------|
| business_id | varchar(128)    | NO       | UNI |                   | utf8mb4       | utf8mb4_bin | 业务id   |
| create_time | datetime        | NO       |     | CURRENT_TIMESTAMP |               |             | 创建时间 |
| description | varchar(255)    | NO       |     |                   | utf8mb4       | utf8mb4_bin | 描述     |
| id          | bigint unsigned | NO       | PRI |                   |               |             | 自增主键 |
| max_id      | bigint unsigned | NO       |     |                 0 |               |             | 最大id   |
| step        | int unsigned    | NO       |     |              1000 |               |             | 步长     |
| update_time | datetime        | NO       |     | CURRENT_TIMESTAMP |               |             | 更新时间 |

### `goadmin_menu`


|   COLUMN    |    DATA TYPE     | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |     COLLATION      | COMMENT |
|-------------|------------------|----------|-----|-------------------|---------------|--------------------|---------|
| created_at  | timestamp        | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| header      | varchar(150)     | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| icon        | varchar(50)      | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| id          | int unsigned     | NO       | PRI |                   |               |                    |         |
| order       | int unsigned     | NO       |     |                 0 |               |                    |         |
| parent_id   | int unsigned     | NO       |     |                 0 |               |                    |         |
| plugin_name | varchar(150)     | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| title       | varchar(50)      | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| type        | tinyint unsigned | NO       |     |                 0 |               |                    |         |
| updated_at  | timestamp        | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| uri         | varchar(3000)    | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| uuid        | varchar(150)     | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |

### `goadmin_operation_log`


|   COLUMN   |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |     COLLATION      | COMMENT |
|------------|--------------|----------|-----|-------------------|---------------|--------------------|---------|
| created_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| id         | int unsigned | NO       | PRI |                   |               |                    |         |
| input      | text         | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| ip         | varchar(15)  | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| method     | varchar(10)  | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| path       | varchar(255) | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| updated_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| user_id    | int unsigned | NO       | MUL |                   |               |                    |         |

### `goadmin_permissions`


|   COLUMN    |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |     COLLATION      | COMMENT |
|-------------|--------------|----------|-----|-------------------|---------------|--------------------|---------|
| created_at  | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| http_method | varchar(255) | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| http_path   | text         | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| id          | int unsigned | NO       | PRI |                   |               |                    |         |
| name        | varchar(50)  | NO       | UNI |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| slug        | varchar(50)  | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| updated_at  | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |

### `goadmin_role_menu`


|   COLUMN   |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET | COLLATION | COMMENT |
|------------|--------------|----------|-----|-------------------|---------------|-----------|---------|
| created_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |
| menu_id    | int unsigned | NO       |     |                   |               |           |         |
| role_id    | int unsigned | NO       | MUL |                   |               |           |         |
| updated_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |

### `goadmin_role_permissions`


|    COLUMN     |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET | COLLATION | COMMENT |
|---------------|--------------|----------|-----|-------------------|---------------|-----------|---------|
| created_at    | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |
| permission_id | int unsigned | NO       | PRI |                   |               |           |         |
| role_id       | int unsigned | NO       | PRI |                   |               |           |         |
| updated_at    | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |

### `goadmin_role_users`


|   COLUMN   |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET | COLLATION | COMMENT |
|------------|--------------|----------|-----|-------------------|---------------|-----------|---------|
| created_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |
| role_id    | int unsigned | NO       | PRI |                   |               |           |         |
| updated_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |
| user_id    | int unsigned | NO       | PRI |                   |               |           |         |

### `goadmin_roles`


|   COLUMN   |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |     COLLATION      | COMMENT |
|------------|--------------|----------|-----|-------------------|---------------|--------------------|---------|
| created_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| id         | int unsigned | NO       | PRI |                   |               |                    |         |
| name       | varchar(50)  | NO       | UNI |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| slug       | varchar(50)  | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| updated_at | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |

### `goadmin_session`


|   COLUMN   |   DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |     COLLATION      | COMMENT |
|------------|---------------|----------|-----|-------------------|---------------|--------------------|---------|
| created_at | timestamp     | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| id         | int unsigned  | NO       | PRI |                   |               |                    |         |
| sid        | varchar(50)   | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| updated_at | timestamp     | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| values     | varchar(3000) | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |

### `goadmin_site`


|   COLUMN    |    DATA TYPE     | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |     COLLATION      | COMMENT |
|-------------|------------------|----------|-----|-------------------|---------------|--------------------|---------|
| created_at  | timestamp        | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| description | varchar(3000)    | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| id          | int unsigned     | NO       | PRI |                   |               |                    |         |
| key         | varchar(100)     | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| state       | tinyint unsigned | NO       |     |                 0 |               |                    |         |
| updated_at  | timestamp        | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| value       | longtext         | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |

### `goadmin_user_permissions`


|    COLUMN     |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET | COLLATION | COMMENT |
|---------------|--------------|----------|-----|-------------------|---------------|-----------|---------|
| created_at    | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |
| permission_id | int unsigned | NO       | PRI |                   |               |           |         |
| updated_at    | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |           |         |
| user_id       | int unsigned | NO       | PRI |                   |               |           |         |

### `goadmin_users`


|     COLUMN     |  DATA TYPE   | NULLABLE | KEY |      DEFAULT      | CHARACTER SET |     COLLATION      | COMMENT |
|----------------|--------------|----------|-----|-------------------|---------------|--------------------|---------|
| avatar         | varchar(255) | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| created_at     | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| id             | int unsigned | NO       | PRI |                   |               |                    |         |
| name           | varchar(100) | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| password       | varchar(100) | NO       |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| remember_token | varchar(100) | YES      |     |                   | utf8mb4       | utf8mb4_unicode_ci |         |
| updated_at     | timestamp    | YES      |     | CURRENT_TIMESTAMP |               |                    |         |
| username       | varchar(100) | NO       | UNI |                   | utf8mb4       | utf8mb4_unicode_ci |         |

### `group`
群组

|    COLUMN    |    DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET |      COLLATION      | COMMENT  |
|--------------|-----------------|----------|-----|---------|---------------|---------------------|----------|
| atavar_url   | varchar(1024)   | NO       |     |         | utf8mb4       | utf8mb4_croatian_ci | 头像     |
| created_at   | datetime        | YES      |     |         |               |                     | 创建时间 |
| deleted_at   | datetime        | YES      |     |         |               |                     | 删除时间 |
| extra        | varchar(1024)   | NO       |     |         | utf8mb4       | utf8mb4_croatian_ci | 附加属性 |
| id           | bigint unsigned | NO       | PRI |         |               |                     | 自增主键 |
| introduction | varchar(255)    | NO       |     |         | utf8mb4       | utf8mb4_croatian_ci | 群组简介 |
| name         | varchar(50)     | NO       |     |         | utf8mb4       | utf8mb4_croatian_ci | 群组名称 |
| updated_at   | datetime        | YES      |     |         |               |                     | 更新时间 |
| user_id      | bigint          | NO       | MUL |       0 |               |                     | 群主ID   |

### `group_user`
群组成员关系

|     COLUMN     |    DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET |  COLLATION  |     COMMENT      |
|----------------|-----------------|----------|-----|---------|---------------|-------------|------------------|
| created_at     | datetime        | YES      |     |         |               |             | 创建时间         |
| deleted_at     | datetime        | YES      | MUL |         |               |             | 删除时间         |
| extra          | varchar(1024)   | NO       |     |         | utf8mb4       | utf8mb4_bin | 附加属性         |
| group_id       | bigint unsigned | NO       | MUL |         |               |             | 组id             |
| id             | bigint unsigned | NO       | PRI |         |               |             | 自增主键         |
| updated_at     | datetime        | YES      |     |         |               |             | 更新时间         |
| user_id        | bigint unsigned | NO       | MUL |         |               |             | 用户id           |
| user_show_name | varchar(20)     | NO       |     |         | utf8mb4       | utf8mb4_bin | 用户在群组的昵称 |

### `message`
消息

|    COLUMN     |    DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      |                   COMMENT                    |
|---------------|-----------------|----------|-----|---------|---------------|--------------------|----------------------------------------------|
| at_user_id    | varchar(255)    | NO       |     |         | utf8mb4       | utf8mb4_0900_ai_ci | 需要@的用户，多个用,分割                     |
| created_at    | datetime        | YES      |     |         |               |                    | 创建时间                                     |
| data          | text            | NO       |     |         | utf8mb4       | utf8mb4_0900_ai_ci | 内容                                         |
| deleted_at    | datetime        | YES      |     |         |               |                    | 删除时间                                     |
| id            | bigint unsigned | NO       | PRI |         |               |                    |                                              |
| message_type  | int             | NO       |     |       0 |               |                    | 消息类型                                     |
| receiver_id   | bigint unsigned | NO       |     |       0 |               |                    | 接收人ID，单聊则为user_id,群聊则为group_id   |
| receiver_type | tinyint         | NO       |     |       0 |               |                    | 接收人类型                                   |
| sender_id     | bigint unsigned | NO       |     |       0 |               |                    | 发送人ID                                     |
| sender_type   | tinyint         | NO       |     |       0 |               |                    | 发送人类型                                   |
| status        | tinyint         | NO       |     |       1 |               |                    | 状态：1-待送达，2-已送到，3-已确认，0-已删除 |
| updated_at    | datetime        | YES      |     |         |               |                    | 更新时间                                     |

### `students`


|  COLUMN   |    DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET |    COLLATION    | COMMENT |
|-----------|-----------------|----------|-----|---------|---------------|-----------------|---------|
| age       | int             | YES      |     |         |               |                 |         |
| id        | bigint unsigned | NO       | PRI |         |               |                 |         |
| name      | varchar(16)     | YES      | MUL |         | utf8          | utf8_general_ci |         |
| school_id | int             | YES      |     |         |               |                 |         |

### `tb_seller`


|   COLUMN   |  DATA TYPE   | NULLABLE | KEY | DEFAULT | CHARACTER SET |     COLLATION      | COMMENT |
|------------|--------------|----------|-----|---------|---------------|--------------------|---------|
| address    | varchar(100) | YES      |     |         | utf8mb4       | utf8mb4_0900_ai_ci |         |
| createtime | datetime     | YES      |     |         |               |                    |         |
| name       | varchar(100) | YES      | MUL |         | utf8mb4       | utf8mb4_0900_ai_ci |         |
| nickname   | varchar(50)  | YES      |     |         | utf8mb4       | utf8mb4_0900_ai_ci |         |
| password   | varchar(60)  | YES      |     |         | utf8mb4       | utf8mb4_0900_ai_ci |         |
| sellerid   | varchar(100) | NO       | PRI |         | utf8mb4       | utf8mb4_0900_ai_ci |         |
| status     | varchar(1)   | YES      |     |         | utf8mb4       | utf8mb4_0900_ai_ci |         |

### `user`
用户

|   COLUMN   |    DATA TYPE    | NULLABLE | KEY | DEFAULT | CHARACTER SET |  COLLATION  |         COMMENT          |
|------------|-----------------|----------|-----|---------|---------------|-------------|--------------------------|
| avatar_url | varchar(1024)   | NO       |     |         | utf8mb4       | utf8mb4_bin | 用户头像链接             |
| created_at | datetime        | YES      | MUL |         |               |             | 创建时间                 |
| deleted_at | datetime        | YES      | MUL |         |               |             | 删除时间                 |
| extra      | varchar(1024)   | NO       |     |         | utf8mb4       | utf8mb4_bin | 附加属性                 |
| gender     | enum('1','2')   | NO       |     |         | utf8mb4       | utf8mb4_bin | 性别，0:未知；1:男；2:女 |
| id         | bigint unsigned | NO       | PRI |         |               |             | 自增主键                 |
| nickname   | varchar(50)     | NO       |     |         | utf8mb4       | utf8mb4_bin | 昵称                     |
| password   | varchar(100)    | NO       |     |         | utf8mb4       | utf8mb4_bin | 密码                     |
| updated_at | datetime        | YES      |     |         |               |             | 更新时间                 |

### `users`


| COLUMN |   DATA TYPE   | NULLABLE | KEY | DEFAULT | CHARACTER SET |    COLLATION    | COMMENT |
|--------|---------------|----------|-----|---------|---------------|-----------------|---------|
| age    | int           | NO       |     |       0 |               |                 |         |
| gender | int           | NO       |     |       0 |               |                 |         |
| id     | int unsigned  | NO       | PRI |         |               |                 |         |
| name   | varchar(50)   | NO       | MUL |         | utf8          | utf8_general_ci |         |
| remark | varchar(5000) | YES      |     |       0 | utf8          | utf8_general_ci |         |
| status | int           | NO       |     |       0 |               |                 |         |

### `users1`


| COLUMN |   DATA TYPE   | NULLABLE | KEY | DEFAULT | CHARACTER SET |    COLLATION    | COMMENT |
|--------|---------------|----------|-----|---------|---------------|-----------------|---------|
| age    | int           | NO       |     |       0 |               |                 |         |
| gender | int           | NO       | MUL |       0 |               |                 |         |
| id     | int unsigned  | NO       | PRI |         |               |                 |         |
| name   | varchar(50)   | NO       | MUL |         | utf8          | utf8_general_ci |         |
| remark | varchar(5000) | YES      |     |       0 | utf8          | utf8_general_ci |         |
| status | int           | NO       |     |       0 |               |                 |         |

### `users1_role`


|  COLUMN   | DATA TYPE | NULLABLE | KEY | DEFAULT | CHARACTER SET | COLLATION | COMMENT |
|-----------|-----------|----------|-----|---------|---------------|-----------|---------|
| id        | int       | NO       | PRI |         |               |           |         |
| role_id   | int       | NO       | MUL |         |               |           |         |
| users1_id | int       | NO       |     |         |               |           |         |

