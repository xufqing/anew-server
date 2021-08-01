import React, { useEffect, useState } from 'react';
import { queryDepts, createDept } from '@/services/anew/dept';
import ProForm, { ModalForm, ProFormText, ProFormDigit } from '@ant-design/pro-form';
import { message, TreeSelect, Form } from 'antd';
import type { ActionType } from '@ant-design/pro-table';


const loopTreeItem = (tree: API.DeptList[]): API.DeptList[] =>
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
            title="新建部门"
            visible={modalVisible}
            onVisibleChange={handleChange}
            onFinish={async (v) => {
                createDept(v as any).then((res) => {
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
                <ProFormDigit name="sort" label="排序" width="md" fieldProps={{ precision: 0 }} />
            </ProForm.Group>
            <ProForm.Group>
                <Form.Item label="上级部门" name="parent_id" >
                    <TreeSelect
                        style={{ width: 330 }}
                        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                        treeData={treeData}
                        placeholder="请选择部门"
                    />
                </Form.Item>
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;