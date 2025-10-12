const request = require("supertest");
const express = require("express");
const bodyParser = require("body-parser");
const submitRoute = require("../src/routes/submit");
const path = require("path");

// fake express app
const app = express();
app.use(bodyParser.json());
app.use("/submit", submitRoute);

describe("POST /submit", () => {
    it("should submit form successfully", async () => {
        const response = await request(app).post("/submit").send({
            name: "Test User",
            email: "test@example.com",
            message: "Testing path fix!",
        });

        // DEBUG: Log the full response
        console.log("Status Code:", response.statusCode);
        console.log("Response Body:", JSON.stringify(response.body, null, 2));
        console.log("Response Error:", response.error);

        // expect 200 OK
        expect(response.statusCode).toBe(200);

        // expect success field true
        expect(response.body.success).toBe(true);

        // expect data to have an ID
        expect(response.body.data).toHaveProperty("id");
    });

    it("should reject invalid form", async () => {
        const response = await request(app).post("/submit").send({
            name: "",
            email: "nope",
            message: "",
        });

        expect(response.statusCode).toBe(400);
        expect(response.body.errors).toBeDefined();
    });
});
