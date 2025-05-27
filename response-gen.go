package apigo

import "errors"

type ModelID struct {
	ID int64 `json:"id" format:"int64"`
}

type ModelName struct {
	Name string `json:"name"`
}

type ModelIDName struct {
	ID   int64  `json:"id" format:"int64"`
	Name string `json:"name"`
}

type ModelRowsAffected struct {
	RowsAffected int64 `json:"rows_affected" format:"int64"`
}

type ModelValue struct {
	Value any `json:"value"`
}

func RespCauseErr(cause error) map[string]any {
	var webErr *Err
	if errors.As(cause, &webErr) {
		return map[string]any{
			"cause": webErr.Cause,
		}
	}
	return map[string]any{
		"cause": cause,
	}
}

func RespCause(cause string) map[string]any {
	return map[string]any{
		"cause": cause,
	}
}

func RespID(id int64) map[string]any {
	return map[string]any{
		"id": id,
	}
}

func RespRowsAffected(rowsAff int64) map[string]any {
	return map[string]any{
		"rows_affected": rowsAff,
	}
}

func RespIDName(id int64, name string) map[string]any {
	return map[string]any{
		"id":   id,
		"name": name,
	}
}

func RespValue(val any) map[string]any {
	return map[string]any{
		"value": val,
	}
}
