-- 菜单动作关联实体(menu_action)
DROP TABLE IF EXISTS `menu_action`;
CREATE TABLE `menu_action`  (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `menu_id` varchar(64) NOT NULL DEFAULT '' COMMENT '菜单ID',
  `code` varchar(64) NOT NULL DEFAULT '' COMMENT '动作编号',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '动作名称',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;