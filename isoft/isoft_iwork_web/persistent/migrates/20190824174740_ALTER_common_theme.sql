<SqlMigrate>
	<Id>6</Id>
	<MigrateName>20190824174740_ALTER_common_theme.sql</MigrateName>
	<MigrateSql>alter table comment_theme change comment_id theme_pk int(11) NOT NULL DEFAULT &#39;0&#39;;&#xA;alter table comment_theme change comment_content theme_desc  varchar(4000) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;;</MigrateSql>
	<MigrateHash>4UMklc0Th81Km8hiJhR4w7VN4xptuvri142Kf2VXUkc=</MigrateHash>
	<Effective>true</Effective>
	<CreatedBy>SYSTEM</CreatedBy>
	<CreatedTime>2019-08-29T22:25:59+08:00</CreatedTime>
	<LastUpdatedBy>SYSTEM</LastUpdatedBy>
	<LastUpdatedTime>2019-08-29T22:25:59+08:00</LastUpdatedTime>
</SqlMigrate>