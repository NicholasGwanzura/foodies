package httpapi

import (
	"bytes"
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"net/http"
	"reflect"
	"strconv"
	"strings"

	"github.com/google/uuid"

	"foodies/api/internal/service"
)

// Database rows and API schemas share field names by design (docs/data-model.md),
// so they convert through JSON. Only money differs: the database keeps
// <name>_amount and <name>_currency, and the API keeps <name>: {amount, currency}.

const (
	defaultLimit = 25
	maxLimit     = 100
)

func toMap(v any) (map[string]any, error) {
	raw, err := json.Marshal(v)
	if err != nil {
		return nil, fmt.Errorf("encode: %w", err)
	}
	out := map[string]any{}
	dec := json.NewDecoder(bytes.NewReader(raw))
	dec.UseNumber()
	if err := dec.Decode(&out); err != nil {
		return nil, fmt.Errorf("decode: %w", err)
	}
	return out, nil
}

func fromMap[T any](m map[string]any) (T, error) {
	var out T
	raw, err := json.Marshal(m)
	if err != nil {
		return out, fmt.Errorf("encode: %w", err)
	}
	if err := json.Unmarshal(raw, &out); err != nil {
		return out, fmt.Errorf("decode into %T: %w", out, err)
	}
	return out, nil
}

// foldMoney turns <name>_amount and <name>_currency into <name>.
func foldMoney(m map[string]any) {
	for key, amount := range m {
		name, ok := strings.CutSuffix(key, "_amount")
		if !ok {
			continue
		}
		currency, ok := m[name+"_currency"]
		if !ok {
			continue
		}
		delete(m, key)
		delete(m, name+"_currency")
		if amount != nil && currency != nil {
			m[name] = map[string]any{"amount": amount, "currency": currency}
		}
	}
}

// unfoldMoney is the reverse, for request bodies.
func unfoldMoney(m map[string]any) {
	for key, v := range m {
		obj, ok := v.(map[string]any)
		if !ok || len(obj) != 2 || obj["amount"] == nil || obj["currency"] == nil {
			continue
		}
		delete(m, key)
		m[key+"_amount"] = obj["amount"]
		m[key+"_currency"] = obj["currency"]
	}
}

// output converts a row, and optionally extra fields, to an API schema.
func output[T any](row any, extra ...map[string]any) (T, error) {
	m, err := toMap(row)
	if err != nil {
		var zero T
		return zero, err
	}
	foldMoney(m)
	for _, e := range extra {
		for k, v := range e {
			if rows, ok := v.([]map[string]any); ok {
				for _, r := range rows {
					foldMoney(r)
				}
			}
			m[k] = v
		}
	}
	return fromMap[T](m)
}

// childInputs converts the items of a child array in a request body.
func childInputs[T any](items any, parentColumn string, parent uuid.UUID) ([]T, error) {
	list, ok := items.([]any)
	if !ok {
		return nil, nil
	}
	out := make([]T, 0, len(list))
	for _, item := range list {
		m, ok := item.(map[string]any)
		if !ok {
			continue
		}
		unfoldMoney(m)
		m[parentColumn] = parent
		row, err := fromMap[T](m)
		if err != nil {
			return nil, service.Fail(http.StatusUnprocessableEntity, "validation_failed", "Validation failed", "A row is not in the expected format.")
		}
		out = append(out, row)
	}
	return out, nil
}

// outputs converts rows to a slice of an API schema.
func outputs[T any, R any](rows []R) ([]T, error) {
	out := make([]T, 0, len(rows))
	for _, row := range rows {
		item, err := output[T](row)
		if err != nil {
			return nil, err
		}
		out = append(out, item)
	}
	return out, nil
}

// input validates a request body and converts it to query parameters. Extra
// holds fields that come from the path.
func input[T any](body any, extra map[string]any) (T, error) {
	var zero T
	if err := checkRequired(body); err != nil {
		return zero, err
	}
	m, err := toMap(body)
	if err != nil {
		return zero, err
	}
	unfoldMoney(m)
	for k, v := range extra {
		m[k] = v
	}
	out, err := fromMap[T](m)
	if err != nil {
		return zero, service.Fail(http.StatusUnprocessableEntity, "validation_failed", "Validation failed", "A value is not in the expected format.")
	}
	return out, nil
}

