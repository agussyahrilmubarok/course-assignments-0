# 🤖 Machine Learning Project — Bank Transaction Fraud Detection
### Dicoding | Building Machine Learning Projects (BMLP) | Final Submission

---

## 📌 Project Overview

This project integrates **Unsupervised Learning (Clustering)** and **Supervised Learning (Classification)** to analyze bank transaction data originally sourced from the Kaggle "Bank Transaction Dataset for Fraud Detection" (modified version).

The workflow:
1. Apply **K-Means Clustering** to generate labels from unlabeled data
2. Use those labels as the target variable to train a **Classification model**

---

## 🗂️ Project Structure

```
BMLP_Your-Name.zip
├── [Clustering]_Submission_Akhir_BMLP_Your_Name.ipynb
├── [Klasifikasi]_Submission_Akhir_BMLP_Your_Name.ipynb
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
| 3 | **EDA** | `head()`, `info()`, `describe()`, correlation matrix, histograms |
| 4 | **Data Cleaning** | Check nulls (`isnull().sum()`), duplicates (`duplicated().sum()`), drop them |
| 5 | **Preprocessing** | Drop ID/Date columns, `LabelEncoder`, outlier handling, `StandardScaler`, binning |
| 6 | **Elbow Method** | Use `KElbowVisualizer()` to find optimal K |
| 7 | **K-Means** | Train model → save as `model_clustering.h5` |
| 8 | **Evaluation** | Silhouette Score + cluster visualization |
| 9 | **PCA** *(Optional)* | Build PCA model → save as `PCA_model_clustering.h5` |
| 10 | **Interpretation** | Descriptive statistics per cluster (mean, min, max) |
| 11 | **Inverse Transform** | Reverse encoding & scaling back to original values |
| 12 | **Export** | Save `data_clustering.csv` and `data_clustering_inverse.csv` |

---

## 📒 Notebook 2: Classification

### Workflow
```
Load Clustering Result → Train-Test Split → 
Decision Tree → Other Algorithms → Evaluate → Tuning → Export
```

### Steps

| # | Step | Description |
|---|------|-------------|
| 1 | **Import Library** | Standard sklearn imports |
| 2 | **Load Dataset** | Load `data_clustering.csv` (contains `Target` column) |
| 3 | **EDA** | `head()` to verify data |
| 4 | **Train-Test Split** | `train_test_split()` |
| 5 | **Decision Tree** | Build model → save as `decision_tree_model.h5` |
| 6 | **Other Algorithm** *(Optional)* | e.g., `RandomForestClassifier` → save as `explore_<Name>_classification.h5` |
| 7 | **Evaluation** | Accuracy, Precision, Recall, F1-Score for all models |
| 8 | **Hyperparameter Tuning** *(Optional)* | `GridSearchCV` or `RandomizedSearchCV` → save as `tuning_classification.h5` |

---

## 📊 Grading Criteria

| Criteria | Basic (2pts) | Skilled (3pts) | Advanced (4pts) |
|----------|-------------|----------------|-----------------|
| **1. EDA** | head, info, describe | + correlation, histograms | + no overlapping labels |
| **2. Preprocessing** | null/dup check, drop, encode | + outlier drop, StandardScaler | + binning |
| **3. Clustering** | Elbow + KMeans + save model | + Silhouette Score + viz | + PCA model |
| **4. Interpretation** | Descriptive stats + export | + inverse transform + analysis | + save inverse CSV |
| **5. Classification** | Decision Tree + save model | + extra algorithm + evaluation | + hyperparameter tuning |

### Score Formula
```
Final Score = Total Points / Number of Criteria
```

| Final Score | Stars | Grade | Level |
|-------------|-------|-------|-------|
| < 1 | ❌ Rejected | E | — |
| 1 – <2 | ⭐⭐ | D | Below Basic |
| 2 – <3 | ⭐⭐⭐ | C | Basic |
| 3 – <4 | ⭐⭐⭐⭐ | B | Skilled |
| 4 | ⭐⭐⭐⭐⭐ | A | Advanced |

---

## 🛠️ Tech Stack

- **Python** 3.x
- **scikit-learn** 1.7.0 *(recommended)*
- **pandas**, **numpy**
- **matplotlib**, **seaborn**
- **yellowbrick** (KElbowVisualizer)
- **joblib**

---

## ⚠️ Important Rules

- ❌ Do **not** add extra code cells or imports beyond what is instructed
- ❌ Do **not** use AutoML tools (PyCaret, Auto-sklearn, TPOT, etc.)
- ✅ Run all cells before submitting — output must be visible without re-running
- ✅ Use the `Target` column name exactly as specified
- ✅ Package everything into **1 ZIP file** before submitting

---

## 📦 Submission Checklist

- [ ] `[Clustering]_Submission_Akhir_BMLP_Your_Name.ipynb` — fully run
- [ ] `[Klasifikasi]_Submission_Akhir_BMLP_Your_Name.ipynb` — fully run
- [ ] `model_clustering.h5`
- [ ] `decision_tree_model.h5`
- [ ] `data_clustering.csv`
- [ ] *(Optional)* `PCA_model_clustering.h5`
- [ ] *(Optional)* `explore_<Algorithm>_classification.h5`
- [ ] *(Optional)* `tuning_classification.h5`
- [ ] *(Optional)* `data_clustering_inverse.csv`
- [ ] All files zipped as `BMLP_Your-Name.zip`

---

## 👤 Author

**Agus Syahril Mubarok**  
Dicoding — Building Machine Learning Projects (BMLP)