package users

type MockStorage struct {
	dataMock   []User
	errOnWrite error
	errOnRead  error
}

func (m *MockStorage) Read(data interface{}) (err error) {
	if m.errOnRead != nil {
		return m.errOnRead
	}

	castedData := data.(*[]User)
	*castedData = m.dataMock
	return nil
}

func (m *MockStorage) Write(data interface{}) (err error) {
	if m.errOnWrite != nil {
		return m.errOnWrite
	}

	castedData := data.(*User)
	m.dataMock = append(m.dataMock, *castedData)
	return nil
}
