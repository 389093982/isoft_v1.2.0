<SqlMigrate>
	<Id>41</Id>
	<MigrateName>20191214115322_ALTER_user.sql</MigrateName>
	<MigrateSql>alter table user add column verify_code int(11) NOT NULL DEFAULT -1 COMMENT &#39;验证码&#39;;&#xA;alter table user add column verify_code_expired datetime DEFAULT CURRENT_TIMESTAMP COMMENT &#39;验证码过期时间&#39;;</MigrateSql>
	<MigrateHash>1uZE8VxFbAFHI8jeHygyNViNCYb00EefQnZkhf2cQD8=</MigrateHash>
	<Effective>true</Effective>
	<CreatedBy>SYSTEM</CreatedBy>
	<CreatedTime>2019-12-14T11:53:23+08:00</CreatedTime>
	<LastUpdatedBy>SYSTEM</LastUpdatedBy>
	<LastUpdatedTime>2019-12-14T11:53:23+08:00</LastUpdatedTime>
</SqlMigrate>