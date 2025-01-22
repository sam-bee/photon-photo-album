FROM ruby:3.3-alpine

# Install required packages
RUN apk add --no-cache \
    build-base \
    postgresql-dev \
    sqlite-dev \
    tzdata \
    sudo

# Create user with specified UID/GID
ARG DOCKERUID
ARG DOCKERGID
RUN addgroup -g ${DOCKERGID} appuser && \
    adduser -D -u ${DOCKERUID} -G appuser appuser && \
    echo "appuser ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/appuser

WORKDIR /app

# Install Rails
RUN gem install rails

# Create new Rails API
USER appuser
RUN rails new . --api --database=sqlite3 --skip-test --skip-bundle

# Copy Gemfile and install dependencies
COPY api/Gemfile api/Gemfile.lock ./
RUN bundle install

# Copy application code
COPY api/ .

# Set ownership
RUN sudo chown -R appuser:appuser /app

EXPOSE 3000

CMD ["rails", "server", "-b", "0.0.0.0"] 
