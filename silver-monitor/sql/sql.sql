CREATE TABLE `log` (
  `id` int(11) NOT NULL AUTO_INCREMENT COMMENT '主键id',
  `price_bid` decimal(10,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '买入价',
  `price_sell` decimal(10,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '卖出价',
  `price_middle` decimal(10,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '中间价',
  `price_middle_high` decimal(10,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '最高中间价',
  `price_middle_low` decimal(10,3) unsigned NOT NULL DEFAULT '0.000' COMMENT '最低中间价',
  `insert_time` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00' COMMENT '插入时间',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;