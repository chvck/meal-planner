PHONY: proto-gen
proto-gen:
	prototool generate

	protoc --plugin=protoc-gen-ts=./frontend/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:$(PWD)/frontend/src/ \
	 --ts_out=service=true:$(PWD)/frontend/src/ -I $(PWD) $(PWD)/proto/model/model.proto

	protoc --plugin=protoc-gen-ts=./frontend/node_modules/.bin/protoc-gen-ts --js_out=import_style=commonjs,binary:$(PWD)/frontend/src/ \
	    --ts_out=service=true:$(PWD)/frontend/src/ -I $(PWD) $(PWD)/proto/service/service.proto
