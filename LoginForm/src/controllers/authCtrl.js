const validateForm = require("../utils/validateForm");
const fs = require("fs");
const path = require("path");
const bcrypt = require("bcryptjs");
const jwt = require("jsonwebtoken");

//const DATA_PATH = path.join(__dirname, "..", "data.json");

const USERS_PATH = path.join(__dirname, "..", "..", "users.json");
const JWT_SECRET = process.env.JWT_SECRET || "default_secret";

function loadUsers() {
    if (!fs.existsSync(USERS_PATH)) return [];
    const raw = fs.readFileSync(USERS_PATH, "utf-8");
    return raw.trim() ? JSON.parse(raw) : [];
}

//function loadUsers() {
//    if (!fs.existsSync(USERS_PATH)) return [];
//    const content = fs.readFileSync(USERS_PATH, "utf-8");
//    return content.trim() ? JSON.parse(content) : [];
//}

function saveUsers(users) {
    fs.writeFileSync(USERS_PATH, JSON.stringify(users, null, 2));
}

const authController = {
    async Register(req, res) {
        const { name, email, password, role } = req.body;
        //Input validation
        if (!email || !password)
            return res.status(400).json({ error: "Missing fields" });

        //load existing users
        const users = loadUsers();
        // if (users.find((u) => u.email === email)) {
        //   return res.status(409).json({ error: "User already exists" });
        // }

        const exists = users.find(
            (u) => u.email.toLowerCase() === email.toLowerCase,
        );
        if (exists) {
            return res.status(409).json({ error: "Email already registered" });
        }

        const hashed = await bcrypt.hash(password, 10);
        const newUser = {
            id: Date.now(),
            name,
            email,
            password: hashed,
            role: role || "user",
        };

        users.push(newUser);
        saveUsers(users);

        // Send success
        res.status(201).json({ message: "User registered!" });
    },

    async Login(req, res) {
        const { email, password } = req.body;
        const users = loadUsers();
        const user = users.find((u) => u.email === email);
        if (!user || !(await bcrypt.compare(password, user.password))) {
            return res.status(401).json({ error: "Invalid credentials" });
        }

        const token = jwt.sign(
            { id: user.id, email: user.email, role: user.role },
            JWT_SECRET,
            { expiresIn: "2h" },
        );

        res.json({ message: "Login successful", token });
    },
};

module.exports = authController;
