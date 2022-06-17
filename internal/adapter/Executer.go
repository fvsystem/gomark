package adapter

type Executer interface {
	Execute(
		testCreator TestCreator,
		requester Requester,
	)
}
