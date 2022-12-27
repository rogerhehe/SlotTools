CURPATH=$(cd "$(dirname "$0")"; pwd)
cd $CURPATH

POJECT_PATH="../UnityClient"
GEN_PATH_LUA='./lua'
GEN_PATH_JSON='./json'
OUTPUT_PATH_LUA='../UnityClient/Assets/_LuaScripts/Game/Config'
OUTPUT_PATH_JSON='../UnityClient/Assets/_Json'
rm -rf $GEN_PATH_JSON $GEN_PATH_LUA

python3 gen_csv.py
echo "GENERATE CONFIG DONE"

rm -rf $OUTPUT_PATH_JSON $OUTPUT_PATH_LUA
cp -r $GEN_PATH_JSON $OUTPUT_PATH_JSON
cp -r $GEN_PATH_LUA $OUTPUT_PATH_LUA

echo "CPOY CONFIG DONE"