package args

import (
	"flag"
	"fmt"

	"github.com/slowlorizz/gokain/worker/src/exc"
)

var JobPath string

func Load() *exc.Exception {
	jn := flag.String("job", "", "name of the Job to run, its the file name <job-name>.job.yml, without the extensions")

	flag.Parse()

	if *jn == "" {
		return exc.NewFatal(exc.UserArgsException, "no_job_filename", "No Job-Filename has been submitted", nil)
	}

	JobPath = fmt.Sprintf("./jobs/%s.job.yml", *jn)

	return nil
}
