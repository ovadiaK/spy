package app

type Api interface {
	Set(key, value string)
	Get(key string) string
}

type Memory struct {
	storage map[string]string
}

func New() Memory {
	store := make(map[string]string)
	return Memory{storage: store}
}

func (m *Memory) Set(key, value string) {
	m.storage[key] = value
}

func (m *Memory) Get(key string) string {
	if v, ok := m.storage[key]; ok {
		return v
	}
	return ""
}
