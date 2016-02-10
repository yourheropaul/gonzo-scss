package scss

import (
	"bytes"
	"errors"
	"io/ioutil"
	"path/filepath"
	"strings"

	sass "github.com/yourheropaul/gonzo-scss/internal/libsass"

	"github.com/omeid/gonzo"
	"github.com/omeid/gonzo/context"
)

// Config doesn't exist. Code over Configuration please.
type Config struct {
}

func Compile() gonzo.Stage {
	return func(ctx context.Context, in <-chan gonzo.File, out chan<- gonzo.File) error {

		if err := func() error {
			for {
				select {
				case file, ok := <-in:
					if !ok {
						return nil
					}

					source, err := ioutil.ReadAll(file)
					if err != nil {
						return err
					}

					dc := sass.NewDataContext(string(source))
					defer dc.Destroy()
					opts := dc.Context().Options()

					if string(filepath.Base(file.FileInfo().Name())[0]) == "_" {
						break
					}

					base := filepath.Join(
						file.FileInfo().Base(),
					)

					opts.SetIncludePath(base)

					compiler := dc.Compiler()
					defer compiler.Destroy()

					err = compiler.Parse()
					if err != nil {
						ctx.Error(err)
						//ctx.Error(dc.Context().Error().Error())
						return errors.New(dc.Context().Error().Error())
					}

					err = compiler.Execute()
					if err != nil {
						ctx.Error(err)
						return errors.New(dc.Context().Error().Error())
					}

					output := dc.Context().OutputString()

					buff := bytes.NewBufferString(output)

					name := strings.TrimSuffix(file.FileInfo().Name(), ".scss") + ".css"

					ctx.Infof("Compiling %s to %s", file.FileInfo().Name(), name)

					file = gonzo.NewFile(ioutil.NopCloser(buff), file.FileInfo())
					file.FileInfo().SetSize(int64(buff.Len()))
					file.FileInfo().SetName(name)

					out <- file
				case <-ctx.Done():
					return nil
				}
			}
		}(); err != nil {
			return err
		}

		return nil
	}
}
