import { writable } from 'svelte/store';

// Default to localhost:8080 if PUBLIC_API_URL is not set
export const apiUrl = writable(import.meta.env.PUBLIC_API_URL || 'http://localhost:8080');

export const backendAvailable = writable(true); 