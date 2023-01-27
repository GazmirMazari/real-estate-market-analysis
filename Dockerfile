# Use an official Redis image as the base image
FROM redis:latest

# Set the working directory
WORKDIR /app

# Copy the configuration file
COPY redis.conf /usr/local/etc/redis.conf

# Expose the Redis port
EXPOSE 6379

# Start Redis
CMD ["redis-server", "/usr/local/etc/redis.conf"]

# Use the official Dgraph image as the base image
FROM dgraph/dgraph:latest

# Set the working directory
WORKDIR /app

# Copy the configuration file
COPY dgraph.conf /dgraph.conf

# Expose the Dgraph ports
EXPOSE 8080
EXPOSE 9080

# Start Dgraph
CMD ["dgraph", "--config", "/dgraph.conf"]
