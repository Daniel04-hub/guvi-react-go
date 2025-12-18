# Connecting WSL to Cloud DBs & Local Windows MySQL

This guide explains how to configure your project running in **WSL (Ubuntu)** to connect to:
1.  **MongoDB Atlas** (Cloud)
2.  **Redis Cloud** (Cloud)
3.  **MySQL** (Running on your Windows Laptop)

---

## 1. MongoDB Atlas (Cloud) Setup

1.  **Create Account/Cluster**: Go to [MongoDB Atlas](https://www.mongodb.com/cloud/atlas) and create a free cluster.
2.  **Network Access**: Go to **Network Access** > **Add IP Address** > Select **Allow Access from Anywhere** (`0.0.0.0/0`) since your WSL IP changes.
3.  **Database User**: Go to **Database Access** > Create a user (e.g., `guvi_user`) and password.
4.  **Get Connection String**:
    *   Click **Connect** > **Drivers**.
    *   Copy the string. It looks like:
        `mongodb+srv://guvi_user:<password>@cluster0.mongodb.net/?retryWrites=true&w=majority`
5.  **Update Project**:
    *   Open `server/.env` in your project.
    *   Set `MONGO_URI` to this string (replace `<password>` with your actual password).

---

## 2. Redis Cloud Setup

1.  **Create Account**: Go to [Redis Cloud](https://redis.com/try-free/) and create a free subscription.
2.  **Get Details**:
    *   Copy the **Public Endpoint** (e.g., `redis-12345.c1.us-east1-2.gce.cloud.redislabs.com:12345`).
    *   Copy the **Default User Password**.
3.  **Formulate Address**:
    *   Format: `redis://:<password>@<endpoint>`
    *   *Note*: Our Go code currently expects just `host:port` in `REDIS_ADDR` and no password logic is currently hardcoded in `db.go` for the *basic* client options, but for cloud we usually need a URL.
    *   **Action**: We will update your `server/.env` to use the address `host:port` and if needed, we might need a small code tweak if your Redis Cloud enforces a password (which it usually does).
    *   *For now*, set `REDIS_ADDR` to the endpoint: `redis-12345.c1.us-east1-2.gce.cloud.redislabs.com:12345`.

---

## 3. Connect to Windows MySQL from WSL

Since MySQL is on Windows and you are in WSL, `localhost` in WSL refers to the Ubuntu VM, NOT Windows.

### Step A: Find Windows IP
In your WSL terminal, run:
```bash
grep nameserver /etc/resolv.conf
# Output example: nameserver 172.25.160.1
```
Copy that IP address (e.g., `172.25.160.1`). This is your Windows Host IP.

### Step B: Allow MySQL Access
By default, MySQL `root` user only allows access from `localhost`. You need to allow connections from WSL.

1.  Open your **MySQL Command Line Client** (or Workbench) on Windows.
2.  Run these commands:
    ```sql
    -- Create a user that can connect from anywhere (or specifically your WSL IP range)
    CREATE USER 'guvi_remote'@'%' IDENTIFIED BY 'yourpassword';
    GRANT ALL PRIVILEGES ON guvi_db.* TO 'guvi_remote'@'%';
    FLUSH PRIVILEGES;
    ```

### Step C: Configure Windows Firewall
1.  Search for **Windows Defender Firewall with Advanced Security**.
2.  Click **Inbound Rules** > **New Rule**.
3.  **Port** > **TCP** > Specific Local Ports: `3306`.
4.  **Allow the connection**.
5.  Name it "MySQL for WSL".

---

## 4. Final Configuration (.env)

Edit your `server/.env` file:

```bash
# MySQL: Use the Windows IP found in Step 3A and the new user from 3B
MYSQL_DSN=guvi_remote:yourpassword@tcp(172.25.160.1:3306)/guvi_db?parseTime=true

# Redis Cloud Endpoint
REDIS_ADDR=redis-12345.your-cloud:12345

# MongoDB Atlas
MONGO_URI=mongodb+srv://guvi_user:password@cluster0.mongodb.net/?retryWrites=true&w=majority

# Server Port
PORT=8080
```

Now, run your Go server in WSL:
```bash
cd server
go run .
```
It should connect to all three services!
