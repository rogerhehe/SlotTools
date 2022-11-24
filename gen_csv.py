# %%
import pandas as pd
import numpy as np
import json
import os
import os.path

# %%


def space_str(layer):
	lua_str = ""
	for i in range(0,layer):
		lua_str += '\t'
	return lua_str
 
def dic_to_lua_str(data,layer=0):
	d_type = type(data)
	if  d_type is str :
		return "'" + data + "'"
	elif d_type is bool:
		if data:
			return 'true'
		else:
			return 'false'
	elif d_type in (int,float):
		return str(data)
	elif d_type is list:
		lua_str = "{\n"
		lua_str += space_str(layer+1)
		for i in range(0,len(data)):
			lua_str += dic_to_lua_str(data[i],layer+1)
			if i < len(data)-1:
				lua_str += ','
		lua_str += '\n'
		lua_str += space_str(layer)
		lua_str +=  '}'
		return lua_str
	elif d_type is dict:
		lua_str = ''
		lua_str += "\n"
		lua_str += space_str(layer)
		lua_str += "{\n"
		data_len = len(data)
		data_count = 0
		for k,v in data.items():
			data_count += 1
			lua_str += space_str(layer+1)
			if type(k) in (int,float):
				lua_str += '[' + str(k) + ']'
			else:
				lua_str += k 
			lua_str += ' = '
			try:
				lua_str += dic_to_lua_str(v,layer +1)
				if data_count < data_len:
					lua_str += ',\n'
 
			except Exception as e:
				print( 'error in ',k,v)
				raise
		lua_str += ',\n'
		lua_str += space_str(layer)
		lua_str += '}'
		return lua_str
	else:
		print( d_type , 'is error')
		return None
 

# %%
def to_depth(data):
	root = {}
	keys = list(data.keys())
	cols = [str.split(c,'.') for c in keys]
	for i,c in enumerate(cols):
		base_ds = root
		if not pd.isna(data[keys[i]]):
			for j,d in enumerate(c):
				if d not in base_ds:
					base_ds[d] = {}
				if j < len(c)-1:
					base_ds = base_ds[d]
				else:
					base_ds[c[-1]] = data[keys[i]]
	return root

for fn in os.listdir('./csv'):
	if fn.find('.csv') > 0:
		df = pd.read_csv(f"./csv/{fn}")
		df = df.where(df.notnull(), None)

		if fn.find(".const.")>0:
			data = {r[0]:r[1] for r in df.values}
			with open(f'./json/{fn[:-4]}.json', 'w') as f:
				json.dump(data, f, ensure_ascii=True, indent='\t')
			with open(f'./lua/{str.split(fn,".")[0]}.lua', 'w') as f:
				f.write('local config = ' + dic_to_lua_str(data) + '\nreturn config')
				f.close()
			continue

		dtype = df.iloc[0]
		df = df[2:]

		for c in df.columns:
			if dtype[c] == 'int':
				df[c] = df[c].apply(lambda x: int(x) if x is not None else None)
			if dtype[c] == 'float':
				df[c] = df[c].apply(lambda x: float(x) if x is not None else None)
			if dtype[c] == 'list':
				df[c] = df[c].apply(lambda x: eval(x) if x is not None else None)
			if dtype[c] == 'dict':
				df[c] = df[c].apply(lambda x: eval(x) if x is not None else None)

		if fn.find('.list.') > 0:
			data = [to_depth(data) for data in list(df.T.to_dict().values())]
			with open(f'./json/{fn[:-4]}.json', 'w') as f:
				json.dump(data, f, ensure_ascii=True, indent='\t')
				f.close()
			with open(f'./lua/{str.split(fn,".")[0]}.lua', 'w') as f:
				f.write('local config = '+ dic_to_lua_str(data) + '\nreturn config')
				f.close()
		else:
			df.set_index(df.columns[0],inplace=True)
			data = {key:to_depth(data) for key,data in df.T.to_dict().items()}
			with open(f'./json/{fn[:-4]}.json', 'w') as f:
				json.dump(data, f, ensure_ascii=True, indent='\t')
				f.close()
			with open(f'./lua/{str.split(fn,".")[0]}.lua', 'w') as f:
				f.write('local config = '+ dic_to_lua_str(data) + '\nreturn config')
				f.close()



