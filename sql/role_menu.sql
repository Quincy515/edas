-- 角色菜单关联实体(role_menu)
DROP TABLE IF EXISTS `role_menu`;
CREATE TABLE `role_menu` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `role_id` varchar(64) NOT NULL DEFAULT '' COMMENT '角色ID',
  `menu_id` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单ID',
  `action` varchar(255) NOT NULL DEFAULT '' COMMENT '动作权限 动作编号(多个以英文逗号分隔)',
  `resource` varchar(255) NOT NULL DEFAULT '' COMMENT '资源权限 资源编号(多个以英文逗号分隔)',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;