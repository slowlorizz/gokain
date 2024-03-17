# gokain

 A Hash Cracker Utility written in Go


## Usage

### v1.0.0
```shell
.\gokain.exe -th <amount of threads to run> --type <sha1 | sha256 | sha512 | md5> --hash <your hash-string>
```

### v2.0.0
```shell
.\gokain.exe --job <job-name>
```

... needs coresponding job-file in /worker/jobs with name `<job-name>.job.yml`
