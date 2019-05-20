const default_work_step_types=[
  {"name":"work_start","icon":"ios-arrow-dropright"},
  {"name":"work_end","icon":"ios-arrow-dropleft"},
  {"name":"empty","icon":"ios-mail-open-outline"},
  {"name":"work_sub","icon":"logo-buffer"},
  // 数据库相关
  {"name":"sql_query","icon":"ios-cube-outline"},
  {"name":"sql_query_page","icon":"md-cube"},
  {"name":"sql_execute","icon":"ios-crop-outline"},
  {"name":"db_parser","icon":"ios-map-outline"},
  {"name":"json_render","icon":"ios-git-branch"},
  {"name":"json_parser","icon":"ios-git-compare"},
  {"name":"http_request","icon":"ios-globe-outline"},
  {"name":"mapper","icon":"ios-infinite"},
  {"name":"file_read","icon":"ios-book-outline"},
  {"name":"file_write","icon":"ios-create-outline"},
  {"name":"file_delete","icon":"ios-log-out"},
  {"name":"file_sync","icon":"md-paper"},
  {"name":"href_parser","icon":"ios-ionitron-outline"},
  {"name":"entity_parser","icon":"ios-refresh-circle-outline"},
  {"name":"memorymap_cache","icon":"ios-color-filter-outline"},
  {"name":"set_env","icon":"ios-nuclear-outline"},
  {"name":"get_env","icon":"ios-nuclear"},
  {"name":"cal_hash","icon":"ios-flower-outline"},
  {"name":"run_cmd","icon":"md-bonfire"},
  {"name":"sftp_upload","icon":"md-arrow-up"},
  {"name":"ssh_shell","icon":"ios-cloud-upload-outline"},
  {"name":"targz_uncompress","icon":"ios-aperture"},
  {"name":"targz_compress","icon":"ios-aperture-outline"},
  {"name":"ini_read","icon":"ios-fastforward"},
  {"name":"ini_write","icon":"ios-aperture-outline"},
  {"name":"base64_encode","icon":"ios-magnet"},
  {"name":"base64_decode","icon":"ios-magnet-outline"},
  {"name":"if","icon":"md-code-working"},
  {"name":"elif","icon":"md-code-working"},
  {"name":"else","icon":"md-code-working"},
  {"name":"define_var","icon":"md-hammer"},
  {"name":"assign_var","icon":"md-hammer"},
  {"name":"map","icon":"md-hammer"},
  {"name":"foreach","icon":"md-hammer"},
];

const mysql_datatypes = ["varchar","char","text","float","int","date","datetime","decimal","double",
  "bigint","binary","bit","blob","bool","boolean","enum","longblob","longtext","mediumblob","mediumint",
  "mediumtext","numeric","real","set","smallint","time","timestamp","tinyblob","tinyint","tinytext","varbinary","year"];

export default {
  default_work_step_types,
  mysql_datatypes,
}