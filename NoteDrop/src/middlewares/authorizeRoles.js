export const authorizationRoles = (...allowedroles) => {
    return (req, res, next) => {
        if (!req.user || !allowedroles.includes(req.user.role)) {
            return res.status(403).json({ error: "Forbidden: Insufficient role" });
        }
        next();
    };
};
