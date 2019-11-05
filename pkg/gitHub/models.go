package gitHub

//baseIssueModel Базовая модель для создания и обновления Issue
type baseIssueModel struct {
	Title       string   `json:"title"`
	Description string   `json:"body"`
	Labels      []string `json:"labels"`
	Assignees   []string `json:"assignees"`
}

// createIssueModel Модель для создания Issue
type createIssueModel struct {
	baseIssueModel
}

// createIssueWithMilestoneModel Модель для создания Issue со спринтом
type createIssueWithMilestoneModel struct {
	baseIssueModel
	Milestone int `json:"milestone"`
}

// Credentials Модель для идентификации
type Credentials struct {
	Owner string
	Repo  string
	Token string
}

// milestoneModel Модель для создания спринта
type milestoneModel struct {
	Title       string `json:"title"`
	State       string `json:"state"`
	Description string `json:"description"`
	DueOn       string `json:"due_on"`
}

// numberModel Модель для получения номера созданного объекта
type numberModel struct {
	Value int `json:"number"`
}

// operationCollaboratorsModel Модель для выполнение операции по проверки юзера/коллаборатора
// или по назначению юзера коллаборатором
type operationCollaboratorsModel struct {
	name   string
	status bool
	err    error
}

// collaboratorModel Модель для назначения юзера коллаборатором (тело запроса)
type collaboratorModel struct {
	permission string
}
