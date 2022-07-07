package app

import (
	"bufio"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"

	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"golang.org/x/exp/maps"
)

type Plugins map[string]Versions

type Versions []string

type Options struct {
	Plugin  string
	Version string
}

const (
	ASDFBin = "/bin/asdf"
	ASDFEnv = "export ASDF_%v_VERSION=%v\n"
)

func Shell(opts ...func(*Options)) error {
	var options Options

	for _, o := range opts {
		o(&options)
	}

	plugin := options.Plugin
	version := options.Version

	plugins := listPlugins()

	if plugin == "" {
		plugin = selectPlugin(maps.Keys(plugins))
	}

	if version == "" {
		version = selectVersion(plugins[plugin], currentVersion(plugin))
	}

	fmt.Print(asdfEnv(plugin, version))

	return nil
}

func WithPlugin(plugin string) func(*Options) {
	return func(o *Options) {
		o.Plugin = plugin
	}
}

func WithVersion(version string) func(*Options) {
	return func(o *Options) {
		o.Version = version
	}
}

func selectPlugin(ps []string) string {
	var p string

	prompt := &survey.Select{
		Message: "Choose a plugin:",
		Options: ps,
	}

	if err := survey.AskOne(prompt, &p, survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)); err == terminal.InterruptErr {
		os.Exit(1)
	}

	return p
}

func selectVersion(vs Versions, cv string) string {
	var v string

	prompt := &survey.Select{
		Message: "Choose a version:",
		Options: vs,
		Default: cv,
	}
	if err := survey.AskOne(prompt, &v, survey.WithStdio(os.Stdin, os.Stderr, os.Stderr)); err == terminal.InterruptErr {
		os.Exit(1)
	}

	return v
}

func listPlugins() Plugins {
	p := make(Plugins)

	var name string

	r := strings.NewReader(asdfList())
	s := bufio.NewScanner(r)
	for s.Scan() {
		line := s.Text()
		if !strings.HasPrefix(line, " ") {
			name = line
			continue
		}

		p[name] = append(p[name], strings.TrimSpace(line))
	}

	return p
}

func currentVersion(p string) string {
	c := strings.Fields(asdfCurrent(p))

	return c[1]
}

func asdfList() string {
	out, err := exec.Command(asdfBin(), "list").Output()
	if err != nil {
		panic(err)
	}

	return string(out)
}

func asdfCurrent(p string) string {
	out, err := exec.Command(asdfBin(), "current", p).Output()
	if err != nil {
		fmt.Println(err)
	}

	return string(out)
}

func asdfBin() string {
	return path.Join(os.Getenv("ASDF_DIR"), ASDFBin)
}

func asdfEnv(p string, v string) string {
	return fmt.Sprintf(ASDFEnv, strings.ToUpper(p), v)
}
