package gipr

import (
    "context"
    "github.com/Masterminds/semver/v3"
    "github.com/google/go-github/v32/github"
    "golang.org/x/oauth2"
    "log"
)

type ClientImpl struct {
    client *github.Client
}

type Client interface {
    GetLatestVersion(owner, repo string) (string, error)
    NextTag(owner, repo, currentTag string) (string, error)
}

func NewClient(token string) Client {
    ctx := context.Background()

    ts := oauth2.StaticTokenSource(&oauth2.Token{AccessToken: token})
    tc := oauth2.NewClient(ctx, ts)
    cl := github.NewClient(tc)
    return &ClientImpl{client: cl}
}

func (c *ClientImpl) GetLatestVersion(owner, repo string) (string, error) {
    release, _, err := c.client.Repositories.GetLatestRelease(context.TODO(), owner, repo)
    if err != nil {
        return "", err
    }
    return release.GetTagName(), nil
}

func (c *ClientImpl) NextTag(owner, repo, currentTag string) (string, error) {
    v, err := semver.NewVersion(currentTag)
    log.Printf("latest %v", currentTag)
    if err != nil {
        return "", err
    }

    nextVersion := v.IncPatch()
    log.Printf("nextVersion %v", nextVersion)
    return nextVersion.String(), nil
}