// checkRequired reports required fields left empty. The generated decoders do
// not enforce `required`, so it comes from the struct tags: a required field is
// one whose JSON tag has no omitempty.
func checkRequired(body any) error {
	v := reflect.ValueOf(body)
	for v.Kind() == reflect.Pointer {
		if v.IsNil() {
			return &service.Error{Status: http.StatusUnprocessableEntity, Code: "validation_failed", Title: "Validation failed", Detail: "The request has no body."}
		}
		v = v.Elem()
	}
	if v.Kind() != reflect.Struct {
		return nil
	}
	var issues []service.FieldIssue
	t := v.Type()
	for i := range t.NumField() {
		tag := t.Field(i).Tag.Get("json")
		name, opts, _ := strings.Cut(tag, ",")
		if name == "" || name == "-" || strings.Contains(opts, "omitempty") {
			continue
		}
		f := v.Field(i)
		switch f.Kind() {
		case reflect.Bool, reflect.Int, reflect.Int32, reflect.Int64, reflect.Float32, reflect.Float64:
			continue
		case reflect.Struct:
			if name == "amount" || !f.IsZero() {
				continue
			}
		}
		if f.IsZero() {
			issues = append(issues, service.FieldIssue{Pointer: "/" + name, Code: "required", Message: "This field is required."})
		}
	}
	if len(issues) > 0 {
		return &service.Error{Status: http.StatusUnprocessableEntity, Code: "validation_failed", Title: "Validation failed",
			Detail: "Some fields need attention.", Fields: issues}
	}
	return nil
}

func encodeCursor(id string) *string {
	s := base64.RawURLEncoding.EncodeToString([]byte(id))
	return &s
}

func decodeCursor(s string) (uuid.UUID, error) {
	raw, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		return uuid.Nil, fmt.Errorf("decode cursor: %w", err)
	}
	id, err := uuid.ParseBytes(raw)
	if err != nil {
		return uuid.Nil, fmt.Errorf("parse cursor: %w", err)
	}
	return id, nil
}

// listPage runs a list query and shapes the result as { data, page }. Filters
// come from params by name. It fetches one extra row to learn if more exist.
func listPage[T any, P any, R any](ctx context.Context, params any, extra map[string]any, call func(context.Context, P) ([]R, error)) (T, error) {
	var zero T
	m, err := toMap(params)
	if err != nil {
		return zero, err
	}
	limit := defaultLimit
	if raw, ok := m["limit"]; ok {
		n, err := strconv.Atoi(fmt.Sprint(raw))
		if err != nil || n < 1 || n > maxLimit {
			return zero, service.Fail(http.StatusBadRequest, "validation_failed", "Malformed request", "limit must be from 1 to 100.")
		}
		limit = n
	}
	if raw, ok := m["cursor"].(string); ok && raw != "" {
		id, err := decodeCursor(raw)
		if err != nil {
			return zero, service.Fail(http.StatusBadRequest, "validation_failed", "Malformed request", "The cursor is not valid.")
		}
		m["cursor_id"] = id
	}
	delete(m, "limit")
	delete(m, "cursor")
	delete(m, "sort")
	if v, ok := m["from"]; ok {
		m["from_date"] = v
	}
	if v, ok := m["to"]; ok {
		m["to_date"] = v
	}
	m["row_limit"] = limit + 1
	for k, v := range extra {
		m[k] = v
	}
	args, err := fromMap[P](m)
	if err != nil {
		return zero, service.Fail(http.StatusBadRequest, "validation_failed", "Malformed request", "A filter is not in the expected format.")
	}
	rows, err := call(ctx, args)
	if err != nil {
		return zero, err
	}

	more := len(rows) > limit
	if more {
		rows = rows[:limit]
	}
	data := make([]map[string]any, 0, len(rows))
	for _, row := range rows {
		item, err := toMap(row)
		if err != nil {
			return zero, err
		}
		foldMoney(item)
		data = append(data, item)
	}
	var next *string
	if more {
		next = encodeCursor(fmt.Sprint(data[len(data)-1]["id"]))
	}
	return fromMap[T](map[string]any{
		"data": data,
		"page": map[string]any{"limit": limit, "has_more": more, "next_cursor": next, "prev_cursor": nil},
	})
}

func etag(version int32) *string {
	s := service.Etag(version)
	return &s
}

// ifMatch reads a row version from an If-Match header.
func ifMatch(header string) (int32, error) {
	v := strings.Trim(strings.TrimPrefix(strings.TrimSpace(header), "W/"), `"`)
	n, err := strconv.ParseInt(v, 10, 32)
	if err != nil {
		return 0, service.Fail(http.StatusPreconditionFailed, "stale_version", "Changed since you opened it", "If-Match does not hold a row version.")
	}
	return int32(n), nil
}

// str reads an optional field. The generator makes writeOnly fields pointers
// even when they are required, so empty strings fail in the service instead.
func str(p *string) string {
	if p == nil {
		return ""
	}
	return *p
}
