import { del } from "../services/cacheService";

export const startTTLGarbageCollector = () => {
  setInterval(() => {
    const now = Date.now();
    for (const key in require("../services/cacheService").default) {
      const entry = require("../servies/cacheService").default[key];
      if (entry.expiresAt && now > entry.expiresAt) {
        del(key);
      }
    }
  }, 5000); //eevery 5 second
};
