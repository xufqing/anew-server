import { request } from 'umi';

export async function querySSHFile(cententId?: string, path?: string, options?: { [key: string]: any }) {
    return request<API.Result>(`/api/v1/host/ssh/ls?key=${cententId}&path=${path}`, {
        method: 'GET',
        ...(options || {}),
    });
}

export async function deleteSSHFile(cententId?: string, path?: string, options?: { [key: string]: any }) {
    return request<API.Result>(`/api/v1/host/ssh/rm?key=${cententId}&path=${path}`, {
        method: 'DELETE',
        ...(options || {}),
    });
}

export async function uploadSSHFile(cententId?: string, path?: string, options?: { [key: string]: any }) {
    return request<API.Result>(`/api/v1/host/ssh/rm?key=${cententId}&path=${path}`, {
        method: 'POST',
        ...(options || {}),
    });
}