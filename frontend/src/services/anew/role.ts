import { request } from 'umi';

export async function queryRoles(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/role/list', {
      method: 'GET',
      ...(options || {}),
    });
  }