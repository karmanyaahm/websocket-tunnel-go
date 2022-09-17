package config

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"os"

	"github.com/BurntSushi/toml"
)

var MyConfig Config = Config{}

type (
	URL struct {
		*url.URL
	}
	Addr struct {
		HostPort string
		Net      string
	}

	Config struct {
		Peers []Peer `toml:"peer"`
	}

	Peer struct {
		Url    URL
		Pass   string
		Listen Conns
		Forw   Conns `toml:"forward"`
	}

	Conns = map[string]Addr
)

func (u *URL) UnmarshalText(text []byte) (err error) {
	u.URL, err = url.Parse(string(text))
	return
}

func (u URL) MarshalText() (text []byte, err error) {
	if u.URL == nil {
		return []byte("null"), nil
	}
	return []byte(u.String()), nil
}

func (a *Addr) UnmarshalText(text []byte) (err error) {
	url, err := url.Parse(string(text))
	fmt.Println(string(text), url.Scheme)
	if err != nil {
		return err
	}
	a.Net = url.Scheme
	a.HostPort = url.Host
	return nil
}

func (a Addr) MarshalText() (text []byte, err error) {
	if a.HostPort == "" || a.Net == "" {
		return []byte(`aaa`), errors.New("addr nil")
	}
	u := url.URL{Host: a.HostPort, Scheme: a.Net}
	return []byte(u.String()), nil
}

func Parse(content string) error {
	_, err := toml.Decode(content, &MyConfig)
	return err
}

func ParseFile(file string) error {
	fp, err := os.Open(file)
	if err != nil {
		return err
	}
	defer fp.Close()
	r, err := io.ReadAll(fp)
	return Parse(string(r))
}
