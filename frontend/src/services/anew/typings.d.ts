// @ts-ignore
/* eslint-disable */

declare namespace API {
  type Result = {
    code: number,
    status: boolean;
    data: any;
    message: string;
  };

  type LoginParams = {
    username: string;
    password: string;
  };

  type ChangePwdParams = {
    old_password: string;
    new_password: string;
  };

  type ChangePwdParams = {
    old_password: string;
    new_password: string;
  };

  type UserParams = {
    username?: string,
    password: string,
    mobile?: string,
    avatar?: string,
    name: string,
    email?: string,
    dept_id?: number,
    role_id: number,
    creator?: string,
  }

  type DeptParams = {
    parent_id?: number,
    name: string,
    sort?: number,
    status?: boolean,
  }

  type MenuParams = {
    name: string,
    icon?: string,
    path?: string,
    sort?: number,
    status?: boolean,
    parent_id?: number,
  }

  type RoleParams = {
    name: string,
    keyword?: string,
    desc?: string,
  };

  type ApiParams = {
    name: string,
    method?: string,
    path?: string,
    perms_tag: string,
    desc?: string,
    parent_id?: number,
  };

  type DictParams = {
    dict_key: string,
    dict_value: string,
    sort?: number,
    desc?: string,
    parent_id?: number,
  };

  type HostParams = {
    host_name: string,
    ip_address: string,
    host_type: string,
    port: string,
    auth_type: string,
    user: string,
    password: string,
    privatekey: string,
    key_passphrase: string,
  };

  /////////////////////////////

  type UserInfo = {
    id: number,
    username: string,
    mobile: string,
    avatar: string,
    name: string,
    email: string,
    dept: { id: number, name: string, status: boolean },
    role: { id: number, name: string, desc: string, keyword: string, status: boolean },
    perms: string[],
  }

  type Ids = {
    ids: number[]
  }

  type UserList = {
    id: number,
    username: string,
    name: string,
    mobile: string,
    email: string,
    dept: { id: number, name: string, status: boolean },
    role: { id: number, name: string, desc: string, keyword: string, status: boolean },
    creator: string,
    status: boolean,
  };

  type DeptList = {
    id: number,
    parent_id?: number,
    name: string,
    sort?: number,
    status?: boolean,
    creator?: string,
    children?: DeptList[],
  };

  type MenuList = {
    id: number,
    parent_id?: number,
    name: string,
    icon?: string,
    path?: string,
    sort?: number,
    creator?: string,
    children?: MenuList[],
  };

  type RoleList = {
    id: number,
    name: string,
    keyword?: string,
    desc?: string,
    creator?: string,
  };

  type ApiList = {
    id: number,
    name: string,
    method?: string,
    path?: string,
    desc?: string,
    perms_tag?: string,
    parent_id?: number,
    creator?: string,
    children?: ApiList[],
  };

  type DictList = {
    id: number,
    dict_key: string,
    dict_value: any,
    sort?: number,
    desc?: string,
    parent_id?: number,
    creator?: string,
    children?: DictList[],
  };

  type OperLogList = {
    id: number,
    name: string,
    path: string,
    method: string,
    params: string,
    body: string,
    data: string,
    status: number,
    username: string,
    ip: string,
    ip_location: string,
    latency: string,
    user_agent: string,
    created_at: string,
  }

  type HostInfo = {
    id: number,
    host_name: string,
    ip_address: string,
    port: string,
    os_version: string,
    host_type: string,
    auth_type: string,
    user: string,
    privatekey: string,
    creator: string,
  }
 
  type TtyList = {
    id:string,
    hostname:string,
    ipaddr:string,
    port:string,
    actKey:string,
    secKey:string|null,
  }

  type HostList = {
    id: number,
    host_name: string,
    ip_address: string,
    port: string,
    os_version: string,
    host_type: string,
    auth_type: string,
    user: string,
    privatekey?: string,
    creator?: string,
  }

  type HostGroupList = {
    id: number,
    name: string,
    desc?: string,
    hosts_id: number[],
    creator?: string,
  }

  type RecordList = {
    id: number,
    connect_id: string,
    user_name: string,
    host_name: string,
    ip_address: string,
    connect_time: string,
    logout_time: string,
  }

  type SSHFileList = {
    name: string,
    path: string,
    isDir: boolean,
    mode: string,
    size: string,
    mtime: string,
    isLink: boolean,
  }

  type SessionList = {
    connect_id: string,
    user_name: string,
    host_name: string,
    ip_address: string,
    connect_time: string,
  }
}
