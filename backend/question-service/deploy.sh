#!/bin/bash

# Параметры по умолчанию
REMOTE_USER=""
REMOTE_HOST=""
LOCAL_CONFIG="./config/config.yaml"
REMOTE_PATH="~/app/bin"
REMOTE_CONFIG="~/app/config/config.yaml"
LOCAL_APP_PATH="./build/app"

# Функция для вывода справки
usage() {
    echo "Usage: $0 [OPTIONS]"
    echo "Options:"
    echo "  -u, --username      Username for SSH connection"
    echo "  -h, --host          Hostname or IP address of the remote server"
    echo "  -c, --config        Local config path"
    echo "  -r, --remote-path   Remote path to deploy the application"
    echo "  -l, --local-path    Local path to your web application"
    echo "  -help               Display this help and exit"
}

# Обработка аргументов командной строки
while getopts ":u:h:c:r:l:" opt; do
    case ${opt} in
        u | --username )
            REMOTE_USER=$OPTARG
            ;;
        h | --host )
            REMOTE_HOST=$OPTARG
            ;;
        c | --config )
            LOCAL_CONFIG=$OPTARG
            ;;
        r | --remote-path )
            REMOTE_PATH=$OPTARG
            ;;
        l | --local-path )
            LOCAL_APP_PATH=$OPTARG
            ;;
        \? | : | * )
            usage
            exit 1
            ;;
    esac
done
shift $((OPTIND -1))

# Проверка обязательных параметров
if [ -z "$REMOTE_USER" ] || [ -z "$REMOTE_HOST" ]; then
    echo "Error: Username and host are required."
    usage
    exit 1
fi

# Функция для сборки приложения
build_app() {
    CGO_ENABLED=0 GOOS=linux go build -o ./build/app ./cmd/question/main.go
}

# Функция для копирования собранного приложения на удаленный сервер
deploy_app() {
    scp $LOCAL_APP_PATH $REMOTE_USER@$REMOTE_HOST:$REMOTE_PATH
    scp $LOCAL_CONFIG $REMOTE_USER@$REMOTE_HOST:$REMOTE_CONFIG
}


# Сборка приложения
build_app

# Доставка на удаленный сервер
deploy_app


echo "Deployment completed successfully."