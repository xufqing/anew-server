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

  type UpdateUserInfoParams = {
    mobile?: string;
    name?: string;
    email?: string;
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
    perms?: string[],
  }
}
