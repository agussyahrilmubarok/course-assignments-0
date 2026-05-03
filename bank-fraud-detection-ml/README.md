# 🤖 Machine Learning Project — Bank Transaction Fraud Detection
### Dicoding | Building Machine Learning Projects (BMLP) | Final Submission

---

## 📌 Project Overview

This project integrates **Unsupervised Learning (Clustering)** and **Supervised Learning (Classification)** to analyze bank transaction data originally sourced from the Kaggle "Bank Transaction Dataset for Fraud Detection" (modified version provided by Dicoding).

The workflow:
1. Apply **K-Means Clustering** to generate labels from unlabeled data
2. Use those labels as the target variable to train a **Classification model**

---

## 🗂️ Project Structure

```
BMLP_Agus-Syahril-Mubarok.zip
├── [Clustering]_Submission_Akhir_BMLP_Agus_Syahril_Mubarok.ipynb
├── [Klasifikasi]_Submission_Akhir_BMLP_Agus_Syahril_Mubarok.ipynb
├── model_clustering.h5
├── PCA_model_clustering.h5          # Optional
├── decision_tree_model.h5
├── explore_RandomForest_classification.h5   # Optional
├── tuning_classification.h5         # Optional
├── data_clustering.csv
└── data_clustering_inverse.csv      # Optional
```

---

## 📒 Notebook 1: Clustering

### Workflow
```
Load Dataset → EDA → Data Cleaning → Preprocessing →
Elbow Method → K-Means → Evaluate → Interpret → Export
```

### Steps

| # | Step | Description |
|---|------|-------------|
| 1 | **Import Library** | Do not modify — run as provided |
| 2 | **Load Dataset** | Load the modified Bank Transaction dataset |
| 3 | **EDA** | `head()`, `info()`, `describe()`, correlation matrix, histograms, boxplot |
| 4 | **Data Cleaning** | Check nulls (`isnull().sum()`), duplicates (`duplicated().sum()`), drop them |
| 5 | **Preprocessing** | Drop ID/Date columns, `LabelEncoder`, outlier handling (IQR), `StandardScaler`, binning |
| 6 | **Elbow Method** | Use `KElbowVisualizer()` to find optimal K |
| 7 | **K-Means** | Train model → save as `model_clustering.h5` |
| 8 | **Evaluation** | Silhouette Score + cluster visualization (PCA 2D) |
| 9 | **PCA** *(Advanced)* | Build PCA model → save as `PCA_model_clustering.h5` |
| 10 | **Interpretation** | Descriptive statistics per cluster (mean, min, max) — before & after inverse |
| 11 | **Inverse Transform** | Reverse `StandardScaler` and `LabelEncoder` back to original values |
| 12 | **Export** | Save `data_clustering.csv` and `data_clustering_inverse.csv` |

---

## 📒 Notebook 2: Classification

### Workflow
```
Load Clustering Result → Feature Encoding → Train-Test Split →
Decision Tree → Random Forest → Evaluate → Tuning → Export
```

### Steps

| # | Step | Description |
|---|------|-------------|
| 1 | **Import Library** | Standard sklearn imports — do not modify |
| 2 | **Load Dataset** | Load `data_clustering_inverse.csv` (contains `Target` column) |
| 3 | **Feature Encoding** | One-Hot Encoding via `pd.get_dummies()` for categorical columns |
| 4 | **Train-Test Split** | `train_test_split()` — 80% train, 20% test |
| 5 | **Decision Tree** | Build model → evaluate → save as `decision_tree_model.h5` |
| 6 | **Random Forest** *(Skilled)* | Extra algorithm → save as `explore_RandomForest_classification.h5` |
| 7 | **Evaluation** | Accuracy, Precision, Recall, F1-Score for all models |
| 8 | **Hyperparameter Tuning** *(Advanced)* | `RandomizedSearchCV` on Random Forest → save as `tuning_classification.h5` |

---

## 📊 Grading Criteria

