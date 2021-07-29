import { request } from 'umi';

export async function GetMenuTree(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/menu/tree', {
        ...(options || {}),
    });
}

export async function queryMenus(options?: { [key: string]: any }) {
    return request<API.Result>('/api/v1/menu/list', {
      method: 'GET',
      ...(options || {}),
    });
  }
  
export async function createMenu(body: API.MenuParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/menu/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function updateMenu(body: API.MenuParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/menu/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteMenu(body: API.Ids, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/menu/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}