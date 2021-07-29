import React, { useEffect, useState } from 'react';
import { queryDepts, updateDept } from '@/services/anew/dept';
import ProForm, { ModalForm, ProFormText, ProFormSelect, ProFormDigit } from '@ant-design/pro-form';
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
    values: API.DeptList | undefined;
};

const UpdateForm: React.FC<UpdateFormProps> = (props) => {
    const { actionRef, modalVisible, onChange, values } = props;
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
            title="更新部门"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async (v) => {
                updateDept(v as any, values?.id).then((res) => {
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
                <ProFormDigit
                    name="sort"
                    label="排序"
                    width="md"
                    initialValue={values?.sort}
                    fieldProps={{ precision: 0 }}
                />
            </ProForm.Group>
            <ProForm.Group>
                <Form.Item label="上级部门" name="parent_id" initialValue={values?.parent_id}>
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

export default UpdateForm;