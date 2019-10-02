<SqlMigrate>
	<Id>25</Id>
	<MigrateName>20191002082256_CREATE_GOOD.sql</MigrateName>
	<MigrateSql>CREATE TABLE IF NOT EXISTS `good` (&#xA;  `id` int(11) NOT NULL AUTO_INCREMENT,&#xA;  `good_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  `good_desc` varchar(4000) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  `good_price` decimal(19,2) NOT NULL,&#xA;  `good_images` varchar(4000) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  `good_seller` varchar(400) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  `seller_contact` varchar(400) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  PRIMARY KEY (`id`)&#xA;) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;&#xA;CREATE TABLE IF NOT EXISTS `order_info` (&#xA;  `id` int(11) NOT NULL AUTO_INCREMENT,&#xA;  `good_id` int(11) NOT NULL DEFAULT &#39;-1&#39;,&#xA;  `payment_status` int(11) NOT NULL DEFAULT &#39;-1&#39;,&#xA;  `created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  `created_time` datetime NOT NULL,&#xA;  `last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  `last_updated_time` datetime NOT NULL,&#xA;  `order_code` varchar(100) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;  PRIMARY KEY (`id`)&#xA;) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;</MigrateSql>
	<MigrateHash>DPlwzHlOvGgfe/nRFv+4gJ/j5s9mYcHKeoG10PnBdDE=</MigrateHash>
	<Effective>true</Effective>
	<CreatedBy>SYSTEM</CreatedBy>
	<CreatedTime>2019-10-02T08:22:56+08:00</CreatedTime>
	<LastUpdatedBy>SYSTEM</LastUpdatedBy>
	<LastUpdatedTime>2019-10-02T09:07:30+08:00</LastUpdatedTime>
</SqlMigrate>