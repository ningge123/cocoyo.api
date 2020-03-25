CREATE TABLE `users` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL COMMENT '用户名',
    `email` varchar(255) NOT NULL COMMENT '邮箱',
    `phone` varchar(255) NOT NULL default '' COMMENT '手机号码',
    `password` varchar(255) NOT NULL  COMMENT '密码',
    `avatar` varchar(255) not null default '' COMMENT '头像',
    `gender` enum('male', 'female') not null default 'male' comment '性别',
    `bio` varchar(255) not null default '描述',
    `extends` json default null comment '扩展',
    `settings` json default null comment '设置资料',
    `energy` int(11) not null default 0 comment '',
    `level` int(11) not null default 0 comment '等级',
    `is_admin` boolean not null default false comment '管理员',
    `cache` json default null comment '缓存',
    `last_active_at` timestamp null default null comment '最后登录时间',
    `banned_at` timestamp null default null comment '被禁止时间',
    `activated_at` timestamp null default null comment '活跃时间',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    PRIMARY KEY (`id`),
    UNIQUE KEY `index_unique_username` (`username`),
    UNIQUE KEY `index_unique__email` (`email`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;