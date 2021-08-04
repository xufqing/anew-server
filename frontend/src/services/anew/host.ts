import { request } from "umi";

export async function queryHosts(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/host/list', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function queryHostByGroupId(id?: number,options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/host/list?group_id=${id}`, {
    method: 'GET',
    ...(options || {}),
  });
}


export async function queryHostByID(id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/host/info/${id}`, {
    method: 'GET',
    ...(options || {}),
  });
}


export async function createHost(body: API.HostParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/host/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function updateHost(body: API.HostParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/host/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteHost(body: API.Ids, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/host/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function queryRecords(id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/host/record/list/${id}`, {
    method: 'GET',
    ...(options || {}),
  });
}


export async function deleteRecord(body: API.Ids, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/host/record/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}


export async function queryHostGroups(options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/host/group/list', {
    method: 'GET',
    ...(options || {}),
  });
}

export async function createHostGroup(body: API.HostGroupParams, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/host/group/create', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function updateHostGroup(body: API.HostGroupParams, id?: number, options?: { [key: string]: any }) {
  return request<API.Result>(`/api/v1/host/group/update/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}

export async function deleteHostGroup(body: API.Ids, options?: { [key: string]: any }) {
  return request<API.Result>('/api/v1/host/group/delete', {
    method: 'DELETE',
    headers: {
      'Content-Type': 'application/json',
    },
    data: body,
    ...(options || {}),
  });
}