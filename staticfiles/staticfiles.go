package staticfiles

import (
    "fmt"
    "net/http"
    "github.com/julienschmidt/httprouter"

    "path/filepath"
    "os"
    "strings"

    "github.com/firedrake969/karts/views"
    "mime"
)

const StaticDirectories = "/static"


// STATICFILES

type Staticfile struct {
    Servedpath string
    Realpath string
}

func (file Staticfile) Serve(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
    filetype := strings.SplitAfter(file.Realpath, ".")
    mimetype := mime.TypeByExtension("." + filetype[len(filetype)-1])
    w.Header().Set("Content-Type", mimetype + ", charset=UTF-8")
    fmt.Fprint(w, GetStaticfile(file.Realpath))
}

func GetStaticfile(path string) string {
    return views.GetTemplateContents(path)
}

func GetStaticfiles() []Staticfile {
    searchdir, _ := filepath.Abs(filepath.Dir(os.Args[0]))
    staticfiles := []string{}
    staticstructs := []Staticfile{}
    filepath.Walk(searchdir, func(fp string, fi os.FileInfo, err error) error {
        if err != nil {
            fmt.Println(err) // can't walk here,
            return nil       // but continue walking elsewhere
        }
        if !!fi.IsDir() {
            return nil // not a file.  ignore.
        }
        fp = strings.Replace(fp, "\\", "/", -1)
        matched, err := filepath.Match("*" + StaticDirectories + "/*.*", fp)
        if err != nil {
            fmt.Println(err) // malformed pattern
            return err       // this is fatal.
        }
        if matched {
            staticfiles = append(staticfiles, fp)
        }
        return nil
    })
    for file := range staticfiles {
        servedpath := Staticfile{
            Servedpath: "/static" + strings.SplitAfterN(staticfiles[file], "/static", 2)[1],
            Realpath: staticfiles[file],
        }
        staticstructs = append(staticstructs, servedpath)
    }
    return staticstructs
}