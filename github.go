package gitio

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

type API struct {
	token string
	Temp  Temp
}

func NewAPI(token string) *API {
	api := &API{token: token}
	if _, err := os.Stat(defaultTmpName); err == nil {
		api.Temp.Read(defaultTmpName)
		if token == "" {
			api.token = api.Temp.Token //使用本地token
			return api
		} else if api.Temp.Token != token {
			goto query
		} else {
			return api
		}
	} else {
		goto query
	}
query:
	if token == "" {
		println("token is null, login first")
		return nil
	}
	api.Temp.Owner = api.QueryOwner()
	if api.Temp.Owner == (Owner{}) {
		return nil
	}
	api.Temp.UrlMap = api.QueryUrl()
	api.Temp.Repositories = api.QueryOwnerRepositories()
	api.Temp.Token = token
	api.Temp.Save()

	return api
}

type Repository struct {
	Id               int64       `json:"id"`
	NodeId           string      `json:"node_id"`
	Name             string      `json:"name"`
	FullName         string      `json:"full_name"`
	Private          bool        `json:"private"`
	Owner            Owner       `json:"owner"`
	HtmlUrl          string      `json:"html_url"`
	Description      string      `json:"description"`
	Fork             bool        `json:"fork"`
	Url              string      `json:"url"`
	ForksUrl         string      `json:"forks_url"`
	KeysUrl          string      `json:"keys_url"`
	CollaboratorsUrl string      `json:"collaborators_url"`
	TeamsUrl         string      `json:"teams_url"`
	HooksUrl         string      `json:"hooks_url"`
	IssueEventsUrl   string      `json:"issue_events_url"`
	EventsUrl        string      `json:"events_url"`
	AssigneesUrl     string      `json:"assignees_url"`
	BranchesUrl      string      `json:"branches_url"`
	TagsUrl          string      `json:"tags_url"`
	BlobsUrl         string      `json:"blobs_url"`
	GitTagsUrl       string      `json:"git_tags_url"`
	GitRefsUrl       string      `json:"git_refs_url"`
	TreesUrl         string      `json:"trees_url"`
	StatusesUrl      string      `json:"statuses_url"`
	LanguagesUrl     string      `json:"languages_url"`
	StargazersUrl    string      `json:"stargazers_url"`
	ContributorsUrl  string      `json:"contributors_url"`
	SubscribersUrl   string      `json:"subscribers_url"`
	SubscriptionUrl  string      `json:"subscription_url"`
	CommitsUrl       string      `json:"commits_url"`
	GitCommitsUrl    string      `json:"git_commits_url"`
	CommentsUrl      string      `json:"comments_url"`
	IssueCommentUrl  string      `json:"issue_comment_url"`
	ContentsUrl      string      `json:"contents_url"`
	CompareUrl       string      `json:"compare_url"`
	MergesUrl        string      `json:"merges_url"`
	ArchiveUrl       string      `json:"archive_url"`
	DownloadsUrl     string      `json:"downloads_url"`
	IssuesUrl        string      `json:"issues_url"`
	PullsUrl         string      `json:"pulls_url"`
	MilestonesUrl    string      `json:"milestones_url"`
	NotificationsUrl string      `json:"notifications_url"`
	LabelsUrl        string      `json:"labels_url"`
	ReleasesUrl      string      `json:"releases_url"`
	DeploymentsUrl   string      `json:"deployments_url"`
	CreatedAt        string      `json:"created_at"`
	UpdatedAt        string      `json:"updated_at"`
	PushedAt         string      `json:"pushed_at"`
	GitUrl           string      `json:"git_url"`
	SshUrl           string      `json:"ssh_url"`
	CloneUrl         string      `json:"clone_url"`
	SvnUrl           string      `json:"svn_url"`
	Homepage         string      `json:"homepage"`
	Size             int64       `json:"size"`
	StargazersCount  int         `json:"stargazers_count"`
	WatchersCount    int         `json:"watchers_count"`
	Language         string      `json:"language"`
	HasIssues        bool        `json:"has_issues"`
	HasProjects      bool        `json:"has_projects"`
	HasDownloads     bool        `json:"has_downloads"`
	HasWiki          bool        `json:"has_wiki"`
	HasPages         bool        `json:"has_pages"`
	ForksCount       int         `json:"forks_count"`
	MirrorUrl        string      `json:"mirror_url"`
	Archived         bool        `json:"archived"`
	Disabled         bool        `json:"disabled"`
	OpenIssuesCount  int         `json:"open_issues_count"`
	License          License     `json:"license"`
	Forks            string      `json:"forks"`
	OpenIssues       string      `json:"open_issues"`
	Watchers         string      `json:"watchers"`
	DefaultBranch    string      `json:"default_branch"`
	Permissions      Permissions `json:"permissions"`
}

type License struct {
	Key    string `json:"key"`
	Name   string `json:"name"`
	SpdxId string `json:"spdx_id"`
	Url    string `json:"url"`
	NodeId string `json:"node_id"`
}

