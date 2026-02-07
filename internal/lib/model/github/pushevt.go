package gh

/*
Github push event schema
	see: https://docs.github.com/en/webhooks/webhook-events-and-payloads#push
*/

type GhRepository struct {
	Id                          int    `json:"id"`
	Node_id                     string `json:"node_id"`
	Name                        string `json:"name"`
	Full_name                   string `json:"full_name"`
	Private                     bool   `json:"private"`
	Owner                       GhUser `json:"owner"`
	Html_url                    string `json:"html_url"`
	Description                 string `json:"description"`
	Fork                        bool   `json:"fork"`
	Url                         string `json:"url"`
	Forks_url                   string `json:"forks_url"`
	Keys_url                    string `json:"keys_url"`
	Collaborators_url           string `json:"collaborators_url"`
	Teams_url                   string `json:"teams_url"`
	Hooks_url                   string `json:"hooks_url"`
	Issue_events_url            string `json:"issue_events_url"`
	Events_url                  string `json:"events_url"`
	Assignees_url               string `json:"assignees_url"`
	Branches_url                string `json:"branches_url"`
	Tags_url                    string `json:"tags_url"`
	Blobs_url                   string `json:"blobs_url"`
	Git_tags_url                string `json:"git_tags_url"`
	Git_refs_url                string `json:"git_refs_url"`
	Trees_url                   string `json:"trees_url"`
	Statuses_url                string `json:"statuses_url"`
	Languages_url               string `json:"languages_url"`
	Stargazers_url              string `json:"stargazers_url"`
	Contributors_url            string `json:"contributors_url"`
	Subscribers_url             string `json:"subscribers_url"`
	Subscription_url            string `json:"subscription_url"`
	Commits_url                 string `json:"commits_url"`
	Git_commits_url             string `json:"git_commits_url"`
	Comments_url                string `json:"comments_url"`
	Issue_comment_url           string `json:"issue_comment_url"`
	Contents_url                string `json:"contents_url"`
	Compare_url                 string `json:"compare_url"`
	Merges_url                  string `json:"merges_url"`
	Archive_url                 string `json:"archive_url"`
	Downloads_url               string `json:"downloads_url"`
	Issues_url                  string `json:"issues_url"`
	Pulls_url                   string `json:"pulls_url"`
	Milestones_url              string `json:"milestones_url"`
	Notifications_url           string `json:"notifications_url"`
	Labels_url                  string `json:"labels_url"`
	Releases_url                string `json:"releases_url"`
	Deployments_url             string `json:"deployments_url"`
	Created_at                  int    `json:"created_at"`
	Updated_at                  string `json:"updated_at"`
	Pushed_at                   int    `json:"pushed_at"`
	Git_url                     string `json:"git_url"`
	Ssh_url                     string `json:"ssh_url"`
	Clone_url                   string `json:"clone_url"`
	Svn_url                     string `json:"svn_url"`
	Homepage                    string `json:"homepage"`
	Size                        int    `json:"size"`
	Stargazers_count            int    `json:"stargazers_count"`
	Watchers_count              int    `json:"watchers_count"`
	Language                    string `json:"language"`
	Has_issues                  bool   `json:"has_issues"`
	Has_projects                bool   `json:"has_projects"`
	Has_downloads               bool   `json:"has_downloads"`
	Has_wiki                    bool   `json:"has_wiki"`
	Has_pages                   bool   `json:"has_pages"`
	Has_discussions             bool   `json:"has_discussions"`
	Forks_count                 int    `json:"forks_count"`
	Mirror_url                  string `json:"mirror_url"`
	Archived                    bool   `json:"archived"`
	Disabled                    bool   `json:"disabled"`
	Open_issues_count           int    `json:"open_issues_count"`
	License                     string `json:"license"`
	Allow_forking               bool   `json:"allow_forking"`
	Is_template                 bool   `json:"is_template"`
	Web_commit_signoff_required bool   `json:"web_commit_signoff_required"`
	Visibility                  string `json:"visibility"`
	Forks                       int    `json:"forks"`
	Open_issues                 int    `json:"open_issues"`
	Watchers                    int    `json:"watchers"`
	Default_branch              string `json:"default_branch"`
	Stargazers                  int    `json:"stargazers"`
	Master_branch               string `json:"master_branch"`
	/* Topics [] `json:"topics"` */
}

type GhCommitUser struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Date     string `json:"date"`
	Username string `json:"username"`
}

type GhCommit struct {
	Id        string       `json:"id"`
	Tree_id   string       `json:"tree_id"`
	Distinct  bool         `json:"distinct"`
	Message   string       `json:"message"`
	Timestamp string       `json:"timestamp"`
	Url       string       `json:"url"`
	Author    GhCommitUser `json:"author"`
	Committer GhCommitUser `json:"committer"`
	Added     []string     `json:"added"`
	Removed   []string     `json:"removed"`
	Modified  []string     `json:"modified"`
}

type GhUserBasic struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}
type GhUser struct {
	Login               string `json:"login"`
	Id                  int    `json:"id"`
	Node_id             string `json:"node_id"`
	Avatar_url          string `json:"avatar_url"`
	Gravatar_id         string `json:"gravatar_id"`
	Url                 string `json:"url"`
	Html_url            string `json:"html_url"`
	Followers_url       string `json:"followers_url"`
	Following_url       string `json:"following_url"`
	Gists_url           string `json:"gists_url"`
	Starred_url         string `json:"starred_url"`
	Subscriptions_url   string `json:"subscriptions_url"`
	Organizations_url   string `json:"organizations_url"`
	Repos_url           string `json:"repos_url"`
	Events_url          string `json:"events_url"`
	Received_events_url string `json:"received_events_url"`
	Type                string `json:"type"`
	User_view_type      string `json:"user_view_type"`
	Site_admin          bool   `json:"site_admin"`
}

type GhPushPayload struct {
	Ref         string       `json:"ref"`
	Before      string       `json:"before"`
	After       string       `json:"after"`
	Repository  GhRepository `json:"repository"`
	Pusher      GhUserBasic  `json:"pusher"`
	Sender      GhUser       `json:"sender"`
	Created     bool         `json:"created"`
	Deleted     bool         `json:"deleted"`
	Forced      bool         `json:"forced"`
	Base_ref    string       `json:"base_ref"`
	Compare     string       `json:"compare"`
	Commits     []GhCommit   `json:"commits"`
	Head_commit GhCommit     `json:"head_commit"`
}
