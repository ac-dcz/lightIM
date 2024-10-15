use lightim;

drop table if exists user_infos;

create table if not exists user_infos(
    uid bigint not null auto_increment comment "用户id" ,
    nickname varchar(50) not null comment "用户昵称",
    gender enum("M","F") comment "性别",
    tel char(11) not null comment "手机号",
    password char(32) not null comment "密码",
    create_time TIMESTAMP not null default CURRENT_TIMESTAMP comment "用户创建时间",
    update_time TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment "用户更新时间",
    status tinyint unsigned not null default 0 comment "用户状态: 0正常/1注销/2禁用",
    primary key (uid),
    unique key ind_tel (tel)
)ENGINE = InnoDB default charset utf8mb4 COLLATE = utf8mb4_bin COMMENT = '用户表';

