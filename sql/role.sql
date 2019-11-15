-- 角色实体(role)
DROP TABLE IF EXISTS `role`;
CREATE TABLE `role` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `record_id` varchar(64) NOT NULL DEFAULT '' COMMENT '记录ID',
  `name` varchar(64) NOT NULL DEFAULT '' COMMENT '角色名称',
  `sequence` int(11) NOT NULL DEFAULT '0' COMMENT '排序值',
  `memo` varchar(64) NOT NULL DEFAULT '' COMMENT '备注',
  `creator` varchar(64) NOT NULL DEFAULT '' COMMENT '创建人',
  `created_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '创建时间',
  `deleted_at` datetime DEFAULT CURRENT_TIMESTAMP COMMENT '删除时间',
  `updated_at` datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '更新时间',
  PRIMARY KEY (`id`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;