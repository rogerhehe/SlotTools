@echo off
@REM csv to json lua
set POJECT_PATH=..\UnityClient
set S_PROJECT_PATH=..\SlotServer


set CURPATH= %~dp0

cd %CURPATH%
set GEN_PATH=%CURPATH%out\table
set GEN_PATH_LUA=%GEN_PATH%\lua
set GEN_PATH_JSON=%GEN_PATH%\json
@REM client
set C_OUTPUT_PATH_LUA=%POJECT_PATH%\Assets\_LuaScripts\Game\Config\
set C_OUTPUT_PATH_JSON=%POJECT_PATH%\Assets\_Json\
@REM server
set S_OUTPUT_PATH_JSON=%S_PROJECT_PATH%\common\fileconfig\json\


rd /s /q %GEN_PATH%

python bin\gen_csv.py --input_dir %CURPATH%csv --output_dir %GEN_PATH%
echo "GENERATE CONFIG DONE"

@REM copy files to repo
rd /s /q %C_OUTPUT_PATH_JSON% %C_OUTPUT_PATH_LUA%
xcopy /Y /E %GEN_PATH_JSON% %C_OUTPUT_PATH_JSON%
xcopy /Y /E %GEN_PATH_LUA% %C_OUTPUT_PATH_LUA%
rd /s /q %S_OUTPUT_PATH_JSON%
xcopy /Y /E %GEN_PATH_JSON% %S_OUTPUT_PATH_JSON%

echo "CPOY CONFIG DONE"