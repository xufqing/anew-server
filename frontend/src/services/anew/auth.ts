import { request } from 'umi';


// login
export async function AuthLogin(body: API.LoginParams, options?: { [key: string]: any }) {
    return request<API.Result>('/api/auth/login', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        data: body,
        ...(options || {}),
    });
}

// logout
export async function AuthLogout(options?: { [key: string]: any }) {
    return request<API.Result>('/api/auth/logout', {
        method: 'POST',
        ...(options || {}),
    });
}