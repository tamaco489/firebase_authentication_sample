import { API_HOST } from '@/constants/api';

export const apiClient = {
  get: async (endpoint: string, token: string) => {
    const response = await fetch(`${API_HOST}${endpoint}`, {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${token}`,
      },
    });

    if (!response.ok) {
      throw new Error(`API request failed: ${response.status} ${response.statusText}`);
    };

    return response.json();
  },

  post: async (endpoint: string, token: string, body: object) => {
    const response = await fetch(`${API_HOST}${endpoint}`, {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        Accept: 'application/json',
        Authorization: `Bearer ${token}`,
      },
      body: JSON.stringify(body),
    });

    if (!response.ok) {
      throw new Error(`API request failed: ${response.status} ${response.statusText}`);
    };

    return response.json();
  },
};
