package mockhttp

import (
	"fmt"
	"github.com/thorsager/mockdev/util"
	"gopkg.in/yaml.v2"
	"os"
	"path/filepath"
)

type Configuration struct {
	Name              string   `yaml:"name"`
	BindAddr          string   `yaml:"bind-addr"`
	ConversationFiles []string `yaml:"conversation-files"`
	Conversations     []Conversation
}

type Conversation struct {
	Name     string   `yaml:"name"`
	Request  Request  `yaml:"request"`
	Response Response `yaml:"response"`
}

type Request struct {
	UrlMatcher     UrlMatcher `yaml:"url-matcher"`
	MethodMatcher  string     `yaml:"method-matcher"`
	HeaderMatchers []string   `yaml:"header-matchers,omitempty"`
	BodyMatcher    string     `yaml:"body-matcher,omitempty"`
}

type Response struct {
	StatusCode int      `yaml:"status-code"`
	Headers    []string `yaml:"headers"`
	Body       string   `yaml:"body"`
	BodyFile   string   `yaml:"body-file,omitempty"`
}

type UrlMatcher struct {
	Path  string `yaml:"path,omitempty"`
	Query string `yaml:"query,omitempty"`
}

func DecodeConversationFile(filename string) (*Conversation, error) {
	cff, err := os.Open(filename)
	if err != nil {
		return nil, fmt.Errorf("unable to load file '%s':, %w", filename, err)
	}
	defer func() { _ = cff.Close() }()
	var v Conversation
	err = yaml.NewDecoder(cff).Decode(&v)
	if err != nil {
		return nil, fmt.Errorf("unable to decode conversation in file '%s': %w", filename, err)
	}
	if v.Response.BodyFile != "" {
		v.Response.BodyFile = util.MakeFileAbsolute(filepath.Dir(filename), v.Response.BodyFile)
	}
	return &v, nil
}
