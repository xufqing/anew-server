import { request } from "umi";

export async function queryOperLogs(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/operlog/list', {
        method: 'GET',
        ...(options || {}),
    });
}

export async function deleteOperLog(body: API.Ids, options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/operlog/delete', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        data: body,
        ...(options || {}),
    });
}