drop table if exists group_member;

create table if not exists group_member(
    gid bigint not null comment "群id",
    member bigint not null comment "成员id",
    type tinyint not null  default 0 comment "0： 成员 1：管理员 2：群主",
    create_time TIMESTAMP not null default CURRENT_TIMESTAMP comment "用户创建时间",
    update_time TIMESTAMP not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment "更新时间",
    foreign key (gid) references `group`(gid),
    foreign key (member) references user_infos(uid)
)ENGINE = InnoDB default charset utf8mb4 comment "群成员表";

-- 插入数据时，增加计数
create trigger tr_insert_group_members after insert on group_member for each row
    update `group` set member_cnts = member_cnts + 1 where gid = new.gid;

-- 删除数据时，减少计数
create trigger tr_delete_group_members after delete on group_member for each row
    update `group` set member_cnts = member_cnts - 1 where gid = old.gid;