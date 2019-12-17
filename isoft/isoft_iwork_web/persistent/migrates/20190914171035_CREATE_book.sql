<SqlMigrate>
	<Id>19</Id>
	<MigrateName>20190914171035_CREATE_book.sql</MigrateName>
	<MigrateSql>CREATE TABLE IF NOT EXISTS `book` (&#xA;`id` int(11) NOT NULL AUTO_INCREMENT,&#xA;`book_name` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;`book_author` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;`created_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;`created_time` datetime NOT NULL,&#xA;`last_updated_by` varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;,&#xA;`last_updated_time` datetime NOT NULL,&#xA;PRIMARY KEY (`id`)&#xA;) ENGINE=InnoDB AUTO_INCREMENT=32 DEFAULT CHARSET=utf8 COLLATE=utf8_bin;</MigrateSql>
	<MigrateHash>mcaiNDd0lgEfp/nIdOaEbYRL81OeTxvyQOOx2Pygvto=</MigrateHash>
	<Effective>true</Effective>
	<CreatedBy>SYSTEM</CreatedBy>
	<CreatedTime>2019-09-14T17:10:35+08:00</CreatedTime>
	<LastUpdatedBy>SYSTEM</LastUpdatedBy>
	<LastUpdatedTime>2019-09-14T17:10:35+08:00</LastUpdatedTime>
</SqlMigrate>