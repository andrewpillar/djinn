package github

import (
	"bytes"
	"crypto/hmac"
	"crypto/rand"
	"crypto/sha1"
	"database/sql"
	"encoding/hex"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"

	"github.com/andrewpillar/djinn/database"
	"github.com/andrewpillar/djinn/errors"
	"github.com/andrewpillar/djinn/provider"
	"github.com/andrewpillar/djinn/runner"

	"golang.org/x/oauth2"
)

type GitHub struct {
	provider.Client
}

type Error struct {
	Message string
	Errors  []map[string]string
}

type User struct {
	ID    int64
	Login string
}

type Repo struct {
	ID          int64
	Owner       User
	FullName    string `json:"full_name"`
	HTMLURL     string `json:"html_url"`
	ContentsURL string `json:"contents_url"`
}

type PushEvent struct {
	Ref        string
	Repo       Repo `json:"repository"`
	HeadCommit struct {
		ID      string
		URL     string
		Message string
		Author  map[string]string
	} `json:"head_commit"`
}

type PullRequestEvent struct {
	Action      string
	Number      int64
	PullRequest struct {
		Title   string
		User    User
		HTMLURL string `json:"html_url"`
		Head    struct {
			Sha  string
			User User
			Repo Repo
		}
		Base    struct {
			Ref  string
			User User
			Repo Repo
		}
	} `json:"pull_request"`
}

var (
	_ provider.Interface = (*GitHub)(nil)

	signatureLength = 45

	states = map[runner.Status]string{
		runner.Queued:             "pending",
		runner.Running:            "pending",
		runner.Passed:             "success",
		runner.PassedWithFailures: "success",
		runner.Failed:             "failure",
		runner.Killed:             "failure",
		runner.TimedOut:           "failure",
	}

	PullRequestActions = map[string]struct{}{
		"opened":      {},
		"reopened":    {},
		"unlocked":    {},
		"synchronize": {},
	}
)

func decodeError(r io.Reader) error {
	err := &Error{}
	json.NewDecoder(r).Decode(err)
	return err
}

func (g *GitHub) VerifyRequest(r io.Reader, signature string) ([]byte, error) {
	if len(signature) != signatureLength {
		return nil, provider.ErrInvalidSignature
	}

	b, err := ioutil.ReadAll(r)

	if err != nil {
		return nil, errors.Err(err)
	}

	actual := make([]byte, 20)

	hex.Decode(actual, []byte(signature[5:]))

	expected := hmac.New(sha1.New, []byte(g.Secret))
	expected.Write(b)

	if !hmac.Equal(expected.Sum(nil), actual) {
		return nil, provider.ErrInvalidSignature
	}
	return b, nil
}

func New(host, endpoint, secret, clientId, clientSecret string) *GitHub {
	if endpoint == "" {
		endpoint = "https://api.github.com"
	}

	url := strings.Replace(endpoint, "api.", "", 1)

	b := make([]byte, 16)
	rand.Read(b)

	cfg := oauth2.Config{
		ClientID:     clientId,
		ClientSecret: clientSecret,
		RedirectURL:  host + "/oauth/github",
		Scopes:       []string{"admin:repo_hook", "read:org", "repo"},
		Endpoint:     oauth2.Endpoint{
			AuthURL:  url + "/login/oauth/authorize",
			TokenURL: url + "/login/oauth/access_token",
		},
	}

	return &GitHub{
		Client: provider.Client{
			Config:      cfg,
			Host:        host,
			APIEndpoint: endpoint,
			State:       hex.EncodeToString(b),
			Secret:      secret,
		},
	}
}

func (g *GitHub) Repos(tok string, page int64) ([]*provider.Repo, database.Paginator, error) {
	resp, err := g.Get(tok, "/user/repos?sort=updated&page=" + strconv.FormatInt(page, 10))

	if err != nil {
		return nil, database.Paginator{}, errors.Err(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return  nil, database.Paginator{}, errors.New("unexpected http status: " + resp.Status)
	}

	repos := make([]Repo, 0)

	json.NewDecoder(resp.Body).Decode(&repos)

	rr := make([]*provider.Repo, 0, len(repos))

	for _, r := range repos {
		rr = append(rr, &provider.Repo{
			RepoID: r.ID,
			Name:   r.FullName,
			Href:   r.HTMLURL,
		})
	}

	next, prev := provider.GetNextAndPrev(resp.Header.Get("Link"))

	p := database.Paginator{
		Next:  next,
		Prev:  prev,
		Page:  page,
		Pages: []int64{next, prev},
	}
	return rr, p, nil
}

func (g *GitHub) Groups(tok string) ([]int64, error) {
	resp, err := g.Get(tok, "/user/orgs")

	if err != nil {
		return nil, errors.Err(err)
	}

	defer resp.Body.Close()

	orgs := make([]struct {
		ID int64
	}, 0)

	json.NewDecoder(resp.Body).Decode(&orgs)

	ids := make([]int64, 0, len(orgs))

	for _, org := range orgs {
		ids = append(ids, org.ID)
	}
	return ids, nil
}

func (g *GitHub) ToggleRepo(tok string, r *provider.Repo) error {
	if !r.Enabled {
		body := map[string]interface{}{
			"config": map[string]interface{}{
				"url":          g.Host + "/hook/github",
				"secret":       g.Secret,
				"content_type": "json",
				"insecure_ssl": 0,
			},
			"events": []string{"push", "pull_request"},
		}

		buf := &bytes.Buffer{}

		json.NewEncoder(buf).Encode(body)

		resp, err := g.Post(tok, "/repos/" + r.Name + "/hooks", buf)

		if err != nil {
			return errors.Err(err)
		}

		defer resp.Body.Close()

		if resp.StatusCode != http.StatusCreated {
			return errors.Err(decodeError(resp.Body))
		}

		hook := struct {
			ID int64
		}{}

		json.NewDecoder(resp.Body).Decode(&hook)

		r.HookID = sql.NullInt64{
			Int64: hook.ID,
			Valid: true,
		}
		r.Enabled = true
		return nil
	}

	resp, err := g.Delete(tok, "/repos/" + r.Name + "/hooks/" + strconv.FormatInt(r.HookID.Int64, 10))

	if err != nil {
		return errors.Err(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return decodeError(resp.Body)
	}

	r.HookID = sql.NullInt64{
		Int64: 0,
		Valid: false,
	}
	r.Enabled = false
	return nil
}

func (g *GitHub) SetCommitStatus(tok string, r *provider.Repo, status runner.Status, url, sha string) error {
	body := map[string]string{
		"state":       states[status],
		"target_url":  url,
		"description": provider.StatusDescriptions[status],
		"context":     "continuous-integration/djinn-ci",
	}

	buf := &bytes.Buffer{}

	json.NewEncoder(buf).Encode(body)

	resp, err := g.Post(tok, "/repos/" + r.Name + "/statuses/" + sha, buf)

	if err != nil {
		return errors.Err(err)
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return errors.Err(decodeError(resp.Body))
	}
	return nil
}

func (e *Error) Error() string {
	if len(e.Errors) > 0 {
		s := e.Message + ": "

		for i, err := range e.Errors {
			s += err["message"]

			if i != len(e.Errors) - 1 {
				s += ", "
			}
		}
		return s
	}
	return e.Message
}
