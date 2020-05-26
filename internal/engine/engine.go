package engine

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"path/filepath"
	"sync"

	"github.com/mpragliola/victor/internal/content"
	"github.com/mpragliola/victor/internal/filesystem"
	"github.com/mpragliola/victor/internal/parser/parsers"
)

// Engine ...
type Engine struct {
	namingStrategy NamingStrategy
	fileSystem     filesystem.FileSystem
}

// NewEngine ...
func NewEngine(s NamingStrategy, f filesystem.FileSystem) *Engine {
	return &Engine{
		namingStrategy: s,
		fileSystem:     f,
	}
}

// Build ...
func (e *Engine) Build() {
	e.fileSystem.InitPublicFolder()

	funcMap := template.FuncMap{
		"foo": func(s string) string {
			return s + "Carlo"
		},
	}

	layouts, err := e.fileSystem.GatherAllLayouts()

	if err != nil {
		log.Fatal(err)
	}

	layout, err := template.New("main.html").
		Funcs(funcMap).
		ParseFiles(layouts...)

	if err != nil {
		log.Fatal(err)
	}

	files, err := e.fileSystem.GatherAllContent()

	if err != nil {
		log.Fatal(err)
	}

	wg := &sync.WaitGroup{}
	c := make(chan string)

	for _, file := range files {
		wg.Add(1)
		go e.process(file, layout, c, wg)
	}

	go func() {
		wg.Wait()
		close(c)
	}()

	for i := range c {
		fmt.Println(i)
	}
}

func (e *Engine) process(file string, layout *template.Template, c chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	outpath := e.parseFile(layout, file)
	relpath, _ := filepath.Rel(e.fileSystem.GetDataPath(), file)
	relOutPath, _ := filepath.Rel(e.fileSystem.GetPublicPath(), outpath)

	c <- fmt.Sprintf("%s -> %s", relpath, relOutPath)
}

func (e *Engine) writeDestFile(layout *template.Template, page content.Page) string {
	outpath := e.namingStrategy.GetDestinationFilename(page)

	f, err := e.fileSystem.Create(outpath)
	defer f.Close()

	if err != nil {
		log.Panic(err)
	}

	err = layout.ExecuteTemplate(f, "main.html", page)

	if err != nil {
		log.Panic(err)
	}

	return outpath
}

func (e *Engine) parseFile(layout *template.Template, sourceFileName string) string {
	content, err := ioutil.ReadFile(sourceFileName)

	if err != nil {
		log.Fatal(err)
	}

	p := parsers.NewDefault()
	page := p.Parse(content, sourceFileName)

	return e.writeDestFile(layout, page)
}
