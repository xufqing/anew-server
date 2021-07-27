import { request } from 'umi';


export async function getUserInfo(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/user/info', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function updateUserInfo(body: API.UpdateUserInfoParams, id?: string, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/user/info/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}