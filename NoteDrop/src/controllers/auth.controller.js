import bcrypt from "bcrypt";
import jwt from "jsonwebtoken";
import { PrismaClient } from "@prisma/client";
import generateToken from "../utils/generateToken.js";

const prisma = new PrismaClient();

//Register functionality
export const registerUser = async (req, res) => {
    const { email, password } = req.body;

    //input validation
    if (!email || !password) {
        return res.status(400).json({ error: "Email and password are required" });
    }

    try {
        //check if user already exists
        const isUserExist = await prisma.user.findUnique({
            where: { email },
        });

        if (isUserExist) {
            return res.status(400).json({ error: "User already exists" });
        }

        //password hashed with 10 solt
        const hashedPassword = await bcrypt.hash(password, 10);

        //Create a new user
        const user = await prisma.user.create({
            data: {
                email,
                password: hashedPassword,
                role: "USER", //default role
            },
        });
        //generate new token
        const token = generateToken(user.id, user.role);

        // Return result
        res.status(201).json({
            message: "User registerd successfully",
            user: {
                id: user.id,
                email: user.email,
                role: user.role,
            },
            token,
        });
    } catch (err) {
        console.error("Register Error", err);
        res.status(500).json({ error: "->Server Error" });
    }
};

//Login functionality
export const loginUser = async (req, res) => {
    const { email, password } = req.body;

    //input validation
    if (!email || !password) {
        return res.status(400).json({ error: "Email and password are required" });
    }

    try {
        //Chekck if user exists
        const user = await prisma.user.findUnique({
            where: { email },
        });

        if (!email) {
            return res.status(401).json({ error: "Invalid email or password" }); //generic message
        }

        //compaire password
        const isMatch = await bcrypt.compare(password, user.password);

        if (!isMatch) {
            return res.status(401).json({ error: "Invalid email or password" });
        }

        //Generate token
        const token = generateToken(user.id, user.role);

        //send response
        res.status(200).json({
            message: "Login successfully",
            user: {
                id: user.id,
                email: user.email,
                role: user.role,
            },
            token,
        });
    } catch (err) {
        console.error("Login Error:", err);
        res.status(500).json({ error: "Server error" });
    }
};

//Promote a user
export const promoteUserToAdmin = async (req, res) => {
    const userIdToPromote = req.params.id;

    try {
        const updateUser = await prisma.user.update({
            where: { id: userIdToPromote },
            data: { krole: "Admin" },
        });
        res.json({
            message: `User promoted to Admin`,
            user: {
                id: updateUser.id,
                email: updateUser.email,
                role: updateUser.role,
            },
        });
    } catch (err) {
        console.error("Promote Error:", err);
        res.status(500).json({ error: "Failed  to promote user" });
    }
};
