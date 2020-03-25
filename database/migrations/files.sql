CREATE TABLE `files` (
    `id` int(11) NOT NULL AUTO_INCREMENT,
    `sha1` char(40) NOT NULL DEFAULT '' COMMENT '文件hash',
    `name` varchar(255) NOT NULL DEFAULT '' COMMENT '文件名',
    `size` bigint(20) DEFAULT '0' COMMENT '文件大小',
    `location` varchar(1024) DEFAULT '' COMMENT '存储路径',
    `created_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
    `updated_at` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新日期',
    `status` int(11) NOT NULL DEFAULT '0' COMMENT '状态 0可用 1禁用 2已删除',
    PRIMARY KEY (`id`),
    UNIQUE KEY `file_index_hash` (`sha1`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;