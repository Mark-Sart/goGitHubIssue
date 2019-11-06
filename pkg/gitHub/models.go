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

// createCollaboratorModel Модель для назначения юзера коллаборатором (тело запроса)
type createCollaboratorModel struct {
	Permission string `json:"permission"`
}

// issueModel Модель issueModel
type issueModel struct {
	Number      int              `json:"number"`
	URL         string           `json:"html_url"`
	State       string           `json:"state"`
	Title       string           `json:"title"`
	Description string           `json:"body"`
	Creator     userModel        `json:"user"`
	Labels      []labelModel     `json:"labels"`
	Assignees   []userModel      `json:"assignees"`
	Milestone   milestoneModel   `json:"milestone"`
	Comments    int              `json:"comments"`
	PullRequest pullRequestModel `json:"pull_request"`
	CreatedAt   string           `json:"created_at"`
	UpdatedAt   string           `json:"updated_at"`
	ClosedAt    string           `json:"closed_at"`
	ClosedBy    userModel        `json:"closed_by"`
}

// userModel Модель юзера
type userModel struct {
	Login string `json:"login"`
	URL   string `json:"html_url"`
}

// labelModel Модель тега
type labelModel struct {
	Name        string `json:"name"`
	Description string `json:"description"`
}

// milestoneModel Модель спринта
type milestoneModel struct {
	Number       int       `json:"number"`
	State        string    `json:"state"`
	Title        string    `json:"title"`
	Description  string    `json:"description"`
	Creator      userModel `json:"creator"`
	OpenIssues   int       `json:"open_issues"`
	ClosedIssues int       `json:"closed_issues"`
	CreatedAt    string    `json:"created_at"`
	UpdatedAt    string    `json:"updated_at"`
	ClosedAt     string    `json:"closed_at"`
	DueOn        string    `json:"due_on"`
}

// pullRequestModel Модель pull-request
type pullRequestModel struct {
	URL      string `json:"html_url"`
	DiffURL  string `json:"diff_url"`
	PatchURL string `json:"patch_url"`
}
