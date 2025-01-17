SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
TEMP_DIR=$SCRIPT_DIR/../temp
BIN_DIR=$SCRIPT_DIR/../bin

mkdir -p $TEMP_DIR
mkdir -p $BIN_DIR

cd $TEMP_DIR
# curl
curl -O https://artifacts.elastic.co/downloads/kibana/kibana-8.17.0-linux-x86_64.tar.gz
curl -O https://artifacts.elastic.co/downloads/kibana/kibana-8.17.0-linux-x86_64.tar.gz.sha512

shasum -a 512 -c kibana-8.17.0-linux-x86_64.tar.gz.sha512
tar -xzf ${TEMP_DIR}/kibana-8.17.0-linux-x86_64.tar.gz -C ${BIN_DIR}
rm -rf $TEMP_DIR
