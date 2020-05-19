# Use the official Golang image to create a build artifact.
    FROM golang:1.14

    # Create and change to the app directory.
    WORKDIR /go/src

    # Copy go.mod & go.sum
    COPY go.* ./
    # GoModuleを使うため、念の為 GO111MODULE をonにする.
    RUN export GO111MODULE=on && go mod download

    # 必要なもののみUP.
    COPY ./app ./app

    # app配下を実行ファイルに変換？
    RUN go install ./app

    # 実行
    CMD ["/go/bin/app"]
