import { error } from '@sveltejs/kit';

const backendPath = import.meta.env.VITE_BACKEND_URL;

/**
 * Send any type of requests, expecting to get json response.
 */
const send = async (options: RequestInit, path: string, token?: string) => {
    const h = new Headers(options.headers);
    if (token) {
        options.headers = [...h, 
			["Authorization", `Bearer ${token}`],
			["Accept", "application/json"],
			["Accept-Encoding", "gzip"]
		];
	}
    return fetch(`${backendPath}/${path}`, options)
		.then(async res => {
			if (res?.ok) {
				const text = await res.text();
				return text ? JSON.parse(text) : {};
			}
		
			throw error(res.status);
		})
		.catch(e => {
			console.log(e);
		});
};

export const get = (path: string, headers?: HeadersInit, token?: string) => {
    return send({ method: "GET", headers }, path, token);
};

export const post = (path: string, headers?: HeadersInit, token?: string) => {
	return send({ method: "POST", headers }, path, token);
};

export const put = (path: string, headers?: HeadersInit, token?: string) => {
	return send({ method: "PUT", headers }, path, token);
};
