#Use official Node image
FROM node:18

#Create app directory
WORKDIR /app

#Copy package.json and install
COPY package*.json ./
RUN npm install

#Copy source code
COPY . .

#Expose port (matches Express)
EXPOSE 3600

# Start app
CMD ["npm", "start"]


