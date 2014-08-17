all: test

serve:
	cd appengine; goapp serve

deploy:
	cd appengine; goapp deploy -oauth

go_tests: proto
	goapp test bitbucket.org/bugloco/bugloco/...

test: go_tests

convey:
	cd goapp; goconvey -gobin=goapp -port=8081

.PHONY: all proto serve deploy test go_tests convey
