-- 创建用户表 用户实体(user)
CREATE TABLE user (
  user_id INT(11) AUTO_INCREMENT NOT NULL COMMENT '类型id',
  record_id char(50) NOT NULL DEFAULT '' COMMENT '用户的记录id',
  user_name char(50) NOT NULL DEFAULT '' COMMENT '用户名称',
  password char(50) NOT NULL DEFAULT '' COMMENT '用户的密码',
  email char(50) NOT NULL DEFAULT '' COMMENT '用户的email',
  phone char(12) NOT NULL DEFAULT 0 COMMENT '用户联系方式',
  profile text NOT NULL DEFAULT '' COMMENT '用户属性备注',
  create_at datetime DEFAULT CURRENT_TIMESTAMP  COMMENT '用户的注册时间',
  last_active datetime DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP COMMENT '最后活跃时间戳',
  status int(11) NOT NULL DEFAULT '0' COMMENT '账户状态(启用/禁用/锁定/标记删除等)',
  PRIMARY KEY (`user_id`),
  KEY `idx_status` (`status`)
) ENGINE = InnoDB DEFAULT CHARSET =utf8;