package tools

//go:generate docker run --rm -v $PROJECT_HOME:/app/project --platform linux/amd64 gitea.linuxcode.net/linuxcode/dbml-go ./dbml-go-generator -f ./project/resources/dbml/database.dbml -p models -o ./project/pkg/models
