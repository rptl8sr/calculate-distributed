.PHONY: git-init git-checkout golangci-lint-install lint blueprint api-generate

REPO_NAME := $(shell basename $(CURDIR))
PROJECT := $(CURDIR)
LOCAL_BIN := $(CURDIR)/bin

ifneq (,$(wildcard ./.env))
ENV_FILE := .env
else
ENV_FILE :=
endif

ifneq ($(ENV_FILE),)
include $(ENV_FILE)
export
endif

# GIT
git-init:
	gh repo create $(USER)/$(REPO_NAME) --private
	git init
	git config user.name "$(USER)"
	git config user.email "$(EMAIL)"
	git add Makefile go.mod
	git commit -m "Init commit"
	git remote add origin git@github.com:$(USER)/$(REPO_NAME).git
	git remote -v
	git push -u origin master

BN ?= dev
.PHONY:
git-checkout:
	git checkout -b $(BN)

# LINT
lint-install:
	GOBIN=$(LOCAL_BIN) go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.62.2

lint:
	$(LOCAL_BIN)/golangci-lint run ./...


# PROJECT
blueprint:
	mkdir -p $(LOCAL_BIN)
	mkdir -p api
	mkdir -p cmd/orchestrator && echo 'package main' > cmd/orchestrator/main.go
	mkdir -p cmd/agent && echo 'package main' > cmd/agent/main.go
	mkdir -p internal/model && echo 'package model' > internal/model/model.go
	mkdir -p internal/logger && echo 'package logger' > internal/logger/logger.go

api-generate:
	go mod tidy
	go mod download
	go generate ./...

