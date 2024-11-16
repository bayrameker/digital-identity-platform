import pandas as pd
import mlflow
from mlflow import tensorflow
import tensorflow as tf
from utils.data_preprocessing import preprocess_data

# Veri yükleme
df = pd.read_csv("data/user_data.csv")

# Veri işleme
processed_df = preprocess_data(df)

# Özellikler ve hedef değişkenlerin ayrılması
X = processed_df.drop('target', axis=1)
y = processed_df['target']

# Veriyi eğitim ve test setlerine ayırma
from sklearn.model_selection import train_test_split
X_train, X_val, y_train, y_val = train_test_split(X, y, test_size=0.2)

# Model oluşturma
model = tf.keras.Sequential([
    tf.keras.layers.Dense(128, activation='relu', input_shape=(X_train.shape[1],)),
    tf.keras.layers.Dropout(0.2),
    tf.keras.layers.Dense(64, activation='relu'),
    tf.keras.layers.Dense(1, activation='sigmoid')
])

# Model eğitimi
model.compile(optimizer='adam', loss='binary_crossentropy', metrics=['accuracy'])

# MLflow ile model izleme
with mlflow.start_run():
    model.fit(X_train, y_train, epochs=10, validation_data=(X_val, y_val))
    # Modeli MLflow'a kaydet
    mlflow.tensorflow.log_model(model, "RecommendationModel")
    mlflow.log_param("epochs", 10)
    mlflow.log_metric("val_accuracy", model.evaluate(X_val, y_val)[1])
