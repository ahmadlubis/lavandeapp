# -*- coding: utf-8 -*-
"""
Spyder Editor

This is a temporary script file.
"""

import pandas as pd
import seaborn as sns
import xlwings as xw
from sklearn.model_selection import cross_val_score, train_test_split
from sklearn.preprocessing import OneHotEncoder, MinMaxScaler
from sklearn.metrics import confusion_matrix
from sklearn import svm
import matplotlib.pyplot as plt


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

df_target = df.iloc[:, 11] #iloc[element_of_feature, key_of_feature]
del df['Kategori']

#print(df_target)

# ============================ One-Hot Encoding ===============================

df = pd.get_dummies(df, 
            columns=["punya_bisnis", "status_perkawinan", "jenis_kelamin", "pendidikan", "pekerjaan"], 
            prefix=["biz", "mariage", "gender", "edu", "job"])

# ============================ LEARNING PART===================================

clf = svm.SVC()

# ============================= Split-Test ====================================
X_train, X_test, y_train, y_test = train_test_split(df, df_target, test_size=0.3, random_state=0)

y_pred = clf.fit(df, df_target).predict(df)

print(confusion_matrix(df_target, y_pred))

cm = confusion_matrix(df_target, y_pred)

#cm = [[6664, 209], [1398, 25267]]

TN = cm[0][0]
TP = cm[1][1]
FP = cm[0][1]
FN = cm[1][0]

Akurasi = (TP+TN)/(TP+TN+FP+FN)
Presisi = (TP)/(TP+FP)
Recall = (TP)/(TP+FN)
Spesitifitas = (TN)/(TN+FP)

print("Akurasi: ", Akurasi)
print("Presisi: ", Presisi)
print("Recall: ", Recall)
print("Spesitifitas: ", Spesitifitas)

#print(clf.score(df, df_target))


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

