import { request } from 'umi';

export async function queryDepts(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/dept/list', {
      method: 'GET',
      ...(options || {}),
    });
  }