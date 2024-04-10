# Use the Nginx image from Docker Hub
FROM nginx:alpine

# Copy the static site into place.
COPY index.html /usr/share/nginx/html

# Copy the Nginx configuration file (you'll create this next)
COPY nginx.conf /etc/nginx/nginx.conf
