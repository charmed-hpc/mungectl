# Copyright 2024 Canonical Ltd.
#
# This program is free software: you can redistribute it and/or modify
# it under the terms of the GNU General Public License as published by
# the Free Software Foundation, either version 3 of the License, or
# (at your option) any later version.
#
# This program is distributed in the hope that it will be useful,
# but WITHOUT ANY WARRANTY; without even the implied warranty of
# MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
# GNU General Public License for more details.
#
# You should have received a copy of the GNU General Public License
# along with this program.  If not, see <https://www.gnu.org/licenses/>.

project := "mungectl"
org_path := "github.com/charmed-hpc"
repo_path := org_path / project

build_dir := "_build"

# build mungectl binary
build:
  @rm -rf {{build_dir}}
  @mkdir {{build_dir}}
  go build -o {{build_dir}}/mungectl {{repo_path}}/cmd/mungectl

# apply formatting to project
fmt:
  golangci-lint -v run --fix

# check project against coding style standards
lint:
  golangci-lint -v run

# run all mungectl tests
test: unit

# run mungectl unit tests
unit:
  go test -v ./...

# clean up project
clean:
  rm -rf {{build_dir}}

# show available recipes
help:
  @just --list --unsorted
