<SqlMigrate>
	<Id>38</Id>
	<MigrateName>20191027161250_ALTER_blog.sql</MigrateName>
	<MigrateSql>alter table blog_catalog drop column catalog_source;&#xA;alter table blog_article drop column book_id;&#xA;alter table blog_article drop column book_type;</MigrateSql>
	<MigrateHash>NEQl/H0j4E/ANeSpjVP8hPWq+TH3q3tqcTXKJFhe3hk=</MigrateHash>
	<Effective>true</Effective>
	<CreatedBy>SYSTEM</CreatedBy>
	<CreatedTime>2019-10-27T16:12:50.0259326+08:00</CreatedTime>
	<LastUpdatedBy>SYSTEM</LastUpdatedBy>
	<LastUpdatedTime>2019-10-27T16:12:50.0259326+08:00</LastUpdatedTime>
</SqlMigrate>