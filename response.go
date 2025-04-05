package apigo

import "errors"

func RespCauseErr(cause error) map[string]any {
	var webErr *WebErr
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

func RespID(id int64) map[string]int64 {
	return map[string]int64{
		"id": id,
	}
}

func RespIDName(id int64, name string) map[string]any {
	return map[string]any{
		"id":   id,
		"name": name,
	}
}
