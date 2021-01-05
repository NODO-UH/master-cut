package api

import (
	"fmt"
	"io/ioutil"
	"strings"

	"net/http"
	"os"
	"os/exec"

	"errors"

	"github.com/NODO-UH/master-cut/src/conf"
	"github.com/gin-gonic/gin"
)

const (
	ErrOpenFile = "open file err"
	ErrReadAll  = "read all error"
)

// Cut handle POST /cust?group=[group]&user=[user]
func Cut(ctx *gin.Context) {
	if user, ok := ctx.GetQuery("user"); !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("user required"))
	} else if group, ok := ctx.GetQuery("group"); !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("group required"))
	} else if group := conf.GetGroup(group); group != nil {
		f, err := os.OpenFile(*group.File, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0600)
		if err != nil {
			fmt.Println(err)
			return
		}
		defer f.Close()

		f.WriteString(fmt.Sprintf("%s\n", user))
		f.Sync()
		// Run cut script
		cmd := exec.Command("sh", *group.Script)
		cmd.Run()
	}
}

func Uncut(ctx *gin.Context) {
	if user, ok := ctx.GetQuery("user"); !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("user required"))
	} else if group, ok := ctx.GetQuery("group"); !ok {
		ctx.AbortWithError(http.StatusBadRequest, errors.New("group required"))
	} else if group := conf.GetGroup(group); group != nil {
		f, err := os.OpenFile(*group.File, os.O_CREATE|os.O_RDWR, 0666)
		if err != nil {
			fmt.Println(err)
			ctx.AbortWithError(http.StatusInternalServerError, errors.New(ErrOpenFile))
			return
		}
		defer f.Close()

		if data, err := ioutil.ReadAll(f); err != nil {
			fmt.Println(err)
			ctx.AbortWithError(http.StatusInternalServerError, errors.New(ErrReadAll))
		} else {
			lines := strings.Split(string(data), "\n")
			fmt.Println(lines)
			newLines := []string{}
			for _, l := range lines {
				if l != user {
					newLines = append(newLines, l)
				}
			}
			f.Truncate(0)
			f.Seek(0, 0)
			f.WriteString(strings.Join(newLines, "\n"))

			// Run cut script
			cmd := exec.Command("sh", *group.Script)
			cmd.Run()
		}
	}
}
