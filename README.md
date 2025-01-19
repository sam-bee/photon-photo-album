# Photon - the Machine Learning Photo Album

Classify your photos locally with machine learning, and curate albums with a local web interface.

Keep your photos private and secure, and your data local.

Uses EfficientNet/TensorFlow for the machine learning model.

## Setup

```bash
make setup
```

## Usage

Go to http://localhost:3000 and start classifying your photos.

## System Design Overview

### Frontend
- React web application
- Provides a way to curate your albums, and classify photos using machine learning

### Backend
- Go service
- Provides a REST API for the frontend to interact with
- Manages the database
- Manages the machine learning models

### ML
- Python service
- Uses a machine learning model to classify photos
- Runs on a GPU


### MQ
- RabbitMQ
- Used to queue up photos for the machine learning model to classify
