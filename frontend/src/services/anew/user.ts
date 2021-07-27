import { request } from 'umi';


export async function GetUserInfo(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/user/info', {
      method: 'GET',
      ...(options || {}),
    });
  }