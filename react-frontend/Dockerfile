# Stage 1: Build the React application
FROM node:18-alpine AS builder

# Set the working directory
WORKDIR /app

# Copy package.json and package-lock.json (or yarn.lock)
COPY package*.json ./

# Install dependencies
RUN npm install

# Copy the rest of the application code
COPY . .

# Build the React app for production
# This creates the 'build' folder with static files
RUN npm run build

# Stage 2: Serve the static files with Nginx
FROM nginx:alpine

# Copy the built React app from the builder stage into Nginx's public directory
COPY --from=builder /app/build /usr/share/nginx/html

# Expose port 80 where Nginx serves the content
EXPOSE 80

# Command to run Nginx
CMD ["nginx", "-g", "daemon off;"]