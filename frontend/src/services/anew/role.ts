import { request } from 'umi';
import type { CheckboxValueType } from 'antd/lib/checkbox/Group'

export async function queryRoles(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/role/list', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function createRole(body: API.RoleParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/role/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function updateRole(body: API.RoleParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/role/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteRole(body: API.Ids, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/role/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function getRolePermsByID(id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/role/perms/${id}`, {
    method: 'GET',
    ...(options || {}),
  });
}

export type PermsProps = {
  menus_id: React.Key[];
  apis_id: CheckboxValueType[];
}

export async function updatePermsRole(body: PermsProps, id?: number, options?: { [key: string]: any }) {
  return request(`/api/v1/role/perms/update/${id}`, {
    method: 'PATCH',
    data: body,
    ...(options || {}),
  });
}