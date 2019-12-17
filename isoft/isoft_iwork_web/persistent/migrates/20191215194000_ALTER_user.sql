<SqlMigrate>
	<Id>43</Id>
	<MigrateName>20191215194000_ALTER_user.sql</MigrateName>
	<MigrateSql>alter table user add column role_name varchar(255) COLLATE utf8_bin NOT NULL DEFAULT &#39;common&#39; COMMENT &#39;用户角色名称&#39;;&#xA;alter table user add column vip_level int(11) NOT NULL DEFAULT 0 COMMENT &#39;会员等级&#39;;&#xA;alter table user add column vip_expired_time datetime DEFAULT CURRENT_TIMESTAMP COMMENT &#39;会员过期时间&#39;;</MigrateSql>
	<MigrateHash>mTzQXjhwno0tMAiA5Hiywi3tULcVRWKpMcfVVKBVnkA=</MigrateHash>
	<Effective>true</Effective>
	<CreatedBy>SYSTEM</CreatedBy>
	<CreatedTime>2019-12-15T19:40:01+08:00</CreatedTime>
	<LastUpdatedBy>SYSTEM</LastUpdatedBy>
	<LastUpdatedTime>2019-12-15T19:40:01+08:00</LastUpdatedTime>
</SqlMigrate>