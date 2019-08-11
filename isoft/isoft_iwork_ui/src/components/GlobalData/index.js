const default_work_step_types=[
  {"name":"work_start","icon":"ios-arrow-dropright"},
  {"name":"work_end","icon":"ios-arrow-dropleft"},
  {"name":"empty","icon":"ios-mail-open-outline"},
  {"name":"work_sub","icon":"logo-buffer"},
  // 数据库相关
  {"name":"sql_query","icon":"ios-cube-outline"},
  {"name":"sql_execute","icon":"ios-crop-outline"},
  {"name":"db_parser","icon":"ios-map-outline"},
  {"name":"json_render","icon":"ios-git-branch"},
  {"name":"json_parser","icon":"ios-git-compare"},
  {"name":"http_request","icon":"ios-globe-outline"},
  {"name":"http_request_parser","icon":"ios-globe-outline"},
  {"name":"mapper","icon":"ios-infinite"},
  {"name":"do_receive_file","icon":"ios-book-outline"},
  {"name":"do_response_receive_file","icon":"ios-book-outline"},
  {"name":"file_read","icon":"ios-book-outline"},
  {"name":"file_write","icon":"ios-create-outline"},
  {"name":"file_delete","icon":"ios-log-out"},
  {"name":"file_sync","icon":"md-paper"},
  {"name":"href_parser","icon":"ios-ionitron-outline"},
  {"name":"entity_parser","icon":"ios-refresh-circle-outline"},
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
  {"name":"panic_error","icon":"md-hammer"},
  {"name":"template","icon":"md-hammer"},
  {"name":"create_jwt","icon":"md-hammer"},
  {"name":"parse_jwt","icon":"md-hammer"},
  {"name":"do_error_filter","icon":"md-hammer"},
];

const mysql_datatypes = ["varchar","char","text","float","int","date","datetime","decimal","double",
  "bigint","binary","bit","blob","bool","boolean","enum","longblob","longtext","mediumblob","mediumint",
  "mediumtext","numeric","real","set","smallint","time","timestamp","tinyblob","tinyint","tinytext","varbinary","year"];


const quick_funcs = [
    {funcDemo:"StringsEq($str1,$str2)",funcDesc:"字符串转大写函数",},
    {funcDemo:"stringsToUpper($str)",funcDesc:"字符串相等比较",},
    {funcDemo:"stringsToLower($str)",funcDesc:"字符串转小写函数",},
    {funcDemo:"stringsJoin($str1,$str2)",funcDesc:"字符串拼接函数",},
    {funcDemo:"stringsJoinWithSep($str1,$str2)",funcDesc:"字符串拼接函数",},
    {funcDemo:"int64Add($int1,$int2)",funcDesc:"数字相加函数",},
    {funcDemo:"int64Sub($int1,$int2)",funcDesc:"数字相减函数",},
    {funcDemo:"int64Multi($int1,$int2)",funcDesc:"数字相乘函数",},
    {funcDemo:"stringsContains($str1,$str2)",funcDesc:"字符串包含函数",},
    {funcDemo:"stringsHasPrefix($str1,$str2)",funcDesc:"字符串前缀判断函数",},
    {funcDemo:"stringsHasSuffix($str1,$str2)",funcDesc:"字符串后缀判断函数",},
    {funcDemo:"stringsTrimSuffix($str1,$suffix)",funcDesc:"字符串去除后缀",},
    {funcDemo:"stringsTrimPrefix($str1,$prefix)",funcDesc:"字符串去除前缀",},
    {funcDemo:"int64Gt($int1,$int2)",funcDesc:"判断数字1是否大于数字2",},
    {funcDemo:"int64Lt($int1,$int2)",funcDesc:"判断数字1是否小于数字2",},
    {funcDemo:"int64Eq($int1,$int2)",funcDesc:"判断数字1是否等于数字2",},
    {funcDemo:"and($bool1,$bool2)",funcDesc:"判断bool1和bool2同时满足",},
    {funcDemo:"or($bool,$bool2)",funcDesc:"判断bool1和bool2只要一个满足即可",},
    {funcDemo:"not($bool)",funcDesc:"bool值取反",},
    {funcDemo:"uuid()",funcDesc:"生成随机UUID信息",},
    {funcDemo:"isempty($var)",funcDesc:"判断变量或者字符串是否为空",},
    {funcDemo:"getDirPath($filepath)",funcDesc:"获取当前文件父级目录的绝对路径",},
    {funcDemo:"pathJoin($path1,$path2)",funcDesc:"文件路径拼接",},
    {funcDemo:"ifThenElse($condition,$var1,$var2)",funcDesc:"三目运算符,条件满足返回$var1,不满足返回$var2",},
    {funcDemo:"getRequestParameter($url,$paramName)",funcDesc:"从url地址中根据参数名获取参数值",},
    {funcDemo:"getRequestParameters($url,$paramName)",funcDesc:"从url地址中根据参数名获取参数值,返回的是数组",},
    {funcDemo:"getDomain($url)",funcDesc:"从url地址中获取 domain 信息",},
    {funcDemo:"getNotEmpty($var1,$var2)",funcDesc:"从参数列表中获取第一个非空值",},
  ]

export default {
  default_work_step_types,
  mysql_datatypes,
  quick_funcs,
}
