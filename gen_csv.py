# %%
import pandas as pd
import json
import os


# %%
os.makedirs('json',exist_ok = True)

for fn in os.listdir('./csv'):
    if fn.find('.csv') > 0:
        df = pd.read_csv(f"./csv/{fn}")

        if fn.find(".const.")>0:
            data = {r[0]:r[1] for r in df.values}
            with open(f'./json/{fn[:-4]}.json', 'w') as f:
                json.dump(data, f, ensure_ascii=True, indent='\t')
            continue

        dtype = df.iloc[0]
        df = df[2:]

        for c in df.columns:
            if dtype[c] == 'int':
                df[c] = df[c].apply(int)
            if dtype[c] == 'float':
                df[c] = df[c].astype(float)
            if dtype[c] == 'list':
                df[c] = df[c].apply(eval)
            if dtype[c] == 'dict':
                df[c] = df[c].apply(eval)

        if fn.find('.list.') > 0:
            with open(f'./json/{fn[:-4]}.json', 'w') as f:
                json.dump(list(df.T.to_dict().values()), f, ensure_ascii=True, indent='\t')
        else:
            df.set_index(df.columns[0],inplace=True)
            with open(f'./json/{fn[:-4]}.json', 'w') as f:
                json.dump(df.T.to_dict(), f, ensure_ascii=True, indent='\t')
