# csv to json lua
POJECT_PATH="../UnityClient"
S_PROJECT_PATH="../SlotServer"


CURPATH=$(cd "$(dirname "$0")"; pwd)
cd $CURPATH
GEN_PATH=$CURPATH/out/table
GEN_PATH_LUA=$GEN_PATH'/lua'
GEN_PATH_JSON=$GEN_PATH'/json'
# client
C_OUTPUT_PATH_LUA=$POJECT_PATH'/Assets/_LuaScripts/Game/Config'
C_OUTPUT_PATH_JSON=$POJECT_PATH'/Assets/_Json'
# server
S_OUTPUT_PATH_JSON=$S_PROJECT_PATH'/common/fileconfig/json'


rm -rf $GEN_PATH

python3 bin/gen_csv.py --input_dir $CURPATH/csv --output_dir $GEN_PATH
echo "GENERATE CONFIG DONE"

# copy files to repo
rm -rf $C_OUTPUT_PATH_JSON $C_OUTPUT_PATH_LUA
cp -r $GEN_PATH_JSON $C_OUTPUT_PATH_JSON
cp -r $GEN_PATH_LUA $C_OUTPUT_PATH_LUA
cp -r $GEN_PATH_JSON $S_OUTPUT_PATH_JSON

echo "CPOY CONFIG DONE"