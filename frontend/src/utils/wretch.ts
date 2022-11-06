import wretch from 'wretch';

export const api = wretch(process.env.API_URL).resolve(r => r.json());
