# Deploying NexEntry to AWS EC2 (Ubuntu/WSL)

This guide explains how to deploy the **Go Backend** and **React Frontend** to an Ubuntu environment (such as WSL or AWS EC2).

## Prerequisites

1.  **Ubuntu Server** (AWS EC2 `t2.micro` or `t3.micro` with Ubuntu 24.04/22.04 LTS).
2.  **Go** installed (v1.23+).
3.  **Note**: Ensure Security Groups allow ports `80`, `443`, `8080`, `3306`, `27017` (restrict DB ports to localhost if possible).

## 1. Install Dependencies

Run these commands on your Ubuntu server:

```bash
# Update system
sudo apt update && sudo apt upgrade -y

# Install Go
sudo snap install go --classic

# Install MySQL
sudo apt install mysql-server -y
sudo systemctl start mysql
sudo systemctl enable mysql

# Install Redis
sudo apt install redis-server -y
sudo systemctl start redis-server

# Install MongoDB (Follow official MongoDB docs for verified Ubuntu versions)
# Example for Ubuntu 22.04:
sudo apt install -y gnupg curl
curl -fsSL https://www.mongodb.org/static/pgp/server-7.0.asc | sudo gpg -o /usr/share/keyrings/mongodb-server-7.0.gpg --dearmor
echo "deb [ arch=amd64,arm64 signed-by=/usr/share/keyrings/mongodb-server-7.0.gpg ] https://repo.mongodb.org/apt/ubuntu jammy/mongodb-org/7.0 multiverse" | sudo tee /etc/apt/sources.list.d/mongodb-org-7.0.list
sudo apt update
sudo apt install -y mongodb-org
sudo systemctl start mongod

# Install Nginx (Web Server)
sudo apt install nginx -y
```

## 2. Setup Database

Access MySQL and create the database:
```bash
sudo mysql -u root
# Inside MySQL shell:
CREATE DATABASE guvi_db;
CREATE TABLE guvi_db.users (
    id INT AUTO_INCREMENT PRIMARY KEY,
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL UNIQUE,
    password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);
EXIT;
```

## 3. Deploy Backend (Go)

1.  Copy the project files to the server (e.g., using `git clone` or `scp`).
2.  Navigate to `server/`:
    ```bash
    cd guvi-project/server
    ```
3.  Create a `.env` file based on `.env.example`:
    ```bash
    nano .env
    ```
    *Paste your real credentials.*
4.  Build and Run:
    ```bash
    go build -o server
    # Run in background (simple method)
    nohup ./server > server.log 2>&1 &
    ```

## 4. Deploy Frontend (React)

1.  Navigate to `client/`:
    ```bash
    cd guvi-project/client
    ```
2.  Install dependencies and build:
    ```bash
    sudo apt install npm
    npm install
    # Set the API URL to your EC2 Public IP or Domain
    export VITE_API_BASE_URL=http://<YOUR_EC2_IP>:8080/api
    npm run build
    ```
3.  Copy build to Nginx:
    ```bash
    sudo cp -r dist/* /var/www/html/
    ```
4.  Configure Nginx to handle React Router (SPA):
    Edit `/etc/nginx/sites-available/default`:
    ```nginx
    server {
        listen 80;
        root /var/www/html;
        index index.html;

        colocation / {
            try_files $uri $uri/ /index.html;
        }
    }
    ```
5.  Restart Nginx:
    ```bash
    sudo systemctl restart nginx
    ```

## 5. Verify

Visit `http://<YOUR_EC2_IP>` in your browser.
*   The **Landing Page** should appear.
*   **Sign Up** should connect to your Go backend (port 8080).
*   **Dark/Bright Mode** should work perfectly.
