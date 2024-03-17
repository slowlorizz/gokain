package args

import (
	"flag"

	"github.com/slowlorizz/gokain/worker/src/exc"
)

var JobPath string

func Load() *exc.Exception {
	jn := flag.String("job", "", "file path to the <name>.job.yml file")

	flag.Parse()

	if *jn == "" {
		return exc.NewFatal(exc.UserArgsException, "no_job_filepath", "No Job-Filepath has been submitted", nil)
	}

	JobPath = *jn

	return nil
}
