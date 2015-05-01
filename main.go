package main

import (
	"github.com/lunny/log"
	"github.com/lunny/tango"
	"github.com/tango-contrib/renders"
)

type BuildAction struct {
	tango.Ctx
	renders.Renderer
}

func (b *BuildAction) Get() error {
	pkg, version := b.Params().Get("*pkg"), b.Params().Get(":version")
	go build(pkg, version)
	return b.Render("build.html", renders.T{
		"pkg":     pkg,
		"version": version,
	})
}

type AboutAction struct {
	tango.Ctx
	renders.Renderer
}

func (b *AboutAction) Get() error {
	return b.Render("about.html", renders.T{})
}

type HomeAction struct {
	tango.Ctx
	renders.Renderer
}

func (b *HomeAction) Get() error {
	return b.Render("home.html", renders.T{})
}

func main() {
	log.SetOutputLevel(log.Lall)

	t := tango.Classic(log.Std)
	t.Use(renders.New(renders.Options{
		Reload: true,
	}))
	t.Get("/", new(HomeAction))
	t.Get("/about", new(AboutAction))
	t.Get("/*pkg/commit/:version", new(BuildAction)) // github.com/lunny/tango/commit/92f09e5ad97407cca51be096654a0857d87fa087
	t.Get("/*pkg", new(BuildAction))                 // github.com/lunny/tango
	t.Run()
}
