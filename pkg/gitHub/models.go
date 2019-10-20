package gitHub

//baseIssueModel Базовая модель для создания и обновления Issue
type baseIssueModel struct {
	Title     string   `json:"title"`
	Body      string   `json:"body"`
	Labels    []string `json:"labels"`
	Assignees []string `json:"assignees"`
}

// CreateIssueModel Модель для создания Issue
type CreateIssueModel struct {
	baseIssueModel
}

// CreateIssueWithMilestoneModel Модель для создания Issue с вехой
type CreateIssueWithMilestoneModel struct {
	baseIssueModel
	Milestone int `json:"milestone"`
}

// Credentials Модель для идентификации
type Credentials struct {
	Owner string
	Repo  string
	Token string
}

// MilestoneModel Модель для создания milestone
type MilestoneModel struct {
	Title       string `json:"title"`
	State       string `json:"state"`
	Description string `json:"description"`
	DueOn       string `json:"due_on"`
}
