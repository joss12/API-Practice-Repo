import jwt from "jsonwebtoken";

const generateToken = (userId, role) => {
    return jwt.sign(
        { userId, role },
        process.env.JWT_SECRET,
        { expiresIn: "7d" }, // Token expires in 7 days
    );
};

export default generateToken;
