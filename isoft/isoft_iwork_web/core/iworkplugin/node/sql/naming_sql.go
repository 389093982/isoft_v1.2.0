package sql

import "isoft/isoft/common/stringutil"

func parseNamingSql(sqlStr string) (string, []string) {
	namings := stringutil.GetSubStringWithRegexp(sqlStr, ":[a-zA-Z0-9_]+")
	sqlStr, _ = stringutil.ReplaceAllString(sqlStr, ":[a-zA-Z0-9_]+", "?")
	return sqlStr, namings
}
