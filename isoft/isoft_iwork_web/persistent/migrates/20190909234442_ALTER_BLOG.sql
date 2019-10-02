<SqlMigrate>
	<Id>14</Id>
	<MigrateName>20190909234442_ALTER_BLOG.sql</MigrateName>
	<MigrateSql>alter table blog drop  column catalog_id; &#xA;alter table blog add  column catalog_name varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;&#39;;</MigrateSql>
	<MigrateHash>nOjNA9OZnGCyF6q4BS6S1TsUabaY8I8yptXpIBe6cWY=</MigrateHash>
	<Effective>true</Effective>
	<CreatedBy>SYSTEM</CreatedBy>
	<CreatedTime>2019-09-09T23:44:42+08:00</CreatedTime>
	<LastUpdatedBy>SYSTEM</LastUpdatedBy>
	<LastUpdatedTime>2019-09-09T23:44:42+08:00</LastUpdatedTime>
</SqlMigrate>