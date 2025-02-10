-- 创建数据库（如果不存在）
CREATE DATABASE IF NOT EXISTS `go-vue` DEFAULT CHARACTER SET utf8mb4 COLLATE utf8mb4_0900_ai_ci;

USE `go-vue`;

-- 创建表（如果不存在）
CREATE TABLE IF NOT EXISTS `sys_menu` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `parent_id` INT UNSIGNED DEFAULT NULL COMMENT '父菜单id',
  `menu_name` VARCHAR(100) NOT NULL COMMENT '菜单名称',
  `icon` VARCHAR(100) DEFAULT NULL COMMENT '菜单图标',
  `value` VARCHAR(100) DEFAULT NULL COMMENT '权限值',
  `menu_type` INT UNSIGNED NOT NULL COMMENT '菜单类型: 1->目录: 2->菜单: 3->按钮',
  `url` VARCHAR(100) DEFAULT NULL COMMENT '菜单url',
  `menu_status` INT UNSIGNED NOT NULL COMMENT '启用状态: 1->启用: 2->禁用',
  `sort` INT UNSIGNED NOT NULL DEFAULT '0' COMMENT '排序',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_0900_ai_ci COMMENT='系统菜单表';

CREATE TABLE IF NOT EXISTS `sys_role_menu` (
  `role_id` INT UNSIGNED NOT NULL COMMENT '角色id',
  `menu_id` INT UNSIGNED NOT NULL COMMENT '菜单id',
  PRIMARY KEY (`role_id`, `menu_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COMMENT='角色菜单关联表';

CREATE TABLE IF NOT EXISTS `sys_role` (
  `id` INT UNSIGNED NOT NULL AUTO_INCREMENT COMMENT '主键',
  `role_name` VARCHAR(64) NOT NULL COMMENT '角色名称',
  `role_key` VARCHAR(64) NOT NULL COMMENT '权限字符串',
  `status` INT NOT NULL DEFAULT 1 COMMENT '账号启动状态: 1->启用,2->禁用',
  `description` VARCHAR(500) DEFAULT NULL COMMENT '描述',
  `create_time` DATETIME NOT NULL COMMENT '创建时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='系统角色表';