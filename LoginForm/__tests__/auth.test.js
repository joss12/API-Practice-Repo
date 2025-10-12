const request = require("supertest");
const express = require("express");
const bodyParser = require("body-parser");
const authRoutes = require("../src/routes/auth");

//fake express app
const app = express();
app.use(bodyParser.json());
app.use("/auth", authRoutes);

describe("POST /auth/register + /auth/ling", () => {
    const email = `test${Date.now()}@example.com`;
    const password = "123456";

    it("should register a new user", async () => {
        const response = await request(app).post("/auth/register").send({
            name: "Test",
            email,
            password,
            role: "user",
        });

        expect(response.statusCode).toBe(201);
        expect(response.body.message).toMatch(/User registered/i);
    });

    it("should login the user", async () => {
        const response = await request(app).post("/auth/login").send({
            email,
            password,
        });

        expect(response.statusCode).toBe(200);
        expect(response.body).toHaveProperty("token");
    });
});
