export interface NormalizedURL {
  raw: string;
  protocol: string;
  host: string;
  path: string;
}

export function NormalizeURL(u: string): NormalizedURL {
  const res = new URL(u);
  return {
    raw: res.href,
    protocol: res.protocol,
    host: res.host,
    path: res.pathname,
  };
}
