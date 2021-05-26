package filter

import (
	"errors"
	"fmt"
	"net/url"
	"reflect"
	"strconv"
	"strings"

	"github.com/jmoiron/sqlx"
)

func BuildWhereClauseFromQuery(q url.Values, model interface{}) (string, []interface{}, error) {
	var conditions []string
	var arguments []interface{}

	r, err := castQueryParameters(q, model)
	if err != nil {
		return "", nil, err
	}

	for k, vs := range r {
		clause, args, err := sqlx.In(fmt.Sprintf("%v IN (?)", k), vs)
		if err != nil {
			return "", nil, err
		}
		conditions = append(conditions, clause)
		arguments = append(arguments, args...)
	}

	whereClause := strings.Join(conditions, " AND ")

	if len(whereClause) > 0 {
		whereClause = "WHERE " + whereClause
	}

	whereClause = sqlx.Rebind(sqlx.DOLLAR, whereClause)

	return whereClause, arguments, nil
}

func castQueryParameters(q url.Values, model interface{}) (map[string][]interface{}, error) {
	result := make(map[string][]interface{})
	m := reflect.ValueOf(model)
	mt := reflect.TypeOf(model)

	for k, vs := range q {
		rs := make([]interface{}, len(vs))
		f := m.FieldByName(k)
		ft, ok := mt.FieldByName(k)

		if !ok {
			return nil, errors.New("Field not found")
		}

		dbName := ft.Tag.Get("filter-alias") + "." + ft.Tag.Get("filter-column-name")

		for i, v := range vs {
			if r, err := castStringToType(v, f.Type()); err != nil {
				return nil, err
			} else {
				rs[i] = r
			}
		}

		result[dbName] = rs
	}

	return result, nil
}

func castStringToType(str string, t reflect.Type) (interface{}, error) {
	result := reflect.New(t).Elem()
	switch result.Kind() {
	case reflect.Int:
		if i, err := strconv.ParseInt(str, 0, 64); err != nil {
			return nil, err
		} else {
			result.SetInt(i)
		}
	case reflect.String:
		result.SetString(str)
	default:
		return nil, errors.New("Unrecognized type")
	}
	return result.Interface(), nil
}
