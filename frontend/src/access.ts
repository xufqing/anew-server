/**
 * @see https://umijs.org/zh-CN/plugins/plugin-access
 * */
export default function access(initialState: { currentUser?: API.UserInfo | undefined }) {
  const { currentUser } = initialState || {};
  return {
    hasPerms: (tags: string[]) => {
      for (let tag of tags) {
        if (currentUser?.perms.includes(tag)) {
          return true;
        }
      }
      return false;
    },
  }
}
