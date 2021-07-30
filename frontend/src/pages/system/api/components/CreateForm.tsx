import React, { useEffect, useState } from 'react';
import { queryApis, createApi } from '@/services/anew/api';
import ProForm, { ModalForm, ProFormText } from '@ant-design/pro-form';
import { message, TreeSelect, Form } from 'antd';
import type { ActionType } from '@ant-design/pro-table';

// 处理返回的树数据
const loopTreeItem = (tree: API.ApiList[]): API.ApiList[] =>
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
    const [treeData, setTreeData] = useState<API.ApiList[]>([]);

    useEffect(() => {
        queryApis().then((res) => {
            const top: API.ApiList = { id: 0, name: '暂无所属' };
            res.data.unshift(top)
            const menus = loopTreeItem(res.data);
            setTreeData(menus);
        });
    }, []);

    return (
        <ModalForm
            title="新建接口"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async (v) => {
                createApi(v as any).then((res) => {
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
                <ProFormText name="perms_tag" label="权限标识" width="md" rules={[{ required: true }]} />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText name="method" label="请求方式" width="md" />
                <ProFormText name="path" label="路径" width="md" />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText name="desc" label="说明" width="md" />
                <Form.Item label="上级接口" name="parent_id">
                    <TreeSelect
                        style={{ width: 330 }}
                        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                        treeData={treeData}
                        placeholder="请选择接口"
                    />
                </Form.Item>
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;