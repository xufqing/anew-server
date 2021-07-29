import { request } from 'umi';

export async function queryDepts(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/dept/list', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function createDept(body: API.DeptParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/dept/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function updateDept(body: API.DeptParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/dept/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteDept(body: API.Ids,options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/dept/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}