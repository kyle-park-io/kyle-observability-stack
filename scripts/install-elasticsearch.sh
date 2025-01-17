SCRIPT_DIR=$(dirname "$(readlink -f "$0")")
TEMP_DIR=$SCRIPT_DIR/../temp
BIN_DIR=$SCRIPT_DIR/../bin

mkdir -p $TEMP_DIR
mkdir -p $BIN_DIR

# wget
wget -P $TEMP_DIR https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.17.0-linux-x86_64.tar.gz
wget -P $TEMP_DIR https://artifacts.elastic.co/downloads/elasticsearch/elasticsearch-8.17.0-linux-x86_64.tar.gz.sha512

cd ${TEMP_DIR}
shasum -a 512 -c elasticsearch-8.17.0-linux-x86_64.tar.gz.sha512
tar -xzf ${TEMP_DIR}/elasticsearch-8.17.0-linux-x86_64.tar.gz -C ${BIN_DIR}
rm -rf $TEMP_DIR