type Owner struct {
	//normal
	Login             string `json:"login"`
	Id                int64  `json:"id"`
	NodeId            string `json:"node_id"`
	AvatarUrl         string `json:"avatar_url"`
	GravatarId        string `json:"gravatar_id"`
	Url               string `json:"url"`
	HtmlUrl           string `json:"html_url"`
	FollowersUrl      string `json:"followers_url"`
	FollowingUrl      string `json:"following_url"`
	GistsUrl          string `json:"gists_url"`
	StarredUrl        string `json:"starred_url"`
	SubscriptionsUrl  string `json:"subscriptions_url"`
	OrganizationsUrl  string `json:"organizations_url"`
	ReposUrl          string `json:"repos_url"`
	EventsUrl         string `json:"events_url"`
	ReceivedEventsUrl string `json:"received_events_url"`
	Type              string `json:"type"`
	SiteAdmin         bool   `json:"site_admin"`
	//current owner have
	Name        string `json:"name"`
	Company     string `json:"company"`
	Blog        string `json:"blog"`
	Location    string `json:"location"`
	Email       string `json:"email"`
	Hireable    string `json:"hireable"`
	Bio         string `json:"bio"`
	PublicRepos string `json:"public_repos"`
	PublicGists string `json:"public_gists"`
	Followers   string `json:"followers"`
	Following   string `json:"following"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

type Permissions struct {
	Admin bool `json:"admin"`
	Push  bool `json:"push"`
	Pull  bool `json:"pull"`
}

type File struct {
	Name        string `json:"name"`
	Path        string `json:"path"`
	Sha         string `json:"sha"`
	Size        int64  `json:"size"`
	Url         string `json:"url"`
	HtmlUrl     string `json:"html_url"`
	GitUrl      string `json:"git_url"`
	DownloadUrl string `json:"download_url"`
	Type        string `json:"type"`
	Links       Links  `json:"_links"`
}

type Links struct {
	Self string `json:"self"`
	Git  string `json:"git"`
	Html string `json:"html"`
}

func (a *API) QueryUrl() map[string]string {
	if a.Temp.UrlMap != nil {
		return a.Temp.UrlMap
	}
	urlMap := map[string]string{}
	a.Query("https://api.github.com", &urlMap)
	return urlMap
}

func (a *API) QueryOwnerRepositories() []Repository {
	if a.Temp.Repositories != nil {
		return a.Temp.Repositories
	}
	r := []Repository{}
	a.Query("https://api.github.com/user/repos", &r)
	return r
}

func (a *API) QueryOwner() Owner {
	if a.Temp.Owner != (Owner{}) {
		return a.Temp.Owner
	}
	o := Owner{}
	a.Query("https://api.github.com/user", &o)
	return o
}

func (a *API) Query(url string, data interface{}) interface{} {
	cli := NewHttpClient(map[string]string{"Authorization": fmt.Sprintf("token %s", a.token)}, nil, "")
	status, b := cli.Get(url)
	if status != http.StatusOK {
		println(status, string(b))
		return nil
	}
	if data != nil {
		_ = json.Unmarshal(b, data)
	}
	return data
}

func (a *API) OAuth() {

}

const (
	actionAdd    = http.MethodPut
	actionDelete = http.MethodDelete
	actionUpdate = http.MethodPut
)

type PushForm struct {
	Message string `json:"message"` //Required. The commit message.
	Content string `json:"content"` //Required. The updated file content, Base64 encoded.
	Sha     string `json:"sha"`     //Required. The blob SHA of the file being replaced.
	Branch  string `json:"branch"`  //The branch name. Default: the repository’s default branch (usually master)
}

func (a *API) Push(action string, url string, pushForm PushForm) (int, []byte) {
	header := map[string]string{"Authorization": fmt.Sprintf("token %s", a.token)}
	body, _ := json.Marshal(pushForm)
	cli := NewHttpClient(header, body, "application/json")
	return cli.Do(action, url)
}

func (a *API) QueryPath(url string, path string) []File {
	fs := []File{}
	url = strings.Replace(url, "{+path}", path, -1)
	a.Query(url, &fs)
	return fs
}

type Temp struct {
	Token        string
	UrlMap       map[string]string
	Owner        Owner
	Repositories []Repository
}

var defaultTmpName = Home() + PathSeparator() + ".gitio.tmp"

func (t *Temp) Save() {
	b, err := json.Marshal(t)
	if err != nil {
		return
	}
	WriteFile(defaultTmpName, []byte(Base64Encode(b))) //base64加密
}

func (t *Temp) Read(name string) {
	if name == "" {
		name = defaultTmpName
	}
	f, err := os.Open(name)
	if err != nil {
		println(err.Error())
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		println(err.Error())
		return
	}
	db := Base64Decode(string(b))
	err = json.Unmarshal(db, &t)
	if err != nil {
		println(err.Error())
	}

}

func (t *Temp) FindRepo(name string) Repository {
	for _, repo := range t.Repositories {
		if repo.Name == name {
			return repo
		}
	}
	return Repository{}
}
