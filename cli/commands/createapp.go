package commands

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/template"
)

func Createapp(name, destdir, templatedir string) error {
	/*err := utils.CopyDir(templatedir, destdir, "")
	if err != nil {
		return err
	}*/
	dest, err := os.Open(destdir)
	defer dest.Close()
	if err != nil {
		fmt.Println("Error reading destination directory ", destdir)
		return err
	} else {
		fils, _ := dest.Readdir(1)
		if fils != nil && len(fils) > 0 {
			return errors.New("Directory is not empty")
		}
	}
	temp, err := findAndParseTemplates(destdir, templatedir, template.FuncMap{})
	if err != nil {
		return err
	}
	anon := struct {
		Appname string
	}{
		Appname: name,
	}
	if temp != nil {
		for _, templ := range temp.Templates() {
			fileToCreate := filepath.Join(destdir, templ.Name())
			if err == nil {
				fil, err := os.Create(fileToCreate)
				if err != nil {
					return err
				}
				defer fil.Close()
				err = templ.Execute(fil, anon)
				if err != nil {
					return err
				}
			} else {
				return err
			}
		}
	}
	return nil
}

func findAndParseTemplates(destDir, rootDir string, funcMap template.FuncMap) (*template.Template, error) {
	cleanRoot := filepath.Clean(rootDir)
	pfx := len(cleanRoot) + 1
	root := template.New("")

	err := filepath.Walk(cleanRoot, func(path string, info os.FileInfo, e1 error) error {
		if path == cleanRoot {
			return nil
		}
		name := path[pfx:]
		if !info.IsDir() {
			if e1 != nil {
				return e1
			}

			b, e2 := ioutil.ReadFile(path)
			if e2 != nil {
				return e2
			}

			t := root.New(name).Funcs(funcMap)
			t, e2 = t.Parse(string(b))
			if e2 != nil {
				fmt.Println("Template not created", name)
				return e2
			}
		} else {
			fileToCreate := filepath.Join(destDir, name)
			err := os.Mkdir(fileToCreate, 0755)
			if err != nil {
				fmt.Println("Failed to create directory", fileToCreate)
				return err
			}
		}

		return nil
	})

	return root, err
}
