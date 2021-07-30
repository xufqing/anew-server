import { request } from 'umi';

export async function queryDicts(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/dict/list', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function createDict(body: API.DictParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/dict/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function updateDict(body: API.DictParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/dict/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteDict(body: API.Ids, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/dict/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}