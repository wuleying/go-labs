SET NAMES utf8;
SET FOREIGN_KEY_CHECKS = 0;

CREATE TABLE `calendar` (
  `id` int(10) unsigned NOT NULL AUTO_INCREMENT COMMENT '主键ID',
  `title` varchar(255) NOT NULL DEFAULT '' COMMENT '标题',
  `summary` varchar(4096) NOT NULL DEFAULT '' COMMENT '描述',
  `url` varchar(4096) NOT NULL DEFAULT '' COMMENT '新闻链接',
  `origin_url` varchar(4096) NOT NULL DEFAULT '' COMMENT '来源链接',
  `origin_name` varchar(128) NOT NULL DEFAULT '' COMMENT '来源名称',
  `image_url` varchar(1024) NOT NULL DEFAULT '' COMMENT '新闻图片地址',
  `image_title` varchar(255) NOT NULL DEFAULT '' COMMENT '新闻图片标题',
  `input_date` varchar(32) NOT NULL DEFAULT '' COMMENT '新闻发表时间',
  `insert_date` date NOT NULL DEFAULT '0000-00-00' COMMENT '日期',
  `insert_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '数据写入时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;