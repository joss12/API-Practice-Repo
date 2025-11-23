type CacheEntry = {
  value: any;
  expiresAt: number | null; //Unix timestamp
};

const cache: Record<string, CacheEntry> = {};

export const set = (key: string, value: any, ttl?: number) => {
  const now = Date.now();
  const expiresAt = ttl ? now + ttl * 1000 : null;
  cache[key] = { value, expiresAt };
};

export const get = (key: string) => {
  const entry = cache[key];
  if (!entry) return null;
  if (entry.expiresAt && Date.now() > entry.expiresAt) {
    del(key);
    return null;
  }
  return entry.value;
};

export const del = (key: string) => {
  delete cache[key];
};

export const stats = () => {
  const keys = Object.keys(cache);
  const memorySize = Buffer.byteLength(JSON.stringify(cache), "utf8");
  return {
    keys: keys.length,
    memorySizeKB: (memorySize / 1024).toFixed(2),
  };
};

export const getAll = () => {
  return Object.entries(cache).map(([key, entry]) => ({
    key,
    value: entry.value,
    expiresAt: entry.expiresAt,
    expired: entry.expiresAt !== null && Date.now() > entry.expiresAt,
  }));
};
