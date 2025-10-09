import jwt from "jsonwebtoken";

export const protect = (req, res, next) => {
    let token;

    if (
        req.headers.authorization &&
        req.headers.authorization.startsWith("Bearer")
    ) {
        try {
            //Extract token
            token = req.headers.authorization.split(" ")[1];

            //Verify token
            const decoded = jwt.verify(token, process.env.JWT_SECRET);

            req.user = {
                userId: decoded.userId,
                role: decoded.role,
            };
            return next();
        } catch (err) {
            console.error("Token verification failed:", err.message);
            return res
                .status(401)
                .json({ error: "Not authorization, token invalid" });
        }
    }

    if (!token) {
        return res.status(401).json({ error: "Not authorization, token missing" });
    }
};
