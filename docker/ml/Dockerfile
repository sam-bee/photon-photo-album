FROM nvidia/cuda:12.6.3-cudnn-runtime-ubuntu24.04

# Install Python and pip
RUN apt-get update && apt-get install -y \
    python3 \
    python3-pip \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

# Install Python dependencies
COPY ml/requirements.txt .
RUN pip3 install --no-cache-dir -r requirements.txt

# Copy application code
COPY ml/src/ .

CMD ["python3", "mq_handler.py"] 