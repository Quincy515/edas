-- 菜单实体(menu)
DROP TABLE IF EXISTS `menu`;
CREATE TABLE `menu`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `record_id` varchar(64) NOT NULL DEFAULT '' COMMENT '记录ID',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单名称',
  `sequence` int(10) NOT NULL DEFAULT '0' COMMENT '排序值',
  `icon` varchar(64) NOT NULL DEFAULT '' COMMENT '图标',
  `router` varchar(64) NOT NULL DEFAULT '' COMMENT '访问路由',
  `hidden` int(10) NOT NULL DEFAULT '0' COMMENT '0：不隐藏 1:隐藏',
  `parent_id` varchar(64) NOT NULL DEFAULT '' COMMENT '父级ID',
  `parent_path` varchar(64) NOT NULL DEFAULT '' COMMENT '父级路径',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP  COMMENT '创建时间',
  `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP  COMMENT '删除时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`),
  UNIQUE KEY `idx_name` (`name`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;