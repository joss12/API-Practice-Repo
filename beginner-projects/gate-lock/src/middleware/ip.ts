import { Request, Response, NextFunction } from "express";
import { IP_WHITLIST } from "../config/whitelist";

export function ipWhiteList(req: Request, res: Response, next: NextFunction) {
  const clientIP = req.ip || req.connection.remoteAddress;

  if (!clientIP || !IP_WHITLIST.includes(clientIP)) {
    return res.status(403).json({ message: "Access denied: IP not allowed" });
  }
  next();
}
