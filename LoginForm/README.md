# 🛡️ LogiForm – Secure Form Submission Platform

A production-ready Node.js project that features secure form submission with rate limiting, NGINX reverse proxy, email notifications via Nodemailer, and testing with Jest + Supertest.

---

## 🚀 Features

- ✅ Form submission with name, email, and message
- ✅ Rate limiting (to prevent spam or abuse)
- ✅ Email notifications via Gmail (Nodemailer)
- ✅ JWT-based Authentication (Login & Register)
- ✅ Simple dashboard with submission history
- ✅ NGINX reverse proxy setup
- ✅ Dockerized (Node.js + NGINX)
- ✅ Jest & Supertest for backend testing

---

## 📦 Tech Stack

- **Node.js** + **Express**
- **Nodemailer**
- **express-rate-limit**
- **jsonwebtoken**
- **bcryptjs**
- **Docker** + **NGINX**
- **Jest** + **Supertest** (for testing)

---

## 📁 Project Structure

```
loginform/
├── public/               # Frontend (HTML/CSS/JS)
├── src/
│   ├── controllers/
│   ├── middleware/
│   ├── routes/
│   └── index.js          # Express entry point
├── .env
├── Dockerfile
├── docker-compose.yml
├── nginx/
│   └── default.conf
├── package.json
└── README.md
```

---

## ⚙️ Commands

### 🛠 Development

```bash
npm install
npm run dev
```

### 🧪 Run Tests

```bash
npm test
```

### 🚀 Run with Docker

```bash
docker-compose up --build
```
## Copy the links and past it in the browser 
http://localhost:8080/login.html -> Login,
http://localhost:8080/register.html -> Register,
http://localhost:8080/dashboard.html -> Dashboard


### 🛑 Stop Docker

```bash
docker-compose down
```

---

## 🔐 Environment Variables (.env)

```
PORT=8080
JWT_SECRET=your_jwt_secret
EMAIL_USER=your_email@gmail.com
EMAIL_PASS=your_email_password_or_app_pass
```
> ⚠️ Use a Gmail App Password if you have 2FA enabled.

---

## 🌐 NGINX Reverse Proxy

Located in `nginx/default.conf`:

```
server {
    listen 80;
    location / {
        proxy_pass http://app:8080;
        proxy_http_version 1.1;
        proxy_set_header Upgrade $http_upgrade;
        proxy_set_header Connection 'upgrade';
        proxy_set_header Host $host;
        proxy_cache_bypass $http_upgrade;
    }
}
```

---

## ✉️ Sending Email

Emails are sent on form submission using the Gmail SMTP server. Configured via Nodemailer using environment variables.

---

## ✅ Todo

- [x] Add email confirmation UI
- [x] Add dashboard (HTML-based)
- [x] Protect /my-submissions route
- [x] Add RBAC (Admin/User roles)
- [x] Add Supertest + Jest tests
- [x] Add HTML <-> Express integration

---

## 🧑‍💻 Author

Built ❤️ by Eddy Mouity
