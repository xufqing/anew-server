import React, { useEffect, useState } from 'react';
import { queryMenus, updateMenu } from '@/services/anew/menu';
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

export type UpdateFormProps = {
    modalVisible: boolean;
    handleChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
    values: API.MenuList | undefined;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
    const { actionRef, modalVisible, handleChange, values } = props;
    const [treeData, setTreeData] = useState<API.MenuList[]>([]);
    useEffect(() => {
        queryMenus().then((res) => {
            const top: API.MenuList = { id: 0, name: '顶级菜单' };
            res.data.unshift(top)
            const menus = loopTreeItem(res.data);
            setTreeData(menus);
        });
    }, []);

    return (
        <ModalForm
            title="更新菜单"
            visible={modalVisible}
            onVisibleChange={handleChange}
            onFinish={async (v) => {
                updateMenu(v as any, values?.id).then((res) => {
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
                <ProFormText
                    name="name"
                    label="名称"
                    width="md"
                    initialValue={values?.name}
                    rules={[{ required: true }]}
                />
                <ProFormText name="icon" label="图标" width="md" initialValue={values?.icon} />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText name="path" label="路径" width="md" initialValue={values?.path} />
                <ProFormDigit
                    name="sort"
                    label="排序"
                    width="md"
                    initialValue={values?.sort}
                    fieldProps={{ precision: 0 }}
                />
            </ProForm.Group>
            <ProForm.Group>
                <Form.Item label="上级菜单" name="parent_id" initialValue={values?.parent_id}>
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

export default UpdateForm;