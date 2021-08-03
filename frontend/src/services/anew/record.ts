import { request } from "umi";

export async function queryRecords(id?: number, options?: { [key: string]: any }) {
    return request<API.Result>(`/api/v1/host/record/list/${id}`, {
        method: 'GET',
        ...(options || {}),
    });
}

export async function deleteRecord(body: API.Ids, options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/host/record/delete', {
        method: 'DELETE',
        headers: {
            'Content-Type': 'application/json',
        },
        data: body,
        ...(options || {}),
    });
}