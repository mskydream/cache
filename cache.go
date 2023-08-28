package cache

func New() *Memory {
	return &Memory{
		memory: make(map[string]interface{}),
	}
}

type Memory struct {
	memory map[string]interface{}
}

func (m *Memory) Set(key string, value interface{}) {
	m.memory[key] = value
}

func (m *Memory) Get(key string) interface{} {
	return m.memory[key]
}

func (m *Memory) Delete(key string) {
	delete(m.memory, key)
}
