import { Request, Response, NextFunction } from "express";
import jwt from "jsonwebtoken";

export function verifyToken(req: Request, res: Response, next: NextFunction) {
  const header = req.headers.authorization;

  if (!header)
    return res.status(401).json({ message: "Missing Authorizationt header" });

  const token = header.split(" ")[1];
  if (!token) return res.status(401).json({ message: "Malformed token" });

  try {
    const decoded = jwt.verify(token, process.env.JWT_SECRET as string);
    (req as any).user = decoded;
    next();
  } catch {
    return res.status(401).json({ message: "Invalid token" });
  }
}
