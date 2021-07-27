import { request } from 'umi';

export async function GetMenuTree(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/menu/tree', {
        ...(options || {}),
    });
}