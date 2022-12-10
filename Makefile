.PHONY: proto

proto:
	rm -rf web/src/gen
	buf generate

proto_lint:
	buf lint

proto_fmt:
	buf format -w
