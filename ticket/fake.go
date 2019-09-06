package ticket

type FakeTicket struct {
	api TicketApi
}

type FakeApi struct {
}

func NewFakeApi() TicketApi {
	return &FakeApi{}
}

func (t *FakeTicket) Valid(username string) bool {
	return true
}

func (t *FakeTicket) AddComment(comment string) error {
	return nil
}

func (j *FakeApi) GetTicket(id string) (Ticket, error) {
	return &FakeTicket{api: j}, nil
}
