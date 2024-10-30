drop table if exists `group`;

create table if not exists `group`(
    gid bigint not null auto_increment comment "group id",
    owner bigint not null comment "创建者",
    group_num char(11) not null comment "群号",
    group_name varchar(50) not null comment "群名",
    `desc`  varchar(255) not null default '' comment  "群简介",
    member_cnts int unsigned not null comment "成员数",
    create_time timestamp not null default CURRENT_TIMESTAMP comment "创建时间",
    update_time timestamp not null default CURRENT_TIMESTAMP on update CURRENT_TIMESTAMP comment "上次跟新时间",
    primary key (gid),
    unique key (group_num),
    foreign key (owner) references user_infos(uid)
)Engine = InnoDB default charset utf8mb4 comment "群";

# alter table `group` add column group_num char(11) not null;
#
# create unique index ind_group_num_unique on `group`(group_num);

# -- 当创建群时，向group_member中插入成员
# create trigger tr_group_insert after insert on `group` for each row
# begin
#     declare id bigint;
#     declare mid bigint;
#     set id = new.gid;
#     set mid = new.owner;
#     insert into group_member(gid,member,type) values(id,mid,2);
# end;
#
# drop trigger tr_group_insert;
#
# insert into `group`(owner,group_num,group_name,member_cnts) values (1,"1331231","test",0);