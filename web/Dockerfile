# Build stage
FROM node:18-alpine as build

WORKDIR /app

# Copy package files
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy source code
COPY . .

# Build the app
RUN npm run build

# Production stage - просто копируем build файлы
FROM nginx:alpine

# Copy built app to nginx
COPY --from=build /app/build /usr/share/nginx/html

# Create a simple nginx config that just serves static files
RUN echo 'server { listen 80; root /usr/share/nginx/html; index index.html; location / { try_files $uri $uri/ /index.html; } }' > /etc/nginx/conf.d/default.conf

EXPOSE 80

CMD ["nginx", "-g", "daemon off;"]