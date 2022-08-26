#!/bin/sh
# shellcheck disable=SC2039
set -o errexit -eo pipefail

cd internal/tools

go install \
  -tags 'postgres' github.com/golang-migrate/migrate/v4/cmd/migrate


go install \
  github.com/deepmap/oapi-codegen/cmd/oapi-codegen \
   github.com/golang/mock/mockgen@v1.4.4
  github.com/fdaines/spm-go \
  github.com/golangci/golangci-lint/cmd/golangci-lint \
  github.com/kyleconroy/sqlc/cmd/sqlc@latest \
  github.com/maxbrunsfeld/counterfeiter/v6 \
  goa.design/model/cmd/mdl \
  goa.design/model/cmd/stz