# Base image
FROM python:3.9-slim

# Çalışma dizini
WORKDIR /app

# Bağımlılıkları kopyala ve yükle
COPY requirements.txt .

RUN pip install --no-cache-dir -r requirements.txt

# Uygulama dosyalarını kopyala
COPY . .

# Portları aç
EXPOSE 8000

# Komutu çalıştır
CMD ["uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
