import React, { useEffect, useState } from 'react';
import { queryDicts, createDict } from '@/services/anew/dict';
import ProForm, { ModalForm, ProFormText, ProFormDigit } from '@ant-design/pro-form';
import { message, TreeSelect, Form } from 'antd';
import type { ActionType } from '@ant-design/pro-table';

// 处理返回的树数据
const loopTreeItem = (tree: API.DictList[]): API.DictList[] =>
    tree.map(({ children, ...item }) => ({
        ...item,
        title: item.dict_value,
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
    const [treeData, setTreeData] = useState<API.DictList[]>([]);

    useEffect(() => {
        queryDicts().then((res) => {
            const top: API.DictList = { id: 0, dict_value: '暂无所属', dict_key: '' };
            res.data.unshift(top)
            const dicts = loopTreeItem(res.data);
            setTreeData(dicts);
        });
    }, []);

    return (
        <ModalForm
            title="新建字典"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async (v) => {
                createDict(v as any).then((res) => {
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
                <ProFormText name="dict_key" label="字典标签" width="md" rules={[{ required: true }]} />
                <ProFormText name="dict_value" label="字典键值" width="md" rules={[{ required: true }]} />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormDigit name="sort" label="排序" width="md" fieldProps={{ precision: 0 }} />
                <ProFormText name="desc" label="说明" width="md" />
            </ProForm.Group>
            <ProForm.Group>
                <Form.Item label="上级字典" name="parent_id">
                    <TreeSelect
                        style={{ width: 330 }}
                        dropdownStyle={{ maxHeight: 400, overflow: 'auto' }}
                        treeData={treeData}
                        placeholder="请选择字典"
                    />
                </Form.Item>
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;