| Criteria | Basic (2pts) | Skilled (3pts) | Advanced (4pts) |
|----------|-------------|----------------|-----------------|
| **1. EDA** | head, info, describe | + correlation matrix, histograms | + no overlapping labels, boxplot |
| **2. Preprocessing** | null/dup check, drop ID/Date, LabelEncoder | + outlier drop (IQR), StandardScaler | + binning 1-2 features |
| **3. Clustering** | Elbow + KMeans + save model | + Silhouette Score + viz | + PCA model |
| **4. Interpretation** | Descriptive stats + export CSV | + inverse transform + num & cat analysis | + save inverse CSV |
| **5. Classification** | Decision Tree + save model | + extra algorithm + evaluation metrics | + hyperparameter tuning |

### Score Formula
```
Final Score = Total Points / Number of Criteria (5)
```

| Final Score | Stars | Grade | Level |
|-------------|-------|-------|-------|
| < 1 | ❌ Rejected | E | — |
| 1 – <2 | ⭐⭐ | D | Below Basic |
| 2 – <3 | ⭐⭐⭐ | C | Basic |
| 3 – <4 | ⭐⭐⭐⭐ | B | Skilled |
| 4 | ⭐⭐⭐⭐⭐ | A | Advanced |

---

## 📊 Cluster Analysis Results

### Cluster Distribution
| Cluster | Jumlah Data |
|---------|------------|
| Cluster 0 | 555 |
| Cluster 1 | 690 |
| Cluster 2 | 700 |

### Cluster Characteristics (After Inverse Transform)

| Feature | Cluster 0 | Cluster 1 | Cluster 2 |
|---------|-----------|-----------|-----------|
| **Persona** | Nasabah Standar — Usia Tua | Nasabah Aktif — Usia Muda | Nasabah Profesional — Saldo Tertinggi |
| **TransactionAmount (mean)** | 254.56 | 258.21 | 257.29 |
| **CustomerAge (mean)** | 44.84 thn | 43.85 thn | 45.41 thn |
| **TransactionDuration (mean)** | 119.74 dtk | 117.31 dtk | 120.70 dtk |
| **AccountBalance (mean)** | 5,001.21 | 5,109.80 | 5,170.92 |
| **LoginAttempts (mean)** | 1.0 | 1.0 | 1.0 |
| **TransactionType (mode)** | Debit | Debit | Debit |
| **Channel (mode)** | Branch | Branch | Branch |
| **CustomerOccupation (mode)** | Student | Student | Engineer |
| **CustomerAge_bin (mode)** | Low | Low | Medium |

---

## 🛠️ Tech Stack

- **Python** 3.x
- **scikit-learn** 1.7.0 *(recommended)*
- **pandas**, **numpy**
- **matplotlib**, **seaborn**
- **yellowbrick** (`KElbowVisualizer`)
- **joblib**
- **Google Colab**

---

## ⚠️ Important Rules

- ❌ Do **not** add extra code cells or imports beyond what is instructed
- ❌ Do **not** use AutoML tools (PyCaret, Auto-sklearn, TPOT, etc.)
- ✅ Run **all cells** before submitting — output must be visible without re-running
- ✅ Use the `Target` column name exactly as specified
- ✅ Package everything into **1 ZIP file** before submitting
- ✅ Use **scikit-learn 1.7.0** to avoid version conflicts during review

---

## 📦 Submission Checklist

- [x] `[Clustering]_Submission_Akhir_BMLP_Agus_Syahril_Mubarok.ipynb` — fully run
- [x] `[Klasifikasi]_Submission_Akhir_BMLP_Agus_Syahril_Mubarok.ipynb` — fully run
- [x] `model_clustering.h5`
- [x] `decision_tree_model.h5`
- [x] `data_clustering.csv`
- [x] `PCA_model_clustering.h5` *(Advanced)*
- [x] `explore_RandomForest_classification.h5` *(Skilled)*
- [x] `tuning_classification.h5` *(Advanced)*
- [x] `data_clustering_inverse.csv` *(Advanced)*

---

## 👤 Author

**Agus Syahril Mubarok**
Dicoding — Building Machine Learning Projects (BMLP)