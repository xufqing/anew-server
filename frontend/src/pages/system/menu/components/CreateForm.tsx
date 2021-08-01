import React, { useEffect, useState } from 'react';
import { queryMenus, createMenu } from '@/services/anew/menu';
import ProForm, { ModalForm, ProFormText, ProFormDigit } from '@ant-design/pro-form';
import { message, TreeSelect, Form } from 'antd';
import type { ActionType } from '@ant-design/pro-table';

// 处理返回的树数据
const loopTreeItem = (tree: API.MenuList[]): API.MenuList[] =>
    tree.map(({ children, ...item }) => ({
        ...item,
        title: item.name,
        value: item.id,
        children: children && loopTreeItem(children),
    }));

export type CreateFormProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
};

const CreateForm: React.FC<CreateFormProps> = (props) => {
    const { actionRef, modalVisible, handleChange } = props;
    const [treeData, setTreeData] = useState<API.MenuList[]>([]);

    useEffect(() => {
        queryMenus().then((res) => {
            const top: API.MenuList = { id: 0, name: '顶级菜单'};
            res.data.unshift(top)
            const menus = loopTreeItem(res.data);
            setTreeData(menus);
        });
    }, []);

    return (
        <ModalForm
            title="新建菜单"
            visible={modalVisible}
            onVisibleChange={handleChange}
            onFinish={async (v) => {
                createMenu(v as any).then((res) => {
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
                <ProFormText name="name" label="名称" width="md" rules={[{ required: true }]} />
                <ProFormText name="icon" label="图标" width="md" />
                <ProFormText name="path" label="路径" width="md" />
                <ProFormDigit name="sort" label="排序" width="md" fieldProps={{ precision: 0 }} />
            </ProForm.Group>
            <ProForm.Group>
                <Form.Item label="上级菜单" name="parent_id">
                    <TreeSelect
                        style={{ width: 330 }}
                        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                        treeData={treeData}
                        placeholder="请选择菜单"
                    />
                </Form.Item>
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;