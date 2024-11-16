# Base image
FROM node:14-alpine

# Çalışma dizini
WORKDIR /app

# Paketleri kopyala
COPY package.json tsconfig.json ./

# Bağımlılıkları yükle
RUN npm install

# Uygulama dosyalarını kopyala
COPY src ./src

# Uygulamayı derle
RUN npm run build

# Portları aç
EXPOSE 4000

# Komutu çalıştır
CMD ["npm", "start"]
