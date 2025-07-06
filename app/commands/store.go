package commands

type Store struct{
    store map[string]string
}

func NewStore() *Store{
    return &Store{store: make(map[string]string)}
}

func (s *Store) Set(key string, value string) error{
    s.store[key] = value
    return nil
}

func (s *Store) Get(key string) (string, error){
    if val, exists := s.store[key]; exists{
        return val, nil
    }
    return "", nil
}