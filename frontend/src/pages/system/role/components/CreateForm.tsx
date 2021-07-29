import React from 'react';
import { createRole } from '@/services/anew/role';
import ProForm, { ModalForm, ProFormText } from '@ant-design/pro-form';
import { message } from 'antd';
import type { ActionType } from '@ant-design/pro-table';

export type CreateFormProps = {
    modalVisible: boolean;
    onChange: (modalVisible: boolean) => void;
    actionRef: React.MutableRefObject<ActionType | undefined>;
};

const CreateForm: React.FC<CreateFormProps> = (props) => {
    const { actionRef, modalVisible, onChange } = props;

    return (
        <ModalForm
            title="新建角色"
            visible={modalVisible}
            onVisibleChange={onChange}
            onFinish={async (v) => {
                createRole(v as any).then((res) => {
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
                <ProFormText name="keyword" label="关键字" width="md" />
            </ProForm.Group>
            <ProForm.Group>
                <ProFormText name="desc" label="说明" width="md" />
            </ProForm.Group>
        </ModalForm>
    );
};

export default CreateForm;