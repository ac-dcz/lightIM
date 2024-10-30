drop table if exists relation_ship;

create table if not exists relation_ship(
    rid bigint not null auto_increment comment "关系id",
    uid_1 bigint not null comment "user-1",
    uid_2 bigint not null comment "user-2",
    status tinyint not null default 0 comment "好友状态(0-normal,1-(1->2 拉黑),2-(1<-2 拉黑))",
    create_time TIMESTAMP not null default CURRENT_TIMESTAMP comment "用户创建时间",
    update_time TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment "用户更新时间",
    primary key (rid),
    foreign key (uid_1) references user_infos(uid),
    foreign key (uid_2) references user_infos(uid),
    unique key (uid_1,uid_2)
)Engine = InnoDB default charset utf8mb4 comment "用户关系表"