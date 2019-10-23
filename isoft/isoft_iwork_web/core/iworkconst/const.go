package iworkconst

const BINARY_STREAM_PREFIX = "binstream_"

const BYTE_ARRAY_PREFIX = "bytearr_"

const BASE64STRING_PREFIX = "base64str_"

const STRING_PREFIX = "str_"

const BOOL_PREFIX = "bool_"

const FOREACH_PREFIX = "iter_"

const MULTI_PREFIX = "multi_"

const NUMBER_PREFIX = "number_"

const FIELD_PREFIX = "field_"

const COMPLEX_PREFIX = "complex_"

const NODE_TYPE_WORK_START = "work_start"
const NODE_TYPE_WORK_END = "work_end"
const NODE_TYPE_WORK_SUB = "work_sub"

const LOG_LEVEL_INFO = "INFO"
const LOG_LEVEL_SUCCESS = "SUCCESS"
const LOG_LEVEL_ERROR = "ERROR"

const HTTP_REQUEST_OBJECT = "____request"
const HTTP_REQUEST_IFILE_UPLOAD = "____request_ifile_upload"
const DO_ERROR_FILTER = "__doErrorFilter__"
const DO_RESPONSE_RECEIVE_FILE = "__do_response_receive_file__"
const TRACKING_ID = "__trackingId"
const FILTER_TRACKING_ID_STACK = "filter__trackingId_stack"

var FORBIDDEN_WORK_NAMES = []string{"WORK", "RESOURCE", "GLOBAL"}

const PARENT_STEP_ID_FOR_START_END = -1
