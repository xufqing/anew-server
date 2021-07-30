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
    key: string,
    value: string,
    desc?: string,
    parent_id?: number,
  };

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
    key: string,
    value: any,
    desc?: string,
    parent_id?: number,
    creator?: string,
    children?: DictList[],
  };
}
