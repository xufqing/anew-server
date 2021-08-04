import { request } from "umi";

export async function querySessions(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/host/session/list', {
        method: 'GET',
        ...(options || {}),
    });
}


export async function deleteSession(body: any, options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/host/session/delete', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        data: body,
        ...(options || {}),
    });
}