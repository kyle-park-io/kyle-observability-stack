SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
BIN_DIR=$SCRIPT_DIR/../bin/elasticsearch-8.17.0/bin

${BIN_DIR}/elasticsearch-create-enrollment-token --scope kibana
