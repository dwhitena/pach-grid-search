import numpy as np
import pandas as pd
import json
from os import listdir
from os.path import isfile, join
import sklearn.cross_validation as cv
import sklearn.metrics as metrics
from sklearn.ensemble import RandomForestClassifier

# Read the data in from PFS.
df = pd.read_csv("/pfs/training/training.csv", delimiter=";")

# List the files this container has access to.
param_dir = "/pfs/filter"
param_files = [f for f in listdir(param_dir) if isfile(join(param_dir, f))]

# Loop over the given parameters.
for param_file in param_files:
    
    # Unmarshal the parameters.
    with open(join(param_dir, param_file), "r") as f:    
        params_str = f.read()
        params = json.loads(params_str)

    # Construct the model.
    rf = RandomForestClassifier(n_estimators=int(params["n_estimators"]), 
            max_features=int(params["max_features"]))
    
    # Perform cross validation.
    score_list = cv.cross_val_score(rf, df[list(df.columns.values)[0:-1]], 
            df[list(df.columns.values)[-1]], scoring="accuracy", cv=10)
    score = np.mean(score_list)

    # Output the results.
    with open("/pfs/out/results.csv", "a") as myfile:
        output = "{0}, {1}\n".format(params_str.rstrip("\n"), score)
        myfile.write(output)

