#!/bin/bash

source "$(dirname "${BASH_SOURCE[0]}")/setup.sh"

PROTOC_GEN_DOC_VERSION="v1.5.1"
GOMARKDOC_VERSION="v1.1.0"

FOX_ROOT="${REPO_ROOT}/../fox"
KIT_DOCS_GO="${DOCS_SRC}/reference/kit/go"

rm -rf ${DOCS_OUT}

### Install tools. If wrong version is installed, tool will be overwritten.
${TOOLS_DIR}/protoc-gen-doc -version | grep -q ${PROTOC_GEN_DOC_VERSION} ||
    go install github.com/pseudomuto/protoc-gen-doc/cmd/protoc-gen-doc@${PROTOC_GEN_DOC_VERSION}

${TOOLS_DIR}/gomarkdoc --version | grep -q ${GOMARKDOC_VERSION} ||
    go install github.com/princjef/gomarkdoc/cmd/gomarkdoc@${GOMARKDOC_VERSION}
###

# Generate docs from proto files.
protoc \
    --proto_path=./${PROTO_SRC} \
    --doc_out=./${DOCS_SRC}/reference/ --doc_opt=./${DOCS_SRC}/protobuf.tmpl,protobuf.md \
    protobuf_msgs.proto broker_svc.proto

if [ -d "${FOX_ROOT}" ]; then
    (
        cd "${FOX_ROOT}"
        make docs
    )
    rm -rf ${DOCS_SRC}/reference/fox
    cp -r ${FOX_ROOT}/docs ${DOCS_SRC}/reference/fox
fi

# Generate docs from source files.
rm -rf "${KIT_DOCS_GO}"
mkdir -p "${KIT_DOCS_GO}"

find "${REPO_ROOT}/kit" -type d -exec bash -c '
gomarkdoc --output "${0}/$(basename "${1}").md" "${1}"
' "${KIT_DOCS_GO}" {} \;

pipenv install
pipenv run mkdocs build
