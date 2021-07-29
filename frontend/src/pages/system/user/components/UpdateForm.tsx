import React, { useEffect, useState } from 'react';
import { queryRoles } from '@/services/anew/role';
import { queryDepts } from '@/services/anew/dept';
import { updateUser } from '@/services/anew/user';
import ProForm, { ModalForm, ProFormText, ProFormSelect } from '@ant-design/pro-form';
import { message, TreeSelect, Form } from 'antd';
import type { ActionType } from '@ant-design/pro-table';

// 处理返回的树数据
const loopTreeItem = (tree: API.DeptList[]): API.DeptList[] =>
    tree.map(({ children, ...item }) => ({
        ...item,
        title: item.name,
        value: item.id,
        children: children && loopTreeItem(children),
    }));

export type UpdateFormProps = {
    modalVisible: boolean;
    onChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
    values: API.UserList | undefined;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
    const { actionRef, modalVisible, onChange, values } = props;
    const [treeData, setTreeData] = useState<API.DeptList[]>([]);
    useEffect(() => {
        queryDepts().then((res) => {
            const top: API.DeptList = { id: 0, name: '暂无所属'};
            res.data.unshift(top)
            const depts = loopTreeItem(res.data);
            setTreeData(depts);
        });
    }, []);

    return (
        <ModalForm
            title="更新用户"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async (v) => {
                updateUser(v as any, values?.id).then((res) => {
                    if (res.code === 200 && res.status === true) {
                        message.success(res.message);
                        if (actionRef.current) {
                            actionRef.current.reload();
                        }
                    }
                });
                return true;
            }}
        >
            <ProForm.Group>
                <ProFormText name="username" label="用户名" width="md" initialValue={values?.username} rules={[{ required: true }]} />
                <ProFormText name="name" label="姓名" width="md" initialValue={values?.name} rules={[{ required: true }]} />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText
                    name="mobile"
                    label="手机"
                    width="md"
                    initialValue={values?.mobile}
                    rules={[
                        {
                            pattern: /^1(?:70\d|(?:9[89]|8[0-24-9]|7[135-8]|66|5[0-35-9])\d|3(?:4[0-8]|[0-35-9]\d))\d{7}$/,
                            message: '请输入正确的手机号码',
                        },
                    ]}
                />
                <ProFormText
                    name="email"
                    label="邮箱"
                    width="md"
                    initialValue={values?.email}
                    rules={[
                        {
                            type: 'email',
                            message: '请输入正确的邮箱地址',
                        },
                    ]}
                />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormSelect
                    name="role_id"
                    label="角色"
                    width="md"
                    hasFeedback
                    initialValue={values?.role.id}
                    request={() =>
                        queryRoles().then((res) =>
                            res.data.data.map((item: any) => ({
                                label: item.name,
                                value: item.id,
                            })),
                        )
                    }
                    rules={[{ required: true, message: '请选择角色' }]}
                />
                <Form.Item label="部门" name="dept_id" initialValue={values?.dept.id}>
                    <TreeSelect
                        style={{ width: 330 }}
                        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                        treeData={treeData}
                        placeholder="请选择部门"
                    />
                </Form.Item>

            </ProForm.Group>
            <ProForm.Group>
                <ProFormSelect
                    name="status"
                    label="状态"
                    width="md"
                    hasFeedback
                    initialValue={values?.status}
                    options={[
                        {
                            value: true as any,
                            label: '激活',
                        },
                        {
                            value: false as any,
                            label: '禁用',
                        },
                    ]}
                    rules={[{ required: true, message: '请选择状态' }]}
                />
                <ProFormText.Password label="重置密码" name="password" width="md" placeholder="输入则修改密码" />
            </ProForm.Group>

        </ModalForm>
    );
};

export default UpdateForm;