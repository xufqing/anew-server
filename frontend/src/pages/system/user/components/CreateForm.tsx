import React, { useEffect, useState } from 'react';
import { queryRoles } from '@/services/anew/role';
import { queryDepts } from '@/services/anew/dept';
import { createUser } from '@/services/anew/user';
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

export type CreateFormProps = {
    modalVisible: boolean;
    onChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
};

const CreateForm: React.FC<CreateFormProps> = (props) => {
    const { actionRef, modalVisible, onChange } = props;
    const [treeData, setTreeData] = useState<API.DeptList[]>([]);

    useEffect(() => {    
        queryDepts().then((res) => {
            const top: API.DeptList = { id: 0, name: '暂无所属' };
            res.data.unshift(top)
            const depts = loopTreeItem(res.data);
            setTreeData(depts);
        });
    }, []);

    return (
        <ModalForm
            title="新建用户"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async (v) => {
                createUser(v as any).then((res) => {
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
                <ProFormText name="username" label="用户名" width="md" rules={[{ required: true }]} />
                <ProFormText name="name" label="姓名" width="md" rules={[{ required: true }]} />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText
                    name="mobile"
                    label="手机"
                    width="md"
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
                {/* <Form.Item label="部门" name="dept_id" width="md"> */}
                <Form.Item label="部门" name="dept_id">
                    <TreeSelect
                        style={{ width: 330 }}
                        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                        treeData={treeData}
                        placeholder="请选择部门"
                    />
                </Form.Item>
                <ProFormText.Password label="密码" name="password" width="md" rules={[{ required: true }]} />
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;