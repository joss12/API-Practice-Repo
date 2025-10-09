
# 1. Base image
FROM node:18

# 2. Set working directory
WORKDIR /app

# 3. Copy project files
COPY package*.json ./
COPY prisma ./prisma/
COPY src ./src/

# 4. Install dependencies
RUN npm install

# 5. Generate Prisma client
RUN npx prisma generate

# 6. Run server (migrate + start dev)
CMD ["sh", "-c", "npx prisma migrate dev --name init && npm run dev"]

