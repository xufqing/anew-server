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

  type CreateUserParams = {
    username: string,
    password: string,
    mobile?: string,
    avatar?: string,
    name: string,
    email?: string,
    dept_id?: number,
    role_id: number,
    creator: string,

  }

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
    parent_id: number,
    name: string,
    sort: number,
    status: boolean,
    creator: string,
    title?: string,
    value?: number,
    children: DeptList[],
  };
}
