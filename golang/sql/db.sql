


CREATE TABLE `oe_meta` (
                            `id` bigint(24) NOT NULL AUTO_INCREMENT COMMENT '主键',
                            `file_Id` varchar(36) Not NULL COMMENT '文档Id',
                            `file_name` varchar(128) Not NULL COMMENT '管理员账户',

                            `cell` varbinary(128) DEFAULT NULL COMMENT '文件方格信息',
                            `data` varbinary(1024) DEFAULT NULL COMMENT '文件数据',
                            `row_data` varbinary(256) DEFAULT NULL COMMENT '行（标题）数据',
                            `version` varchar(128) DEFAULT NULL COMMENT '版本',

                            `author` varchar(100) DEFAULT NULL COMMENT '创建人',
                            `create_stime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP COMMENT '行插入时间',
                            `modifier` varchar(100) DEFAULT NULL COMMENT '修改人',
                            `update_stime` datetime NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '行修改时间',
                            PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='ES管理员表';