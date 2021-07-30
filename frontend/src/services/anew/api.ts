import { request } from 'umi';

export async function queryApis(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/api/list', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function createApi(body: API.ApiParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/api/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function updateApi(body: API.ApiParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/api/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteApi(body: API.Ids, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/api/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}