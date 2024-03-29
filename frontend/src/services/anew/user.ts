import { request } from 'umi';


export async function getUserInfo(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/user/info', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function updateUserInfo(body: API.UserInfo, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/user/info/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function changePassword(body: API.ChangePwdParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/user/changePwd', {
    method: 'PUT',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function queryUsers(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/user/list', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function deleteUser(body: API.Ids,options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/user/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function createUser(body: API.UserParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/user/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function updateUser(body: API.UserParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/user/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}