FROM node:20-alpine

# Install required packages including sudo
RUN apk add --no-cache sudo

# Create user with specified UID/GID
ARG DOCKERUID
ARG DOCKERGID
RUN addgroup -g ${DOCKERGID} appuser && \
    adduser -D -u ${DOCKERUID} -G appuser appuser && \
    echo "appuser ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/appuser

WORKDIR /app

# Set ownership of the working directory
RUN chown -R appuser:appuser /app

# Create new React TypeScript project
USER appuser
RUN npx create-react-app . --template typescript

# Install dependencies
RUN npm install @emotion/react @emotion/styled

# Copy source code
COPY frontend/ .

# Set ownership of all copied files
RUN chown -R appuser:appuser /app

EXPOSE 3000

CMD ["npm", "start"] 