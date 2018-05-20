# -*- coding: utf-8 -*-
"""
Spyder Editor

"Valar Morghulis"
"""

import pandas as pd
import seaborn as sns
import xlwings as xw
from sklearn.model_selection import cross_val_score, train_test_split, GridSearchCV
from sklearn.preprocessing import OneHotEncoder, MinMaxScaler
from sklearn.metrics import confusion_matrix
from sklearn import svm
import matplotlib.pyplot as plt
import math, time, pickle
from imblearn.over_sampling import SMOTE

start = time.time()

# ======================== Decimal-to-Percent =================================

def percentage (decimal):
# Mengubah Decimal menjadi percent dalam string
    return ("%.1f" % (decimal * 100));

# =============================================================================




# ============================ LOAD Excel w/ Pass =============================
PATH = "SIJEKH.xlsx"
wb = xw.Book(PATH)
sheet = wb.sheets['SIJEKH']

df = sheet['A1:O34925'].options(pd.DataFrame, index=False, header=True).value

del df['nilai_likuidasi']
#del df['setia']
del df['pemasukan_tambahan']

# Drop rows that contain NaN, infinite, or too large. 
# Inplace to directly change object if True or throw results if False
df.dropna(axis=0, how='any', inplace=True)
df.fillna(0, inplace=True)

# ============================ VISUALISASI DATA ===============================

#sns.set(style="darkgrid")
#plt.figure(figsize=(14,7))
#dfx = df.umur
#x = sns.distplot(dfx, color='g')
#x = sns.distplot(dfx[dfx < 5000000], color='navy')
#print(df['jenis_kelamin'].value_counts())
#flatui = ["#9b59b6", "#3498db", "#95a5a6", "#e74c3c", "#34495e", "#2ecc71"]
#ax = sns.countplot(x='jenis_kelamin', data=df, order=["F", "M"],  orient="v", palette="husl")


# ========================== Separasi Data Target =============================

df_target = df.iloc[:, 12] #iloc[element_of_feature, key_of_feature]
del df['Kategori']

#print(df_target)

# ============================ One-Hot Encoding ===============================

df = pd.get_dummies(df, 
            columns=["punya_bisnis", "status_perkawinan", "jenis_kelamin", "pendidikan", "pekerjaan"], 
            prefix=["biz", "mariage", "gender", "edu", "job"])

# ============================ LEARNING PART===================================

parameters = [
  {'C': [1, 10, 100, 1000], 'gamma': [1/11, 0.001, 0.0001], 'kernel': ['rbf']}
 ]

#print(len(list(df)))

#svc = svm.SVC()
#clf = GridSearchCV(svc, parameters)

clf = svm.SVC(C= 1, cache_size= 200, class_weight= None, coef0= 0.0, decision_function_shape= 'ovr', degree= 3, gamma= 0.1, kernel= 'rbf', max_iter= -1, probability= False, random_state= None, shrinking= True, tol= 0.001, verbose= False)
# ============================= Split-Test ====================================
X_train, X_test, y_train, y_test = train_test_split(df, df_target, test_size=0.3, random_state=0)

#clf.fit(df, df_target)
sm = SMOTE(random_state=42)

X_res, y_res = sm.fit_sample(df, df_target)

y_pred = clf.fit(X_res, y_res).predict(df)

cm = confusion_matrix(df_target, y_pred)
#cm = [[6768, 105], [1181, 25484]]

print(cm)

TN = cm[0][0]
TP = cm[1][1]
FP = cm[0][1]
FN = cm[1][0]

Akurasi = (TP+TN)/(TP+TN+FP+FN)
Presisi = (TP)/(TP+FP)
Recall = (TP)/(TP+FN)
Spesitifitas = (TN)/(TN+FP)

print("Akurasi: ", percentage(Akurasi), "%")
print("Presisi: ", percentage(Presisi), "%")
print("Recall: ", percentage(Recall), "%")
print("Spesitifitas: ", percentage(Spesitifitas), "%")

#print(clf.score(df, df_target))

end = time.time()
time_taken = int(math.ceil(end - start))
print("Time: ", int(math.floor(time_taken / 60)), " minutes & ", time_taken % 60, " seconds.")
# =============================================================================

# ===================== Save / Load Classifier Object =========================
# 

file = open("SVM.obj", "wb")
pickle.dump(clf,file)

#file = open("SVM.obj",'rb')
#object_file = pickle.load(file)

file.close()

# =============================================================================

# ====================== K-Fold Cross Validation ==============================
# clf.fit(df, df_target)
# 
# result = cross_val_score(clf, df, df_target, cv=10)
#     
# print("Hello, World!\n")
# print(result)
# print("\n")
# print("Accuracy: %0.2f (+/- %0.2f)" % (result.mean(), result.std() * 2))
# =============================================================================


# ============================ UNTUK CEK FILES ================================
# import os
# 
# for root, dirs, files in os.walk("."):  
#     for filename in files:
#         print(filename)
# =============================================================================

#Best parameters set:
#(C= 1, cache_size= 200, class_weight= None, coef0= 0.0, decision_function_shape= 'ovr', degree= 3, gamma= 0.09090909090909091, kernel= 'rbf', max_iter= -1, probability= False, random_state= None, shrinking= True, tol= 0.001, verbose= False)

#print("Best parameters set:")
#best_parameters = clf.best_estimator_.get_params()
#print(best_parameters)