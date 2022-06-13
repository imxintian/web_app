select * from user;

show create table user;
CREATE TABLE `user` (
                        `id` bigint NOT NULL AUTO_INCREMENT,
                        `user_id` bigint NOT NULL,
                        `username` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `password` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci NOT NULL,
                        `email` varchar(64) CHARACTER SET utf8mb4 COLLATE utf8mb4_general_ci  default NULL,
                        `gender` tinyint NOT NULL DEFAULT '0',
                        `create_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP,
                        `update_time` timestamp NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
                        PRIMARY KEY (`id`),
                        UNIQUE KEY `idx_username` (`username`) USING BTREE,
                        UNIQUE KEY `idx_user_id` (`user_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;




select * from user;


create table community (
    id int(11) not null auto_increment,
    community_id int(10) unsigned not null,
    community_name varchar(128) collate utf8mb4_general_ci not null,
    introduction varchar(256) collate utf8mb4_general_ci not null,
    create_time timestamp not null default current_timestamp,
    update_time timestamp not null default current_timestamp on update current_timestamp,
    primary key (id),
    unique key idx_community_id (community_id),
    unique key idx_community_name (community_name)
) engine=innodb default charset=utf8mb4 collate=utf8mb4_general_ci;

insert into community (community_id, community_name, introduction) values (1, 'test', 'test');
insert into community (community_id, community_name, introduction) values (2, 'go', 'golang');
insert into community (community_id, community_name, introduction) values (3, 'python', 'python');
insert into community (community_id, community_name, introduction) values (4, 'java', 'java');

select * from community;
