SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
BIN_DIR=$SCRIPT_DIR/../bin/kibana-8.17.0/bin

${BIN_DIR}/kibana
