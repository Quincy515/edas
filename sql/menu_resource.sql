-- 菜单资源关联实体(menu_resource)
DROP TABLE IF EXISTS `menu_resource`;
CREATE TABLE `menu_resource` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_id` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单ID',
  `code` varchar(64) NOT NULL DEFAULT '' COMMENT '资源编号',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '资源名称',
  `method` varchar(64) NOT NULL DEFAULT '' COMMENT '资源请求方式',
  `path` varchar(64) NOT NULL DEFAULT '' COMMENT '资源请求路径',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;