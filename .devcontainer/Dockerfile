# to run member oauth
ARG GOLANG_VERSION=1.22.4
FROM golang:${GOLANG_VERSION}
EXPOSE 8888
RUN apt-get update && apt-get install -y graphviz hey
# gopls - go language server (https://pkg.go.dev/golang.org/x/tools/gopls)
RUN go install -v golang.org/x/tools/gopls@latest
# dlv -  debugger for golang (https://github.com/go-delve/delve)
RUN go install -v github.com/go-delve/delve/cmd/dlv@latest
# staticcheck - linter for golang (https://staticcheck.io/)
RUN go install -v honnef.co/go/tools/cmd/staticcheck@latest
# digraph - used for analysing dependency graph
RUN go install golang.org/x/tools/cmd/digraph@latest
# modgraphviz - used for visualizing dependency graph 
RUN go install golang.org/x/exp/cmd/modgraphviz@latest
# mockgen - used for generating mocks
RUN go install github.com/golang/mock/mockgen@v1.6.0
# govulncheck - used for checking vulnerabilities in go modules
RUN go install golang.org/x/vuln/cmd/govulncheck@latest
# pprof - used for profiling go programs
RUN go install github.com/google/pprof@latest
