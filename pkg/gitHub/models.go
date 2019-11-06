package gitHub

//baseCreateIssueModel Базовая модель для создания и обновления Issue
type baseCreateIssueModel struct {
	Title       string   `json:"title"`
	Description string   `json:"body"`
	Labels      []string `json:"labels"`
	Assignees   []string `json:"assignees"`
}

// createIssueModel Модель для создания Issue
type createIssueModel struct {
	baseCreateIssueModel
}

// createIssueWithMilestoneModel Модель для создания Issue со спринтом
type createIssueWithMilestoneModel struct {
	baseCreateIssueModel
	Milestone int `json:"milestone"`
}

// CredentialsModel Модель для идентификации
type CredentialsModel struct {
	Owner string
	Repo  string
	Token string
}

// createMilestoneModel Модель для создания спринта
type createMilestoneModel struct {
	Title       string `json:"title"`
	State       string `json:"state"`
	Description string `json:"description"`
	DueOn       string `json:"due_on"`
}

// numberModel Модель для получения номера созданного объекта
type numberModel struct {
	Value int `json:"number"`
}

// operationModel Модель для выполнение операции по проверки юзера/коллаборатора
// или по назначению юзера коллаборатором
type operationModel struct {
	name   string
	status bool
	err    error
}

// collaboratorModel Модель для назначения юзера коллаборатором (тело запроса)
type collaboratorModel struct {
	permission string
}
