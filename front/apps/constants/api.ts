export const API_HOST = process.env.NEXT_PUBLIC_API_BASE_URL ?? '';
export const API_KEY = `${process.env.API_KEY ?? 'ABCDEFG123456789'}`;
export const API_REQUEST_OPTIONS = {
  headers: {
    'Content-Type': 'application/json',
    Accept: 'application/json',
    Authorization: `Bearer ${API_KEY}`,
  },
  withCredentials: true,
};
