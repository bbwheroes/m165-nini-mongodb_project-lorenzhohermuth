# M165 MongoDB Quiz

## How to Install

```bash
# clone repo
git clone https://github.com/bbwheroes/m165-nini-mongodb_quiz-lorenzhohermuth.git
# If is not set
# path has to be in $PATH env var
export PATH=$PATH:/path/to/your/install/directory #Linux
set PATH=%PATH%;C:\path\to\your\install\directory #Windows
# or add to your enviroment var
# or in .bashrc
# set GOBIN env var
go env -w GOBIN=/path/to/your/bin
# to install
cd m165-nini-mongodb_quiz-lorenzhohermuth && go install cmd/mongodb_quiz/pokemon-quiz.go
# db
docker compose up
# run app
pokemon-quiz
```
